<div class="header">
    <div class="container">
        <div class="title-bar">
            <div>
                <a href="/" class="no-style">
                    <h1 class="title static">SocialLoot</h1>
                </a>
            </div>
        </div>
    </div>
</div>

<div class="error-page">
    <div class="info">
        <h1 class="code">{{.ErrorCode}}</h1>
        <h2 class="message">{{.Message}}</h2>
    </div>
    {{if eq .ErrorCode 404}}
        {{template "pages/search.tpl" .}}
    {{end}}

    <div class="home">
        <a href="/" class="no-underline">Go back to the main page</a>
    </div>
</div>