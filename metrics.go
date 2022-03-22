package main

import (
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

	id := ""
	envPlatform := ""

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
	// TODO
	return ""
}

// Helper for GeneratePayload()
func GetEnvPlatform() Platform {
	// TODO
	return Unknown
}

// Send to Pubsub
// https://cloud.google.com/pubsub/docs/reference/libraries#client-libraries-install-go
func PublishMetricsPayload(p MetricsPayload) error {
	// ctx := context.Background()
	// projectID := "hello-telemetry"
	// client, err := pubsub.NewClient(ctx, projectID)
	// if err != nil {
	// 	return err
	// }
	// defer client.Close()

	// Sets the id for the new topic.
	// topicID := "bq-writer"

	// Serialize payload as JSON
	// TODO
	// Publish payload to pub/sub
	// TODO
	return nil
}

// Run PublishMetricsPayload every day - make sure to modify payload inline to set IS_STARTUP to false!
func PeriodicPing(p MetricsPayload) {
	// TODO
}
