package consumer

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"learn/mail"
	"learn/rabbitmq"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Consumer struct {
	ctx        context.Context
	wg         *sync.WaitGroup
	chanel     *amqp.Channel
	exChange   string
	exchType   string
	bindingKey string
	queue      string
	mailer     mail.Mailer
	db         *sql.DB
}

func NewConsumer(ctx context.Context, wg *sync.WaitGroup, chanel *amqp.Channel, rmqConfig rabbitmq.RabbitMQConfig, mailer mail.Mailer, db *sql.DB) *Consumer {
	return &Consumer{
		ctx:        ctx,
		wg:         wg,
		chanel:     chanel,
		exChange:   rmqConfig.Exch,
		exchType:   rmqConfig.ExchType,
		bindingKey: rmqConfig.RoutingKey,
		queue:      rmqConfig.Queue,
		mailer:     mailer,
		db: db,
	}
}

func (c *Consumer) Start() {
	if c.chanel == nil || c.queue == "" || c.mailer == nil || c.db == nil {
		logrus.Error("wrong consumer config")
		return
	}

	c.declare()

	logrus.Info("Queue is bound to exchange. Consuming data now")
	msgs, err := c.chanel.Consume(
		c.queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		logrus.Error("Queue consume error", err)
	}

	for {
		select {
		case em := <-msgs:
			em.Ack(false)
			var sendMail mail.EmailContent
			_ = json.Unmarshal(em.Body, &sendMail)
			err := c.mailer.Send(&sendMail)
			if err != nil {
				fmt.Println("Cannot send email due to error: ", err)
				continue
			}
			// update sql data
			_, err = c.db.Exec("UPDATE ecommerce.orders SET thank_you_email_sent = true WHERE id = ?", true, sendMail.ID)
			if err != nil {
				fmt.Println("Cannot update thankyou_email_sent to true")
			}
		case <-c.ctx.Done():
			logrus.Info("Exiting consumer")
			c.wg.Done()
			return
		}
	}
}

//declare exchange and queue, also bind queue to exchange
func (c *Consumer) declare() error {
	//declare exchange
	logrus.Info("Binding exchange: ", c.exChange)
	if err := c.chanel.ExchangeDeclare(
		c.exChange,
		c.exchType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	//declare queue
	logrus.Info("Declare queue : ", c.queue)
	queue, err := c.chanel.QueueDeclare(
		c.queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	// binding queue
	logrus.Info("Binding queue ", c.queue, "to exchange ", c.exChange)
	if err := c.chanel.QueueBind(
		queue.Name,
		c.bindingKey,
		c.exChange,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (c *Consumer) Close() error {
	return c.chanel.Close()
}
