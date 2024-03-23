package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
)

func NewSQS() *sqs.SQS {
	config := &aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_SQS_ACCESSKEYID"),
			os.Getenv("AWS_SQS_SECRETKEY"),
			"",
		),
	}
	sess, err := session.NewSession(
		config,
	)
	if err != nil {
		panic(err)
	}

	sqsService := sqs.New(sess)

	return sqsService
}
