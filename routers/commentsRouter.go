package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountVerificationController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountVerificationController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountVerificationController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountVerificationController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountVerificationController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountVerificationController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountVerificationController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountVerificationController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountVerificationController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountVerificationController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AccountsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupPermissionsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupPermissionsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupPermissionsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupPermissionsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupPermissionsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupPermissionsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupPermissionsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupPermissionsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupPermissionsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthGroupPermissionsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthPermissionController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthPermissionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthPermissionController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthPermissionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthPermissionController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthPermissionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthPermissionController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthPermissionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthPermissionController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthPermissionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserGroupsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserGroupsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserGroupsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserGroupsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserGroupsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserGroupsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserGroupsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserGroupsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserGroupsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserGroupsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserUserPermissionsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserUserPermissionsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserUserPermissionsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserUserPermissionsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserUserPermissionsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserUserPermissionsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserUserPermissionsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserUserPermissionsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserUserPermissionsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:AuthUserUserPermissionsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsBulkrecipientsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsBulkrecipientsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsBulkrecipientsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsBulkrecipientsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsBulkrecipientsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsBulkrecipientsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsBulkrecipientsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsBulkrecipientsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsBulkrecipientsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsBulkrecipientsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:BulkSmsCampaignsController"],
        beego.ControllerComments{
            Method: "SendBulkSMS",
            Router: `/send-bulk-sms/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomerCategoriesController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomerCategoriesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomerCategoriesController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomerCategoriesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomerCategoriesController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomerCategoriesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomerCategoriesController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomerCategoriesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomerCategoriesController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomerCategoriesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:CustomersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoAdminLogController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoAdminLogController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoAdminLogController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoAdminLogController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoAdminLogController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoAdminLogController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoAdminLogController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoAdminLogController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoAdminLogController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoAdminLogController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoContentTypeController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoContentTypeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoContentTypeController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoContentTypeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoContentTypeController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoContentTypeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoContentTypeController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoContentTypeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoContentTypeController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoContentTypeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoMigrationsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoMigrationsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoMigrationsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoMigrationsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoMigrationsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoMigrationsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoMigrationsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoMigrationsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoMigrationsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoMigrationsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoSessionController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoSessionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoSessionController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoSessionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoSessionController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoSessionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoSessionController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoSessionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoSessionController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:DjangoSessionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:MigrationsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:MigrationsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:NewsletterCustomersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:NewsletterCustomersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:NewsletterCustomersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:NewsletterCustomersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:NewsletterCustomersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:NewsletterCustomersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:NewsletterCustomersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:NewsletterCustomersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:NewsletterCustomersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:NewsletterCustomersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardBusinessdetailsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardBusinessdetailsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardBusinessdetailsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardBusinessdetailsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardBusinessdetailsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardBusinessdetailsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardBusinessdetailsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardBusinessdetailsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardBusinessdetailsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardBusinessdetailsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardDirectoridsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardDirectoridsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardDirectoridsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardDirectoridsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardDirectoridsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardDirectoridsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardDirectoridsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardDirectoridsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardDirectoridsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:OnboardDirectoridsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:ShopsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:ShopsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:ShopsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:ShopsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:ShopsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:ShopsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserOtpsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserOtpsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserOtpsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserOtpsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserOtpsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserOtpsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserOtpsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserOtpsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserOtpsController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserOtpsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserTypesController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserTypesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserTypesController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserTypesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserTypesController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserTypesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserTypesController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserTypesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserTypesController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UserTypesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UsersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UsersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UsersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UsersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UsersController"] = append(beego.GlobalControllerRouter["mes_bulk_sms_processor/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
