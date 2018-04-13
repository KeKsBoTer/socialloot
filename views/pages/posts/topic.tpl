<div class="page">
    <h1 class="topic-title">{{.Topic.Title}}<span>/topic/{{.Topic.Name}}</span></h1>
    <h3 class="topic-description">{{.Topic.Description}}</h3>
    <hr/>
    <div>
        <ul class="post-list">
            {{range $post := .Posts}} {{template "components/post-item.tpl" $post}} {{end}}
        </ul>
    </div>
</div>