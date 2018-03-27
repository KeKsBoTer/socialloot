<h1>Welcome to socailloot</h1>
<div>
    {{range $topic := .Topics}}
        <a href="{{URL $topic}}">{{$topic.Title}} ({{$topic.Name}})</a>
    {{end}}
</div>