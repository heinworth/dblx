/*
	Contains functions to interact with AWS lambda functions, as well as 
	providing a mocking framework
*/

package lambda

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"

	_ "github.com/aws/aws-lambda-go/lambda"

	"os"
)

type AWS interface{
	invoke(input *lambda.InvokeInput) (*lambda.InvokeOutput, error)
}


type MockAWS struct{}
type AWSImplementation struct {}
var Client AWS

func (MockAWS) invoke(input *lambda.InvokeInput) (*lambda.InvokeOutput, error){
	return nil, nil
}

func (AWSImplementation) invoke(input *lambda.InvokeInput) (*lambda.InvokeOutput, error){
	
	if input == nil {
		input = &lambda.InvokeInput{FunctionName: aws.String("my-function"), Payload: nil}
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	
	client := lambda.New(sess, &aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	return client.Invoke(input)

}
