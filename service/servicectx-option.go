package servicectx

import (
	"io"
	"net/http"

	l "github.com/stdutil/log"
)

type ServiceOption func(so *ServiceContext) error

// LogFunc sets the Log function
func LogFunc(f func(msgType l.LogType, message ...string)) ServiceOption {
	return func(sc *ServiceContext) error {
		sc.logFunc = f
		return nil
	}
}

// RespondFunc sets the Respond function
func RespondFunc(f func(data any, w http.ResponseWriter, r *http.Request) error) ServiceOption {
	return func(sc *ServiceContext) error {
		sc.respondFunc = f
		return nil
	}
}

// RespondBytesFunc sets the respondBytes function
func RespondBytesFunc(f func(data []byte, fileExt string, w http.ResponseWriter, r *http.Request) error) ServiceOption {
	return func(sc *ServiceContext) error {
		sc.respondBytesFunc = f
		return nil
	}
}

// RespondDirectFunc sets the responseDirect function
func RespondDirectFunc(f func(src io.ReadCloser, w http.ResponseWriter, gzipped bool, mime string) error) ServiceOption {
	return func(sc *ServiceContext) error {
		sc.respondDirectFunc = f
		return nil
	}
}
