 {{/* Comment form, needs the id of the item to comment on as data */}}
<div class="comment-form">
    <form method="POST" action='{{urlfor "APIController.Comment"}}' class="clear-on-submit" onsuccess="location.reload();">
        <textarea placeholder="Write a comment..." name="comment" rows="1"></textarea>
        <input type="hidden" name="item" value="{{.}}" />
        <div class="send">
            <p class="message"></p>
            <input type="submit" value="Publish" />
        </div>
    </form>
</div>