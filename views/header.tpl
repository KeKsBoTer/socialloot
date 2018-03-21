 {{if .IsLogin}}
  <a itemprop="url" href='{{urlfor "LoginController.Logout"}}'>Logout</a>
{{else}}
  <a itemprop="url" href='{{urlfor "LoginController.Login"}}'>Login</a>
{{end}}