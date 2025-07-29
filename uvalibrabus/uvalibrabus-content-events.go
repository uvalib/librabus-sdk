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

var EventContentView = "content.object.view"         // view event
var EventContentDownload = "content.object.download" // download event

//
// corresponding schema for these events
//

type UvaContentEvent struct {
	SourceIp       string `json:"source-ip"`       // source ip address
	Referrer       string `json:"referrer"`        // referrer url
	UserAgent      string `json:"user-agent"`      // user agent
	AcceptLanguage string `json:"accept-language"` // accept language
}

// standard behavior
func (impl UvaContentEvent) String() string {
	return fmt.Sprintf("<%s|%s|%s|%s>",
		impl.SourceIp, impl.Referrer, impl.UserAgent, impl.AcceptLanguage)
}

func (impl UvaContentEvent) Serialize() ([]byte, error) {
	// serialize the event object
	buf, err := json.Marshal(impl)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, ErrEventSerialize)
	}
	return buf, nil
}

func MakeContentEvent(buf []byte) (*UvaContentEvent, error) {
	var event UvaContentEvent
	err := json.Unmarshal(buf, &event)
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, ErrEventDeserialize)
	}
	return &event, nil
}

//
// end of file
//
