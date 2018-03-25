<div>
  <a href="/">Home</a>
  {{if .IsLogin}}
  <a itemprop="url" href='{{urlfor "SubmitController.Submit"}}'>Submit</a>
  <a itemprop="url" href='{{urlfor "SubmitController.CreateTopic"}}'>Create Topic</a>
  {{end}}
  <div class="right">
    {{if .IsLogin}}
    <span>User: {{.User.Name}}</span>
    <form action="{{urlfor "LoginController.Logout"}}" method="get" id="logout" style="display:inline">
      <input type="hidden" name="dest" value="/" />
      <a href="javascript:void(0)" onclick="document.getElementById('logout').submit()">(logout)</a>
    </form>
    {{else}}
    <a itemprop="url" href='{{urlfor "LoginController.Login"}}'>Login</a>
    <a itemprop="url" href='{{urlfor "LoginController.Signup"}}'>Signup</a>
    {{end}}
  </div>
</div> 