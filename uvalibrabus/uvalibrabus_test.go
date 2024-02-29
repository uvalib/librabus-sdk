//
//
//

package uvalibrabus

import (
	"testing"
)

var sourceName = "testing.unit.automated"
var goodBusName = "uva-libra-bus-staging"
var badBusName = "xxx"
var goodNamespace = "libraopen"
var goodIdentifier = "xxx"

func TestPublishHappyDay(t *testing.T) {
	cfg := UvaBusConfig{sourceName, goodBusName, nil}
	bus, err := NewUvaBus(cfg)
	if err != nil {
		t.Fatalf("expected 'OK' but got '%s'\n", err)
	}

	ev := NewUvaBusEvent(goodNamespace, goodIdentifier)
	err = bus.PublishBusEvent(EventTest, ev)
	if err != nil {
		t.Fatalf("expected 'OK' but got '%s'\n", err)
	}
}

//
// end of file
//
