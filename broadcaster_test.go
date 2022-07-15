package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/assert"
)

// set default values for env variables
func setup() {
	if os.Getenv("PROJECT") == "" {
		os.Setenv("PROJECT", "local-test")
	}
	if os.Getenv("TOPIC") == "" {
		os.Setenv("TOPIC", "dummy")
	}
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		os.Setenv("PUBSUB_EMULATOR_HOST", "0.0.0.0:8262")
	}
	if os.Getenv("PUBSUB_PROJECT_ID") == "" {
		os.Setenv("PUBSUB_PROJECT_ID", "local-test")
	}
}

func TestBroadcastJobCompletedEventarc(t *testing.T) {
	setup()
	ctx = context.Background()
	config = BroadcastConfig{}
	envconfig.MustProcess("", &config)

	client, err := pubsub.NewClient(ctx, config.Project)
	if err != nil {
		t.Errorf("failed to create pubsub client: %s", err)
	}
	defer client.Close()

	topicID := os.Getenv("TOPIC")
	topic := client.Topic(topicID)
	topicExists, err := topic.Exists(ctx)
	if err != nil {
		t.Errorf("failed to check if pubsub topic exists: %s", err)
	}
	if !topicExists {
		topic, err = client.CreateTopic(ctx, topicID)
		if err != nil {
			t.Errorf("Error creating the topic: %s", err)
		}
	}
	assert.NotNil(t, topic)

	// Open our jsonFile
	jsonFile, err := os.Open("sample-payload.json")
	if err != nil {
		t.Errorf("Error loading sample payload for test. Err: %s", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var eventarcPayload EventarcPayload

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &eventarcPayload)

	// create subscription to test the received result
	subsc := client.Subscription("subscMock")
	if err != nil {
		t.Errorf("error resolving subscription: %v", err)
	}
	subscExists, err := subsc.Exists(ctx)
	if err != nil {
		t.Errorf("error checking if subscriptions exists: %v", err)
	}
	if !subscExists {
		subsc, err = client.CreateSubscription(ctx, "subscMock", pubsub.SubscriptionConfig{
			Topic: topic,
		})
		if err != nil {
			t.Errorf("error creating subscription to validate: %v", err)
		}
	}

	err = publishJobCompletedEventarcBroadcast(ctx, eventarcPayload, topicID)
	assert.NoError(t, err)

	// as this is a test, receive messages for 1 second to make sure we are grabbing
	// the test message and then keep going
	ctxTime, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	var received EventarcPayload
	err = subsc.Receive(ctxTime, func(_ context.Context, msg *pubsub.Message) {
		json.Unmarshal(msg.Data, &received)
		assert.Equal(t, received.ProtoPayload.Type, EventarcTypeAuditLog)
		assert.Equal(t, received.ProtoPayload.MethodName, EventarcMethodNameJobCompleted)
		msg.Ack()
	})
	if err != nil {
		t.Errorf("sub.Receive: %v", err)
	}
	assert.NotNil(t, received)
}
