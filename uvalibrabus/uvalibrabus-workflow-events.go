//
// Events related to workflow lifecycle
//

package uvalibrabus

import (
	"encoding/json"
	"fmt"
)

//
// event names
//

var EventWorkCreate = "workflow.work.create"       // work created
var EventWorkDelete = "workflow.work.delete"       // work deleted
var EventWorkPublish = "workflow.work.publish"     // work published
var EventWorkUnpublish = "workflow.work.unpublish" // work unpublished

//
// corresponding schema for these events
//

type UvaWorkflowEvent struct {
}

// standard behavior
func (impl UvaWorkflowEvent) String() string {
	return "<none>"
}

func (impl UvaWorkflowEvent) Serialize() ([]byte, error) {
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
