package ctx

type IContext interface {
	JSON(code int, obj any)
	BodyParser(obj any) error
	ReadBody() ([]byte, error)
	Param(key string) string
	Query(key string) string
}
