{{/* post item in list, data needs to be of type model.Post */}}
<li class="post">
    {{template "components/vote.tpl" .}}
    {{if eq .Type "image"}}
    <a class="title" href="{{URL .}}">
        <div class="preview">
            <img src="/media/image/small/{{.Content}}" alt=".Title"/>
        </div>
    </a>
    {{end}}
    <div class="post-details">
        <a class="title" href="{{URL .}}">{{.Title}}</a>
        <div class="tagline">
            <span>submitted {{dateformat .Date}} by</span>
            <a class="user" href="{{URL .User}}">{{.User.Name}}</a>
        </div>
    </div>
</li>