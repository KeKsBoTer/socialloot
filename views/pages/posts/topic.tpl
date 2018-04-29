<h1 class="topic-title">{{.Topic.Title}}<span>{{URL .Topic}}</span></h1>
<h3 class="topic-description">{{.Topic.Description}}</h3>
{{if isempty .Posts}}
    {{template "components/nothing.tpl" .}}
{{else}}
    <ul class="tab-list small">
        <li class="{{if or (eq .Choice "hot") (eq .Choice "")}} active{{end}}">
            <a href="{{urlfor "TopicController.Get" ":topic" .Topic.Name ":choice" "hot"}}" class="uppercase">hot</a>
        </li>
        <li class="{{if eq .Choice "new"}} active{{end}}">
            <a href="{{urlfor "TopicController.Get" ":topic" .Topic.Name ":choice" "new"}}" class="uppercase">new</a>
        </li>
    </ul>
    <ul class="item-list">
        {{range $post := .Posts}} 
            {{template "components/post_item.tpl" $post}} 
        {{end}}
    </ul>
{{end}}