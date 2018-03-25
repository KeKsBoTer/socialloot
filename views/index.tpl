<h1>Welcome to socailloot</h1>
{{template "alert.tpl" .}} 
<div>
    {{range $topic := .Topics}}
        <a href="{{URL $topic}}">{{$topic.Title}} ({{$topic.Name}})</a>
    {{end}}
</div>