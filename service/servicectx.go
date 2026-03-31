package servicectx

import (
	"io"
	"net/http"

	"github.com/eaglebush/appctx"
	ck "github.com/eaglebush/cachekit"
	l "github.com/stdutil/log"
)

type ServiceContext struct {
	appctx.Meta
	Cache             ck.Cache
	logFunc           func(msgType l.LogType, message ...string)
	respondFunc       func(data any, w http.ResponseWriter, r *http.Request) error                    // Respond function
	respondBytesFunc  func(data []byte, fileExt string, w http.ResponseWriter, r *http.Request) error // RespondBytes function
	respondDirectFunc func(src io.ReadCloser, w http.ResponseWriter, gzipped bool, mime string) error // RespondDirect function
}

func NewServiceContext() *ServiceContext {
	return &ServiceContext{
		
	}
}
