<div>
  <a href="/">Home</a>
  {{if .IsLogin}}
  <a itemprop="url" href='{{urlfor "SubmitController.Submit"}}'>Submit</a>
  <a itemprop="url" href='{{urlfor "SubmitController.CreateTopic"}}'>Create Topic</a>
  {{end}}
  <div class="right">
    {{if .IsLogin}}
    <a itemprop="url" href='{{urlfor "LoginController.Logout"}}'>Logout</a>
    {{else}}
    <a itemprop="url" href='{{urlfor "LoginController.Login"}}'>Login</a>
    {{end}}
  </div>
</div> 