package main

import (
	"context"
	"database/sql"
	"fmt"
	"learn/consumer"
	"learn/mail"
	"learn/producer"
	"learn/scheduler"
	"os"
	"os/signal"
	"sync"

	"learn/rabbitmq"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func main() {
	// prepare params
	apiKey := "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"
	msgExchange := make(chan *mail.EmailContent)

	// prepare db
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/ecommerce")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// rmq
	rmqConfig := rabbitmq.RabbitMQConfig{
		RmqURL:     "amqp://guest:guest@localhost:5672/",
		Exch:       "order",
		ExchType:   "direct",
		Queue:      "order_processor",
		RoutingKey: "",
	}

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

	wg := &sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())

	// sql
	mailer := mail.NewSendGrid(apiKey)
	sched := scheduler.NewScheduler(ctx, db, msgExchange)

	producer := producer.NewProducer(ctx, wg, pCh, rmqConfig, msgExchange)
	consumer := consumer.NewConsumer(ctx, wg, cCh, rmqConfig, mailer, db)

	// graceful shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c
		fmt.Printf("Got %s signal. Exiting...\n", sig)
		sched.Stop()
		cancelFunc()
	}()

	///
	wg.Add(2)
	go producer.Start()
	go consumer.Start()
	sched.Start()
	wg.Wait()
}
