package scheduler

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"worker/mail"

	"github.com/robfig/cron/v3"
)

const (
	DefaultThankyouSubject   = "Thank you vinamilk for purchasing from myshop.com"
	DefaultThankyouBodyPlain = "Thank you for purchasing from our shop. Here's your order details:"
	DefaultThankyouBodyHtml  = "<strong>Thank you for purchasing from our store. Here's your order details:</strong>"
	DefaultFromName          = "Hieu hoc code"
	DefaultFromEmail         = "duchieuctk41@gmail.com"
)

type Scheduler struct {
	db      *sql.DB
	c       *cron.Cron
	outChan chan<- *mail.EmailContent
	ctx     context.Context
}

func NewScheduler(ctx context.Context, db *sql.DB, ch chan<- *mail.EmailContent) *Scheduler {
	return &Scheduler{
		ctx:     ctx,
		db:      db,
		c:       cron.New(cron.WithSeconds()),
		outChan: ch,
	}
}

func (sched *Scheduler) Start() {
	sched.c.AddFunc("0 * * * * *", sched.scheduleJob)
	sched.c.Start()
}

func (sched *Scheduler) Stop() {
	fmt.Println("Stopping scheduler")
	sched.c.Stop()
}

func (sched *Scheduler) scheduleJob() {
	fmt.Printf("Scanning for new order(s) at %v\n", time.Now().Format("2006-Jan-02 15-04-05"))
	resp, err := sched.getEmailForSending()
	if err != nil {
		return
	}

	fmt.Printf("Scheduling %v email(s) at %v\n", len(resp), time.Now().Format("2006-Jan-02 15-04-05"))
	for _, em := range resp {
		sched.outChan <- em
	}

}

func (sched *Scheduler) getEmailForSending() ([]*mail.EmailContent, error) {
	resp, err := sched.scanFromDB()
	if err != nil {
		return resp, err
	}

	for _, emailContent := range resp {
		emailContent.FromUser = &mail.EmailUser{
			Name:  DefaultFromName,
			Email: DefaultFromEmail,
		}
	}
	return resp, err
}

func (sched *Scheduler) scanFromDB() ([]*mail.EmailContent, error) {
	var resp []*mail.EmailContent
	fromTime := time.Now().Add(-time.Minute * 2) // subtract by 2 minutes - why not one?
	// What is prepared statement? Why we should know and use that? is the below usage right? Why not?
	stmt, err := sched.db.Prepare("SELECT id, customer_name, email FROM `order` WHERE created_at >= ? AND thankyou_email_sent = ?;")
	if err != nil {
		fmt.Println("Cannot prepare statement, ", err)
		return nil, err
	}

	rows, err := stmt.Query(fromTime, false)
	if err != nil || rows == nil {
		fmt.Printf("Cannot query from db due to error: %v, %v\n", err, rows == nil)
		return nil, err
	}

	// MUST to call this function at the end to free connection to mysql
	defer rows.Close()

	var id int64
	var email, name string
	for rows.Next() {
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println("Cannot scan row due to error: ", err)
			continue
		}
		resp = append(resp, &mail.EmailContent{
			Id:               id,
			Subject:          DefaultThankyouSubject,
			PlainTextContext: DefaultThankyouBodyPlain,
			HtmlContent:      DefaultThankyouBodyHtml,
			ToUser: &mail.EmailUser{
				Name:  name,
				Email: email,
			},
		})
	}
	return resp, nil
}
