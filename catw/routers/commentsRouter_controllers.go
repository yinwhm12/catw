package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetType",
			Router: `/type`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetPlayTheme",
			Router: `/palytime/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:LevelTypeController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:LevelTypeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ObjectController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:Root1TypeController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:Root1TypeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:Root2TypeController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:Root2TypeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:SessionController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:SessionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:SessionController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:SessionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:SessionController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:SessionController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
