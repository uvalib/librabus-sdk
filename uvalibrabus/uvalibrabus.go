//
//
//

package uvalibrabus

import (
	"fmt"
	"log"
)

// errors
var ErrBadParameter = fmt.Errorf("bad parameter")
var ErrConfig = fmt.Errorf("configuration error")
var ErrEventSerialize = fmt.Errorf("serialize error")
var ErrEventPublish = fmt.Errorf("publish error")

// event names
var EventTest = "test.event"                // used for testing
var EventObjectCreate = "object.create"     // new object created, typically from easystore
var EventObjectDelete = "object.create"     // object deleted, typically from easystore
var EventMetadataUpdate = "metadata.update" // metadata updated, typically from easystore
var EventFileUpdate = "file.update"         // file updated, typically from easystore

type UvaBusEvent interface {
	GetEventId() string    // some sort of opaque event identifier
	GetNamespace() string  // object namespace
	GetIdentifier() string // object identifier
}

type UvaBus interface {
	PublishBusEvent(string, UvaBusEvent) error // publish the specified event
}

// UvaBusConfig -- our configuration structure
type UvaBusConfig struct {
	Source  string      // the event source
	BusName string      // the name of the bus
	log     *log.Logger // logger if we want to log
}

// NewUvaBus -- factory for our bus interface
func NewUvaBus(config UvaBusConfig) (UvaBus, error) {
	bus, err := newUvaBus(config)
	return bus, err
}

// NewUvaBusEvent -- factory for our bus event interface
func NewUvaBusEvent(namespace string, identifier string) UvaBusEvent {
	return newUvaBusEvent(namespace, identifier)
}

//
// end of file
//
