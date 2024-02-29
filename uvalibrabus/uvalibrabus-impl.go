//
//
//

package uvalibrabus

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/google/uuid"
)

// this is our bus interface implementation
type uvaBusImpl struct {
	config UvaBusConfig
	client *cloudwatchevents.Client
}

// this is our s3 object implementation
type uvaBusEventImpl struct {
	EventId    string `json:"eventId"`    // opaque event id
	Namespace  string `json:"namespace"`  // namespace
	Identifier string `json:"identifier"` // identifier
}

// newUvaBus -- factory for our bus interface
func newUvaBus(busConfig UvaBusConfig) (UvaBus, error) {

	// validate inbound config
	if len(busConfig.Source) == 0 {
		return nil, fmt.Errorf("%q: %w", "source is blank", ErrBadParameter)
	}
	if len(busConfig.BusName) == 0 {
		return nil, fmt.Errorf("%q: %w", "bus name is blank", ErrBadParameter)
	}

	// Set up EventBridge client
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("%q: %w", err, ErrConfig)
	}
	client := cloudwatchevents.NewFromConfig(cfg)

	var impl uvaBusImpl
	impl.config = busConfig
	impl.client = client

	return &impl, nil
}

// newUvaBusEvent -- factory for our bus event interface
func newUvaBusEvent(namespace string, identifier string) UvaBusEvent {
	ev := uvaBusEventImpl{
		EventId:    uuid.New().String(),
		Namespace:  namespace,
		Identifier: identifier}
	return ev
}

func (impl uvaBusImpl) PublishBusEvent(eventName string, event UvaBusEvent) error {

	// validate inbound parameters
	if len(eventName) == 0 {
		return fmt.Errorf("%q: %w", "eventName is blank", ErrBadParameter)
	}
	if len(event.GetNamespace()) == 0 {
		return fmt.Errorf("%q: %w", "event namespace is blank", ErrBadParameter)
	}
	if len(event.GetIdentifier()) == 0 {
		return fmt.Errorf("%q: %w", "event object identifier is blank", ErrBadParameter)
	}

	impl.logInfo(fmt.Sprintf("publish event [%s], id/ns/oid: [%s/%s/%s]", eventName, event.GetEventId(), event.GetNamespace(), event.GetIdentifier()))

	// serialize the event object
	buf, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("%q: %w", err, ErrEventSerialize)
	}

	// publish it
	_, err = impl.client.PutEvents(context.Background(),
		&cloudwatchevents.PutEventsInput{
			Entries: []types.PutEventsRequestEntry{
				{
					EventBusName: aws.String(impl.config.BusName),
					Source:       aws.String(impl.config.Source),

					DetailType: aws.String(eventName),
					Detail:     aws.String(string(buf)),
				},
			},
		})
	if err != nil {
		return fmt.Errorf("%q: %w", err, ErrEventPublish)
	}
	return nil
}

//
// helpers
//

func (impl uvaBusImpl) logInfo(message string) {
	if impl.config.log != nil {
		log.Printf("INFO: %s", message)
	}
}

//
// uvaBusEventImpl implementation methods
//

func (impl uvaBusEventImpl) GetEventId() string {
	return impl.EventId
}

func (impl uvaBusEventImpl) GetNamespace() string {
	return impl.Namespace
}

func (impl uvaBusEventImpl) GetIdentifier() string {
	return impl.Identifier
}

//
// end of file
//
