package lib

import (
	"errors"
	"strings"
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
		Name: strings.ToLower(p.TopicName),
	}
	if err = p.Topic.Read("Name"); err != nil {
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
	t.Name = strings.ToLower(t.Name)
	return t.Insert()
}

func VoteOnPost(action models.UserVote, postId string, user *models.User) error {
	vote := models.Vote{
		User:   user,
		Action: action,
		Item:   postId,
		Type:   "post",
	}
	if err := models.IsValid(vote); err != nil {
		return err
	}
	return vote.InsertOrUpdate()
}
