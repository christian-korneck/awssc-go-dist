package awssolutionsconstructsawskinesisfirehoses3

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-kinesisfirehose-s3.KinesisFirehoseToS3",
		reflect.TypeOf((*KinesisFirehoseToS3)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "kinesisFirehose", GoGetter: "KinesisFirehose"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisFirehoseLogGroup", GoGetter: "KinesisFirehoseLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "kinesisFirehoseRole", GoGetter: "KinesisFirehoseRole"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "s3Bucket", GoGetter: "S3Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketInterface", GoGetter: "S3BucketInterface"},
			_jsii_.MemberProperty{JsiiProperty: "s3LoggingBucket", GoGetter: "S3LoggingBucket"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisFirehoseToS3{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-kinesisfirehose-s3.KinesisFirehoseToS3Props",
		reflect.TypeOf((*KinesisFirehoseToS3Props)(nil)).Elem(),
	)
}
