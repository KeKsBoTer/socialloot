package lib

import (
	"errors"
	"log"
	"time"

	"github.com/KeKsBoTer/socialloot/models"
)

func Submit(p *models.Post) error {
	var (
		err error
		msg string
	)

	if p.User == nil {
		msg = "no user provided"
		return errors.New(msg)
	}
	p.Date = time.Now()
	p.Id = GetRandomString(11)
	p.Topic = &models.Topic{
		Name: p.TopicName,
	}

	if err = models.Topics().Filter("name", p.TopicName).One(p.Topic); err != nil {
		log.Println(err)
		msg = "topic does not exist"
		return errors.New(msg)
	}

	err = p.Insert()
	if err != nil {
		return err
	}

	return nil
}

func CreateTopic(t *models.Topic) error {
	return t.Insert()
}
