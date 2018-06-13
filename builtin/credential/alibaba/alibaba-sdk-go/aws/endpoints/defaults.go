package endpoints

import (
	"regexp"
)

// Partition identifiers
const (
	AcsPartitionID = "acs" // ACS Standard partition.
)

// AWS Standard partition's regions.
const (
	ChinaNorth1RegionID    = "cn-qingdao"     // Qingdao, China.
	ChinaNorth2RegionID    = "cn-beijing"     // Beijing, China.
	ChinaNorth3RegionID    = "cn-zhangjiakou" // Zhangjiakou, China.
	ChinaNorth5RegionID    = "cn-huhehaote"   // Hohhot, China.
	ChinaEast1RegionID     = "cn-hangzhou"    // Hangzhou, China.
	ChinaEast2RegionID     = "cn-shanghai"    // Shanghai, China.
	ChinaSouth1RegionID    = "cn-shenzhen"    // Shenzhen, China.
	HongKongRegionID       = "cn-hongkong"    // Hong Kong, China.
	AsiaPacificSE1RegionID = "ap-southeast-1" // Singapore.
	AsiaPacificSE2RegionID = "ap-southeast-2" // Sydney, Australia.
	AsiaPacificSE3RegionID = "ap-southeast-3" // Kuala Lumpur, Malaysia.
	AsiaPacificSE5RegionID = "ap-southeast-5" // Jakarta, Indonesia.
	AsiaPacificNE1RegionID = "ap-northeast-1" // Tokyo, Japan.
	AsiaPacificSouRegionID = "ap-south-1"     // Mumbai
)

// AWS China partition's regions.
const (
	CnNorth1RegionID     = "cn-north-1"     // China (Beijing).
	CnNorthwest1RegionID = "cn-northwest-1" // China (Ningxia).
)

// AWS GovCloud (US) partition's regions.
const (
	UsGovWest1RegionID = "us-gov-west-1" // AWS GovCloud (US).
)

