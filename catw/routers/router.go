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
	"yinwhm.com/yin/catw/filters"
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
		beego.NSNamespace("/root_1_type",
			beego.NSInclude(
				&controllers.Root1TypeController{},
			)),
		beego.NSNamespace("/root_2_type",
			beego.NSInclude(&controllers.Root2TypeController{},
			)),
		beego.NSNamespace("/level_type",
			beego.NSInclude(&controllers.LevelTypeController{},
			)),
		beego.NSNamespace("/article",
			beego.NSInclude(&controllers.ArticleController{},
			)),
		beego.NSNamespace("/comment_area",
			beego.NSInclude(&controllers.RespondOneController{},
			),
		),
		beego.NSNamespace("/comment_two",
			beego.NSInclude(&controllers.RespondTwoController{},
			),
		),



	)
	beego.AddNamespace(ns)

	//beego.InsertFilter("/v1/*",beego.BeforeRouter,filters.AuthLogin,true)
	beego.InsertFilter("/v1/article/",beego.BeforeRouter,filters.BeforeWrite,true)//写文章认证 是否登录
	beego.InsertFilter("/v1/user/*",beego.BeforeRouter,filters.BeforeWrite,true)//点赞 收藏认证 是否登录
	beego.InsertFilter("/v1/comment_area/",beego.BeforeRouter,filters.BeforeWrite,true)//写1评论认证 是否登录
	beego.InsertFilter("/v1/comment_two/",beego.BeforeRouter,filters.BeforeWrite,true)//写2评论认证 是否登录
}
