package lib

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/nfnt/resize"
)

func Submit(title, content string, postType models.PostType, topic *models.Topic, user *models.User) (*models.Post, error) {
	post := models.Post{
		Id:    GetRandomString(models.ItemIDLength),
		Title: title,
		Topic: topic,
		Type:  postType,
		User:  user,
	}
	if post.User == nil {
		return nil, errors.New("No user provided")
	}
	if post.Topic == nil {
		return nil, errors.New("No topic provided")
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
		var mediaImage *models.Media
		var err error
		// find the right decoder/encoder
		var encodeFunc encodeImgFunc
		var decodeFunc decodeImgFunc
		switch fileType {
		case "image/png":
			encodeFunc = func(w io.Writer, img interface{}) error {
				return png.Encode(w, img.(image.Image))
			}
			decodeFunc = func(r io.Reader) (interface{}, error) {
				return png.Decode(r)
			}
		case "image/jpg", "image/jpeg":
			encodeFunc = func(w io.Writer, img interface{}) error {
				return jpeg.Encode(w, img.(image.Image), nil)
			}
			decodeFunc = func(r io.Reader) (interface{}, error) {
				return jpeg.Decode(r)
			}
		case "image/gif":
			encodeFunc = func(w io.Writer, img interface{}) error {
				return gif.EncodeAll(w, img.(*gif.GIF))
			}
			decodeFunc = func(r io.Reader) (interface{}, error) {
				return gif.DecodeAll(r)
			}
		}
		mediaImage, err = createMedia(&content, decodeFunc, encodeFunc)
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
	if vote.Action != models.VoteDirectionUp && vote.Action != models.VoteDirectionDown {
		return errors.New("direction musst be up or downvote")
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
		if err := post.Read("Id"); err != nil {
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

type decodeImgFunc = func(io.Reader) (interface{}, error)
type encodeImgFunc = func(io.Writer, interface{}) error

const thumbnailSize = 255

func createMedia(file *string, decode decodeImgFunc, encode encodeImgFunc) (*models.Media, error) {
	reader := strings.NewReader(*file)
	buffer := new(bytes.Buffer)
	img, err := decode(reader)
	if err != nil {
		beego.Error(err)
		return nil, errors.New("Cannot decode image")
	}
	mediaImage := models.Media{
		Type: models.MediaImage,
	}
	if err := encode(buffer, img); err != nil {
		beego.Error(err)
		return nil, errors.New("Cannot encode image")
	}
	mediaImage.File = buffer.String()
	buffer.Reset()
	var decImage image.Image
	if i, ok := img.(image.Image); ok {
		decImage = i
	} else if i, ok := img.(*gif.GIF); ok {
		decImage = i.Image[0]
	} else {
		panic("invalid image object")
	}

	// create small png thumbnail
	size := decImage.Bounds().Size()
	var width, height uint
	if size.X > size.Y {
		width = thumbnailSize
		height = uint(thumbnailSize * float64(size.Y) / float64(size.X))
	} else {
		height = thumbnailSize
		width = uint(thumbnailSize * float64(size.X) / float64(size.Y))
	}
	thumbnail := resize.Resize(width, height, decImage, resize.Lanczos3)
	if err := png.Encode(buffer, thumbnail); err != nil {
		beego.Error(err)
		return nil, errors.New("Cannot encode image")
	}
	mediaImage.Thumbnail = buffer.String()
	return &mediaImage, nil
}
