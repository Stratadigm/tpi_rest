package tpi

import (
	"net/http"
)

type AuthHandlerFunc func(http.ResponseWriter, *http.Request, http.HandlerFunc)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type ARoute struct {
	Name    string
	Method  string
	Pattern string
	Auth    AuthHandlerFunc
}

type Routes []Route
type ARoutes []ARoute

var routes = Routes{
	//JSON API
	Route{
		"Login",
		"POST",
		"/auth_token",
		Login,
	},
	Route{
		"Refresh",
		"PUT",
		"/auth_token",
		Refresh,
	},
	Route{
		"CreateUser",
		"POST",
		"/user",
		Create,
	},
	Route{
		"CreateVenue",
		"POST",
		"/venue",
		Create,
	},
	Route{
		"CreateThali",
		"POST",
		"/thali",
		Create,
	},
	Route{
		"RetrieveUsers",
		"GET",
		"/users",
		Retrieve,
	},
	Route{
		"RetrieveVenues",
		"GET",
		"/venues",
		Retrieve,
	},
	Route{
		"RetrieveThalis",
		"GET",
		"/thalis",
		Retrieve,
	},
	Route{
		"RetrieveUser",
		"GET",
		"/user/{id}",
		Retrieve,
	},
	Route{
		"RetrieveVenue",
		"GET",
		"/venue/{id}",
		Retrieve,
	},
	Route{
		"RetrieveThali",
		"GET",
		"/thali/{id}",
		Retrieve,
	},
	Route{
		"RetrieveImage",
		"GET",
		"/thali/{id}/image",
		Retrieve,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/user/{id}",
		Update,
	},
	Route{
		"UpdateVenue",
		"PUT",
		"/venue/{id}",
		Update,
	},
	Route{
		"UpdateThali",
		"PUT",
		"/thali/{id}",
		Update,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/user/{id}",
		Delete,
	},
	Route{
		"DeleteVenue",
		"DELETE",
		"/venue/{id}",
		Delete,
	},
	Route{
		"DeleteThali",
		"DELETE",
		"/thali/{id}",
		Delete,
	},
	//HTML URLS
	Route{
		"Logs",
		"GET",
		"/logs",
		Logs,
	},
	Route{
		"Counters",
		"GET",
		"/counters",
		Counters,
	},
	Route{
		"List",
		"GET",
		"/list/{what}",
		List,
	},
	Route{
		"PostForm",
		"POST",
		"/postform/{what}",
		PostForm,
	},
	Route{
		"GetForm",
		"GET",
		"/getform/{what}",
		GetForm,
	},
	Route{
		"GetFormId",
		"GET",
		"/getform/{what}/{id:[0-9]+}",
		GetForm,
	},
	Route{
		"GetUpload",
		"GET",
		"/upload/{what}",
		GetUpload,
	},
	Route{
		"PostUpload",
		"POST",
		"/upload/{what}",
		PostUpload,
	},
	Route{
		"GetImage",
		"GET",
		"/image/{what}",
		GetImage,
	},
}

var aroutes = ARoutes{
	ARoute{
		"Logout",
		"DELETE",
		"/auth_token",
		Logout,
	},
	ARoute{
		"Hello",
		"GET",
		"/hello",
		Hello,
	},
}
