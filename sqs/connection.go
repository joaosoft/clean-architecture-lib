package sqs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/joaosoft/clean-infrastructure/domain/message"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"github.com/joaosoft/clean-infrastructure/domain"
	"github.com/joaosoft/clean-infrastructure/errors"
)

// Connect connects the sqs
func (c *Connection) Connect() error {
	// create a new aws session
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(c.config.Credentials.Region),
		Endpoint: aws.String(c.config.Credentials.Api),
	})

	if err != nil {
		log.Fatalf("failed to create session: %s", err)
		return err
	}

	conn := sqs.New(sess)
	c.consumer, c.producer = conn, conn

	c.migration, err = NewMigration(c.serviceName, c.config, conn)
	if err != nil {
		message.ErrorMessage(c.serviceName, fmt.Errorf("failed to create migration at AWS SQS: %s", err))
		return err
	}

	return nil
}

// Produce to the sqs
func (c *Connection) Produce(queue string, messageAttributes map[string]*sqs.MessageAttributeValue, messages ...string) error {
	var batch []*sqs.SendMessageBatchRequestEntry
	maskedQueue := c.maskQueue(c.app.Config().Env, queue)

	queueUrl, err := url.JoinPath(c.config.Credentials.Api, c.config.Credentials.IdAccount, maskedQueue)
	if err != nil {
		c.app.Logger().Log().Do(err, &domain.LoggerInfo{Msg: fmt.Sprintf("Failed to publish message to queue %s: %s", maskedQueue, err)})
		return err
	}

	for i := range messages {
		batch = append(batch, &sqs.SendMessageBatchRequestEntry{
			Id:                aws.String(uuid.New().String()),
			MessageBody:       aws.String(messages[i]),
			MessageAttributes: messageAttributes,
		})

		// If batch reaches the MaxNumberOfMessages it means it's full and we should send it
		if len(batch) == int(c.config.MaxNumberOfMessages) {
			_, err = c.producer.SendMessageBatch(&sqs.SendMessageBatchInput{
				QueueUrl: aws.String(queueUrl),
				Entries:  batch,
			})
			if err != nil {
				c.app.Logger().Log().Do(err, &domain.LoggerInfo{Msg: fmt.Sprintf("Failed to publish message to queue %s: %s", maskedQueue, err)})
				return err
			}
			// Setting batch to nil after it was sent
			batch = nil
		}
	}

	// Send messages when batch doesn't reach its limit
	if len(batch) > 0 {
		_, err = c.producer.SendMessageBatch(&sqs.SendMessageBatchInput{
			QueueUrl: aws.String(queueUrl),
			Entries:  batch,
		})
		if err != nil {
			c.app.Logger().Log().Do(err, &domain.LoggerInfo{Msg: fmt.Sprintf("Failed to publish message to queue %s: %s", maskedQueue, err)})
			return err
		}
	}

	return nil
}

// Consume consumes from the sqs
func (c *Connection) Consume(maskedQueue string, consumer domain.ISQSConsumer) {
	queueUrl, err := url.JoinPath(c.config.Credentials.Api, c.config.Credentials.IdAccount, maskedQueue)
	if err != nil {
		c.app.Logger().Log().Do(err, &domain.LoggerInfo{Msg: fmt.Sprintf("Failed to consume from queue %s: %s", maskedQueue, err)})
		return
	}

	for {
		response, err := c.consumer.ReceiveMessage(
			&sqs.ReceiveMessageInput{
				QueueUrl:              aws.String(queueUrl),
				AttributeNames:        consumer.GetAttributeNames(),            // standard attributes of the messages to retrieve when receiving messages
				MessageAttributeNames: consumer.GetMessageAttributeNames(),     // user-defined key-value pairs attached to individual messages
				MaxNumberOfMessages:   aws.Int64(c.config.MaxNumberOfMessages), // maximum number of messages to retrieve from the queue in a single API call.
				VisibilityTimeout:     aws.Int64(c.config.VisibilityTimeout),   // duration for which a message is considered "invisible" after being received by a consumer
				WaitTimeSeconds:       aws.Int64(c.config.WaitTimeSeconds),     // wait time for a message to become available in the queue if no messages are currently present
			},
		)

		if err != nil {
			log.Fatalf("Failed to consume messages at queue %s: %s", maskedQueue, err)
		}

		// if the queue is empty, continue
		if len(response.Messages) == 0 {
			continue
		}

		c.handleMessages(response.Messages, maskedQueue, consumer.GetHandlers())
	}
}

// handleMessages handles the messages
func (c *Connection) handleMessages(messages []*sqs.Message, maskedQueue string, handlers map[string]func(msg *sqs.Message) bool) {
	queueUrl, err := url.JoinPath(c.config.Credentials.Api, c.config.Credentials.IdAccount, maskedQueue)
	if err != nil {
		c.app.Logger().Log().Do(err, &domain.LoggerInfo{Msg: fmt.Sprintf("Failed to handle message from queue %s: %s", maskedQueue, err)})
		return
	}

	wg := sync.WaitGroup{}
	batch := make([]*sqs.DeleteMessageBatchRequestEntry, 0)

	for i := range messages {
		wg.Add(1)

		// Start a goroutine to handle each message received
		f := func(msg *sqs.Message, handleFunc func(*sqs.Message) bool, wg *sync.WaitGroup, batch *[]*sqs.DeleteMessageBatchRequestEntry) {
			defer wg.Done()

			if handleFunc(msg) {
				*batch = append(*batch, &sqs.DeleteMessageBatchRequestEntry{
					Id:            msg.MessageId,
					ReceiptHandle: msg.ReceiptHandle,
				})
			}
		}

		go f(messages[i], c.getHandlerFunc(messages[i], maskedQueue, handlers), &wg, &batch)
	}

	// Wait for all the consumed messages
	wg.Wait()

	// Delete the messages that were successfully handled
	if len(batch) > 0 {
		_, err = c.consumer.DeleteMessageBatch(&sqs.DeleteMessageBatchInput{
			QueueUrl: aws.String(queueUrl),
			Entries:  batch,
		})
		if err != nil {
			c.app.Logger().Log().Do(err)
		}
	}
}

// getHandlerFunc gets the handler function
func (c *Connection) getHandlerFunc(message *sqs.Message, maskedQueue string, handlers map[string]func(message *sqs.Message) bool) func(*sqs.Message) bool {
	for attribute, handlerFunc := range handlers {
		key, ok := message.MessageAttributes["handler"]
		if ok && *key.StringValue == attribute {
			return func(*sqs.Message) bool {
				defer c.recover(message, maskedQueue)
				return handlerFunc(message)
			}
		}
	}

	return func(*sqs.Message) bool {
		c.app.Logger().Log().Do(fmt.Errorf(fmt.Sprintf("failed to handle message from queue %s: handler is nil", maskedQueue)))
		return false
	}
}

// recover recovers from a panic during message consumption
func (c *Connection) recover(message *sqs.Message, maskedQueue string) {
	if r := recover(); r != nil {
		data, _ := json.Marshal(r)
		errorMessage := string(data)

		c.app.Logger().Log().Do(
			errors.ErrorInSQSConsumer().Formats(maskedQueue, message.MessageAttributes, *message.Body),
			&domain.LoggerInfo{
				SubType: domain.SQSSubType,
				Msg:     errorMessage,
				Response: domain.Response{
					Body:       *message.Body,
					StatusCode: http.StatusBadRequest,
				},
			},
		)
	}
}
