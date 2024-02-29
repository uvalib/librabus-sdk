//
//
//

package uvalibrabus

import (
	"encoding/json"
	"fmt"
)

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
