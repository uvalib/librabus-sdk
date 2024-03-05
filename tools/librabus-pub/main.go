package main

import (
	"flag"
	"fmt"
	"github.com/uvalib/librabus-sdk/uvalibrabus"
	"log"
	"os"
)

// main entry point
func main() {

	var eventBus string
	var eventSource string
	var eventName string
	var debug bool
	var logger *log.Logger
	var namespace string
	var oid string

	flag.StringVar(&eventBus, "bus", "", "Event bus name")
	flag.StringVar(&eventSource, "source", "", "Event source name")
	flag.StringVar(&eventName, "event", "", "The name of the event")
	flag.StringVar(&namespace, "ns", "", "The event namespace")
	flag.StringVar(&oid, "oid", "", "The event object identifier")
	flag.BoolVar(&debug, "debug", false, "Log debug information")
	flag.Parse()

	if debug == true {
		logger = log.Default()
	}

	// validate parameters
	if len(eventBus) == 0 ||
		len(eventSource) == 0 ||
		len(eventName) == 0 ||
		len(namespace) == 0 ||
		len(oid) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	cfg := uvalibrabus.UvaBusConfig{
		Source:  eventSource,
		BusName: eventBus,
		Log:     logger,
	}
	bus, err := uvalibrabus.NewUvaBus(cfg)
	if err != nil {
		log.Fatalf("ERROR: creating event bus client (%s)", err.Error())
	}
	fmt.Printf("Using: %s@%s\n", eventSource, eventBus)

	// for now, just 1 event type
	ev := uvalibrabus.UvaBusEvent{
		EventName:  eventName,
		Namespace:  namespace,
		Identifier: oid,
	}
	err = bus.PublishEvent(ev)
	if err != nil {
		log.Fatalf("ERROR: publishing event (%s)", err.Error())
	}

	fmt.Printf("Published: %s\n", ev.String())
	fmt.Printf("Terminating normally\n")
}

//
// end of file
//
