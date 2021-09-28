package producer

import (
	"buoi11/model"
	rabbitmq "buoi11/rmq"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Producer struct {
	ctx        context.Context
	wg         *sync.WaitGroup
	channel    *amqp.Channel
	exchange   string
	exchType   string
	routingKey string
	mailChan   chan *model.EmailContent
}

//Creates new producer
func NewProducer(ctx context.Context, wg *sync.WaitGroup, channel *amqp.Channel, rmqConfig rabbitmq.RabbitMQConfig, mailChan chan *model.EmailContent) *Producer {
	return &Producer{
		ctx:        ctx,
		wg:         wg,
		channel:    channel,
		exchange:   rmqConfig.Exch,
		exchType:   rmqConfig.ExchType,
		routingKey: rmqConfig.RoutingKey,
		mailChan:   mailChan,
	}
}

// Start generating data
func (p *Producer) Start() {
	if p.channel == nil || p.exchange == "" || p.exchType == "" {
		logrus.Error("wrong producer config")
		return
	}

	p.declare()

	for {
		select {
		case d := <-p.mailChan:
			logrus.Info("Sending mail infomation to consumer ", d.ToUser.Name)
			data, _ := json.Marshal(d)
			err := p.publish(p.exchange, p.routingKey, data)
			if err != nil {
				logrus.Error("error when publishing data: ", err)
			}
		case <-p.ctx.Done():
			logrus.Info("Exiting Producer")
			p.wg.Done()
			return
		}
	}

}

func (p *Producer) publish(exch, routingKey string, body []byte) error {
	if err := p.channel.Publish(
		exch,
		routingKey,
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient,
		},
	); err != nil {
		return fmt.Errorf("publish data error: %s", err)
	}
	return nil
}

func (p *Producer) declare() error {
	//declare exchange
	logrus.Info("Binding exchange ", p.exchange)
	if err := p.channel.ExchangeDeclare(
		p.exchange,
		p.exchType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("exchange declare error: %s", err)
	}
	return nil
}

func (c *Producer) Close() error {
	return c.channel.Close()
}
