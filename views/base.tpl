<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0"/>
    <title>{{if .Title}}{{.Title}}{{else}}Socialloot{{end}}</title>
    <link href="/static/css/style.css" rel="stylesheet">
    <script type="text/javascript"  src="/static/js/jquery-3.3.1.min.js"></script>
    <script src="/static/js/script.js"></script>
    {{range .HeadStyles}}
        <link rel="stylesheet" href="{{.}}">
    {{end}}
 </head>
<body id="home">

  {{.BaseHeader}}
  
  <div id="wrap">
      {{.LayoutContent}}
  </div>
  
  {{range .HeadScripts}}
      <script src="{{.}}"></script>
  {{end}}
 
</body>
</html>
