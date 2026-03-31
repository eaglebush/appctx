package datactx

import (
	"context"
	"time"

	"github.com/eaglebush/appctx"
	ck "github.com/eaglebush/cachekit"
	di "github.com/eaglebush/datainfo"
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
) *DataContext {
	dc := DataContext{
		Meta: *mt,
	}
	for _, o := range do {
		o(&dc)
	}
	return &dc
}
