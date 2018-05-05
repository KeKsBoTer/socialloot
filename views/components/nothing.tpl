{{/* This is shown if a post list is empty */}}
<div class="nothing-here">
    <p> There is nothing here yet. 
        <br>
        Be the first person to
        {{if .Topic}}
            <a class="no-underline" href="{{urlfor "SubmitController.Submit" "topic" .Topic.Name}}">submit</a> to <b>{{.Topic.Name}}</b>
        {{else}}
            <a class="no-underline" href="{{urlfor "SubmitController.Submit"}}">submit</a> something
        {{end}}
    </p>
</div>