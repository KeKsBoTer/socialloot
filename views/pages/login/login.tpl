<h1>Login:</h1>
<form method="POST" action='{{urlfor "LoginController.Login"}}'>
  <label for="inputName">Name</label>
  <div>
    <input placeholder="username" name="username" type="text" id="inputName" />
  </div>
  <label for="inputPassword">Password</label>
  <div>
    <input placeholder="password" name="password" type="password" value="" title="Password" id="inputPassword" />
  </div>
  {{if .Dest}}
    <input type="hidden" name="dest" value="{{.Dest}}"/>
  {{end}}
  <p class="message"></p>
  <input type="submit" value="Login">

  <input type="hidden" name="dest" value="/" />
</form>
<a href='{{urlfor "LoginController.Signup"}}'>Signup</a>