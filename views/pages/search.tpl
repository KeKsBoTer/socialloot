<form class="search-form" method="GET" action="{{urlfor "SearchController.Get"}}">
    <input  type="text" 
            name="query" 
            placeholder="What are you looking for?" 
            value="{{.SearchQuery}}" 
            title="search input" 
            autofocus/>
            <input type="hidden" value="{{.Choice}}" name="choice"/>
</form>
{{if not (eq (GetParam .URL "query") "")}}
    <ul class="tab-list small">
        <li class="{{if or (eq .Choice "posts") (eq .Choice "")}} active{{end}}">
            <a href="{{urlfor "SearchController.Get" "query" .SearchQuery "choice" "posts"}}" class="uppercase">posts</a>
        </li>
        <li class="{{if eq .Choice "topics"}} active{{end}}">
            <a href="{{urlfor "SearchController.Get" "query" .SearchQuery "choice" "topics"}}" class="uppercase">topics</a>
        </li>
        <li class="{{if eq .Choice "users"}} active{{end}}">
            <a href="{{urlfor "SearchController.Get" "query" .SearchQuery "choice" "users"}}" class="uppercase">users</a>
        </li>
    </ul>
    {{if isempty .SearchResult}}
        <p class="no-results">No search results.</p>
    {{else}}
    <ul class="item-list">
        {{if or (eq .Choice "posts") (eq .Choice "")}}
            {{range $post := .SearchResult}} 
                {{template "components/post_item.tpl" $post}}
            {{end}}
        {{else if eq .Choice "topics"}}
            {{range $topic := .SearchResult}} 
                <a href="{{URL $topic}}">
                    <p>{{$topic.Name}}</p>
                </a>
            {{end}}
        {{else if eq .Choice "users"}}
            {{range $user := .SearchResult}}
                <div class="user-large">
                    {{template "components/user.tpl" $user}}
                </div>
            {{end}}
        {{end}}
    </ul>
    {{end}}
{{end}}