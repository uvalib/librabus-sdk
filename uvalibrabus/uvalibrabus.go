//
//
//

package uvalibrabus

import (
	"encoding/json"
	"fmt"
	"log"
)

// errors
var ErrBadParameter = fmt.Errorf("bad parameter")
var ErrConfig = fmt.Errorf("configuration error")
var ErrEventSerialize = fmt.Errorf("serialize error")
var ErrEventDeserialize = fmt.Errorf("deserialize error")
var ErrEventPublish = fmt.Errorf("publish error")

// definitions for more specific events are in their respective files
var EventTest = "ignore.test.event" // used for testing, should be ignored

// UvaBusEvent -- our event bus
type UvaBusEvent struct {
	EventName  string          `json:"name"`       // event name
	Namespace  string          `json:"namespace"`  // namespace
	Identifier string          `json:"identifier"` // identifier
	Detail     json.RawMessage // remainder of the schema is event specific

	// other stuff
}

// UvaBus -- abstract event bus
type UvaBus interface {
	PublishEvent(UvaBusEvent) error // publish the specified event
}

// UvaBusConfig -- our configuration structure
type UvaBusConfig struct {
	Source  string      // the event source
	BusName string      // the name of the bus
	Log     *log.Logger // logger if we want to log
}

// NewUvaBus -- factory for our bus interface
func NewUvaBus(config UvaBusConfig) (UvaBus, error) {
	bus, err := newUvaBus(config)
	return bus, err
}

//
// end of file
//
