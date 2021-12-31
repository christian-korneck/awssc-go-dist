package awssolutionsconstructsawslambdadynamodb

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-lambda-dynamodb.LambdaToDynamoDB",
		reflect.TypeOf((*LambdaToDynamoDB)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dynamoTable", GoGetter: "DynamoTable"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaToDynamoDB{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-lambda-dynamodb.LambdaToDynamoDBProps",
		reflect.TypeOf((*LambdaToDynamoDBProps)(nil)).Elem(),
	)
}
