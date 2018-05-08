package lib

import (
	"crypto/rand"
	"reflect"

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego"
)

// URLForItem generates the relative url for the given model
// Valid models are Post,PostMetaData,Topic and User
// If the given data has non of the types above, "/" is returned
func URLForItem(data interface{}) string {
	rv := reflect.ValueOf(data)
	// if data is a pointer, get the value to make type switching possible
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	data = rv.Interface()
	switch data.(type) {
	case models.PostMetaData:
		post := data.(models.PostMetaData).Post
		return urlForPost(post)
	case models.Post:
		post := data.(models.Post)
		return urlForPost(&post)
	case models.Topic:
		topic := data.(models.Topic)
		return beego.URLFor("TopicController.Get", ":topic", topic.Name, ":choice", "")
	case models.User:
		user := data.(models.User)
		return beego.URLFor("UserController.Get", ":user", user.Name, ":choice", "")
	default:
		beego.Error("Cannot create URL for:", reflect.TypeOf(data))
		return "/"
	}
}

// generates the url for a post model
// if a error occures "/" is returned
func urlForPost(post *models.Post) string {
	if post.Topic == nil {
		beego.Error("Cannot create URL for post (empty topic):", *post)
		return "/"
	}
	return beego.URLFor("PostController.Get", ":topic", post.Topic.Name, ":post", post.Id)
}

// GetRandomString generates base62 string of given length
// Available characters are: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
func GetRandomString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// GenerateItemID generates a new unused item id with length of models.ItemIDLength
func GenerateItemID() string {
	id := GetRandomString(models.ItemIDLength)
	// find free id
	for models.Comments().Filter("id", id).Exist() || models.Posts().Filter("id", id).Exist() {
		id = GetRandomString(models.ItemIDLength)
	}
	return id
}
