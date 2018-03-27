<h1>{{.Topic.Title}}
    <small> ({{.Topic.Name}})</small>
</h1>
<h3>{{.Topic.Description}}</h3>
<div>
    <h2>Posts:</h2>
    <ul class="post-list">
        {{range $post := .Posts}}
            {{template "components/post-item.tpl" $post}}
        {{end}}
    </ul>
</div>