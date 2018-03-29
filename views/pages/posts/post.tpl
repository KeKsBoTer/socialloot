{{template "components/vote.tpl" .Post}}
<h1 style="display:inline-block">{{.Post.Title}}</h1>
<span>by <a href="{{URL .Post.User}}">{{.Post.User.Name}}</a> at {{dateformat .Post.Date}}</span>
<div>
    <p>{{.Post.Content}}</p>
</div>
<h3>Comments (number)</h3>
<hr/>
<form method="POST" action='{{urlfor "ApiController.Comment"}}' class="clear-on-submit">
    <textarea placeholder="comment here" name="comment"></textarea>
    <input type="hidden" name="post" value="{{.Post.Id}}"/>
    <br>
    <input type="submit" value="Publish"/>
</form>
<div class="comments-section">
    {{range $c := .Post.Comments}}
    <div class="comment">
        <a href="{{URL $c.User}}">{{$c.User.Name}}</a><span> at {{dateformat $c.Date}}</span>
        <p class="text">{{$c.Text}}</p>
    </div>
    {{end}}
</div>