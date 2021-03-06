// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"app_finder/controllers"

	"github.com/astaxie/beego"
)

func init() {
	/*ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/apple",
			beego.NSInclude(
				&controllers.AppleController{},
			),
		),
	)*/
	beego.BConfig.RunMode = "prod"
	ns := beego.NewNamespace("/apple",
		beego.NSInclude(
			&controllers.AppleController{},
		),
	)
	beego.AddNamespace(ns)

	ns = beego.NewNamespace("/google",
		beego.NSInclude(
			&controllers.GoogleController{},
		),
	)
	beego.AddNamespace(ns)

	ns = beego.NewNamespace("/notfound",
		beego.NSInclude(
			&controllers.NotfoundController{},
		),
	)
	beego.AddNamespace(ns)

}
