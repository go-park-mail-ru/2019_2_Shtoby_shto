package handle

type Handler interface {
	Get()
	Post()
	Put()
	Delete()
}

type HandlerService struct {
}

func (h HandlerService) Get() {

}
func (h HandlerService) Post() {

}
func (h HandlerService) Put() {

}
func (h HandlerService) Delete() {

}
