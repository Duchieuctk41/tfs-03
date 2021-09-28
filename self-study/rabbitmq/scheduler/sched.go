package scheduler

import (
	"buoi11/model"
	"context"
	"database/sql"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Scheduler struct {
	c       *cron.Cron
	db      *sql.DB
	outchan chan<- *model.EmailContent
	ctx     context.Context
}

func NewScheduler(ctx context.Context, db *sql.DB, ch chan<- *model.EmailContent) *Scheduler {
	return &Scheduler{
		ctx:     ctx,
		db:      db,
		c:       cron.New(cron.WithSeconds()),
		outchan: ch,
	}
}

func (sched *Scheduler) Start() {
	sched.c.AddFunc("*/10 * * * * *", sched.scheduleJob)
	sched.c.Start()
}

func (sched *Scheduler) Stop() {
	logrus.Info("Stopping scheduler")
	sched.c.Stop()
}

func (sched *Scheduler) scheduleJob() {
	logrus.Info("Scan db...")
	resp, err := sched.getEmail()
	if err != nil {
		return
	}
	logrus.Info("Send to email Info to chan")
	for _, email := range resp {
		sched.outchan <- email
	}
}

func (sched *Scheduler) getEmail() ([]*model.EmailContent, error) {
	var resp []*model.EmailContent
	rows, err := sched.db.Query("SELECT id, customer_name, email FROM ecommerce.orders WHERE thank_you_email_sent = ?", false)
	if err != nil || rows == nil {
		logrus.Error("cannot query get email", err)
		return nil, err
	}
	defer rows.Close()
	var id int64
	var email, name string
	var User = model.ShopUser{
		Name:  "hieu chu nhat",
		Email: "duchieu.ctk41@gmail.com",
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			logrus.Error("cannot scan row: ", err)
			continue
		}
		resp = append(resp, &model.EmailContent{
			ID:               id,
			Subject:          "thank you for purchasing from our store",
			PlainTextContent: "Thank you for purchasing from our store. Here's your order details: ",
			HtmlContent:      "<strong>Thank you for purchasing from our store. Here's your order details:</strong>",
			FromUser:         User,
			ToUser: model.User{
				ID:    id,
				Name:  name,
				Email: email,
			},
		})
	}
	return resp, nil
}
