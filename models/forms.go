package models

// This file contains models for all forms
// Some models need specific validation, which is done in validators.go
// They represent the data entered by the user and the field tags are used
// to do baic syntax validation.
// The "form" tag is the key name for the field in the form.

// LoginForm is used to authenticate a user
// Both UserName and Password are required.
type LoginForm struct {
	// The unique name for the user
	UserName string `form:"username" valid:"Required;"`

	// The users password as plain text
	Password string `form:"password" valid:"Required;"`
}

// SignUpForm holds the data needed for the registration process.
// The validators check if the entered data is conform with the policy.
type SignUpForm struct {

	// Every user has a unique name with a length between 3 and 15 characters
	// Also the name can only contain letters, numbers, underscores and "-"
	UserName string `form:"username" valid:"Required;MinSize(3);MaxSize(15);Match(^[a-zA-Z0-9_\\-]+$)"`

	// A Password must consist of 8 to 30 characters
	Password string `form:"password" valid:"Required;MinSize(8);MaxSize(30)"`

	// The user needs to enter the password twice to avoid typing mistakes.
	PasswordRe string `form:"passwordre" valid:"Required;"`
}

// SubmitForm is used for creating posts
// A post can be a text, link or an image
type SubmitForm struct {

	// A short title for the post consisting of 3 to 300 characters
	Title string `form:"title" valid:"Required;MinSize(3);MaxSize(300)"`

	// The content of the post.
	// If the post type is text the length is limited to 40000 characters
	Content string `form:"content"`

	// Type of submitwted data
	Type PostType `form:"type" valid:"Required"`

	// Name of the topic the post should be submitted tp
	TopicName string `form:"topic" valid:"Required"`
}

// VoteForm is sent by the browser to register a user vote on an item
// The user votes by clicking on buttons
type VoteForm struct {

	// Up- or downvote
	Direction VoteDirection `form:"dir" valid:"Required"`

	// The item on which is voted
	// Needs to be a valid item id with 11 characters
	Item string `form:"id" valid:"Required;Length(11)"`
}

// CreateTopicForm is the data entered by the user to create a new topic
type CreateTopicForm struct {

	// The short unique identifier for the topic
	// The name can only contain letters, numbers and underscores but can not start with an underscores
	// The length must be between 3 and 21 characters
	Name string `form:"name" valid:"Required;MinSize(3);MaxSize(21);Match(^[[:alnum:]][[:alnum:]_]+$)"`

	// The title for the topic
	Title string `form:"title" valid:"Required;MinSize(3);MaxSize(200)"`

	// A short discription about the topic with up to 2000 characters
	Description string `form:"description" valid:"Required;MinSize(3);MaxSize(2000)"`
}

// CommentForm is used to comment on items
// The user only enters a text and sends this via a button
type CommentForm struct {

	// The item's id to which the comments answers
	// This can be a post or another comment
	Item string `form:"item" valid:"Required;Length(11)"`

	// The non empty comment with up to 40000 characters
	Comment string `form:"comment" valid:"Required;MaxSize(40000)"`
}

// DeleteForm is used to delete a post via the web app
// The user triggers this form by clicking on a delete button.
type DeleteForm struct {
	Item string `form:"item" valid:"Required;Length(11)"`
}
