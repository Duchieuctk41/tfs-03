package main

import (
	"buoi11/consumer"
	"buoi11/db"
	"buoi11/model"
	"buoi11/producer"
	rabbitmq "buoi11/rmq"
	schedule "buoi11/scheduler"
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/sirupsen/logrus"
)

func main() {
	rmqConfig := rabbitmq.RabbitMQConfig{
		RmqURL:     "amqp://guest:guest@localhost:5672/",
		Exch:       "order",
		ExchType:   "direct",
		Queue:      "order_processor",
		RoutingKey: "",
	}
	//Open db
	db := db.Connect()
	defer db.Close()
	//rmq
	rmq := rabbitmq.NewRMQ(rmqConfig.RmqURL)
	pCh, err := rmq.GetChannel()
	if err != nil {
		logrus.Error(err)
		return
	}
	cCh, err := rmq.GetChannel()
	if err != nil {
		logrus.Error(err)
		return
	}

	///
	wg := &sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())
	mailChan := make(chan *model.EmailContent)
	sched := schedule.NewScheduler(ctx, db, mailChan)

	producer := producer.NewProducer(ctx, wg, pCh, rmqConfig, mailChan)
	consumer := consumer.NewConsumer(ctx, wg, cCh, rmqConfig)
	///
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c
		logrus.Info("Got ", sig, " signal. Exiting...")
		sched.Stop()
		cancelFunc()
	}()
	wg.Add(2)
	go consumer.Start()
	go producer.Start()

	sched.Start()
	wg.Wait()
}
