package datactx

import (
	"context"
	"time"
)

type DataOption func(so *DataContext) error

// ResultPrefix sets the prefix of the result
func ResultPrefix(pfx string) DataOption {
	return func(dc *DataContext) error {
		dc.ResultPrefix = pfx
		return nil
	}
}

// CacheKeyPrefix sets the cache prefix of the data
func CacheKeyPrefix(pfx string) DataOption {
	return func(dc *DataContext) error {
		dc.CacheKeyPrefix = pfx
		return nil
	}
}

// CacheDuration sets the cache duration
func CacheDuration(durInMilSecs int) DataOption {
	return func(dc *DataContext) error {
		dc.CacheDuration = time.Duration(durInMilSecs) * time.Millisecond
		return nil
	}
}

// Context sets the context
func Context(ctx context.Context) DataOption {
	return func(dc *DataContext) error {
		dc.Context = ctx
		return nil
	}
}
