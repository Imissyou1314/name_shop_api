// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"name_shop_api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/entity",
			beego.NSInclude(
				&controllers.EntityController{},
			),
		),
		beego.NSNamespace("/book",
			beego.NSInclude(
				&controllers.BookController{},
			),
			beego.NSRouter("/words", &controllers.EntityController{}, "GET:GetWords"),
		),
		// beego.NSNamespace("/user",
		// 	beego.NSInclude(
		// 		&controllers.UserController{},
		// 	),
		// ),
	)
	beego.AddNamespace(ns)
}