// Service identifiers
const (
	AcmServiceID                          = "acm"                          // Acm.
	ApiPricingServiceID                   = "api.pricing"                  // ApiPricing.
	ApigatewayServiceID                   = "apigateway"                   // Apigateway.
	ApplicationAutoscalingServiceID       = "application-autoscaling"      // ApplicationAutoscaling.
	Appstream2ServiceID                   = "appstream2"                   // Appstream2.
	AthenaServiceID                       = "athena"                       // Athena.
	AutoscalingServiceID                  = "autoscaling"                  // Autoscaling.
	AutoscalingPlansServiceID             = "autoscaling-plans"            // AutoscalingPlans.
	BatchServiceID                        = "batch"                        // Batch.
	BudgetsServiceID                      = "budgets"                      // Budgets.
	ClouddirectoryServiceID               = "clouddirectory"               // Clouddirectory.
	CloudformationServiceID               = "cloudformation"               // Cloudformation.
	CloudfrontServiceID                   = "cloudfront"                   // Cloudfront.
	CloudhsmServiceID                     = "cloudhsm"                     // Cloudhsm.
	Cloudhsmv2ServiceID                   = "cloudhsmv2"                   // Cloudhsmv2.
	CloudsearchServiceID                  = "cloudsearch"                  // Cloudsearch.
	CloudtrailServiceID                   = "cloudtrail"                   // Cloudtrail.
	CodebuildServiceID                    = "codebuild"                    // Codebuild.
	CodecommitServiceID                   = "codecommit"                   // Codecommit.
	CodedeployServiceID                   = "codedeploy"                   // Codedeploy.
	CodepipelineServiceID                 = "codepipeline"                 // Codepipeline.
	CodestarServiceID                     = "codestar"                     // Codestar.
	CognitoIdentityServiceID              = "cognito-identity"             // CognitoIdentity.
	CognitoIdpServiceID                   = "cognito-idp"                  // CognitoIdp.
	CognitoSyncServiceID                  = "cognito-sync"                 // CognitoSync.
	ConfigServiceID                       = "config"                       // Config.
	CurServiceID                          = "cur"                          // Cur.
	DatapipelineServiceID                 = "datapipeline"                 // Datapipeline.
	DaxServiceID                          = "dax"                          // Dax.
	DevicefarmServiceID                   = "devicefarm"                   // Devicefarm.
	DirectconnectServiceID                = "directconnect"                // Directconnect.
	DiscoveryServiceID                    = "discovery"                    // Discovery.
	DmsServiceID                          = "dms"                          // Dms.
	DsServiceID                           = "ds"                           // Ds.
	DynamodbServiceID                     = "dynamodb"                     // Dynamodb.
	Ec2ServiceID                          = "ec2"                          // Ec2.
	Ec2metadataServiceID                  = "ec2metadata"                  // Ec2metadata.
	EcrServiceID                          = "ecr"                          // Ecr.
	EcsServiceID                          = "ecs"                          // Ecs.
	ElasticacheServiceID                  = "elasticache"                  // Elasticache.
	ElasticbeanstalkServiceID             = "elasticbeanstalk"             // Elasticbeanstalk.
	ElasticfilesystemServiceID            = "elasticfilesystem"            // Elasticfilesystem.
	ElasticloadbalancingServiceID         = "elasticloadbalancing"         // Elasticloadbalancing.
	ElasticmapreduceServiceID             = "elasticmapreduce"             // Elasticmapreduce.
	ElastictranscoderServiceID            = "elastictranscoder"            // Elastictranscoder.
	EmailServiceID                        = "email"                        // Email.
	EntitlementMarketplaceServiceID       = "entitlement.marketplace"      // EntitlementMarketplace.
	EsServiceID                           = "es"                           // Es.
	EventsServiceID                       = "events"                       // Events.
	FirehoseServiceID                     = "firehose"                     // Firehose.
	GameliftServiceID                     = "gamelift"                     // Gamelift.
	GlacierServiceID                      = "glacier"                      // Glacier.
	GlueServiceID                         = "glue"                         // Glue.
	GreengrassServiceID                   = "greengrass"                   // Greengrass.
	HealthServiceID                       = "health"                       // Health.
	IamServiceID                          = "iam"                          // Iam.
	ImportexportServiceID                 = "importexport"                 // Importexport.
	InspectorServiceID                    = "inspector"                    // Inspector.
	IotServiceID                          = "iot"                          // Iot.
	KinesisServiceID                      = "kinesis"                      // Kinesis.
	KinesisanalyticsServiceID             = "kinesisanalytics"             // Kinesisanalytics.
	KmsServiceID                          = "kms"                          // Kms.
	LambdaServiceID                       = "lambda"                       // Lambda.
	LightsailServiceID                    = "lightsail"                    // Lightsail.
	LogsServiceID                         = "logs"                         // Logs.
	MachinelearningServiceID              = "machinelearning"              // Machinelearning.
	MarketplacecommerceanalyticsServiceID = "marketplacecommerceanalytics" // Marketplacecommerceanalytics.
	MeteringMarketplaceServiceID          = "metering.marketplace"         // MeteringMarketplace.
	MghServiceID                          = "mgh"                          // Mgh.
	MobileanalyticsServiceID              = "mobileanalytics"              // Mobileanalytics.
	ModelsLexServiceID                    = "models.lex"                   // ModelsLex.
	MonitoringServiceID                   = "monitoring"                   // Monitoring.
	MturkRequesterServiceID               = "mturk-requester"              // MturkRequester.
	OpsworksServiceID                     = "opsworks"                     // Opsworks.
	OpsworksCmServiceID                   = "opsworks-cm"                  // OpsworksCm.
	OrganizationsServiceID                = "organizations"                // Organizations.
	PinpointServiceID                     = "pinpoint"                     // Pinpoint.
	PollyServiceID                        = "polly"                        // Polly.
	RdsServiceID                          = "rds"                          // Rds.
	RedshiftServiceID                     = "redshift"                     // Redshift.
	RekognitionServiceID                  = "rekognition"                  // Rekognition.
	Route53ServiceID                      = "route53"                      // Route53.
	Route53domainsServiceID               = "route53domains"               // Route53domains.
	RuntimeLexServiceID                   = "runtime.lex"                  // RuntimeLex.
	S3ServiceID                           = "s3"                           // S3.
	SdbServiceID                          = "sdb"                          // Sdb.
	ServicecatalogServiceID               = "servicecatalog"               // Servicecatalog.
	ShieldServiceID                       = "shield"                       // Shield.
	SmsServiceID                          = "sms"                          // Sms.
	SnowballServiceID                     = "snowball"                     // Snowball.
	SnsServiceID                          = "sns"                          // Sns.
	SqsServiceID                          = "sqs"                          // Sqs.
	SsmServiceID                          = "ssm"                          // Ssm.
	StatesServiceID                       = "states"                       // States.
	StoragegatewayServiceID               = "storagegateway"               // Storagegateway.
	StreamsDynamodbServiceID              = "streams.dynamodb"             // StreamsDynamodb.
	StsServiceID                          = "sts"                          // Sts.
	SupportServiceID                      = "support"                      // Support.
	SwfServiceID                          = "swf"                          // Swf.
	TaggingServiceID                      = "tagging"                      // Tagging.
	WafServiceID                          = "waf"                          // Waf.
	WafRegionalServiceID                  = "waf-regional"                 // WafRegional.
	WorkdocsServiceID                     = "workdocs"                     // Workdocs.
	WorkspacesServiceID                   = "workspaces"                   // Workspaces.
	XrayServiceID                         = "xray"                         // Xray.
)

