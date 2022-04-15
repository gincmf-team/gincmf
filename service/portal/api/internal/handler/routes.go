// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	navItem "gincmf/service/portal/api/internal/handler/navItem"
	portaladminarticle "gincmf/service/portal/api/internal/handler/portal/admin/article"
	portaladmincategory "gincmf/service/portal/api/internal/handler/portal/admin/category"
	portaladmintag "gincmf/service/portal/api/internal/handler/portal/admin/tag"
	portaladmintheme "gincmf/service/portal/api/internal/handler/portal/admin/theme"
	portaladminthemeFile "gincmf/service/portal/api/internal/handler/portal/admin/themeFile"
	portalappbreadcrumb "gincmf/service/portal/api/internal/handler/portal/app/breadcrumb"
	portalappcategory "gincmf/service/portal/api/internal/handler/portal/app/category"
	portalapplist "gincmf/service/portal/api/internal/handler/portal/app/list"
	portalappthemeFile "gincmf/service/portal/api/internal/handler/portal/app/themeFile"
	route "gincmf/service/portal/api/internal/handler/route"
	"gincmf/service/portal/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: portaladmincategory.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: portaladmincategory.ListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: portaladmincategory.ShowHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: portaladmincategory.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/:id",
					Handler: portaladmincategory.EditHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/options",
					Handler: portaladmincategory.OptionsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/category"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: portalapplist.ShowHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/app/list"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: portalappbreadcrumb.BreadcrumbHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/app/breadcrumb"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: portalappcategory.TreeListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/app/category/trees"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: portaladminarticle.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: portaladminarticle.ShowHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: portaladminarticle.StoreHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/:id",
					Handler: portaladminarticle.EditHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/article"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: portalapplist.GetHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/app/list"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: portaladmintheme.InitHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/theme"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: portaladminthemeFile.ListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/:id",
					Handler: portaladminthemeFile.SaveHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/theme_file"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/theme_files",
				Handler: portalappthemeFile.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/theme_file",
				Handler: portalappthemeFile.DetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/app"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: portaladmintag.GetHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/:id",
					Handler: portaladmintag.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/admin/tag"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: route.ListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/app/route"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: navItem.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: navItem.StoreHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/:id",
				Handler: navItem.EditHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: navItem.DelHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/urls",
				Handler: navItem.OptionsUrlsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/options",
				Handler: navItem.OptionsListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/admin/nav_items"),
	)
}
