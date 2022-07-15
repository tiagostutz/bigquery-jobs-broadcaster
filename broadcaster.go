package main

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
)

type BroadcastConfig struct {
	Project   string `envconfig:"PROJECT"`
	TopicName string `envconfig:"TOPIC"`
	Port      string `envconfig:"PORT" default:"8989"`
}

func broadcastJobCompletedEventarc(ctx context.Context, eventarcPayload EventarcPayload, topic string) error {

	pubsubClient, err := pubsub.NewClient(ctx, config.Project)
	if err != nil {
		logrus.Errorf("error creating a pubsub client for a topic different than default. Details: %s", err)
		return err
	}

	defer pubsubClient.Close()

	errorDetails := "none"
	if len(eventarcPayload.ProtoPayload.ServiceData.JobCompletedEvent.Job.JobStatus.Error.Message) > 0 {
		errorDetails = eventarcPayload.ProtoPayload.ServiceData.JobCompletedEvent.Job.JobStatus.Error.Message
	}

	// send the eventarc as pubsub message
	eventarcPayloadBytes, err := json.Marshal(eventarcPayload)
	if err != nil {
		logrus.Errorf("error to marshaling pubsub message body: %s", err)
		return err
	}

	topicClient := pubsubClient.Topic(topic)
	if topicClient == nil {
		logrus.Errorf("could not publish message. Topic %s doesnt exist on project %s. Please provide a topic that exists on the respective project.", topic, config.Project)
		return fmt.Errorf("error publishing message to %s topic. See server logs for details", topic)
	}

	msgId, err := topicClient.Publish(ctx, &pubsub.Message{Data: eventarcPayloadBytes}).Get(ctx)
	if err != nil {
		logrus.Errorf("error publishing the eventarc message to topic %s. Error: %s", topic, err)
		return err
	}

	logrus.WithFields(logrus.Fields{
		"pub-msg-id": msgId,
		"bq-error":   errorDetails,
		"job-id":     eventarcPayload.ProtoPayload.ServiceData.JobCompletedEvent.Job.JobName.JobID,
	}).Debugf("Big Query eventarc successfully processed and forwarded to topic %s/%s", config.Project, config.TopicName)

	return nil
}
