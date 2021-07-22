package constant

import (
	"fmt"
	"os"

	metadata "cloud.google.com/go/compute/metadata"
)

const (
	topicKey = "TOPIC_NAME"
)

var (
	projectId = "dh-gcp-ads-sandbox"
	topicId   = "active-cacher-events"
)

func init() {
	if pId, err := metadata.ProjectID(); err != nil {
		fmt.Println("Error while setting projectId. Using default projectId:", projectId)
	} else {
		projectId = pId
		fmt.Println("Using projectId:", projectId)
	}

	if tId := os.Getenv(topicId); tId == "" {
		fmt.Printf("environment %s has not been set or is empty. Using default topicId: %s\n", topicKey, topicId)
	} else {
		topicId = tId
		fmt.Println("Using topic Id:", topicId)
	}
}

func GetProjectId() string {
	return projectId
}

func GetTopicId() string {
	return topicId
}
