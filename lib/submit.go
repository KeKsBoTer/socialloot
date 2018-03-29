package lib

import (
	"errors"
	"strings"

	"github.com/astaxie/beego"

	"github.com/KeKsBoTer/socialloot/models"
)

func Submit(title, content string, topic *models.Topic, user *models.User) (*models.Post, error) {
	post := models.Post{
		Id:      GetRandomString(models.ItemIDLength),
		Title:   title,
		Content: content,
		Topic:   topic,
		User:    user,
	}
	if post.User == nil {
		return nil, errors.New("No user provided")
	}
	if post.Topic == nil {
		return nil, errors.New("No topic provided")
	}
	if err := post.Insert(); err != nil {
		beego.Error(err)
		return nil, errors.New("Cannot submit post")
	}
	return &post, nil
}

func CreateTopic(name, title, description string) (*models.Topic, error) {
	topic := models.Topic{
		Name:        strings.ToLower(name),
		Title:       title,
		Description: description,
	}
	if err := topic.Insert(); err != nil {
		beego.Error(err)
		return nil, errors.New("Cannot create topic")
	}
	if err := topic.Read("name"); err != nil {
		beego.Error(err)
		return nil, errors.New("An unexpected error occured")
	}
	return &topic, nil
}

func VoteOnPost(dir models.VoteDirection, postId string, user *models.User) error {
	vote := models.Vote{
		User:   user,
		Action: dir,
		Item:   postId,
		Type:   "post",
	}
	if vote.User == nil {
		return errors.New("No user provided")
	}
	if vote.Action != models.VoteDirectionUp && vote.Action != models.VoteDirectionDown {
		return errors.New("direction musst be up or downvote")
	}
	if err := vote.InsertOrUpdate(); err != nil {
		beego.Error(err)
		return errors.New("cannot vote on post")
	}
	return nil
}

func CommentOnPost(text string, post *models.Post, user *models.User) error {
	comment := models.Comment{
		Id:      GetRandomString(models.ItemIDLength),
		User:    user,
		Text:    text,
		Post:    post,
		ReplyTo: nil, //TODO allow reply to comment
	}
	return comment.Insert()
}
