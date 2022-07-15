package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var ctx context.Context
var config BroadcastConfig

func main() {
	logrus.Info("starting bigquery job broadcaster server...")
	ctx = context.Background()
	config = BroadcastConfig{}
	envconfig.MustProcess("", &config)
	logrus.SetLevel(logrus.DebugLevel)

	if config.Project == "" {
		logrus.Fatal("PROJECT env var must be provided")
	}

	// handle request with Eventarc payload
	http.HandleFunc("/", handler)

	log.Printf("bigquery-job-broadcaster server listening on port %s", config.Port)
	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var bqEventarcRequestBody EventarcPayload

	topicToBroadcast := config.TopicName

	// if passed as a query parameter on the URL, it overrides the default (env set) topic path
	topicQueryParam := r.URL.Query()["topic"]
	if len(topicQueryParam) > 0 {
		if topicQueryParam[0] != config.TopicName { // using a different topic than the default one set as env
			logrus.Debugf("using the topic name received as URL param instead of the one configured in env. Topic in env: %s. Topic in URL query param: %s", config.TopicName, topicQueryParam[0])
		}
		topicToBroadcast = topicQueryParam[0]
	}

	if topicToBroadcast == "" {
		http.Error(w, "a topic must be specified to broadcast the eventarc. A topic name can be specified by setting `TOPIC` env var or by passing a `?topic=<topic_name>` query param to this request URL", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&bqEventarcRequestBody)
	if err != nil {
		logrus.Errorf("error decoding eventarc event payload. Details: %s", err)
		http.Error(w, "error decoding payload", http.StatusInternalServerError)
	}

	err = broadcastJobCompletedEventarc(ctx, bqEventarcRequestBody, topicToBroadcast)

	if err != nil {
		logrus.Errorf("error broadcasting eventarc event payload to topic. Details: %s", err)
		http.Error(w, "error broadcasting eventarc to pubsub topic", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"jobID": bqEventarcRequestBody.ProtoPayload.ServiceData.JobCompletedEvent.Job.JobName.JobID,
	})

}
