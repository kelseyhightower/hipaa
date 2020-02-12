package hipaa

import (
	"context"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/mitchellh/cli"
)

func SecureListenAndServe() error {
	return http.ListenAndServe(":443", nil)
}

func initCLI() {
	cli.NewCLI("hipaa", "1.0.0")
}

func CreatePubSubClient() (*pubsub.Client, error) {
	return pubsub.NewClient(context.Background(), "")
}
