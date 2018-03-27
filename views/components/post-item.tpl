{{/* post item in list, data needs to be of type model.Post */}}
<li class="post">
    {{template "components/vote.tpl" .}}
    <div class="post-details">
        <a class="title" href="{{URL .}}">{{.Title}}</a>
        <div class="tagline">
            <span>submitted {{dateformat .Date}} by</span>
            <a class="user" href="{{URL .User}}">{{.User.Name}}</a>
        </div>
    </div>
</li>