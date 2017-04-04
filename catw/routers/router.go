// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"yinwhm.com/yin/catw/controllers"

	"github.com/astaxie/beego"
	//"yinwhm.com/yin/catw/filters"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/*",&controllers.BaseController{},"options:Options"),
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
		beego.NSNamespace("/session",
			beego.NSInclude(
				&controllers.SessionController{},
			)),
	)
	beego.AddNamespace(ns)

	//beego.InsertFilter("/v1/*",beego.BeforeRouter,filters.AuthLogin,true)
}
