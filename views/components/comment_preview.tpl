<div class="comment">
    <div class="content">
        <a href="{{URL .User}}" class="user">{{.User.Name}}</a>
        <span class="time"> at {{dateformat .Date}}</span>
        <p class="text">{{.Text}}</p>
        <div class="controlls">
            {{template "components/vote.tpl" .}}
        </div>
    </div>
</div>