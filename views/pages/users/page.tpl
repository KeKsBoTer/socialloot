<div class="user-large">
    {{template "components/user.tpl" .User}}
</div>
<p>Created on: {{dateformat .User.CreationDate}}</p>
<ul class="tab-list small">
    <li class="{{if or (eq .Choice "posts") (eq .Choice "")}} active{{end}}">
        <a href="{{urlfor "UserController.Get" ":user" .User.Name ":choice" "posts"}}" class="uppercase">posts</a>
    </li>
    <li class="{{if eq .Choice "comments"}} active{{end}}">
        <a href="{{urlfor "UserController.Get" ":user" .User.Name ":choice" "comments"}}" class="uppercase">comments</a>
    </li>
</ul>
{{if or (eq .Choice "posts") (eq .Choice "")}}
<ul class="post-list">
    {{range $post := .Posts}}
        {{template "components/post-item.tpl" $post}}
    {{end}}
</ul>
{{else if eq .Choice "comments"}}
<ul class="comments-list">
    {{range $comment := .Comments}}
        <div class="comment">
            <div class="content">
                <a href="{{URL $comment.User}}" class="user">{{$comment.User.Name}}</a>
                <span class="time"> at {{dateformat $comment.Date}}</span>
                <p class="text">{{$comment.Text}}</p>
                <div class="controlls">
                    {{template "components/vote.tpl" $comment}}
                </div>
            </div>
        </div>
    {{end}}
</ul>
{{end}}