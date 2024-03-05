//
// Events related to storage lifecycle
//

package uvalibrabus

import (
	"encoding/json"
	"fmt"
)

//
// event names
//

var EventObjectCreate = "storage.object.create"     // object created
var EventObjectUpdate = "storage.object.update"     // object updated
var EventObjectDelete = "storage.object.delete"     // object deleted
var EventMetadataUpdate = "storage.metadata.update" // metadata updated
var EventFileCreate = "storage.file.create"         // file created
var EventFileUpdate = "storage.file.update"         // file updated
var EventFileDelete = "storage.file.delete"         // file deleted

//
// corresponding schema for these events
//

type UvaStorageEvent struct {
	VTag string `json:"vtag"` // object vtag
}

// standard behavior
func (impl UvaStorageEvent) String() string {
	return fmt.Sprintf("<%s>", impl.VTag)
}

func (impl UvaStorageEvent) Serialize() ([]byte, error) {
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
