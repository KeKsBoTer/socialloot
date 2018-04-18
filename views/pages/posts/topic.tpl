<div class="page">
    <h1 class="topic-title">{{.Topic.Title}}<span>{{URL .Topic}}</span></h1>
    <h3 class="topic-description">{{.Topic.Description}}</h3>
    <!--<hr/>-->
    <ul class="tab-list small">
        <li class="{{if or (eq .Choice "hot") (eq .Choice "")}} active{{end}}">
          <a href="{{urlfor "TopicController.Get" ":topic" .Topic.Name ":choice" "hot"}}" class="uppercase">hot</a>
        </li>
        <li class="{{if eq .Choice "new"}} active{{end}}">
          <a href="{{urlfor "TopicController.Get" ":topic" .Topic.Name ":choice" "new"}}" class="uppercase">new</a>
        </li>
      </ul>
    <div>
        <ul class="post-list">
            {{range $post := .Posts}} {{template "components/post-item.tpl" $post}} {{end}}
        </ul>
    </div>
</div>