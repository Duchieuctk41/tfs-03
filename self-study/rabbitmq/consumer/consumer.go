package consumer

import (
	email "buoi11/mail"
	rabbitmq "buoi11/rmq"
	"context"
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
}

func NewConsumer(ctx context.Context, wg *sync.WaitGroup, chanel *amqp.Channel, rmqConfig rabbitmq.RabbitMQConfig) *Consumer {
	return &Consumer{
		ctx:        ctx,
		wg:         wg,
		chanel:     chanel,
		exChange:   rmqConfig.Exch,
		exchType:   rmqConfig.ExchType,
		bindingKey: rmqConfig.RoutingKey,
		queue:      rmqConfig.Queue,
	}
}

func (c *Consumer) Start() {
	if c.chanel == nil || c.queue == "" {
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
		case d := <-msgs:
			email.SendMail(d.Body)
			d.Ack(false)
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
