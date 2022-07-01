package sender

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func PublishToPubsub(projectID, topicID string, messageToSend map[string]interface{}) (error, string) {

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err), ""
	}
	defer client.Close()

	t := client.Topic(topicID)
	jstring, _ := json.Marshal(messageToSend)
	result := t.Publish(ctx, &pubsub.Message{
		Data: jstring,
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err), ""
	}
	return nil, id
}
