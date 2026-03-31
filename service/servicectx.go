package servicectx

import (
	"errors"
	"io"
	"net/http"

	"github.com/eaglebush/appctx"
	ck "github.com/eaglebush/cachekit"
	l "github.com/stdutil/log"
)

var (
	ErrLogFuncUnset               = errors.New("log function unset")
	ErrRespondFunctionUnset       = errors.New("respond function unset")
	ErrRespondBytesFunctionUnset  = errors.New("respond bytes function unset")
	ErrRespondDirectFunctionUnset = errors.New("respond direct function unset")
)

type ServiceContext struct {
	appctx.Meta
	Cache             ck.Cache
	logFunc           func(msgType l.LogType, message ...string)
	respondFunc       func(data any, w http.ResponseWriter, r *http.Request) error                    // Respond function
	respondBytesFunc  func(data []byte, fileExt string, w http.ResponseWriter, r *http.Request) error // RespondBytes function
	respondDirectFunc func(src io.ReadCloser, w http.ResponseWriter, gzipped bool, mime string) error // RespondDirect function
}

func NewServiceContext(
	mt *appctx.Meta,
	so ...ServiceOption,
) *ServiceContext {
	sc := ServiceContext{
		Meta: *mt,
	}
	for _, o := range so {
		o(&sc)
	}
	return &sc
}

// Log messages into the server
func (cs *ServiceContext) Log(Type l.LogType, Message ...string) error {
	if cs.logFunc == nil {
		return ErrLogFuncUnset
	}
	(cs.logFunc)(Type, Message...)
	return nil
}

// Respond to API query
func (cs *ServiceContext) Respond(data any, w http.ResponseWriter, r *http.Request) error {
	if cs.respondFunc == nil {
		return ErrRespondFunctionUnset
	}
	return (cs.respondFunc)(data, w, r)
}

// RespondBytes to API query
func (cs *ServiceContext) RespondBytes(data []byte, fileExt string, w http.ResponseWriter, r *http.Request) error {
	if cs.respondBytesFunc == nil {
		return ErrRespondBytesFunctionUnset
	}
	return (cs.respondBytesFunc)(data, fileExt, w, r)
}

// RespondDirect to API query
func (cs *ServiceContext) RespondDirect(src io.ReadCloser, w http.ResponseWriter, gzipped bool, mime string) error {
	if cs.respondDirectFunc == nil {
		return ErrRespondDirectFunctionUnset
	}
	return (cs.respondDirectFunc)(src, w, gzipped, mime)
}
