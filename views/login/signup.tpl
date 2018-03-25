<form method="POST" action='{{urlfor "LoginController.Signup"}}'>
  {{ .xsrfdata }}
   {{template "alert.tpl" .}}
  <label for="inputName">Name</label>
  <div>
    <input placeholder="username" name="Name" value="" type="text" id="inputName" />
  </div>
  <label for="inputPassword">Password</label>
  <div>
    <input placeholder="enter password" name="Password" type="password" value="" title="Password" id="inputPassword" />
    <input placeholder="reenter password" name="Repassword" type="password" title="Password" />
  </div>
  <p class="message"></p>
  <input type="submit" value="signup">
  <input type="hidden" name="dest" value="/" />
</form>