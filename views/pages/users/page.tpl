<div class="user-large">
    {{template "components/user.tpl" .User}}
</div>
<ul class="tab-list small">
    <li class="{{if or (eq .Choice "posts") (eq .Choice "")}} active{{end}}">
        <a href="{{urlfor "UserController.Get" ":user" .User.Name ":choice" "posts"}}" class="uppercase">posts</a>
    </li>
    <li class="{{if eq .Choice "comments"}} active{{end}}">
        <a href="{{urlfor "UserController.Get" ":user" .User.Name ":choice" "comments"}}" class="uppercase">comments</a>
    </li>
</ul>
{{if or (eq .Choice "posts") (eq .Choice "")}}
    {{if isempty .Posts}}
        <p class="no-results">The user has nothing submitted yet.</p>
    {{else}}
        <ul class="item-list">
            {{range $post := .Posts}}
                {{template "components/post_item.tpl" $post}}
            {{end}}
        </ul>
    {{end}}
{{else if eq .Choice "comments"}}
    {{if isempty .Comments}}
        <p class="no-results">The user has not commented anything yet.</p>
    {{else}}
        <ul class="item-list">
            {{range $comment := .Comments}}
                {{template "components/comment_preview.tpl" $comment}}
            {{end}}
        </ul>
    {{end}}
{{end}}