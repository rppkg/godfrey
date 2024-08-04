package errors

type Coder interface {
	Code() string
	String() string
	HTTPStatus() int
}
