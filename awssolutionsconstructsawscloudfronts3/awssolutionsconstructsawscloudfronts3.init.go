package awssolutionsconstructsawscloudfronts3

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-cloudfront-s3.CloudFrontToS3",
		reflect.TypeOf((*CloudFrontToS3)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontFunction", GoGetter: "CloudFrontFunction"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontLoggingBucket", GoGetter: "CloudFrontLoggingBucket"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontWebDistribution", GoGetter: "CloudFrontWebDistribution"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "s3Bucket", GoGetter: "S3Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketInterface", GoGetter: "S3BucketInterface"},
			_jsii_.MemberProperty{JsiiProperty: "s3LoggingBucket", GoGetter: "S3LoggingBucket"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudFrontToS3{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-cloudfront-s3.CloudFrontToS3Props",
		reflect.TypeOf((*CloudFrontToS3Props)(nil)).Elem(),
	)
}
