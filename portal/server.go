package portal



type Server interface {
	Start() error
	Stop() error
}
