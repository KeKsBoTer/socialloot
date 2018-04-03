
{{/* Comment form, needs item id to comment on as data */}}
<form method="POST" action='{{urlfor "ApiController.Comment"}}' class="clear-on-submit" onsuccess="location.reload();">
    <textarea placeholder="comment here" name="comment" rows="7" cols="50"></textarea>
    <input type="hidden" name="item" value="{{.}}" />
    <br>
    <p class="message"></p>
    <input type="submit" value="Publish" />
</form>