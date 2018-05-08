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

// Submit creates a new post and adds it to the database.
// The function checks the content based on the post type.
// Links: Is the inserted link a valid absolute url?
// Images: Does the image have the right format?
// Also a small thumbnail for the image is created.
func Submit(title, content string, postType models.PostType, topic *models.Topic, user *models.User) (*models.Post, error) {
	post := models.Post{
		Title: title,
		Topic: topic,
		Type:  postType,
		User:  user,
	}
	switch postType {
	case models.PostTypeText:
		post.Content = content

	case models.PostTypeLink:
		parsed, err := url.ParseRequestURI(content)
		if err != nil || !parsed.IsAbs() {
			return nil, errors.New("Invalid url format for link")
		}
		post.Content = parsed.String()

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
		// change content to the image id
		post.Content = strconv.Itoa(mediaImage.Id)

	default:
		return nil, errors.New("Invalid post type")
	}
	// get new item id
	post.Id = GenerateItemID()

	if err := post.Insert(); err != nil {
		beego.Error(err)
		return nil, errors.New("Cannot submit post")
	}
	return &post, nil
}

// CreateTopic adds a new topic to the database.
// The topic name is turned to lower case before beeing inserted into the database
// The returned topic object is a copy of the created database entry.
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
		// this should not happen, since we just created the topic
		beego.Error(err)
		return nil, errors.New("An unexpected error occured")
	}
	return &topic, nil
}

// VoteOnItem adds a vote to the database.
// Item can be a comment or post.
// The vote direction is saved as action.
func VoteOnItem(dir models.VoteDirection, itemID string, user *models.User) error {
	vote := models.Vote{
		User:   user,
		Action: dir,
		Item:   itemID,
	}
	if err := vote.InsertOrUpdate(); err != nil {
		beego.Error(err)
		return errors.New("cannot vote on item")
	}
	return nil
}

// CommentOnPost adds a new comment to the database
// First the methode checks if the text is valid and
// the item, which the comment replies to, exists.
// Also the id for the comment is generated.
func CommentOnPost(text string, replyTo string, user *models.User) error {

	if len(text) < 1 {
		return errors.New("Comment cannot be empty")
	}

	// check if replyTo item exists
	replyToComment := models.Comment{
		Id: replyTo,
	}
	if err := replyToComment.Read("Id"); err != nil {
		// if item was not found, check if it's a post
		post := models.Post{
			Id: replyTo,
		}
		if err := post.Read(false); err != nil {
			return errors.New("Cannot comment on non existent item")
		}
	}
	comment := models.Comment{
		Id:      GenerateItemID(),
		User:    user,
		Text:    text,
		ReplyTo: replyTo,
	}
	return comment.Insert()
}
