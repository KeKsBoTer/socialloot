<h1>Login:</h1>
<form method="POST" action='{{urlfor "LoginController.Login"}}'>
  {{ .xsrfdata }}
  {{template "alert.tpl" .}}
  <label for="inputName">Name</label>
  <div>
    <input placeholder="username" name="Name" value="{{index .Params " Name "}}" type="text" id="inputName" />
  </div>
  <label for="inputPassword">Password</label>
  <div>
    <input placeholder="password" name="Password" type="password" value="" title="Password" id="inputPassword" />
  </div>
  <input type="submit" value="Login">
</form>
<a href='{{urlfor "LoginController.Signup"}}'>Signup</a>