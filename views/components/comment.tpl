<div class="comment">
    <div class="wrapper">
        {{template "components/vote.tpl" .}}
        <div class="content">
            <a class="no-underline toggle" onclick="toggleComment(this)"></a>
            <a href="{{URL .User}}">{{.User.Name}}</a>
            <span> at {{dateformat .Date}}</span>
            <p class="text">{{.Text}}</p>
            <div class="controlls">
                <a onclick="showCommentForm(this)" class="uppercase">reply</a>
                <div class="comment-box" style="display:none;">
                    {{template "components/comment_form.tpl" .Id}}
                </div>
            </div>
        </div>
    </div>
    <div class="replies">
        {{range $c := .Replies}} 
            {{template "components/comment.tpl" $c}} 
        {{end}}
    </div>
</div>
