// TODO:: work in progress
package handle

import (
	"github.com/pkg/errors"
	"net/http"
)

const (
	createHandlerErr = "Create handler error"
)

type ModelHandler struct {
	handlerStaticModelFactory map[string]*handlerStaticItem
}

var ModelHandlers *ModelHandler

func init() {
	ModelHandlers = CreateModelHandler()
}

type handlerStaticItem struct {
	t Handler
	u string
}

func CreateModelHandler() *ModelHandler {
	return &ModelHandler{
		handlerStaticModelFactory: make(map[string]*handlerStaticItem),
	}
}

func (p *ModelHandler) InitStaticModel(model string, t interface{}, u string) {
	p.handlerStaticModelFactory[model] = &handlerStaticItem{t: t.(Handler), u: u}
}

func (p *ModelHandler) getHandler(model string) (Handler, error) {
	t := p.handlerStaticModelFactory[model]
	if t != nil {
		return t.t, nil
	}
	return nil, errors.New(createHandlerErr)
}

func (p *ModelHandler) GetHandler(model string) (Handler, error) {
	return p.getHandler(model)
}

func (p *ModelHandler) ModelRequestPost(w http.ResponseWriter, r *http.Request) {

}

func (p *ModelHandler) ModelRequestPut(w http.ResponseWriter, r *http.Request) {

}

func (p *ModelHandler) ModelRequestGet(w http.ResponseWriter, r *http.Request) {

}
