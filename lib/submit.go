package lib

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"github.com/KeKsBoTer/socialloot/models"
)

func Submit(title, content string, postType models.PostType, topic *models.Topic, user *models.User) (*models.Post, error) {
	post := models.Post{
		Id:    GetRandomString(models.ItemIDLength),
		Title: title,
		Topic: topic,
		Type:  postType,
		User:  user,
	}
	switch postType {
	case models.PostTypeText:
		post.Content = content
		// do text stuff, like escaping or stuff...
	case models.PostTypeLink:
		if parsed, err := url.ParseRequestURI(content); err != nil {
			return nil, errors.New("Invalid url format for link")
		} else {
			post.Content = parsed.String()
		}
	case models.PostTypeImage:
		fileType := http.DetectContentType([]byte(content))
		mediaImage, err := parseImage(&content, fileType)
		if err != nil {
			beego.Error(err)
			return nil, errors.New("Cannot decode/encode image")
		}

		if err := mediaImage.Insert(); err != nil {
			beego.Error(err)
			return nil, errors.New("Cannot save image")
		}
		if mediaImage.Id == 0 {
			return nil, errors.New("Image Id not available")
		}
		post.Content = strconv.Itoa(mediaImage.Id)
	default:
		return nil, errors.New("Invalid post type")
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

func VoteOnItem(dir models.VoteDirection, postId string, user *models.User) error {
	vote := models.Vote{
		User:   user,
		Action: dir,
		Item:   postId,
	}
	if vote.User == nil {
		return errors.New("No user provided")
	}
	if err := vote.InsertOrUpdate(); err != nil {
		beego.Error(err)
		return errors.New("cannot vote on post")
	}
	return nil
}

func CommentOnPost(text string, replyTo string, user *models.User) error {

	if len(text) < 1 {
		return errors.New("Comment cannot be empty")
	}

	// check if replyTo item exists
	replyToComment := models.Comment{
		Id: replyTo,
	}
	if err := replyToComment.Read("Id"); err != nil {
		post := models.Post{
			Id: replyTo,
		}
		if err := post.Read(false); err != nil {
			return errors.New("Cannot comment on non existent item")
		}
	}
	comment := models.Comment{
		Id:      GetRandomString(models.ItemIDLength),
		User:    user,
		Text:    text,
		ReplyTo: replyTo,
	}
	return comment.Insert()
}
