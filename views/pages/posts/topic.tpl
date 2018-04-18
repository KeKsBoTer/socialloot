<div class="page">
    <h1 class="topic-title">{{.Topic.Title}}<span>/topic/{{.Topic.Name}}</span></h1>
    <h3 class="topic-description">{{.Topic.Description}}</h3>
    <!--<hr/>-->
    <ul class="tab-list small">
        <li class="{{if or (eq (GetParam .URL "choice") "hot") (eq (GetParam .URL "choice") "")}} active{{end}}">
          <a href="/topic/{{.Topic.Name}}/hot" class="uppercase">hot</a>
        </li>
        <li class="{{if eq (GetParam .URL "choice") "new"}} active{{end}}">
          <a href="/topic/{{.Topic.Name}}/new" class="uppercase">new</a>
        </li>
      </ul>
    <div>
        <ul class="post-list">
            {{range $post := .Posts}} {{template "components/post-item.tpl" $post}} {{end}}
        </ul>
    </div>
</div>