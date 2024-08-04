// Package docs Godfrey API Server API.
//
// Identity and Access Management System.
//
//	 Schemes: http
//	 Host: 127.0.0.1:8080
//	 BasePath: /
//	 Version: 1.0.0
//
//	 Consumes:
//	 - application/json
//
//	 Produces:
//	 - application/json
//
//	 Security:
//	 - basic
//	 - token
//
//	SecurityDefinitions:
//	basic:
//	  type: basic
//	api_key:
//	  type: token
//	  name: Authorization
//	  in: header
//
// swagger:meta
package swagger

import "github.com/rppkg/godfrey/pkg/api/base"

// ErrResponse defines the return messages when an error occurred.
// swagger:response errResponse
type errResponseWrapper struct {
	// in:body
	Body base.ErrResponse
}

// Return nil json object.
// swagger:response okResponse
type okResponseWrapper struct {
}
