package servicectx

import (
	"errors"
	"io"
	"net/http"

	"github.com/eaglebush/appctx"
	ck "github.com/eaglebush/cachekit"
	l "github.com/stdutil/log/v2"
)

var (
	ErrMetaUnset                  = errors.New("meta unset")
	ErrLogFuncUnset               = errors.New("log function unset")
	ErrRespondFunctionUnset       = errors.New("respond function unset")
	ErrRespondBytesFunctionUnset  = errors.New("respond bytes function unset")
	ErrRespondDirectFunctionUnset = errors.New("respond direct function unset")
)

type ServiceContext struct {
	appctx.Meta
	Cache             ck.Cache
	tokenHandling     TokenHandlingInfo // Token handling
	logFunc           func(msgType l.LogType, message ...string)
	respondFunc       func(data any, w http.ResponseWriter, r *http.Request) error                    // Respond function
	respondBytesFunc  func(data []byte, fileExt string, w http.ResponseWriter, r *http.Request) error // RespondBytes function
	respondDirectFunc func(src io.ReadCloser, w http.ResponseWriter, gzipped bool, mime string) error // RespondDirect function
}

type TokenHandlingInfo struct {
	Valid         bool
	ValidateTimes bool
	ValidateAppID bool
}

func NewServiceContext(
	mt *appctx.Meta,
	so ...ServiceOption,
) (*ServiceContext, error) {
	if mt == nil {
		return nil, ErrMetaUnset
	}
	sc := ServiceContext{
		Meta: *mt,
	}
	for _, o := range so {
		o(&sc)
	}
	return &sc, nil
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

// TokenHandling returns the token handling info
func (cs *ServiceContext) TokenHandling() TokenHandlingInfo {
	return cs.tokenHandling
}
