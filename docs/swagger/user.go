package swagger

import v1 "github.com/rppkg/godfrey/pkg/api/v1"

// swagger:route POST /users users regisUserRequest
//
// Regist a user.
//
// Regist users according to input parameters.
//
//     Responses:
//       default: errResponse
//       200: registUserResponse

// swagger:route DELETE /users/{name} users deleteUserRequest
//
// Delete a user.
//
// Delete user according to input parameters.
//
// 	   Security:
//       tokrn:
//
//     Responses:
//       default: errResponse
//       200: okResponse

// swagger:route PUT /users/{name} users updateUserRequest
//
// Update user.
//
// Update user according to input parameters.
//
//     Security:
//       token:
//
//     Responses:
//       default: errResponse
//       200: updateUserResponse

// swagger:route GET /users/{name} users getUserRequest
//
// Get details for specified user.
//
// Get details for specified user according to input parameters.
//
//     Responses:
//       default: errResponse
//       200: getUserResponse

// swagger:parameters regisUserRequest updateUserRequest
type userRequestParamsWrapper struct {
	// User information.
	// in:body
	Body v1.User
}

// User response.
// swagger:response registUserResponse
type registUserResponseWrapper struct {
	// in:body
	Body v1.User
}

// swagger:parameters deleteUserRequest getUserRequest updateUserRequest
type userNameParamsWrapper struct {
	// User name.
	// in:path
	Name string `json:"name"`
}

// User response.
// swagger:response updateUserResponse
type updateUserResponseWrapper struct {
	// in:body
	Body v1.User
}

// User response.
// swagger:response getUserResponse
type getUserResponseWrapper struct {
	// in:body
	Body v1.User
}
