//
//
//

package uvalibrabus

import (
	"encoding/json"
	"fmt"
)

func (impl UvaBusEvent) Serialize() ([]byte, error) {
	// serialize the event object
	buf, err := json.Marshal(impl)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, ErrEventSerialize)
	}
	return buf, nil
}

func (impl UvaBusEvent) String() string {
	return fmt.Sprintf("%s/%s/%s/%s", impl.EventName, impl.Namespace, impl.Identifier, impl.VTag)
}

func MakeBusEvent(buf []byte) (*UvaBusEvent, error) {
	var event UvaBusEvent
	err := json.Unmarshal(buf, &event)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, ErrEventDeserialize)
	}
	return &event, nil
}

//
// end of file
//
