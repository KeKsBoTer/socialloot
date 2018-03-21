<form method="POST" action='{{urlfor "LoginController.Signup"}}'>
  {{ .xsrfdata }}
   {{template "alert.tpl" .}}
  <label for="inputName">Name</label>
  <div>
    <input placeholder="username" name="Name" value="{{index .Params " Name "}}" type="text" id="inputName" />
  </div>
  <label for="inputPassword">Password</label>
  <div>
    <input placeholder="enter password" name="Password" type="password" value="" title="Password" id="inputPassword" />
    <input placeholder="reenter password" name="Repassword" type="password" title="Password" />
  </div>
  <input type="submit" value="signup">
</form>
<a href='{{urlfor "LoginController.Login"}}'>Login</a>