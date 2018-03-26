<h1>{{.Topic.Title}}
    <small> ({{.Topic.Name}})</small>
</h1>
<h3>{{.Topic.Description}}</h3>
<div>
    <h2>Posts:</h2>
    <ul class="post-list">
        {{range $post := .Posts}}
        <li class="post" item="{{$post.Id}}">
            <div class="vote-container">
                <div class="vote-button up {{if eq $post.VoteDir 1}}voted{{end}}" onclick="vote(this)"></div>
                <div class="vote-count">{{$post.Votes}}</div>
                <div class="vote-button down {{if eq $post.VoteDir -1}}voted{{end}}" onclick="vote(this)"></div>
            </div>
            <div class="post-details">
                <a class="title" href="{{URL $post}}">{{$post.Title}}</a>
                <div class="tagline">
                    <span>submitted {{dateformat $post.Date}} by</span>
                    <a class="user" href="{{URL $post.User}}">{{$post.User.Name}}</a>
                </div>
            </div>
            <div class="clear"></div>
        </li>
        {{end}}
    </ul>
</div>