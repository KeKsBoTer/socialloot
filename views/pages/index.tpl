<ul class="tab-list small">
    <li class="{{if or (eq .Choice "hot") (eq .Choice "")}} active{{end}}">
        <a href="{{urlfor "IndexController.Get" ":choice" "hot"}}" class="uppercase">hot</a>
    </li>
    <li class="{{if eq .Choice "new"}} active{{end}}">
        <a href="{{urlfor "IndexController.Get" ":choice" "new"}}" class="uppercase">new</a>
    </li>
</ul>
<div>
    <ul class="post-list">
        {{range $post := .Posts}} {{template "components/post-item.tpl" $post}} {{end}}
    </ul>
</div>