package awssolutionsconstructsawsapigatewaydynamodb

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-apigateway-dynamodb.ApiGatewayToDynamoDB",
		reflect.TypeOf((*ApiGatewayToDynamoDB)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiGateway", GoGetter: "ApiGateway"},
			_jsii_.MemberProperty{JsiiProperty: "apiGatewayCloudWatchRole", GoGetter: "ApiGatewayCloudWatchRole"},
			_jsii_.MemberProperty{JsiiProperty: "apiGatewayLogGroup", GoGetter: "ApiGatewayLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiGatewayRole", GoGetter: "ApiGatewayRole"},
			_jsii_.MemberProperty{JsiiProperty: "dynamoTable", GoGetter: "DynamoTable"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApiGatewayToDynamoDB{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-apigateway-dynamodb.ApiGatewayToDynamoDBProps",
		reflect.TypeOf((*ApiGatewayToDynamoDBProps)(nil)).Elem(),
	)
}
