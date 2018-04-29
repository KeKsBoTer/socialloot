package models

type LoginForm struct {
	UserName string `form:"username" valid:"Required;"`
	Password string `form:"password" valid:"Required;"`
}

type SignUpForm struct {
	UserName   string `form:"username" valid:"Required;MinSize(3);MaxSize(15);Match(^[a-zA-Z0-9_\\-]+$)"`
	Password   string `form:"password" valid:"Required;MinSize(8);MaxSize(30)"`
	PasswordRe string `form:"passwordre" valid:"Required;MinSize(8);MaxSize(30)"`
}

type SubmitForm struct {
	Title     string   `form:"title" valid:"Required"`
	Content   string   `form:"content"`
	Type      PostType `form:"type" valid:"Required"`
	TopicName string   `form:"topic" valid:"Required"`
}

type VoteForm struct {
	Direction VoteDirection `form:"dir" valid:"Required"`
	Item      string        `form:"id" valid:"Required;Length(11)"`
}

type CreateTopicForm struct {
	Name        string `form:"name" valid:"Required"`
	Title       string `form:"title" valid:"Required"`
	Description string `form:"description" valid:"Required"`
}

type CommentForm struct {
	Item    string `form:"item" valid:"Required;Length(11)"`
	Comment string `form:"comment" valid:"Required"`
}

type DeleteForm struct {
	Item string `form:"item" valid:"Required;Length(11)"`
}
