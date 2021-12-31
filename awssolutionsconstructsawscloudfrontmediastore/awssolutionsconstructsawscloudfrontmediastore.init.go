package awssolutionsconstructsawscloudfrontmediastore

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-solutions-constructs/aws-cloudfront-mediastore.CloudFrontToMediaStore",
		reflect.TypeOf((*CloudFrontToMediaStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontFunction", GoGetter: "CloudFrontFunction"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontLoggingBucket", GoGetter: "CloudFrontLoggingBucket"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontOriginAccessIdentity", GoGetter: "CloudFrontOriginAccessIdentity"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontOriginRequestPolicy", GoGetter: "CloudFrontOriginRequestPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cloudFrontWebDistribution", GoGetter: "CloudFrontWebDistribution"},
			_jsii_.MemberProperty{JsiiProperty: "mediaStoreContainer", GoGetter: "MediaStoreContainer"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CloudFrontToMediaStore{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-solutions-constructs/aws-cloudfront-mediastore.CloudFrontToMediaStoreProps",
		reflect.TypeOf((*CloudFrontToMediaStoreProps)(nil)).Elem(),
	)
}
