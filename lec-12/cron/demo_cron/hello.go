package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Order struct {
	Id                  int
	CreatedAt           time.Time
	HasSendEmailConfirm int
	CustomerEmail       string
}

type Scheduler struct {
	
}

func main() {
	fmt.Println("Hello, playground")
}

func NewScheduler(ctx context.Context, db *sql.DB, ch chan<- *email.EmailContent) *Scheduler {
	return &Scheduler{
		ctx:     ctx,
		db:      db,
		c:       cron.New(cron.WithSeconds()),
		outChan: ch,
	}
}

func (sched *Scheduler) Start() {
	sched.c.AddFunc("0 * * * *", sched.scheduleJob)
	sched.c.Start()
}

func (sched *Scheduler) ScheduleJob() {
	fmt.Printf("Scanning for new order(s) at %v\n", time.Now().Format("2006-Jan-02 15:04:05"))
	resp, err := sched.getEmailForSending()
	if err != nil {
		return
	}
	fmt.Printf("Scheduling %v email(s) at %v\n", len(resp), time.Now().Format("2006-Jan-02 15:04:05"))
	for _, em := range resp {
		sched.outChan <- em
	}
}

func (w *Worker) Start() {
	if w.mailer == nil || w.db == nil {
		fmt.Println("cannot start worker since mailer is nil")
		return
	}
	for {
		select {
		case em := <-w.inChan:
			err := w.mailer.Send(em)
			if err != nil {
				fmt.Println("Cannot send email due to error: ", err)
				continue
			}
			// update sql data
			_, err = w.db.Exec("UPDATE `order` SET thankyou_email_sent = ? WHERE id = ?", true, em.ID)
			if err != nil {
				fmt.Println("Cannot update thankyou_email_sen to true")
			}
		case <-w.ctx.Done():
			fmt.Println("Exiting worker")
			w.wg.Done()
			return
		}
	}
}
