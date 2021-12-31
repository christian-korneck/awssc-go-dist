package awssolutionsconstructsawslambdasqslambda

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-lambda-sqs-lambda.LambdaToSqsToLambda",
		reflect.TypeOf((*LambdaToSqsToLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "consumerLambdaFunction", GoGetter: "ConsumerLambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "deadLetterQueue", GoGetter: "DeadLetterQueue"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "producerLambdaFunction", GoGetter: "ProducerLambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "sqsQueue", GoGetter: "SqsQueue"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaToSqsToLambda{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-lambda-sqs-lambda.LambdaToSqsToLambdaProps",
		reflect.TypeOf((*LambdaToSqsToLambdaProps)(nil)).Elem(),
	)
}
