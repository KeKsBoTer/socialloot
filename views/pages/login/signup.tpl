<div class="auth-form">
  <h1>Only one step away.</h1>
  <h2>Sign up to share your intrests,comment on posts and rate other users content. </h2>

  <form method="POST" action='{{urlfor "LoginController.Signup"}}'>
    <label for="inputName">Username</label>
    <div class="reverse-order">
      <input name="username" value="" type="text" id="inputName" autofocus/>
      <label for="inputName" class="message"></label>
      <p class="description">
        Only alpha characters, numerics,"-" and "_" are allowed
        <br>
        Length must be between 3 and 15 characters
      </p> 
    </div>
    <label for="inputPassword">Password</label>
    <div class="reverse-order">
      <input name="password" type="password" value="" title="Password" id="inputPassword" />
      <label for="inputPassword" class="message"></label>
      <p class="description">
        Your password's length must be between 8 and 30 characters
      </p> 
    </div>
    <label for="reinputPassword">Reenter Password</label>
    <div>
      <label for="reinputPassword" class="message"></label>
      <input name="passwordre" type="password" title="Password" id="reinputPassword"/>
    </div>
    {{if .Dest}}
      <input type="hidden" name="dest" value="{{.Dest}}" />
    {{end}}
    <p class="global-message"></p>
    <input type="submit" value="signup">
    <input type="hidden" name="dest" value="/" />
  </form>
  
  <p class="info">Allready have an account? <a class="no-underline" href='{{urlfor "LoginController.Login"}}'>Login</a>.</p>
</div>