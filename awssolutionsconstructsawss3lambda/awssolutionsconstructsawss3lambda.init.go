package awssolutionsconstructsawss3lambda

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-s3-lambda.S3ToLambda",
		reflect.TypeOf((*S3ToLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "s3Bucket", GoGetter: "S3Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketInterface", GoGetter: "S3BucketInterface"},
			_jsii_.MemberProperty{JsiiProperty: "s3LoggingBucket", GoGetter: "S3LoggingBucket"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_S3ToLambda{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-s3-lambda.S3ToLambdaProps",
		reflect.TypeOf((*S3ToLambdaProps)(nil)).Elem(),
	)
}
