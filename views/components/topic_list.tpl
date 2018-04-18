<div class="topic-list">
    <button class="icon-button ion-ios-arrow-back" onclick="scrollList(this,'right')"></button>
    <div class="list-container">
        <ul class="list">
            {{range $topic := .Topics}}
            <li>
                <a href="{{URL $topic}}" class="no-style">{{$topic.Name}}</a>
            </li>
            {{end}}
        </ul>
    </div>
    <button class="icon-button ion-ios-arrow-forward" onclick="scrollList(this,'left')"></button>
</div>