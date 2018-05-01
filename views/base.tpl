<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="_xsrf" content="{{.xsrf_token}}" />
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0"/>
    <title>{{if .Title}}{{.Title}}{{else}}Socialloot{{end}}</title>
    <link href="/static/css/normalize.css" rel="stylesheet">
    <link href="/static/css/style.css" rel="stylesheet">
    <link href="/static/icons/css/ionicons.css" rel="stylesheet">
    <script type="text/javascript"  src="/static/js/jquery-3.3.1.min.js"></script>
    <script src="/static/js/utlis.js"></script>
    <script src="/static/js/script.js"></script>
 </head>
<body>
    {{.BaseHeader}}
    <div class="page">
        {{.LayoutContent}}
    </div>  
</body>
</html>
