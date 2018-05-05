package controllers

type ErrorController struct {
	AuthController
}

func (c *ErrorController) NestPrepare() {
	c.Layout = "plain.tpl"
	c.TplName = "pages/error.tpl"
}

func (c *ErrorController) Error401() {
	c.Data["ErrorCode"] = 403
	c.Data["Message"] = "How are you? Please authorize yourself."
}
func (c *ErrorController) Error403() {
	c.Data["ErrorCode"] = 403
	c.Data["Message"] = "You shall not pass."
}

func (c *ErrorController) Error404() {
	c.Data["ErrorCode"] = 404
	c.Data["Message"] = "We couldn’t find this page."
}

func (c *ErrorController) Error500() {
	c.Data["ErrorCode"] = 500
	c.Data["Message"] = "Something went wrong. This should not happen."
}

func (c *ErrorController) ErrorDb() {
	c.Data["Message"] = "There is a problem with the database."
}
