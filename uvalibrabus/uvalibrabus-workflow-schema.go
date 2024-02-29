//
//
//

package uvalibrabus

import (
	"encoding/json"
	"fmt"
)

// event names published by the easystore service
var EventWorkCreate = "work.create"       // work published
var EventWorkDelete = "work.delete"       // work published
var EventWorkPublish = "work.publish"     // work published
var EventWorkUnpublish = "work.unpublish" // work published

// this is our workflow event object implementation
type uvaWorkflowEvent struct {
	theType    string // the event type (not serialized)
	Namespace  string `json:"namespace"`  // namespace
	Identifier string `json:"identifier"` // identifier
}

func (impl uvaWorkflowEvent) Serialize() ([]byte, error) {
	// serialize the event object
	buf, err := json.Marshal(impl)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, ErrEventSerialize)
	}
	return buf, nil
}

func (impl uvaWorkflowEvent) Type() string {
	return impl.theType
}

// NewWorkflowEvent -- factory for our workflow events
func NewWorkflowEvent(eventType string, namespace string, identifier string) UvaBusEvent {
	return uvaWorkflowEvent{eventType, namespace, identifier}
}

//
// end of file
//
