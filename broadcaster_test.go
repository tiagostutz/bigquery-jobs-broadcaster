package main

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/assert"
)

func TestBroadcastJobCompletedEventarc(t *testing.T) {
	ctx = context.Background()
	config = BroadcastConfig{}
	envconfig.MustProcess("", &config)

	client, err := pubsub.NewClient(ctx, config.Project)
	if err != nil {
		t.Errorf("failed to create pubsub client: %s", err)
	}
	defer client.Close()

	topicID := os.Getenv("TOPIC")
	topic, err := client.CreateTopic(ctx, topicID)
	if err != nil {
		t.Errorf("Error creating the topic: %s", err)
	}
	assert.NotNil(t, topic)
}
