{{/* Header of every web page, needs root data (pass "." as data) */}}
<div class="header">
  <div class="container">
    <div class="title-bar">
      <div>
        <a href="/" class="no-style">
          <h1 class="title">SocialLoot</h1>
        </a>
      </div>
      <div class="lower-container">
          <div>
              {{if .IsLogin}}
              <div class="button-group">
                  {{if .Topic}}
                    <a href="{{urlfor "SubmitController.Submit" "topic" .Topic.Name}}">
                      <button class="text-button"><span>Submit to</span> <span class="limit-chars">{{.Topic.Name}}</span></button>
                    </a>
                  {{else}}
                    <a href="{{urlfor "SubmitController.Submit"}}">
                      <button class="text-button">Submit</button>
                    </a>
                  {{end}}
                  <a href="{{urlfor "SubmitController.CreateTopic"}}">
                    <button class="text-button">Create Topic</button>
                  </a>
                </div>
              {{end}}
          </div>
          <div>
            <div class="button-group">
              <a href="{{urlfor "SearchController.Get"}}">
                <button class="icon-button ion-ios-search"></button>
              </a>
              {{if .IsLogin}}
              <button class="icon-button large ion-ios-person-outline" id="user-popup-button">
              </button>
              {{else}}
              <a href="{{urlfor "LoginController.Login"}}">
                <button class="text-button">Login</button>
              </a>
              <a href="{{urlfor "LoginController.Signup"}}">
                <button class="text-button">Register</button>
              </a>
              {{end}}
            </div>
          </div>
      </div>
    </div>
    {{if .Topics}} {{template "components/topic_list.tpl" .}} {{end}}
  </div>
  {{if .IsLogin}}
  <div class="popup-list" id="user-popup">
    <ul class="inner">
      <li>
        <a href="{{URL .User}}">
          Profile
        </a>
      </li>
      <li>
        <form action="{{urlfor "LoginController.Logout"}}" method="get" id="logout" style="display:inline">
          <input type="hidden" name="dest" value="/" />
          <a href="javascript:void(0)" onclick="document.getElementById('logout').submit()">Logout</a>
        </form>
      </li>
    </ul>
  </div>
  {{end}}
</div>