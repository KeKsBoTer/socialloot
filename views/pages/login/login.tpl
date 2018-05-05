<div class="auth-form">
  <h1>Welcome back.</h1>
  <h2>Sign in to share your intrests,comment on posts and rate other useres content. </h2>

  <form method="POST" action='{{urlfor "LoginController.Login"}}'>
    <label for="inputName">Name</label>
    <div>
      <label for="inputName" class="message"></label>
      <input name="username" type="text" id="inputName" autofocus/>
    </div>

    <label for="inputPassword">Password</label>
    <div>
      <label for="inputPassword" class="message"></label>
      <input name="password" type="password" value="" title="Password" id="inputPassword" />
    </div>

    <input type="hidden" name="dest" value="{{if .Dest}}{{.Dest}}{{else}}/{{end}}" />

    <p class="global-message"></p>
    <input type="submit" value="Login">

  </form>

  <p class="info">
    No account? <a class="no-underline" href='{{urlfor "LoginController.Signup"}}'>Create one</a>.
  </p>
</div>