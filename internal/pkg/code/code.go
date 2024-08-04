package code

var (
	OK = &ErrCode{HTTPCode: 200, ICode: "", Message: ""}

	InternalServerError = &ErrCode{HTTPCode: 500, ICode: "InternalError", Message: "Internal server error."}
)
