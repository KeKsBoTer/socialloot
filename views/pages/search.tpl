<h1>Search</h1>
<div class="search-from">
    <form method="GET" action="{{urlfor "SearchController.Get"}}">
        <input name="query" placeholder="What are you looking for?" value="{{.SearchQuery}}" title="search input"/>
    </form>

    {{if .Posts}}
    <ul class="post-list">
        {{range $post := .Posts}} {{template "components/post-item.tpl" $post}} {{end}}
    </ul>
    {{end}}
</div>