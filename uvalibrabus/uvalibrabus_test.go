//
//
//

package uvalibrabus

import (
	"testing"
)

var sourceName = "testing.unit.automated"
var goodBusName = "uva-libra-bus-staging"
var goodNamespace = "libraopen"
var goodIdentifier = "oid:xx-example-xx"

//var goodVTag = "vtag:xx-example-xx"

func TestPublishHappyDay(t *testing.T) {
	cfg := UvaBusConfig{
		Source:  sourceName,
		BusName: goodBusName,
		Log:     nil}
	bus, err := NewUvaBus(cfg)
	if err != nil {
		t.Fatalf("expected 'OK' but got '%s'\n", err)
	}

	// create a new event
	ev := UvaBusEvent{
		EventName:  EventTest,
		Namespace:  goodNamespace,
		Identifier: goodIdentifier,
	}
	err = bus.PublishEvent(&ev)
	if err != nil {
		t.Fatalf("expected 'OK' but got '%s'\n", err)
	}
}

func TestEventSerialize(t *testing.T) {

	// create a new event
	ev := UvaBusEvent{
		EventName:  EventTest,
		Namespace:  goodNamespace,
		Identifier: goodIdentifier,
	}

	b, err := ev.Serialize()
	if err != nil {
		t.Fatalf("expected 'OK' but got '%s'\n", err)
	}
	evCopy, err := MakeBusEvent(b)
	if err != nil {
		t.Fatalf("expected 'OK' but got '%s'\n", err)
	}
	if ev.String() != evCopy.String() {
		t.Fatalf("expected '%s' but got '%s'\n", ev.String(), evCopy.String())
	}
}

//
// end of file
//
