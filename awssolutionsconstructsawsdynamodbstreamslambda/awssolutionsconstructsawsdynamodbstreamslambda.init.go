package awssolutionsconstructsawsdynamodbstreamslambda

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-dynamodbstreams-lambda.DynamoDBStreamsToLambda",
		reflect.TypeOf((*DynamoDBStreamsToLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dynamoTable", GoGetter: "DynamoTable"},
			_jsii_.MemberProperty{JsiiProperty: "dynamoTableInterface", GoGetter: "DynamoTableInterface"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DynamoDBStreamsToLambda{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-dynamodbstreams-lambda.DynamoDBStreamsToLambdaProps",
		reflect.TypeOf((*DynamoDBStreamsToLambdaProps)(nil)).Elem(),
	)
}
