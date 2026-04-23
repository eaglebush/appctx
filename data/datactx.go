package datactx

import (
	"context"
	"errors"
	"time"

	"github.com/eaglebush/appctx"
	ck "github.com/eaglebush/cachekit"
	di "github.com/eaglebush/datainfo"
)

var (
	ErrMetaUnset = errors.New("meta unset")
)

type DataContext struct {
	appctx.Meta
	di.DataInfo
	Context        context.Context
	Cache          ck.Cache
	CacheKeyPrefix string
	CacheDuration  time.Duration
	ResultPrefix   string
}

func NewDataContext(
	mt *appctx.Meta,
	do ...DataOption,
) (*DataContext, error) {
	if mt == nil {
		return nil, ErrMetaUnset
	}
	dc := DataContext{
		Meta: *mt,
	}
	for _, o := range do {
		o(&dc)
	}
	return &dc, nil
}

// WithContext sets the context of the returned DataContext without affecting the original context
func (obj *DataContext) WithContext(ctx context.Context) *DataContext {
	oc := *obj
	oc.Context = ctx
	return &oc
}
