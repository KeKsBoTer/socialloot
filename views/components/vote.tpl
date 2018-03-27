{{/* shows sum of votes of post and enables voting via buttons, data needs to be of type model.Post */}}
<div class="vote-container" item="{{.Id}}">
    <div class="vote-button up {{if eq .VoteDir 1}}voted{{end}}" onclick="vote(this)"></div>
    <div class="vote-count">{{.Votes}}</div>
    <div class="vote-button down {{if eq .VoteDir -1}}voted{{end}}" onclick="vote(this)"></div>
</div>