package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beego_project_test/controllers:AimsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:AimsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:AimsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:AimsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:AimsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:AimsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:AimsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:AimsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:AimsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:AimsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:ApplicationsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:ApplicationsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:ApplicationsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:ApplicationsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:ApplicationsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:ApplicationsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:ApplicationsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:ApplicationsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:ApplicationsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:ApplicationsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:CommentsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:CommentsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:CommentsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:CommentsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:CommentsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:CommentsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:CommentsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:CommentsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:CommentsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:CommentsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:CompaniesController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:CompaniesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:CompaniesController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:CompaniesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:CompaniesController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:CompaniesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:CompaniesController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:CompaniesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:CompaniesController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:CompaniesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:UsersAimsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:UsersAimsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:UsersAimsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:UsersAimsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:UsersAimsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:UsersAimsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:UsersAimsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:UsersAimsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:UsersAimsController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:UsersAimsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:UsersController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:UsersController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:UsersController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_project_test/controllers:UsersController"] = append(beego.GlobalControllerRouter["beego_project_test/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
