package awssolutionsconstructsawss3stepfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-s3-stepfunctions.S3ToStepfunctions",
		reflect.TypeOf((*S3ToStepfunctions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchAlarms", GoGetter: "CloudwatchAlarms"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "s3Bucket", GoGetter: "S3Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketInterface", GoGetter: "S3BucketInterface"},
			_jsii_.MemberProperty{JsiiProperty: "s3LoggingBucket", GoGetter: "S3LoggingBucket"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachine", GoGetter: "StateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachineLogGroup", GoGetter: "StateMachineLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_S3ToStepfunctions{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-s3-stepfunctions.S3ToStepfunctionsProps",
		reflect.TypeOf((*S3ToStepfunctionsProps)(nil)).Elem(),
	)
}
