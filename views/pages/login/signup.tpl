<form method="POST" action='{{urlfor "LoginController.Signup"}}'>
  {{ .xsrfdata }}
  <label for="inputName">Name</label>
  <div>
    <input placeholder="username" name="username" value="" type="text" id="inputName" />
  </div>
  <label for="inputPassword">Password</label>
  <div>
    <input placeholder="enter password" name="password" type="password" value="" title="Password" id="inputPassword" />
    <input placeholder="reenter password" name="passwordre" type="password" title="Password" />
  </div>
  {{if .Dest}}
    <input type="hidden" name="dest" value="{{.Dest}}"/>
  {{end}}
  <p class="message"></p>
  <input type="submit" value="signup">
  <input type="hidden" name="dest" value="/" />
</form>