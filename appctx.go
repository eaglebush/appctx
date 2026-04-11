package appctx

import (
	"errors"
	"sync"

	cfg "github.com/eaglebush/config"
	vm "github.com/eaglebush/valuemap"
	evt "github.com/stdutil/event"
)

var (
	ErrConfigurationUnset = errors.New("configuration unset")
)

type (
	Meta struct {
		ApplicationID string
		ServiceID     string
		LibraryID     string
		EventSubject  *evt.EventSubject // Event subject for event generation
		Lock          *sync.RWMutex
		miscVar       *vm.ValueMap[string, any]
		config        *cfg.Configuration
	}
	MetaOption func(do *Meta) error
)

func NewMeta(mo ...MetaOption) *Meta {
	mt := Meta{}
	// Create a value map that requires external lock
	// The Lock field will take care of it
	mv := vm.NewExternalLock[string, any]()
	mt.miscVar = mv
	for _, o := range mo {
		o(&mt)
	}
	if mt.Lock == nil {
		mt.Lock = &sync.RWMutex{}
	}
	if mt.ApplicationID == "" {
		mt.ApplicationID = "DEFAULT"
	}
	if mt.ServiceID == "" {
		mt.ServiceID = "DEFAULT"
	}
	if mt.LibraryID == "" {
		mt.LibraryID = "DEFAULT"
	}
	return &mt
}

// Copy copies an existeng meta and then updates members via MetaOptions
func Copy(m *Meta, mo ...MetaOption) *Meta {
	mt := Meta{
		ApplicationID: m.ApplicationID,
		ServiceID:     m.ServiceID,
		LibraryID:     m.LibraryID,
		EventSubject:  m.EventSubject,
		Lock:          m.Lock,
		miscVar:       m.miscVar.Clone(m.Lock),
		config:        m.config,
	}
	for _, o := range mo {
		o(&mt)
	}
	m.SetEventSubject() // Internally update event subject
	return &mt
}

// GetVarToType returns the value to a specified type.
//
// If the map is nil, name does not exist or invalid return type, this will return the zero value of the type indicated
func GetVarToType[T any](rw *sync.RWMutex, miscVar *vm.ValueMap[string, any], name string) T {
	var def T
	if rw == nil {
		return def
	}
	if miscVar == nil {
		return def
	}
	v, ok := miscVar.Get(rw, name)
	if !ok {
		return def
	}
	value, ok := v.(T)
	if !ok {
		return def
	}
	return value
}

// GetVarMap returns the initialized map pointer for static call use
func (mt *Meta) GetVarMap() *vm.ValueMap[string, any] {
	return mt.miscVar
}

// GetVar gets the miscellaneous variable value. If it is not found, nil is returned
func (mt *Meta) GetVar(name string) any {
	value, ok := mt.miscVar.Get(mt.Lock, name)
	if !ok {
		return nil
	}
	return value
}

// SetVar sets a miscellaneous variable to a value
func (mt *Meta) SetVar(name string, value any) {
	mt.miscVar.Set(mt.Lock, name, value)
}

// SetEventSubject sets or refreshes the event subject via parameters
func (mt *Meta) SetEventSubject() {
	if mt.EventSubject == nil {
		mt.EventSubject = &evt.EventSubject{}
	}
	*mt.EventSubject = evt.NewEventSubjectBase(mt.ApplicationID, mt.ServiceID, mt.LibraryID)
}

// GetConfig returns a verified configuration
func (mt *Meta) GetConfig() (*cfg.Configuration, error) {
	if mt.config == nil {
		return nil, ErrConfigurationUnset
	}
	return mt.config, nil
}
