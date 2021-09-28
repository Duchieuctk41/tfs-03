package rabbitmq

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	RmqURL     string
	Exch       string
	ExchType   string
	Queue      string
	RoutingKey string
}

type RMQ struct {
	URL        string `json:"url"`
	connection *amqp.Connection
}

func NewRMQ(url string) *RMQ {
	conn, err := amqp.Dial(url)
	if err != nil {
		logrus.Error("Connet to RabbitMQ fail: ", err)
		return nil
	}

	return &RMQ{
		URL:        url,
		connection: conn,
	}
}

func (rmq *RMQ) GetChannel() (*amqp.Channel, error) {
	return rmq.connection.Channel()
}

func (rmq *RMQ) Close() {
	rmq.connection.Close()
}
