<div>
{{if .flash}}
    {{if (index .flash "warning")}}
        <span>Warning:</span>
        {{i18nja (index .flash "warning")}}
    {{else if (index .flash "error")}}
        <span>Error:</span>
        {{i18nja (index .flash "error")}}
    {{else if (index .flash "success")}}
        <span>Success:</span>
        {{i18nja (index .flash "success")}}
    {{end}}
{{end}}
</div>