package transport

import (
	"github.com/pkg/errors"
	"net/http"
	"reflect"
)

const (
	createHandlerErr = "Create handler error"
)

type ModelHandler struct {
	handlerModelFactory       map[string]*handlerItem
	handlerStaticModelFactory map[string]*handlerStaticItem
}

var ModelHandlers *ModelHandler

func init() {
	ModelHandlers = CreateModelHandler()
}

type handlerItem struct {
	t reflect.Type
	u string
}

type handlerStaticItem struct {
	t Handler
	u string
}

func CreateModelHandler() *ModelHandler {
	return &ModelHandler{
		handlerModelFactory:       make(map[string]*handlerItem),
		handlerStaticModelFactory: make(map[string]*handlerStaticItem),
	}
}

func (p *ModelHandler) InitModel(model string, t reflect.Type, u string) {
	p.handlerModelFactory[model] = &handlerItem{t: t, u: u}
}

func (p *ModelHandler) InitStaticModel(model string, t interface{}, u string) {
	p.handlerStaticModelFactory[model] = &handlerStaticItem{t: t.(Handler), u: u}
}

func (p *ModelHandler) getHandler(model string) (Handler, error) {
	t1 := p.handlerModelFactory[model]
	if t1 != nil {
		v := reflect.New(t1.t).Elem()
		d := v.Interface().(Handler)
		return d, nil
	}
	t2 := p.handlerStaticModelFactory[model]
	if t2 != nil {
		return t2.t, nil
	}
	return nil, errors.New(createHandlerErr)
}

func (p *ModelHandler) GetHandler(model string) (Handler, error) {
	return p.getHandler(model)
}

func (p *ModelHandler) ModelRequestPost(w http.ResponseWriter, r *http.Request) {
	h, err := p.getHandler(r.Context().Value("model").(string))
	if err != nil {
		return
	}
	h.Post(w, r)
}

func (p *ModelHandler) ModelRequestPut(w http.ResponseWriter, r *http.Request) {
	h, err := p.getHandler(r.Context().Value("model").(string))
	if err != nil {
		return
	}
	h.Put(w, r)
}

func (p *ModelHandler) ModelRequestGet(w http.ResponseWriter, r *http.Request) {
	h, err := p.getHandler(r.Context().Value("model").(string))
	if err != nil {
		return
	}
	h.Get(w, r)
}
