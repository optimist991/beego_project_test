// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego_project_test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api/v1",

		beego.NSNamespace("/companies",
			beego.NSInclude(
				&controllers.CompaniesController{},
			),
		),

		beego.NSNamespace("/applications",
			beego.NSInclude(
				&controllers.ApplicationsController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),

		beego.NSNamespace("/aims",
			beego.NSInclude(
				&controllers.AimsController{},
			),
		),

		beego.NSNamespace("/comments",
			beego.NSInclude(
				&controllers.CommentsController{},
			),
		),

		beego.NSNamespace("/users_aims",
			beego.NSInclude(
				&controllers.UsersAimsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
