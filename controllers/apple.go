package controllers

import (
	"app_finder/models"
	"github.com/astaxie/beego"
)

type AppleController struct {
	beego.Controller
}

// @router /:id(.+) [get]
func (this *AppleController) Lookup() {
	bundleId := this.Ctx.Input.Param(":id")

	this.Data["json"] = models.AppleLookup(bundleId)
	this.ServeJSON()
}
