{{/* post item in list, data needs to be of type model.Post */}}
<li class="post-preview {{.Type}}">
    <div class="preview">
        {{if eq .Type "image"}}
        <a class="image" href="{{URL .}}" style="background-image: url('/media/image/small/{{.Content}}')"></a>
        {{else if eq .Type "link"}}
        <a class="link" href="{{.Content}}" target="_blank">
            <img class="favicon" src="{{favicon .Content}}" />
            <div class="host">
                <div class="wrapper">
                    <img class="icon" src="{{favicon .Content}}" onerror="this.src='/static/img/link_icon.png';"/>
                    <span>{{host .Content}}</span>
                </div>
            </div>
        </a>
        {{else if eq .Type "text"}}
        <a class="text" href="{{URL .}}">
            <!--<p>{{.Content}}</p>-->
        </a>
        {{end}}
        <div class="overlay">
            {{template "components/vote.tpl" .}}
        </div>
    </div>
    <div class="post-details">
        <div class="top-container">
            <a class="title" href="{{URL .}}">
                {{.Title}}
                {{if eq .Type "link"}}
                <span class="host">({{host .Content}})</span>
                {{end}}
            </a>
            <p class="date">{{dateformat .Date}}</p>
            {{if eq .Type "text"}}
            <p class="text-preview">{{cut .Content 200}}...</p>
            {{end}}
        </div>
        {{template "components/user.tpl" .User}}
    </div>
</li>