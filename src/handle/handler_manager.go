package handler

import (
	"github.com/pkg/errors"
	"net/http"
)

const (
	createHandlerErr = "Create handler error"
)

type ModelHandler struct {
	Handler                   Handler
	handlerStaticModelFactory map[string]*handlerStaticItem
}

type handlerStaticItem struct {
	t Handler
}

func CreateModelHandler(h Handler) *ModelHandler {
	return &ModelHandler{
		Handler:                   h,
		handlerStaticModelFactory: make(map[string]*handlerStaticItem, 0),
	}
}

func (p *ModelHandler) InitStaticModel(model string, t interface{}) {
	p.handlerStaticModelFactory[model] = &handlerStaticItem{t: t.(Handler)}
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
	h, err := p.getHandler(r.Context().Value("model").(string))
	if err != nil {
		//r.Error = exception.CreateError(http.StatusBadRequest, 10101, createHandlerErr, r.Request.Uri)
		return
	}
	h.Post(w, r)
}

func (p *ModelHandler) ModelRequestPut(w http.ResponseWriter, r *http.Request) {
	h, err := p.getHandler(r.Context().Value("model").(string))
	if err != nil {
		//r.Error = exception.CreateError(http.StatusBadRequest, 10101, createHandlerErr, r.Request.Uri)
		return
	}
	h.Put(w, r)
}

func (p *ModelHandler) ModelRequestGet(w http.ResponseWriter, r *http.Request) {
	h, err := p.getHandler(r.Context().Value("model").(string))
	// resolve duplicate with ModelRequestFetch, for mr. Sonar
	if err != nil {
		//r.Error = exception.CreateError(http.StatusBadRequest, 10101, createHandlerErr, uri)
		return
	}
	h.Get(w, r)
}
