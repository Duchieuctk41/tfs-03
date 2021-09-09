package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"worker/mail"
	"worker/scheduler"
	"worker/worker"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// prepare params
	apiKey := "apiKey"
	msgExchange := make(chan *mail.EmailContent)

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/meow")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	wg := &sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())

	// sql
	mailer := mail.NewSendGrid(apiKey)
	sched := scheduler.NewScheduler(ctx, db, msgExchange)
	worker := worker.NewWorker(ctx, wg, db, mailer, msgExchange)

	// graceful shutdown
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		sig := <-c
		fmt.Printf("Got %s signal. Exiting...\n", sig)
		sched.Stop()
		cancelFunc()
	}()

	//
	wg.Add(1)

	go worker.Start()

	sched.Start()

	wg.Wait()

}
