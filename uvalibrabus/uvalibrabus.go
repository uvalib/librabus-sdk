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

//
// event names and supporting schema
//

// used for testing, should be ignored
var EventTest = "test.event" // used for testing

// emitted by the easystore service
var EventObjectCreate = "object.create"     // object created
var EventObjectUpdate = "object.update"     // object updated
var EventObjectDelete = "object.delete"     // object deleted
var EventMetadataUpdate = "metadata.update" // metadata updated
var EventFileCreate = "file.create"         // file created
var EventFileUpdate = "file.update"         // file updated
var EventFileDelete = "file.delete"         // file deleted

type UvaBusEvent interface {
	Type() string               // the type of event
	Serialize() ([]byte, error) // event serialize
}

type UvaBus interface {
	PublishEvent(UvaBusEvent) error // publish the specified event
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

//
// end of file
//
