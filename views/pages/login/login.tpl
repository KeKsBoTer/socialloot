<div class="auth-form">
  <h1>Wellcome back.</h1>
  <h2>Sign in to share your intrests,comment on posts and rate other useres content. </h2>
  <form method="POST" action='{{urlfor "LoginController.Login"}}'>
    <label for="inputName">Name</label>
    <div>
      <input name="username" type="text" id="inputName" autofocus/>
    </div>
    <label for="inputPassword">Password</label>
    <div>
      <input name="password" type="password" value="" title="Password" id="inputPassword" />
    </div>
    {{if .Dest}}
    <input type="hidden" name="dest" value="{{.Dest}}" /> {{end}}
    <p class="message"></p>
    <input type="submit" value="Login">

    <input type="hidden" name="dest" value="/" />
  </form>
  <p class="info">No account? <a class="no-underline" href='{{urlfor "LoginController.Signup"}}'>Create one</a>.</p>
</div>