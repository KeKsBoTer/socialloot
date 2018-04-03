<h1>Welcome to socialloot</h1>
<div>
    {{range $topic := .Topics}}
        <a href="{{URL $topic}}">{{$topic.Title}} ({{$topic.Name}})</a>
    {{end}}
</div>