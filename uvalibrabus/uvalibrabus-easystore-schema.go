//
//
//

package uvalibrabus

import (
	"encoding/json"
	"fmt"
)

// event names published by the easystore service
var EventObjectCreate = "object.create"     // object created
var EventObjectUpdate = "object.update"     // object updated
var EventObjectDelete = "object.delete"     // object deleted
var EventMetadataUpdate = "metadata.update" // metadata updated
var EventFileCreate = "file.create"         // file created
var EventFileUpdate = "file.update"         // file updated
var EventFileDelete = "file.delete"         // file deleted

// this is our easystore event object implementation
type uvaEasystoreEvent struct {
	theType    string // the event type (not serialized)
	Namespace  string `json:"namespace"`  // namespace
	Identifier string `json:"identifier"` // identifier
}

func (impl uvaEasystoreEvent) Serialize() ([]byte, error) {
	// serialize the event object
	buf, err := json.Marshal(impl)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, ErrEventSerialize)
	}
	return buf, nil
}

func (impl uvaEasystoreEvent) Type() string {
	return impl.theType
}

// NewEasystoreEvent -- factory for our easystore events
func NewEasystoreEvent(eventType string, namespace string, identifier string) UvaBusEvent {
	return uvaEasystoreEvent{eventType, namespace, identifier}
}

//
// end of file
//
