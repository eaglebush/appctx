package datactx

import (
	"context"

	"github.com/eaglebush/appctx"
	ck "github.com/eaglebush/cachekit"
	di "github.com/eaglebush/datainfo"
)

type DataContext struct {
	appctx.Meta
	di.DataInfo
	Cache          ck.Cache
	CacheKeyPrefix string
	CacheDuration  int
	Context        context.Context
}
