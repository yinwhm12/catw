package client

const (
	ucenterUrl string = "http://127.0.0.1:8089/api"
	apiVersion string = "/v1"
	routerUser string = "/user"
	routerCheckLogin string = routerUser + "/login"
	routerChangePwd string = routerUser + "/changepwd"
	routerChangeSomeonePwd string = routerUser + "/changesomeonepwd"
	routerCheckAccessToken string = routerUser + "/check_access_token"
	routerResetAccessToken string = routerUser + "/reset_access_token"
	routerCheckSession string = routerUser + "/check_session"
	routerOffLine string = routerUser + "/off_line"

	token_key = "yinwhm12"
	OK = 1
	TimeOver = 111
	Fail = -1
)
