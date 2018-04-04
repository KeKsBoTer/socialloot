<div class="topic-list">
    <button class="icon-button ion-ios-arrow-back"></button>
    <ul class="list">
        {{range $topic := .Topics}}
        <li>
            <a href="{{URL $topic}}" class="no-style">{{$topic.Name}}</a>
        </li>
        {{end}}
    </ul>
    <button class="icon-button ion-ios-arrow-forward"></button>
</div>