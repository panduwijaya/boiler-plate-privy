// Package handler
package handler

import (
	"context"

	"cake-store/cake-store/internal/appctx"
	"cake-store/cake-store/internal/consts"
	uContract "cake-store/cake-store/internal/ucase/contract"
	"cake-store/cake-store/pkg/awssqs"
)

// SQSConsumerHandler sqs consumer message processor handler
func SQSConsumerHandler(msgHandler uContract.MessageProcessor) awssqs.MessageProcessorFunc {
	return func(decoder *awssqs.MessageDecoder) error {
		return msgHandler.Serve(context.Background(), &appctx.ConsumerData{
			Body:        []byte(*decoder.Body),
			Key:         []byte(*decoder.MessageId),
			ServiceType: consts.ServiceTypeConsumer,
		})
	}
}