// DefaultResolver returns an Endpoint resolver that will be able
// to resolve endpoints for: AWS Standard, AWS China, and AWS GovCloud (US).
//
// Use DefaultPartitions() to get the list of the default partitions.
func DefaultResolver() Resolver {
	return defaultPartitions
}

// DefaultPartitions returns a list of the partitions the SDK is bundled
// with. The available partitions are: AWS Standard, AWS China, and AWS GovCloud (US).
//
//    partitions := endpoints.DefaultPartitions
//    for _, p := range partitions {
//        // ... inspect partitions
//    }
func DefaultPartitions() []Partition {
	return defaultPartitions.Partitions()
}

var defaultPartitions = partitions{
	awsPartition,
	awscnPartition,
	awsusgovPartition,
}

// AwsPartition returns the Resolver for AWS Standard.
func AwsPartition() Partition {
	return awsPartition.Partition()
}

var awsPartition = partition{
	ID:        "aws",
	Name:      "AWS Standard",
	DNSSuffix: "amazonaws.com",
	RegionRegex: regionRegex{
		Regexp: func() *regexp.Regexp {
			reg, _ := regexp.Compile("^(us|eu|ap|sa|ca)\\-\\w+\\-\\d+$")
			return reg
		}(),
	},
	Defaults: endpoint{
		Hostname:          "{service}.{region}.{dnsSuffix}",
		Protocols:         []string{"https"},
		SignatureVersions: []string{"v4"},
	},
	Regions: regions{
		"ap-northeast-1": region{
			Description: "Asia Pacific (Tokyo)",
		},
		"ap-northeast-2": region{
			Description: "Asia Pacific (Seoul)",
		},
		"ap-south-1": region{
			Description: "Asia Pacific (Mumbai)",
		},
		"ap-southeast-1": region{
			Description: "Asia Pacific (Singapore)",
		},
		"ap-southeast-2": region{
			Description: "Asia Pacific (Sydney)",
		},
		"ca-central-1": region{
			Description: "Canada (Central)",
		},
		"eu-central-1": region{
			Description: "EU (Frankfurt)",
		},
		"eu-west-1": region{
			Description: "EU (Ireland)",
		},
		"eu-west-2": region{
			Description: "EU (London)",
		},
		"eu-west-3": region{
			Description: "EU (Paris)",
		},
		"sa-east-1": region{
			Description: "South America (Sao Paulo)",
		},
		"us-east-1": region{
			Description: "US East (N. Virginia)",
		},
		"us-east-2": region{
			Description: "US East (Ohio)",
		},
		"us-west-1": region{
			Description: "US West (N. California)",
		},
		"us-west-2": region{
			Description: "US West (Oregon)",
		},
	},
	Services: services{
		"acm": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"api.pricing": service{
			Defaults: endpoint{
				CredentialScope: credentialScope{
					Service: "pricing",
				},
			},
			Endpoints: endpoints{
				"ap-south-1": endpoint{},
				"us-east-1":  endpoint{},
			},
		},
		"apigateway": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"application-autoscaling": service{
			Defaults: endpoint{
				Hostname:  "autoscaling.{region}.amazonaws.com",
				Protocols: []string{"http", "https"},
				CredentialScope: credentialScope{
					Service: "application-autoscaling",
				},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"appstream2": service{
			Defaults: endpoint{
				Protocols: []string{"https"},
				CredentialScope: credentialScope{
					Service: "appstream",
				},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"athena": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"autoscaling": service{
			Defaults: endpoint{
				Protocols: []string{"http", "https"},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"autoscaling-plans": service{
			Defaults: endpoint{
				Hostname:  "autoscaling.{region}.amazonaws.com",
				Protocols: []string{"http", "https"},
				CredentialScope: credentialScope{
					Service: "autoscaling-plans",
				},
			},
			Endpoints: endpoints{
				"ap-southeast-1": endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"batch": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"budgets": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname: "budgets.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
			},
		},
		"clouddirectory": service{

			Endpoints: endpoints{
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"cloudformation": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"cloudfront": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname:  "cloudfront.amazonaws.com",
					Protocols: []string{"http", "https"},
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
			},
		},
		"cloudhsm": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"cloudhsmv2": service{
			Defaults: endpoint{
				CredentialScope: credentialScope{
					Service: "cloudhsm",
				},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"cloudsearch": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"cloudtrail": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"codebuild": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"codecommit": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"codedeploy": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"codepipeline": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"codestar": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"cognito-identity": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"cognito-idp": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"cognito-sync": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"config": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"cur": service{

			Endpoints: endpoints{
				"us-east-1": endpoint{},
			},
		},
		"datapipeline": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"dax": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-south-1":     endpoint{},
				"eu-west-1":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"devicefarm": service{

			Endpoints: endpoints{
				"us-west-2": endpoint{},
			},
		},
		"directconnect": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"discovery": service{

			Endpoints: endpoints{
				"us-west-2": endpoint{},
			},
		},
		"dms": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"ds": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"dynamodb": service{
			Defaults: endpoint{
				Protocols: []string{"http", "https"},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"local": endpoint{
					Hostname:  "localhost:8000",
					Protocols: []string{"http"},
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"sa-east-1": endpoint{},
				"us-east-1": endpoint{},
				"us-east-2": endpoint{},
				"us-west-1": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"ec2": service{
			Defaults: endpoint{
				Protocols: []string{"http", "https"},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"ec2metadata": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname:  "169.254.169.254/latest",
					Protocols: []string{"http"},
				},
			},
		},
		"ecr": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"ecs": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"elasticache": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"elasticbeanstalk": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"elasticfilesystem": service{

			Endpoints: endpoints{
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"elasticloadbalancing": service{
			Defaults: endpoint{
				Protocols: []string{"https"},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"elasticmapreduce": service{
			Defaults: endpoint{
				SSLCommonName: "{region}.{service}.{dnsSuffix}",
				Protocols:     []string{"http", "https"},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1": endpoint{
					SSLCommonName: "{service}.{region}.{dnsSuffix}",
				},
				"eu-west-1": endpoint{},
				"eu-west-2": endpoint{},
				"eu-west-3": endpoint{},
				"sa-east-1": endpoint{},
				"us-east-1": endpoint{
					SSLCommonName: "{service}.{region}.{dnsSuffix}",
				},
				"us-east-2": endpoint{},
				"us-west-1": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"elastictranscoder": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"email": service{

			Endpoints: endpoints{
				"eu-west-1": endpoint{},
				"us-east-1": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"entitlement.marketplace": service{
			Defaults: endpoint{
				CredentialScope: credentialScope{
					Service: "aws-marketplace",
				},
			},
			Endpoints: endpoints{
				"us-east-1": endpoint{},
			},
		},
		"es": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"events": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"firehose": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"gamelift": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"glacier": service{
			Defaults: endpoint{
				Protocols: []string{"http", "https"},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"glue": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"greengrass": service{
			IsRegionalized: boxedTrue,
			Defaults: endpoint{
				Protocols: []string{"https"},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"us-east-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"health": service{

			Endpoints: endpoints{
				"us-east-1": endpoint{},
			},
		},
		"iam": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname: "iam.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
			},
		},
		"importexport": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname:          "importexport.amazonaws.com",
					SignatureVersions: []string{"v2", "v4"},
					CredentialScope: credentialScope{
						Region:  "us-east-1",
						Service: "IngestionService",
					},
				},
			},
		},
		"inspector": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"iot": service{
			Defaults: endpoint{
				CredentialScope: credentialScope{
					Service: "execute-api",
				},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"kinesis": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"kinesisanalytics": service{

			Endpoints: endpoints{
				"eu-west-1": endpoint{},
				"us-east-1": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"kms": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"lambda": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"lightsail": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"logs": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"machinelearning": service{

			Endpoints: endpoints{
				"eu-west-1": endpoint{},
				"us-east-1": endpoint{},
			},
		},
		"marketplacecommerceanalytics": service{

			Endpoints: endpoints{
				"us-east-1": endpoint{},
			},
		},
		"metering.marketplace": service{
			Defaults: endpoint{
				CredentialScope: credentialScope{
					Service: "aws-marketplace",
				},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"mgh": service{

			Endpoints: endpoints{
				"us-west-2": endpoint{},
			},
		},
		"mobileanalytics": service{

			Endpoints: endpoints{
				"us-east-1": endpoint{},
			},
		},
		"models.lex": service{
			Defaults: endpoint{
				CredentialScope: credentialScope{
					Service: "lex",
				},
			},
			Endpoints: endpoints{
				"us-east-1": endpoint{},
			},
		},
		"monitoring": service{
			Defaults: endpoint{
				Protocols: []string{"http", "https"},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"mturk-requester": service{
			IsRegionalized: boxedFalse,

			Endpoints: endpoints{
				"sandbox": endpoint{
					Hostname: "mturk-requester-sandbox.us-east-1.amazonaws.com",
				},
				"us-east-1": endpoint{},
			},
		},
		"opsworks": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"opsworks-cm": service{

			Endpoints: endpoints{
				"eu-west-1": endpoint{},
				"us-east-1": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"organizations": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname: "organizations.us-east-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
			},
		},
		"pinpoint": service{
			Defaults: endpoint{
				CredentialScope: credentialScope{
					Service: "mobiletargeting",
				},
			},
			Endpoints: endpoints{
				"us-east-1": endpoint{},
			},
		},
		"polly": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"rds": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1": endpoint{
					SSLCommonName: "{service}.{dnsSuffix}",
				},
				"us-east-2": endpoint{},
				"us-west-1": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"redshift": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"rekognition": service{

			Endpoints: endpoints{
				"eu-west-1": endpoint{},
				"us-east-1": endpoint{},
				"us-east-2": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"route53": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname: "route53.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
			},
		},
		"route53domains": service{

			Endpoints: endpoints{
				"us-east-1": endpoint{},
			},
		},
		"runtime.lex": service{
			Defaults: endpoint{
				CredentialScope: credentialScope{
					Service: "lex",
				},
			},
			Endpoints: endpoints{
				"eu-west-1": endpoint{},
				"us-east-1": endpoint{},
			},
		},
		"s3": service{
			PartitionEndpoint: "us-east-1",
			IsRegionalized:    boxedTrue,
			Defaults: endpoint{
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},

				HasDualStack:      boxedTrue,
				DualStackHostname: "{service}.dualstack.{region}.{dnsSuffix}",
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{
					Hostname:          "s3.ap-northeast-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{
					Hostname:          "s3.ap-southeast-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"ap-southeast-2": endpoint{
					Hostname:          "s3.ap-southeast-2.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"ca-central-1": endpoint{},
				"eu-central-1": endpoint{},
				"eu-west-1": endpoint{
					Hostname:          "s3.eu-west-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"eu-west-2": endpoint{},
				"eu-west-3": endpoint{},
				"s3-external-1": endpoint{
					Hostname:          "s3-external-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"sa-east-1": endpoint{
					Hostname:          "s3.sa-east-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"us-east-1": endpoint{
					Hostname:          "s3.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"us-east-2": endpoint{},
				"us-west-1": endpoint{
					Hostname:          "s3.us-west-1.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
				"us-west-2": endpoint{
					Hostname:          "s3.us-west-2.amazonaws.com",
					SignatureVersions: []string{"s3", "s3v4"},
				},
			},
		},
		"sdb": service{
			Defaults: endpoint{
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v2"},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-west-1":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1": endpoint{
					Hostname: "sdb.amazonaws.com",
				},
				"us-west-1": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"servicecatalog": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"shield": service{
			IsRegionalized: boxedFalse,
			Defaults: endpoint{
				SSLCommonName: "Shield.us-east-1.amazonaws.com",
				Protocols:     []string{"https"},
			},
			Endpoints: endpoints{
				"us-east-1": endpoint{},
			},
		},
		"sms": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-3":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"snowball": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"sns": service{
			Defaults: endpoint{
				Protocols: []string{"http", "https"},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"sqs": service{
			Defaults: endpoint{
				SSLCommonName: "{region}.queue.{dnsSuffix}",
				Protocols:     []string{"http", "https"},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1": endpoint{
					SSLCommonName: "queue.{dnsSuffix}",
				},
				"us-east-2": endpoint{},
				"us-west-1": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"ssm": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"states": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"storagegateway": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"streams.dynamodb": service{
			Defaults: endpoint{
				Protocols: []string{"http", "https"},
				CredentialScope: credentialScope{
					Service: "dynamodb",
				},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"local": endpoint{
					Hostname:  "localhost:8000",
					Protocols: []string{"http"},
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"sa-east-1": endpoint{},
				"us-east-1": endpoint{},
				"us-east-2": endpoint{},
				"us-west-1": endpoint{},
				"us-west-2": endpoint{},
			},
		},
		"sts": service{
			PartitionEndpoint: "aws-global",
			Defaults: endpoint{
				Hostname: "sts.amazonaws.com",
				CredentialScope: credentialScope{
					Region: "us-east-1",
				},
			},
			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{
					Hostname: "sts.ap-northeast-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "ap-northeast-2",
					},
				},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"aws-global":     endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-1-fips": endpoint{
					Hostname: "sts-fips.us-east-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
				"us-east-2": endpoint{},
				"us-east-2-fips": endpoint{
					Hostname: "sts-fips.us-east-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-2",
					},
				},
				"us-west-1": endpoint{},
				"us-west-1-fips": endpoint{
					Hostname: "sts-fips.us-west-1.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-west-1",
					},
				},
				"us-west-2": endpoint{},
				"us-west-2-fips": endpoint{
					Hostname: "sts-fips.us-west-2.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-west-2",
					},
				},
			},
		},
		"support": service{

			Endpoints: endpoints{
				"us-east-1": endpoint{},
			},
		},
		"swf": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"eu-west-3":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"tagging": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"waf": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname: "waf.amazonaws.com",
					CredentialScope: credentialScope{
						Region: "us-east-1",
					},
				},
			},
		},
		"waf-regional": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"workdocs": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-west-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"workspaces": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"us-east-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
		"xray": service{

			Endpoints: endpoints{
				"ap-northeast-1": endpoint{},
				"ap-northeast-2": endpoint{},
				"ap-south-1":     endpoint{},
				"ap-southeast-1": endpoint{},
				"ap-southeast-2": endpoint{},
				"ca-central-1":   endpoint{},
				"eu-central-1":   endpoint{},
				"eu-west-1":      endpoint{},
				"eu-west-2":      endpoint{},
				"sa-east-1":      endpoint{},
				"us-east-1":      endpoint{},
				"us-east-2":      endpoint{},
				"us-west-1":      endpoint{},
				"us-west-2":      endpoint{},
			},
		},
	},
}
