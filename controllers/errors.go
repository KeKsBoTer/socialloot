package controllers

// ErrorController creates error pages
type ErrorController struct {
	AuthController
}

// NestPrepare sets common template and layout for all error pages
func (c *ErrorController) NestPrepare() {
	c.Layout = "plain.tpl"
	c.TplName = "pages/error.tpl"
}

// Error401 error page for HTTP code 401
func (c *ErrorController) Error401() {
	c.Data["ErrorCode"] = 403
	c.Data["Message"] = "How are you? Please authorize yourself."
}

// Error403 error page for HTTP code 403
func (c *ErrorController) Error403() {
	c.Data["ErrorCode"] = 403
	c.Data["Message"] = "You shall not pass."
}

// Error404 error page for HTTP code 404
func (c *ErrorController) Error404() {
	c.Data["ErrorCode"] = 404
	c.Data["Message"] = "We couldn’t find this page."
}

// Error405 error page for HTTP code 405
func (c *ErrorController) Error405() {
	c.Data["ErrorCode"] = 405
	c.Data["Message"] = "Method Not Allowed"
}

// Error500 error page for HTTP code 500
func (c *ErrorController) Error500() {
	c.Data["ErrorCode"] = 500
	c.Data["Message"] = "Something went wrong. This should not happen."
}

// ErrorDB error page for database errors
func (c *ErrorController) ErrorDB() {
	c.Data["ErrorCode"] = "DB"
	c.Data["Message"] = "There is a problem with the database."
}
