<div class="auth-form">
  <h1>Only one step away.</h1>
  <h2>Sign up to share your intrests,comment on posts and rate other useres content. </h2>
  <form method="POST" action='{{urlfor "LoginController.Signup"}}'>
    <label for="inputName">Name</label>
    <div>
      <input name="username" value="" type="text" id="inputName" autofocus/>
    </div>
    <label for="inputPassword">Password</label>
    <div>
      <input name="password" type="password" value="" title="Password" id="inputPassword" />
       </div>
    <label for="reinputPassword">Reenter Password</label>
    <div>
      <input name="passwordre" type="password" title="Password" id="reinputPassword"/>
    </div>
    {{if .Dest}}
    <input type="hidden" name="dest" value="{{.Dest}}" /> {{end}}
    <p class="message"></p>
    <input type="submit" value="signup">
    <input type="hidden" name="dest" value="/" />
  </form>
  <p class="info">Allready have an account? <a class="no-underline" href='{{urlfor "LoginController.Login"}}'>Login</a>.</p>
</div>