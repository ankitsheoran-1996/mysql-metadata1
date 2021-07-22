package events

import (
	"context"
	"fmt"
	"mysql-metadata/constant"

	"cloud.google.com/go/pubsub"
)

type pubSubEvent struct {
	ctx    context.Context
	client *pubsub.Client
	topic  *pubsub.Topic
}

var event *pubSubEvent

// Initialize block to setup pubsub clients.
// Create subscription with hostname
func init() {
	event = &pubSubEvent{ctx: context.Background()}
	event.initClientAndTopic()
}

// Create a new client for the project
func (e *pubSubEvent) initClientAndTopic() {
	var err error
	e.client, err = pubsub.NewClient(e.ctx, constant.GetProjectId())
	if err != nil {
		panic(fmt.Errorf("pubsub.NewClient: %v", err))
	}
	e.topic = e.client.Topic(constant.GetTopicId())
	fmt.Printf("Connected to topic name:%s having topicId:%s\n", e.topic, e.topic.ID())
}

func PushMsgToTopic(eventType, bucketId, fileName string, dataType string) {
	var msg = &pubsub.Message{}
	msg.Attributes = make(map[string]string)
	msg.Attributes["eventType"] = eventType
	msg.Attributes["bucketId"] = bucketId
	msg.Attributes["objectId"] = fileName
	msg.Attributes["dataType"] = dataType

	res := event.topic.Publish(event.ctx, msg)
	if _, err := res.Get(event.ctx); err != nil {
		fmt.Println("Error occured while pushing message to topic")
		return
	}
}
