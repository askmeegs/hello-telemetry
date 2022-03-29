package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Enum for ENV_PLATFORM
type Platform string

const (
	CloudRun   Platform = "CLOUD_RUN"
	GKE        Platform = "GKE"
	Kubernetes Platform = "KUBERNETES_NON_GKE"
	Unknown    Platform = "UNKNOWN"
)

type MetricsPayload struct {
	Id          string `json:"id"`
	Timestamp   int64  `json:"timestamp"` //UNIX timestamp
	IsStartup   bool   `json:"is_startup"`
	AppName     string `json:"app_name"`
	AppVersion  string `json:"app_version"`
	EnvPlatform string `json:"env_platform"`
}

func GeneratePayload() (MetricsPayload, error) {
	// Generate unix timestamp - avoid timezones, daylight savings, etc.
	ts := time.Now().Unix()
	fmt.Printf("⏰ TIMESTAMP is: %d\n", ts)

	id := GenerateDemoId()
	envPlatform := "KUBERNETES"

	// Hardcode / use Const
	appName := "HELLO_TELEMETRY"
	appVersion := os.Getenv("APP_VERSION")

	// TODO
	return MetricsPayload{
		Id:          id,
		Timestamp:   ts,
		IsStartup:   true,
		AppName:     appName,
		AppVersion:  appVersion,
		EnvPlatform: envPlatform,
	}, nil
}

// Helper for GeneratePayload()
func GenerateDemoId() string {
	// inputs to hash

	appName := "HELLO_TELEMETRY"
	envPlatform := "KUBERNETES"
	namespaceId := "abcdefg"
	fingerprint := fmt.Sprintf("%s_%s_%s_HELLOWORLD", appName, envPlatform, namespaceId)

	// https://www.rfc-editor.org/rfc/rfc3174.html
	// TODO - determine if this algorithm is cryptographically safe / cannot be decoded
	// other algorithms are here. https://pkg.go.dev/crypto#Hash
	hasher := sha256.New()
	hasher.Write([]byte(fingerprint))

	// https://pkg.go.dev/hash#Hash.Sum
	sha := hasher.Sum(nil)

	// THIS WILL BE 64 DIGITS
	hexSha := fmt.Sprintf("🔑 HASH ID: %x", sha)
	fmt.Println(hexSha)
	return hexSha
}

// TODO - detect GKE. add metadataserver. make not garbage
// Helper for GeneratePayload()
func GetEnvPlatform() Platform {
	// rough version - detect if k8s or not
	// does KUBERNETES_PORT var exist?
	k8sPort := os.Getenv("KUBERNETES_PORT")
	if k8sPort == "" {
		return Unknown
	}
	return Kubernetes
}

// Send to Pubsub
// https://cloud.google.com/pubsub/docs/reference/libraries#client-libraries-install-go
func PublishMetricsPayload(p MetricsPayload) error {
	ctx := context.Background()
	projectID := "hello-telemetry"
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return err
	}
	defer client.Close()

	// Sets the id for the new topic.
	topicID := "bq-writer"

	// Serialize payload as JSON
	message, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	// Publish payload to pub/sub

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}

// Run PublishMetricsPayload every day - make sure to modify payload inline to set IS_STARTUP to false!
func PeriodicPing(p MetricsPayload) {
	// TODO
}
