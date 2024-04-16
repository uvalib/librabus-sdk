//
//
//

package uvalibrabus

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

// this is our bus interface implementation
type uvaBusImpl struct {
	config UvaBusConfig
	client *cloudwatchevents.Client
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

func (impl uvaBusImpl) PublishEvent(event UvaBusEvent) error {

	// validate inbound parameters
	if len(event.EventName) == 0 {
		return fmt.Errorf("%q: %w", "event name is blank", ErrBadParameter)
	}
	impl.logInfo(fmt.Sprintf("publish event [%s]", event.String()))

	// serialize the event object
	buf, err := event.Serialize()
	if err != nil {
		return err
	}

	// publish it
	_, err = impl.client.PutEvents(context.Background(),
		&cloudwatchevents.PutEventsInput{
			Entries: []types.PutEventsRequestEntry{
				{
					EventBusName: aws.String(impl.config.BusName),
					Source:       aws.String(impl.config.Source),

					DetailType: aws.String(event.EventName),
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
	if impl.config.Log != nil {
		log.Printf("INFO: %s", message)
	}
}

//
// end of file
//
