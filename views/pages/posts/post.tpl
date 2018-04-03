{{template "components/vote.tpl" .Post}}

{{if eq .Post.Type "text"}}
    <a href="{{.URL}}" target="_blank">
{{else if eq .Post.Type "image"}}
    <a href="/media/image/original/{{.Post.Content}}" target="_blank">
{{else if eq .Post.Type "link"}}
    <a target="_blank" href="{{.Post.Content}}">
{{end}}
<h1 style="display:inline-block">{{.Post.Title}}</h1>
</a>

<span>by
    <a href="{{URL .Post.User}}">{{.Post.User.Name}}</a> at {{dateformat .Post.Date}}</span>
<div class="post-content">
    {{if eq .Post.Type "text"}}
    <p>{{.Post.Content}}</p>
    {{else if eq .Post.Type "image"}}
    <a href="/media/image/original/{{.Post.Content}}" target="_blank">
        <img src="/media/image/original/{{.Post.Content}}" class="image"/>
    </a>
    {{end}}
</div>
<h3>Comments (number)</h3>
<hr/> 
{{if .User}}
    {{template "components/comment_form.tpl" .Post.Id}}
{{end}}
<div class="comments-section">
    {{range $c := .Post.Comments}}
        {{template "components/comment.tpl" $c}}
    {{end}}
</div>