<div class="comment">
    <div class="content">
        <a class="no-underline toggle" onclick="toggleComment(this)"></a>
        <a href="{{URL .User}}" class="user">{{.User.Name}}</a>
        <span class="time"> at {{dateformat .Date}}</span>
        <p class="text">{{.Text}}</p>
        <div class="controlls">
            {{template "components/vote.tpl" .}}
            <a onclick="showCommentForm(this)" class="reply">reply</a>
        </div>
        {{template "components/comment_form.tpl" .Id}}
    </div>
    <div class="replies">
        {{range $c := .Replies}} {{template "components/comment.tpl" $c}} {{end}}
    </div>
</div>