package main

import (
	"flag"
	"fmt"
	"github.com/uvalib/librabus-sdk/uvalibrabus"
	"log"
	"os"
	"strings"
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
	var bulk string

	flag.StringVar(&eventBus, "bus", "", "Event bus name")
	flag.StringVar(&eventSource, "source", "", "Event source name")
	flag.StringVar(&eventName, "event", "", "The name of the event")
	flag.StringVar(&namespace, "ns", "", "The event namespace")
	flag.StringVar(&oid, "oid", "", "The event object identifier")
	flag.StringVar(&bulk, "bulk", "", "File of events to publish (name/namespace/oid)")
	flag.BoolVar(&debug, "debug", false, "Log debug information")
	flag.Parse()

	if debug == true {
		logger = log.Default()
	}

	// validate required parameters
	if len(eventBus) == 0 ||
		len(eventSource) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// validate optional parameters
	if len(bulk) == 0 && (len(eventName) == 0 ||
		len(namespace) == 0 ||
		len(oid) == 0) {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// create the bus client
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

	count := 0
	if len(bulk) != 0 {

		var bytesRead []byte
		bytesRead, err = os.ReadFile(bulk)
		if err == nil {
			lines := strings.Split(string(bytesRead), "\n")
			for _, line := range lines {
				attrs := strings.Split(line, "/")
				if len(attrs) == 3 {
					err = publishEvent(bus, attrs[0], attrs[1], attrs[2])
					count++
					if err != nil {
						break
					}
				} else {
					if len(line) != 0 {
						fmt.Printf("WARNING: ignoring: %s\n", line)
					}
				}
			}
		}
	} else {
		err = publishEvent(bus, eventName, namespace, oid)
		count++
	}

	if err != nil {
		fmt.Printf("Terminating with error (%s)\n", err.Error())
		os.Exit(1)
	} else {
		fmt.Printf("Terminating normally, %d event(s) published\n", count)
	}
}

func publishEvent(bus uvalibrabus.UvaBus, name string, namespace string, oid string) error {

	ev := uvalibrabus.UvaBusEvent{
		EventName:  name,
		Namespace:  namespace,
		Identifier: oid,
	}
	err := bus.PublishEvent(&ev)
	if err != nil {
		return fmt.Errorf("publishing event (%s)", err.Error())
	}

	fmt.Printf("INFO: published: %s\n", ev.String())
	return nil
}

//
// end of file
//
