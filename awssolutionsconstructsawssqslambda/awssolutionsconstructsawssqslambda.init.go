package awssolutionsconstructsawssqslambda

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-sqs-lambda.SqsToLambda",
		reflect.TypeOf((*SqsToLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deadLetterQueue", GoGetter: "DeadLetterQueue"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "sqsQueue", GoGetter: "SqsQueue"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SqsToLambda{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-sqs-lambda.SqsToLambdaProps",
		reflect.TypeOf((*SqsToLambdaProps)(nil)).Elem(),
	)
}
