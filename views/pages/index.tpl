{{if isempty .Posts}}
    {{template "components/nothing.tpl" .}}
{{else}}
    <ul class="tab-list small">
        <li class="{{if or (eq .Choice "hot") (eq .Choice "")}} active{{end}}">
            <a href="{{urlfor "IndexController.Get" ":choice" "hot"}}" class="uppercase">hot</a>
        </li>
        <li class="{{if eq .Choice "new"}} active{{end}}">
            <a href="{{urlfor "IndexController.Get" ":choice" "new"}}" class="uppercase">new</a>
        </li>
    </ul>
    <ul class="item-list">
        {{range $post := .Posts}} 
            {{template "components/post_item.tpl" $post}}
        {{end}}
    </ul>
{{end}}