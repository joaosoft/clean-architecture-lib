package sqs

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	sqsSdk "github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joaosoft/clean-infrastructure/domain"
	sqsConfig "github.com/joaosoft/clean-infrastructure/sqs/config"
)

// SQS service
type SQS struct {
	// Name
	name string
	// App
	app domain.IApp
	// Configuration
	config *sqsConfig.Config
	// Connections
	connections map[string]*Connection
	// Additional Config Type
	additionalConfigType interface{}
	// Started
	started bool
}

type Connection struct {
	// Service Name
	serviceName string
	// Name
	name string
	// App
	app domain.IApp
	// Config
	config *sqsConfig.Connection
	// Consumer
	consumer *sqs.SQS
	// Producer Connection
	producer *sqs.SQS
	// Migration
	migration *Migration
	// Consumers
	consumers []domain.ISQSConsumer
}

type Migration struct {
	name       string
	config     *sqsConfig.Connection
	connection *sqsSdk.SQS
}
