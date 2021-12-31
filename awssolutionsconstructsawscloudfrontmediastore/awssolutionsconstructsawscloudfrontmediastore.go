// CDK Constructs for Amazon CloudFront to AWS Elemental MediaStore integration.
package awssolutionsconstructsawscloudfrontmediastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscloudfrontmediastore/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediastore"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/christian-korneck/awssc-go-dist/awssolutionsconstructsawscloudfrontmediastore/v2/internal"
)

type CloudFrontToMediaStore interface {
	constructs.Construct
	CloudFrontFunction() awscloudfront.Function
	CloudFrontLoggingBucket() awss3.Bucket
	CloudFrontOriginAccessIdentity() awscloudfront.OriginAccessIdentity
	CloudFrontOriginRequestPolicy() awscloudfront.OriginRequestPolicy
	CloudFrontWebDistribution() awscloudfront.Distribution
	MediaStoreContainer() awsmediastore.CfnContainer
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for CloudFrontToMediaStore
type jsiiProxy_CloudFrontToMediaStore struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CloudFrontToMediaStore) CloudFrontFunction() awscloudfront.Function {
	var returns awscloudfront.Function
	_jsii_.Get(
		j,
		"cloudFrontFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToMediaStore) CloudFrontLoggingBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"cloudFrontLoggingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToMediaStore) CloudFrontOriginAccessIdentity() awscloudfront.OriginAccessIdentity {
	var returns awscloudfront.OriginAccessIdentity
	_jsii_.Get(
		j,
		"cloudFrontOriginAccessIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToMediaStore) CloudFrontOriginRequestPolicy() awscloudfront.OriginRequestPolicy {
	var returns awscloudfront.OriginRequestPolicy
	_jsii_.Get(
		j,
		"cloudFrontOriginRequestPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToMediaStore) CloudFrontWebDistribution() awscloudfront.Distribution {
	var returns awscloudfront.Distribution
	_jsii_.Get(
		j,
		"cloudFrontWebDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToMediaStore) MediaStoreContainer() awsmediastore.CfnContainer {
	var returns awsmediastore.CfnContainer
	_jsii_.Get(
		j,
		"mediaStoreContainer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontToMediaStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewCloudFrontToMediaStore(scope constructs.Construct, id *string, props *CloudFrontToMediaStoreProps) CloudFrontToMediaStore {
	_init_.Initialize()

	j := jsiiProxy_CloudFrontToMediaStore{}

	_jsii_.Create(
		"@aws-solutions-constructs/aws-cloudfront-mediastore.CloudFrontToMediaStore",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCloudFrontToMediaStore_Override(c CloudFrontToMediaStore, scope constructs.Construct, id *string, props *CloudFrontToMediaStoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-solutions-constructs/aws-cloudfront-mediastore.CloudFrontToMediaStore",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CloudFrontToMediaStore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-solutions-constructs/aws-cloudfront-mediastore.CloudFrontToMediaStore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_CloudFrontToMediaStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudFrontToMediaStoreProps struct {
	// Optional user provided props to override the default props for the CloudFront.
	CloudFrontDistributionProps interface{} `json:"cloudFrontDistributionProps"`
	// Optional user provided props to override the default props for the CloudFront Logging Bucket.
	CloudFrontLoggingBucketProps *awss3.BucketProps `json:"cloudFrontLoggingBucketProps"`
	// Existing instance of mediastore.CfnContainer obejct.
	ExistingMediaStoreContainerObj awsmediastore.CfnContainer `json:"existingMediaStoreContainerObj"`
	// Optional user provided props to turn on/off the automatic injection of best practice HTTP security headers in all responses from cloudfront.
	InsertHttpSecurityHeaders *bool `json:"insertHttpSecurityHeaders"`
	// Optional user provided props to override the default props for the MediaStore.
	MediaStoreContainerProps *awsmediastore.CfnContainerProps `json:"mediaStoreContainerProps"`
}

