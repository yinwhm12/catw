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
			Router: `/getOne/:id`,
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
			Router: `/getNine/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetThemesIndex",
			Router: `/getThemes/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/getAll`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:LevelTypeController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:LevelTypeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"],
		beego.ControllerComments{
			Method: "PostOne",
			Router: `/toOwner`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"],
		beego.ControllerComments{
			Method: "PostMany",
			Router: `/toMany`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"],
		beego.ControllerComments{
			Method: "GetNews",
			Router: `/getNews`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:MessageController"],
		beego.ControllerComments{
			Method: "HadReadMessage",
			Router: `/hadReadMessage/:id`,
			AllowHTTPMethods: []string{"put"},
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

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:RespondOneController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:RespondOneController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:RespondOneController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:RespondOneController"],
		beego.ControllerComments{
			Method: "GetAllResponds",
			Router: `/getAll/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:RespondOneController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:RespondOneController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:RespondTwoController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:RespondTwoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
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

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:SessionController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:SessionController"],
		beego.ControllerComments{
			Method: "GetUserInfo",
			Router: `/getUserInfo/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "SayUpArticles",
			Router: `/sayUp/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "SayCollect",
			Router: `/sayCollect/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "SayUser",
			Router: `/sayUser/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUpState",
			Router: `/getUpState/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUpArticles",
			Router: `/getUpArticles`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetCollectArticles",
			Router: `/getCollectArticles`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetMyArticles",
			Router: `/getMyArticles`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "PutPWD",
			Router: `/pwd`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "PutOld",
			Router: `/putOld`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "NewPWD",
			Router: `/newPWD`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetSelf",
			Router: `/getSelf`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "PutInfo",
			Router: `/otherInfo`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAllUser",
			Router: `/getAllUser`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"] = append(beego.GlobalControllerRouter["yinwhm.com/yin/catw/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetCollectUserState",
			Router: `/getCollectUserState/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
