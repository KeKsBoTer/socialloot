<div class="post">
    <span class="date">{{dateformat .Post.Date}}</span>
    <div class="header {{.Post.Type}}">
        {{template "components/vote.tpl" .Post}}
        <div>
            <a {{if eq .Post.Type "link"}} target="_blank"{{end}} 
                  href='{{if eq .Post.Type "text"}}
                            {{.URL}}
                        {{else if eq .Post.Type "image"}}
                            /media/image/original/{{.Post.Content}}
                        {{else if eq .Post.Type "link"}}
                            {{.Post.Content}}
                        {{end}}'>
                <h1 class="title">{{.Post.Title}}</h1>
            </a>
        </div>
        {{if .CanDelete}}
        <form method="POST" action="{{urlfor "APIController.Delete"}}" class="confirm" message="Delete?">
            <input type="hidden" name="item"  value="{{.Post.Id}}"/>
            <button class="delete-post" type="submit">
                <i class="ion-ios-trash"></i>
            </button>
        </form>
        {{end}}
    </div>
    {{template "components/user.tpl" .Post.User}}
    <div class="post-content">
        {{if eq .Post.Type "text"}}
        <p>{{.Post.Content}}</p>
        {{else if eq .Post.Type "image"}}
        <a href="/media/image/original/{{.Post.Content}}" target="_blank">
            <img src="/media/image/original/{{.Post.Content}}" class="image" />
        </a>
        {{end}}
    </div>
    <div class="comments-section">
        <h3>Comments</h3>
        {{if .User}}
            {{template "components/comment_form.tpl" .Post.Id}} 
        {{end}} 
        {{if isempty .Post.Comments}}
            <p style="opacity: .5;" class="no-underline">There are no comments here yet.</p>
        {{else}}
            {{range $c := .Post.Comments}} 
                {{template "components/comment.tpl" $c}}
            {{end}}
        {{end}}
    </div>
</div>