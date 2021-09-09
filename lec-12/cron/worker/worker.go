package worker

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"worker/mail"
)

type Worker struct {
	wg     *sync.WaitGroup
	mailer mail.Mailer
	inChan <-chan *mail.EmailContent
	ctx    context.Context
	db     *sql.DB
}

func NewWorker(ctx context.Context, wg *sync.WaitGroup, db *sql.DB, mailer mail.Mailer, ch <-chan *mail.EmailContent) *Worker {
	return &Worker{
		ctx:    ctx,
		wg:     wg,
		mailer: mailer,
		inChan: ch,
		db:     db,
	}
}

func (w *Worker) Start() {
	if w.mailer == nil || w.db == nil {
		fmt.Println("Cannot start worker since mailer is nil")
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

			_, err = w.db.Exec("UPDATE `order` SET thankyou_email_sent = ? WHERE id = ?", true, em.Id)
			if err != nil {
				fmt.Println("Cannot update thankyou_email_sent to true")
			}

		case <-w.ctx.Done():
			fmt.Println("Exiting worker")
			w.wg.Done()
			return
		}
	}
}
