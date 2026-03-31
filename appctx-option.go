package appctx

import (
	"sync"

	cfg "github.com/eaglebush/config"
)

// LibraryID sets the library id
func LibraryID(id string) MetaOption {
	return func(d *Meta) error {
		d.LibraryID = id
		return nil
	}
}

// ApplicationID sets the application id
func ApplicationID(id string) MetaOption {
	return func(d *Meta) error {
		d.ApplicationID = id
		return nil
	}
}

// ServiceID sets the service id
func ServiceID(id string) MetaOption {
	return func(d *Meta) error {
		d.ServiceID = id
		return nil
	}
}

// Config sets the configuration
func Config(cfgOpts cfg.Configuration) MetaOption {
	return func(d *Meta) error {
		d.Configuration = cfgOpts
		return nil
	}
}

// MiscVar sets a named value
func MiscVar(rw *sync.RWMutex, name string, value any) MetaOption {
	return func(d *Meta) error {
		d.SetVar(name, value)
		return nil
	}
}
