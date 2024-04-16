//
// Events related to work audits
//

package uvalibrabus

import (
	"encoding/json"
	"fmt"
)

//
// event names
//

var EventFieldUpdate = "audit.field.update" // object created

//
// corresponding schema for these events
//

type UvaAuditEvent struct {
	Who       string `json:"who"`       // who did the update (computing ID)
	FieldName string `json:"fieldName"` // name of the field that was updated
	Before    string `json:"before"`    // value before the update
	After     string `json:"after"`     // value after the update
}

// standard behavior
func (impl UvaAuditEvent) String() string {
	return fmt.Sprintf("<%s/%s/%s/%s>",
		impl.Who, impl.FieldName, impl.Before, impl.After)
}

func (impl UvaAuditEvent) Serialize() ([]byte, error) {
	// serialize the event object
	buf, err := json.Marshal(impl)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, ErrEventSerialize)
	}
	return buf, nil
}

//
// end of file
//
