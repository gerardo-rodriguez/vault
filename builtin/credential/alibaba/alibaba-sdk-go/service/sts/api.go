package sts

import (
	"time"

	"github.com/hashicorp/vault/builtin/credential/alibaba/alibaba-sdk-go/aws"
	"github.com/hashicorp/vault/builtin/credential/alibaba/alibaba-sdk-go/aws/awsutil"
	"github.com/hashicorp/vault/builtin/credential/alibaba/alibaba-sdk-go/aws/request"
)

const opAssumeRole = "AssumeRole"

// AssumeRoleRequest generates a "aws/request.Request" representing the
// client's request for the AssumeRole operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See AssumeRole for more information on using the AssumeRole
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the AssumeRoleRequest method.
//    req, resp := client.AssumeRoleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/AssumeRole
func (c *STS) AssumeRoleRequest(input *AssumeRoleInput) (req *request.Request, output *AssumeRoleOutput) {
	op := &request.Operation{
		Name:       opAssumeRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssumeRoleInput{}
	}

	output = &AssumeRoleOutput{}
	req = c.newRequest(op, input, output)
	return
}

// AssumeRole API operation for AWS Security Token Service.
//
// Returns a set of temporary security credentials (consisting of an access
// key ID, a secret access key, and a security token) that you can use to access
// AWS resources that you might not normally have access to. Typically, you
// use AssumeRole for cross-account access or federation. For a comparison of
// AssumeRole with the other APIs that produce temporary credentials, see Requesting
// Temporary Security Credentials (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS APIs (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide.
//
// Important: You cannot call AssumeRole by using AWS root account credentials;
// access is denied. You must use credentials for an IAM user or an IAM role
// to call AssumeRole.
//
// For cross-account access, imagine that you own multiple accounts and need
// to access resources in each account. You could create long-term credentials
// in each account to access those resources. However, managing all those credentials
// and remembering which one can access which account can be time consuming.
// Instead, you can create one set of long-term credentials in one account and
// then use temporary security credentials to access all the other accounts
// by assuming roles in those accounts. For more information about roles, see
// IAM Roles (Delegation and Federation) (http://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html)
// in the IAM User Guide.
//
// For federation, you can, for example, grant single sign-on access to the
// AWS Management Console. If you already have an identity and authentication
// system in your corporate network, you don't have to recreate user identities
// in AWS in order to grant those user identities access to AWS. Instead, after
// a user has been authenticated, you call AssumeRole (and specify the role
// with the appropriate permissions) to get temporary security credentials for
// that user. With those temporary security credentials, you construct a sign-in
// URL that users can use to access the console. For more information, see Common
// Scenarios for Temporary Credentials (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp.html#sts-introduction)
// in the IAM User Guide.
//
// The temporary security credentials are valid for the duration that you specified
// when calling AssumeRole, which can be from 900 seconds (15 minutes) to a
// maximum of 3600 seconds (1 hour). The default is 1 hour.
//
// The temporary security credentials created by AssumeRole can be used to make
// API calls to any AWS service with the following exception: you cannot call
// the STS service's GetFederationToken or GetSessionToken APIs.
//
// Optionally, you can pass an IAM access policy to this operation. If you choose
// not to pass a policy, the temporary security credentials that are returned
// by the operation have the permissions that are defined in the access policy
// of the role that is being assumed. If you pass a policy to this operation,
// the temporary security credentials that are returned by the operation have
// the permissions that are allowed by both the access policy of the role that
// is being assumed, and the policy that you pass. This gives you a way to further
// restrict the permissions for the resulting temporary security credentials.
// You cannot use the passed policy to grant permissions that are in excess
// of those allowed by the access policy of the role that is being assumed.
// For more information, see Permissions for AssumeRole, AssumeRoleWithSAML,
// and AssumeRoleWithWebIdentity (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_assumerole.html)
// in the IAM User Guide.
//
// To assume a role, your AWS account must be trusted by the role. The trust
// relationship is defined in the role's trust policy when the role is created.
// That trust policy states which accounts are allowed to delegate access to
// this account's role.
//
// The user who wants to access the role must also have permissions delegated
// from the role's administrator. If the user is in a different account than
// the role, then the user's administrator must attach a policy that allows
// the user to call AssumeRole on the ARN of the role in the other account.
// If the user is in the same account as the role, then you can either attach
// a policy to the user (identical to the previous different account user),
// or you can add the user as a principal directly in the role's trust policy
//
// Using MFA with AssumeRole
//
// You can optionally include multi-factor authentication (MFA) information
// when you call AssumeRole. This is useful for cross-account scenarios in which
// you want to make sure that the user who is assuming the role has been authenticated
// using an AWS MFA device. In that scenario, the trust policy of the role being
// assumed includes a condition that tests for MFA authentication; if the caller
// does not include valid MFA information, the request to assume the role is
// denied. The condition in a trust policy that tests for MFA authentication
// might look like the following example.
//
// "Condition": {"Bool": {"aws:MultiFactorAuthPresent": true}}
//
// For more information, see Configuring MFA-Protected API Access (http://docs.aws.amazon.com/IAM/latest/UserGuide/MFAProtectedAPI.html)
// in the IAM User Guide guide.
//
// To use MFA with AssumeRole, you pass values for the SerialNumber and TokenCode
// parameters. The SerialNumber value identifies the user's hardware or virtual
// MFA device. The TokenCode is the time-based one-time password (TOTP) that
// the MFA devices produces.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Security Token Service's
// API operation AssumeRole for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeMalformedPolicyDocumentException "MalformedPolicyDocument"
//   The request was rejected because the policy document was malformed. The error
//   message describes the specific error.
//
//   * ErrCodePackedPolicyTooLargeException "PackedPolicyTooLarge"
//   The request was rejected because the policy document was too large. The error
//   message describes how big the policy document is, in packed form, as a percentage
//   of what the API allows.
//
//   * ErrCodeRegionDisabledException "RegionDisabledException"
//   STS is not activated in the requested region for the account that is being
//   asked to generate credentials. The account administrator must use the IAM
//   console to activate STS in that region. For more information, see Activating
//   and Deactivating AWS STS in an AWS Region (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
//   in the IAM User Guide.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/AssumeRole
func (c *STS) AssumeRole(input *AssumeRoleInput) (*AssumeRoleOutput, error) {
	req, out := c.AssumeRoleRequest(input)
	return out, req.Send()
}

// AssumeRoleWithContext is the same as AssumeRole with the addition of
// the ability to pass a context and additional request options.
//
// See AssumeRole for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STS) AssumeRoleWithContext(ctx aws.Context, input *AssumeRoleInput, opts ...request.Option) (*AssumeRoleOutput, error) {
	req, out := c.AssumeRoleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssumeRoleWithSAML = "AssumeRoleWithSAML"

// AssumeRoleWithSAMLRequest generates a "aws/request.Request" representing the
// client's request for the AssumeRoleWithSAML operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See AssumeRoleWithSAML for more information on using the AssumeRoleWithSAML
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the AssumeRoleWithSAMLRequest method.
//    req, resp := client.AssumeRoleWithSAMLRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/AssumeRoleWithSAML
func (c *STS) AssumeRoleWithSAMLRequest(input *AssumeRoleWithSAMLInput) (req *request.Request, output *AssumeRoleWithSAMLOutput) {
	op := &request.Operation{
		Name:       opAssumeRoleWithSAML,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssumeRoleWithSAMLInput{}
	}

	output = &AssumeRoleWithSAMLOutput{}
	req = c.newRequest(op, input, output)
	return
}

// AssumeRoleWithSAML API operation for AWS Security Token Service.
//
// Returns a set of temporary security credentials for users who have been authenticated
// via a SAML authentication response. This operation provides a mechanism for
// tying an enterprise identity store or directory to role-based AWS access
// without user-specific credentials or configuration. For a comparison of AssumeRoleWithSAML
// with the other APIs that produce temporary credentials, see Requesting Temporary
// Security Credentials (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS APIs (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide.
//
// The temporary security credentials returned by this operation consist of
// an access key ID, a secret access key, and a security token. Applications
// can use these temporary security credentials to sign calls to AWS services.
//
// The temporary security credentials are valid for the duration that you specified
// when calling AssumeRole, or until the time specified in the SAML authentication
// response's SessionNotOnOrAfter value, whichever is shorter. The duration
// can be from 900 seconds (15 minutes) to a maximum of 3600 seconds (1 hour).
// The default is 1 hour.
//
// The temporary security credentials created by AssumeRoleWithSAML can be used
// to make API calls to any AWS service with the following exception: you cannot
// call the STS service's GetFederationToken or GetSessionToken APIs.
//
// Optionally, you can pass an IAM access policy to this operation. If you choose
// not to pass a policy, the temporary security credentials that are returned
// by the operation have the permissions that are defined in the access policy
// of the role that is being assumed. If you pass a policy to this operation,
// the temporary security credentials that are returned by the operation have
// the permissions that are allowed by the intersection of both the access policy
// of the role that is being assumed, and the policy that you pass. This means
// that both policies must grant the permission for the action to be allowed.
// This gives you a way to further restrict the permissions for the resulting
// temporary security credentials. You cannot use the passed policy to grant
// permissions that are in excess of those allowed by the access policy of the
// role that is being assumed. For more information, see Permissions for AssumeRole,
// AssumeRoleWithSAML, and AssumeRoleWithWebIdentity (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_assumerole.html)
// in the IAM User Guide.
//
// Before your application can call AssumeRoleWithSAML, you must configure your
// SAML identity provider (IdP) to issue the claims required by AWS. Additionally,
// you must use AWS Identity and Access Management (IAM) to create a SAML provider
// entity in your AWS account that represents your identity provider, and create
// an IAM role that specifies this SAML provider in its trust policy.
//
// Calling AssumeRoleWithSAML does not require the use of AWS security credentials.
// The identity of the caller is validated by using keys in the metadata document
// that is uploaded for the SAML provider entity for your identity provider.
//
// Calling AssumeRoleWithSAML can result in an entry in your AWS CloudTrail
// logs. The entry includes the value in the NameID element of the SAML assertion.
// We recommend that you use a NameIDType that is not associated with any personally
// identifiable information (PII). For example, you could instead use the Persistent
// Identifier (urn:oasis:names:tc:SAML:2.0:nameid-format:persistent).
//
// For more information, see the following resources:
//
//    * About SAML 2.0-based Federation (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_saml.html)
//    in the IAM User Guide.
//
//    * Creating SAML Identity Providers (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_saml.html)
//    in the IAM User Guide.
//
//    * Configuring a Relying Party and Claims (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_saml_relying-party.html)
//    in the IAM User Guide.
//
//    * Creating a Role for SAML 2.0 Federation (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-idp_saml.html)
//    in the IAM User Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Security Token Service's
// API operation AssumeRoleWithSAML for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeMalformedPolicyDocumentException "MalformedPolicyDocument"
//   The request was rejected because the policy document was malformed. The error
//   message describes the specific error.
//
//   * ErrCodePackedPolicyTooLargeException "PackedPolicyTooLarge"
//   The request was rejected because the policy document was too large. The error
//   message describes how big the policy document is, in packed form, as a percentage
//   of what the API allows.
//
//   * ErrCodeIDPRejectedClaimException "IDPRejectedClaim"
//   The identity provider (IdP) reported that authentication failed. This might
//   be because the claim is invalid.
//
//   If this error is returned for the AssumeRoleWithWebIdentity operation, it
//   can also mean that the claim has expired or has been explicitly revoked.
//
//   * ErrCodeInvalidIdentityTokenException "InvalidIdentityToken"
//   The web identity token that was passed could not be validated by AWS. Get
//   a new identity token from the identity provider and then retry the request.
//
//   * ErrCodeExpiredTokenException "ExpiredTokenException"
//   The web identity token that was passed is expired or is not valid. Get a
//   new identity token from the identity provider and then retry the request.
//
//   * ErrCodeRegionDisabledException "RegionDisabledException"
//   STS is not activated in the requested region for the account that is being
//   asked to generate credentials. The account administrator must use the IAM
//   console to activate STS in that region. For more information, see Activating
//   and Deactivating AWS STS in an AWS Region (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
//   in the IAM User Guide.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/AssumeRoleWithSAML
func (c *STS) AssumeRoleWithSAML(input *AssumeRoleWithSAMLInput) (*AssumeRoleWithSAMLOutput, error) {
	req, out := c.AssumeRoleWithSAMLRequest(input)
	return out, req.Send()
}

// AssumeRoleWithSAMLWithContext is the same as AssumeRoleWithSAML with the addition of
// the ability to pass a context and additional request options.
//
// See AssumeRoleWithSAML for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STS) AssumeRoleWithSAMLWithContext(ctx aws.Context, input *AssumeRoleWithSAMLInput, opts ...request.Option) (*AssumeRoleWithSAMLOutput, error) {
	req, out := c.AssumeRoleWithSAMLRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssumeRoleWithWebIdentity = "AssumeRoleWithWebIdentity"

// AssumeRoleWithWebIdentityRequest generates a "aws/request.Request" representing the
// client's request for the AssumeRoleWithWebIdentity operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See AssumeRoleWithWebIdentity for more information on using the AssumeRoleWithWebIdentity
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the AssumeRoleWithWebIdentityRequest method.
//    req, resp := client.AssumeRoleWithWebIdentityRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/AssumeRoleWithWebIdentity
func (c *STS) AssumeRoleWithWebIdentityRequest(input *AssumeRoleWithWebIdentityInput) (req *request.Request, output *AssumeRoleWithWebIdentityOutput) {
	op := &request.Operation{
		Name:       opAssumeRoleWithWebIdentity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssumeRoleWithWebIdentityInput{}
	}

	output = &AssumeRoleWithWebIdentityOutput{}
	req = c.newRequest(op, input, output)
	return
}

// AssumeRoleWithWebIdentity API operation for AWS Security Token Service.
//
// Returns a set of temporary security credentials for users who have been authenticated
// in a mobile or web application with a web identity provider, such as Amazon
// Cognito, Login with Amazon, Facebook, Google, or any OpenID Connect-compatible
// identity provider.
//
// For mobile applications, we recommend that you use Amazon Cognito. You can
// use Amazon Cognito with the AWS SDK for iOS (http://aws.amazon.com/sdkforios/)
// and the AWS SDK for Android (http://aws.amazon.com/sdkforandroid/) to uniquely
// identify a user and supply the user with a consistent identity throughout
// the lifetime of an application.
//
// To learn more about Amazon Cognito, see Amazon Cognito Overview (http://docs.aws.amazon.com/mobile/sdkforandroid/developerguide/cognito-auth.html#d0e840)
// in the AWS SDK for Android Developer Guide guide and Amazon Cognito Overview
// (http://docs.aws.amazon.com/mobile/sdkforios/developerguide/cognito-auth.html#d0e664)
// in the AWS SDK for iOS Developer Guide.
//
// Calling AssumeRoleWithWebIdentity does not require the use of AWS security
// credentials. Therefore, you can distribute an application (for example, on
// mobile devices) that requests temporary security credentials without including
// long-term AWS credentials in the application, and without deploying server-based
// proxy services that use long-term AWS credentials. Instead, the identity
// of the caller is validated by using a token from the web identity provider.
// For a comparison of AssumeRoleWithWebIdentity with the other APIs that produce
// temporary credentials, see Requesting Temporary Security Credentials (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS APIs (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide.
//
// The temporary security credentials returned by this API consist of an access
// key ID, a secret access key, and a security token. Applications can use these
// temporary security credentials to sign calls to AWS service APIs.
//
// The credentials are valid for the duration that you specified when calling
// AssumeRoleWithWebIdentity, which can be from 900 seconds (15 minutes) to
// a maximum of 3600 seconds (1 hour). The default is 1 hour.
//
// The temporary security credentials created by AssumeRoleWithWebIdentity can
// be used to make API calls to any AWS service with the following exception:
// you cannot call the STS service's GetFederationToken or GetSessionToken APIs.
//
// Optionally, you can pass an IAM access policy to this operation. If you choose
// not to pass a policy, the temporary security credentials that are returned
// by the operation have the permissions that are defined in the access policy
// of the role that is being assumed. If you pass a policy to this operation,
// the temporary security credentials that are returned by the operation have
// the permissions that are allowed by both the access policy of the role that
// is being assumed, and the policy that you pass. This gives you a way to further
// restrict the permissions for the resulting temporary security credentials.
// You cannot use the passed policy to grant permissions that are in excess
// of those allowed by the access policy of the role that is being assumed.
// For more information, see Permissions for AssumeRole, AssumeRoleWithSAML,
// and AssumeRoleWithWebIdentity (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_assumerole.html)
// in the IAM User Guide.
//
// Before your application can call AssumeRoleWithWebIdentity, you must have
// an identity token from a supported identity provider and create a role that
// the application can assume. The role that your application assumes must trust
// the identity provider that is associated with the identity token. In other
// words, the identity provider must be specified in the role's trust policy.
//
// Calling AssumeRoleWithWebIdentity can result in an entry in your AWS CloudTrail
// logs. The entry includes the Subject (http://openid.net/specs/openid-connect-core-1_0.html#Claims)
// of the provided Web Identity Token. We recommend that you avoid using any
// personally identifiable information (PII) in this field. For example, you
// could instead use a GUID or a pairwise identifier, as suggested in the OIDC
// specification (http://openid.net/specs/openid-connect-core-1_0.html#SubjectIDTypes).
//
// For more information about how to use web identity federation and the AssumeRoleWithWebIdentity
// API, see the following resources:
//
//    * Using Web Identity Federation APIs for Mobile Apps (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc_manual.html)
//    and Federation Through a Web-based Identity Provider (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_assumerolewithwebidentity).
//
//
//    *  Web Identity Federation Playground (https://web-identity-federation-playground.s3.amazonaws.com/index.html).
//    This interactive website lets you walk through the process of authenticating
//    via Login with Amazon, Facebook, or Google, getting temporary security
//    credentials, and then using those credentials to make a request to AWS.
//
//
//    * AWS SDK for iOS (http://aws.amazon.com/sdkforios/) and AWS SDK for Android
//    (http://aws.amazon.com/sdkforandroid/). These toolkits contain sample
//    apps that show how to invoke the identity providers, and then how to use
//    the information from these providers to get and use temporary security
//    credentials.
//
//    * Web Identity Federation with Mobile Applications (http://aws.amazon.com/articles/4617974389850313).
//    This article discusses web identity federation and shows an example of
//    how to use web identity federation to get access to content in Amazon
//    S3.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Security Token Service's
// API operation AssumeRoleWithWebIdentity for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeMalformedPolicyDocumentException "MalformedPolicyDocument"
//   The request was rejected because the policy document was malformed. The error
//   message describes the specific error.
//
//   * ErrCodePackedPolicyTooLargeException "PackedPolicyTooLarge"
//   The request was rejected because the policy document was too large. The error
//   message describes how big the policy document is, in packed form, as a percentage
//   of what the API allows.
//
//   * ErrCodeIDPRejectedClaimException "IDPRejectedClaim"
//   The identity provider (IdP) reported that authentication failed. This might
//   be because the claim is invalid.
//
//   If this error is returned for the AssumeRoleWithWebIdentity operation, it
//   can also mean that the claim has expired or has been explicitly revoked.
//
//   * ErrCodeIDPCommunicationErrorException "IDPCommunicationError"
//   The request could not be fulfilled because the non-AWS identity provider
//   (IDP) that was asked to verify the incoming identity token could not be reached.
//   This is often a transient error caused by network conditions. Retry the request
//   a limited number of times so that you don't exceed the request rate. If the
//   error persists, the non-AWS identity provider might be down or not responding.
//
//   * ErrCodeInvalidIdentityTokenException "InvalidIdentityToken"
//   The web identity token that was passed could not be validated by AWS. Get
//   a new identity token from the identity provider and then retry the request.
//
//   * ErrCodeExpiredTokenException "ExpiredTokenException"
//   The web identity token that was passed is expired or is not valid. Get a
//   new identity token from the identity provider and then retry the request.
//
//   * ErrCodeRegionDisabledException "RegionDisabledException"
//   STS is not activated in the requested region for the account that is being
//   asked to generate credentials. The account administrator must use the IAM
//   console to activate STS in that region. For more information, see Activating
//   and Deactivating AWS STS in an AWS Region (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
//   in the IAM User Guide.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/AssumeRoleWithWebIdentity
func (c *STS) AssumeRoleWithWebIdentity(input *AssumeRoleWithWebIdentityInput) (*AssumeRoleWithWebIdentityOutput, error) {
	req, out := c.AssumeRoleWithWebIdentityRequest(input)
	return out, req.Send()
}

// AssumeRoleWithWebIdentityWithContext is the same as AssumeRoleWithWebIdentity with the addition of
// the ability to pass a context and additional request options.
//
// See AssumeRoleWithWebIdentity for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STS) AssumeRoleWithWebIdentityWithContext(ctx aws.Context, input *AssumeRoleWithWebIdentityInput, opts ...request.Option) (*AssumeRoleWithWebIdentityOutput, error) {
	req, out := c.AssumeRoleWithWebIdentityRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDecodeAuthorizationMessage = "DecodeAuthorizationMessage"

// DecodeAuthorizationMessageRequest generates a "aws/request.Request" representing the
// client's request for the DecodeAuthorizationMessage operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DecodeAuthorizationMessage for more information on using the DecodeAuthorizationMessage
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DecodeAuthorizationMessageRequest method.
//    req, resp := client.DecodeAuthorizationMessageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/DecodeAuthorizationMessage
func (c *STS) DecodeAuthorizationMessageRequest(input *DecodeAuthorizationMessageInput) (req *request.Request, output *DecodeAuthorizationMessageOutput) {
	op := &request.Operation{
		Name:       opDecodeAuthorizationMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DecodeAuthorizationMessageInput{}
	}

	output = &DecodeAuthorizationMessageOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DecodeAuthorizationMessage API operation for AWS Security Token Service.
//
// Decodes additional information about the authorization status of a request
// from an encoded message returned in response to an AWS request.
//
// For example, if a user is not authorized to perform an action that he or
// she has requested, the request returns a Client.UnauthorizedOperation response
// (an HTTP 403 response). Some AWS actions additionally return an encoded message
// that can provide details about this authorization failure.
//
// Only certain AWS actions return an encoded authorization message. The documentation
// for an individual action indicates whether that action returns an encoded
// message in addition to returning an HTTP code.
//
// The message is encoded because the details of the authorization status can
// constitute privileged information that the user who requested the action
// should not see. To decode an authorization status message, a user must be
// granted permissions via an IAM policy to request the DecodeAuthorizationMessage
// (sts:DecodeAuthorizationMessage) action.
//
// The decoded message includes the following type of information:
//
//    * Whether the request was denied due to an explicit deny or due to the
//    absence of an explicit allow. For more information, see Determining Whether
//    a Request is Allowed or Denied (http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-denyallow)
//    in the IAM User Guide.
//
//    * The principal who made the request.
//
//    * The requested action.
//
//    * The requested resource.
//
//    * The values of condition keys in the context of the user's request.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Security Token Service's
// API operation DecodeAuthorizationMessage for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidAuthorizationMessageException "InvalidAuthorizationMessageException"
//   The error returned if the message passed to DecodeAuthorizationMessage was
//   invalid. This can happen if the token contains invalid characters, such as
//   linebreaks.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/DecodeAuthorizationMessage
func (c *STS) DecodeAuthorizationMessage(input *DecodeAuthorizationMessageInput) (*DecodeAuthorizationMessageOutput, error) {
	req, out := c.DecodeAuthorizationMessageRequest(input)
	return out, req.Send()
}

// DecodeAuthorizationMessageWithContext is the same as DecodeAuthorizationMessage with the addition of
// the ability to pass a context and additional request options.
//
// See DecodeAuthorizationMessage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STS) DecodeAuthorizationMessageWithContext(ctx aws.Context, input *DecodeAuthorizationMessageInput, opts ...request.Option) (*DecodeAuthorizationMessageOutput, error) {
	req, out := c.DecodeAuthorizationMessageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetCallerIdentity = "GetCallerIdentity"

// GetCallerIdentityRequest generates a "aws/request.Request" representing the
// client's request for the GetCallerIdentity operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetCallerIdentity for more information on using the GetCallerIdentity
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetCallerIdentityRequest method.
//    req, resp := client.GetCallerIdentityRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetCallerIdentity
func (c *STS) GetCallerIdentityRequest(input *GetCallerIdentityInput) (req *request.Request, output *GetCallerIdentityOutput) {
	op := &request.Operation{
		Name:       opGetCallerIdentity,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetCallerIdentityInput{}
	}

	output = &GetCallerIdentityOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.URL.RawQuery += "Action=GetCallerIdentity"
	return
}

// GetCallerIdentity API operation for AWS Security Token Service.
//
// Returns details about the IAM identity whose credentials are used to call
// the API.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Security Token Service's
// API operation GetCallerIdentity for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetCallerIdentity
func (c *STS) GetCallerIdentity(input *GetCallerIdentityInput) (*GetCallerIdentityOutput, error) {
	req, out := c.GetCallerIdentityRequest(input)
	return out, req.Send()
}

// GetCallerIdentityWithContext is the same as GetCallerIdentity with the addition of
// the ability to pass a context and additional request options.
//
// See GetCallerIdentity for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STS) GetCallerIdentityWithContext(ctx aws.Context, input *GetCallerIdentityInput, opts ...request.Option) (*GetCallerIdentityOutput, error) {
	req, out := c.GetCallerIdentityRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetFederationToken = "GetFederationToken"

// GetFederationTokenRequest generates a "aws/request.Request" representing the
// client's request for the GetFederationToken operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetFederationToken for more information on using the GetFederationToken
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetFederationTokenRequest method.
//    req, resp := client.GetFederationTokenRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetFederationToken
func (c *STS) GetFederationTokenRequest(input *GetFederationTokenInput) (req *request.Request, output *GetFederationTokenOutput) {
	op := &request.Operation{
		Name:       opGetFederationToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetFederationTokenInput{}
	}

	output = &GetFederationTokenOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetFederationToken API operation for AWS Security Token Service.
//
// Returns a set of temporary security credentials (consisting of an access
// key ID, a secret access key, and a security token) for a federated user.
// A typical use is in a proxy application that gets temporary security credentials
// on behalf of distributed applications inside a corporate network. Because
// you must call the GetFederationToken action using the long-term security
// credentials of an IAM user, this call is appropriate in contexts where those
// credentials can be safely stored, usually in a server-based application.
// For a comparison of GetFederationToken with the other APIs that produce temporary
// credentials, see Requesting Temporary Security Credentials (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS APIs (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide.
//
// If you are creating a mobile-based or browser-based app that can authenticate
// users using a web identity provider like Login with Amazon, Facebook, Google,
// or an OpenID Connect-compatible identity provider, we recommend that you
// use Amazon Cognito (http://aws.amazon.com/cognito/) or AssumeRoleWithWebIdentity.
// For more information, see Federation Through a Web-based Identity Provider
// (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_assumerolewithwebidentity).
//
// The GetFederationToken action must be called by using the long-term AWS security
// credentials of an IAM user. You can also call GetFederationToken using the
// security credentials of an AWS root account, but we do not recommended it.
// Instead, we recommend that you create an IAM user for the purpose of the
// proxy application and then attach a policy to the IAM user that limits federated
// users to only the actions and resources that they need access to. For more
// information, see IAM Best Practices (http://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html)
// in the IAM User Guide.
//
// The temporary security credentials that are obtained by using the long-term
// credentials of an IAM user are valid for the specified duration, from 900
// seconds (15 minutes) up to a maximium of 129600 seconds (36 hours). The default
// is 43200 seconds (12 hours). Temporary credentials that are obtained by using
// AWS root account credentials have a maximum duration of 3600 seconds (1 hour).
//
// The temporary security credentials created by GetFederationToken can be used
// to make API calls to any AWS service with the following exceptions:
//
//    * You cannot use these credentials to call any IAM APIs.
//
//    * You cannot call any STS APIs except GetCallerIdentity.
//
// Permissions
//
// The permissions for the temporary security credentials returned by GetFederationToken
// are determined by a combination of the following:
//
//    * The policy or policies that are attached to the IAM user whose credentials
//    are used to call GetFederationToken.
//
//    * The policy that is passed as a parameter in the call.
//
// The passed policy is attached to the temporary security credentials that
// result from the GetFederationToken API call--that is, to the federated user.
// When the federated user makes an AWS request, AWS evaluates the policy attached
// to the federated user in combination with the policy or policies attached
// to the IAM user whose credentials were used to call GetFederationToken. AWS
// allows the federated user's request only when both the federated user and
// the IAM user are explicitly allowed to perform the requested action. The
// passed policy cannot grant more permissions than those that are defined in
// the IAM user policy.
//
// A typical use case is that the permissions of the IAM user whose credentials
// are used to call GetFederationToken are designed to allow access to all the
// actions and resources that any federated user will need. Then, for individual
// users, you pass a policy to the operation that scopes down the permissions
// to a level that's appropriate to that individual user, using a policy that
// allows only a subset of permissions that are granted to the IAM user.
//
// If you do not pass a policy, the resulting temporary security credentials
// have no effective permissions. The only exception is when the temporary security
// credentials are used to access a resource that has a resource-based policy
// that specifically allows the federated user to access the resource.
//
// For more information about how permissions work, see Permissions for GetFederationToken
// (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_getfederationtoken.html).
// For information about using GetFederationToken to create temporary security
// credentials, see GetFederationToken—Federation Through a Custom Identity
// Broker (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_getfederationtoken).
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Security Token Service's
// API operation GetFederationToken for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeMalformedPolicyDocumentException "MalformedPolicyDocument"
//   The request was rejected because the policy document was malformed. The error
//   message describes the specific error.
//
//   * ErrCodePackedPolicyTooLargeException "PackedPolicyTooLarge"
//   The request was rejected because the policy document was too large. The error
//   message describes how big the policy document is, in packed form, as a percentage
//   of what the API allows.
//
//   * ErrCodeRegionDisabledException "RegionDisabledException"
//   STS is not activated in the requested region for the account that is being
//   asked to generate credentials. The account administrator must use the IAM
//   console to activate STS in that region. For more information, see Activating
//   and Deactivating AWS STS in an AWS Region (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
//   in the IAM User Guide.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetFederationToken
func (c *STS) GetFederationToken(input *GetFederationTokenInput) (*GetFederationTokenOutput, error) {
	req, out := c.GetFederationTokenRequest(input)
	return out, req.Send()
}

// GetFederationTokenWithContext is the same as GetFederationToken with the addition of
// the ability to pass a context and additional request options.
//
// See GetFederationToken for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STS) GetFederationTokenWithContext(ctx aws.Context, input *GetFederationTokenInput, opts ...request.Option) (*GetFederationTokenOutput, error) {
	req, out := c.GetFederationTokenRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetSessionToken = "GetSessionToken"

// GetSessionTokenRequest generates a "aws/request.Request" representing the
// client's request for the GetSessionToken operation. The "output" return
// value will be populated with the request's response once the request complets
// successfuly.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetSessionToken for more information on using the GetSessionToken
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetSessionTokenRequest method.
//    req, resp := client.GetSessionTokenRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetSessionToken
func (c *STS) GetSessionTokenRequest(input *GetSessionTokenInput) (req *request.Request, output *GetSessionTokenOutput) {
	op := &request.Operation{
		Name:       opGetSessionToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSessionTokenInput{}
	}

	output = &GetSessionTokenOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetSessionToken API operation for AWS Security Token Service.
//
// Returns a set of temporary credentials for an AWS account or IAM user. The
// credentials consist of an access key ID, a secret access key, and a security
// token. Typically, you use GetSessionToken if you want to use MFA to protect
// programmatic calls to specific AWS APIs like Amazon EC2 StopInstances. MFA-enabled
// IAM users would need to call GetSessionToken and submit an MFA code that
// is associated with their MFA device. Using the temporary security credentials
// that are returned from the call, IAM users can then make programmatic calls
// to APIs that require MFA authentication. If you do not supply a correct MFA
// code, then the API returns an access denied error. For a comparison of GetSessionToken
// with the other APIs that produce temporary credentials, see Requesting Temporary
// Security Credentials (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS APIs (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide.
//
// The GetSessionToken action must be called by using the long-term AWS security
// credentials of the AWS account or an IAM user. Credentials that are created
// by IAM users are valid for the duration that you specify, from 900 seconds
// (15 minutes) up to a maximum of 129600 seconds (36 hours), with a default
// of 43200 seconds (12 hours); credentials that are created by using account
// credentials can range from 900 seconds (15 minutes) up to a maximum of 3600
// seconds (1 hour), with a default of 1 hour.
//
// The temporary security credentials created by GetSessionToken can be used
// to make API calls to any AWS service with the following exceptions:
//
//    * You cannot call any IAM APIs unless MFA authentication information is
//    included in the request.
//
//    * You cannot call any STS API exceptAssumeRole or GetCallerIdentity.
//
// We recommend that you do not call GetSessionToken with root account credentials.
// Instead, follow our best practices (http://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#create-iam-users)
// by creating one or more IAM users, giving them the necessary permissions,
// and using IAM users for everyday interaction with AWS.
//
// The permissions associated with the temporary security credentials returned
// by GetSessionToken are based on the permissions associated with account or
// IAM user whose credentials are used to call the action. If GetSessionToken
// is called using root account credentials, the temporary credentials have
// root account permissions. Similarly, if GetSessionToken is called using the
// credentials of an IAM user, the temporary credentials have the same permissions
// as the IAM user.
//
// For more information about using GetSessionToken to create temporary credentials,
// go to Temporary Credentials for Users in Untrusted Environments (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_getsessiontoken)
// in the IAM User Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Security Token Service's
// API operation GetSessionToken for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeRegionDisabledException "RegionDisabledException"
//   STS is not activated in the requested region for the account that is being
//   asked to generate credentials. The account administrator must use the IAM
//   console to activate STS in that region. For more information, see Activating
//   and Deactivating AWS STS in an AWS Region (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
//   in the IAM User Guide.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetSessionToken
func (c *STS) GetSessionToken(input *GetSessionTokenInput) (*GetSessionTokenOutput, error) {
	req, out := c.GetSessionTokenRequest(input)
	return out, req.Send()
}

// GetSessionTokenWithContext is the same as GetSessionToken with the addition of
// the ability to pass a context and additional request options.
//
// See GetSessionToken for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STS) GetSessionTokenWithContext(ctx aws.Context, input *GetSessionTokenInput, opts ...request.Option) (*GetSessionTokenOutput, error) {
	req, out := c.GetSessionTokenRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssumeRoleInput struct {
	_ struct{} `type:"structure"`

	// The duration, in seconds, of the role session. The value can range from 900
	// seconds (15 minutes) to 3600 seconds (1 hour). By default, the value is set
	// to 3600 seconds.
	//
	// This is separate from the duration of a console session that you might request
	// using the returned credentials. The request to the federation endpoint for
	// a console sign-in token takes a SessionDuration parameter that specifies
	// the maximum length of the console session, separately from the DurationSeconds
	// parameter on this API. For more information, see Creating a URL that Enables
	// Federated Users to Access the AWS Management Console (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-custom-url.html)
	// in the IAM User Guide.
	DurationSeconds *int64 `min:"900" type:"integer"`

	// A unique identifier that is used by third parties when assuming roles in
	// their customers' accounts. For each role that the third party can assume,
	// they should instruct their customers to ensure the role's trust policy checks
	// for the external ID that the third party generated. Each time the third party
	// assumes the role, they should pass the customer's external ID. The external
	// ID is useful in order to help third parties bind a role to the customer who
	// created it. For more information about the external ID, see How to Use an
	// External ID When Granting Access to Your AWS Resources to a Third Party (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html)
	// in the IAM User Guide.
	//
	// The regex used to validated this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can
	// also include underscores or any of the following characters: =,.@:/-
	ExternalId *string `min:"2" type:"string"`

	// An IAM policy in JSON format.
	//
	// This parameter is optional. If you pass a policy, the temporary security
	// credentials that are returned by the operation have the permissions that
	// are allowed by both (the intersection of) the access policy of the role that
	// is being assumed, and the policy that you pass. This gives you a way to further
	// restrict the permissions for the resulting temporary security credentials.
	// You cannot use the passed policy to grant permissions that are in excess
	// of those allowed by the access policy of the role that is being assumed.
	// For more information, see Permissions for AssumeRole, AssumeRoleWithSAML,
	// and AssumeRoleWithWebIdentity (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_assumerole.html)
	// in the IAM User Guide.
	//
	// The format for this parameter, as described by its regex pattern, is a string
	// of characters up to 2048 characters in length. The characters can be any
	// ASCII character from the space character to the end of the valid character
	// list (\u0020-\u00FF). It can also include the tab (\u0009), linefeed (\u000A),
	// and carriage return (\u000D) characters.
	//
	// The policy plain text must be 2048 bytes or shorter. However, an internal
	// conversion compresses it into a packed binary format with a separate limit.
	// The PackedPolicySize response element indicates by percentage how close to
	// the upper size limit the policy is, with 100% equaling the maximum allowed
	// size.
	Policy *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the role to assume.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`

	// An identifier for the assumed role session.
	//
	// Use the role session name to uniquely identify a session when the same role
	// is assumed by different principals or for different reasons. In cross-account
	// scenarios, the role session name is visible to, and can be logged by the
	// account that owns the role. The role session name is also used in the ARN
	// of the assumed role principal. This means that subsequent cross-account API
	// requests using the temporary security credentials will expose the role session
	// name to the external account in their CloudTrail logs.
	//
	// The regex used to validate this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can
	// also include underscores or any of the following characters: =,.@-
	//
	// RoleSessionName is a required field
	RoleSessionName *string `min:"2" type:"string" required:"true"`

	// The identification number of the MFA device that is associated with the user
	// who is making the AssumeRole call. Specify this value if the trust policy
	// of the role being assumed includes a condition that requires MFA authentication.
	// The value is either the serial number for a hardware device (such as GAHT12345678)
	// or an Amazon Resource Name (ARN) for a virtual device (such as arn:aws:iam::123456789012:mfa/user).
	//
	// The regex used to validate this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can
	// also include underscores or any of the following characters: =,.@-
	SerialNumber *string `min:"9" type:"string"`

	// The value provided by the MFA device, if the trust policy of the role being
	// assumed requires MFA (that is, if the policy includes a condition that tests
	// for MFA). If the role being assumed requires MFA and if the TokenCode value
	// is missing or expired, the AssumeRole call returns an "access denied" error.
	//
	// The format for this parameter, as described by its regex pattern, is a sequence
	// of six numeric digits.
	TokenCode *string `min:"6" type:"string"`
}

// String returns the string representation
func (s AssumeRoleInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AssumeRoleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssumeRoleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssumeRoleInput"}
	if s.DurationSeconds != nil && *s.DurationSeconds < 900 {
		invalidParams.Add(request.NewErrParamMinValue("DurationSeconds", 900))
	}
	if s.ExternalId != nil && len(*s.ExternalId) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("ExternalId", 2))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Policy", 1))
	}
	if s.RoleArn == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(request.NewErrParamMinLen("RoleArn", 20))
	}
	if s.RoleSessionName == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleSessionName"))
	}
	if s.RoleSessionName != nil && len(*s.RoleSessionName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("RoleSessionName", 2))
	}
	if s.SerialNumber != nil && len(*s.SerialNumber) < 9 {
		invalidParams.Add(request.NewErrParamMinLen("SerialNumber", 9))
	}
	if s.TokenCode != nil && len(*s.TokenCode) < 6 {
		invalidParams.Add(request.NewErrParamMinLen("TokenCode", 6))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDurationSeconds sets the DurationSeconds field's value.
func (s *AssumeRoleInput) SetDurationSeconds(v int64) *AssumeRoleInput {
	s.DurationSeconds = &v
	return s
}

// SetExternalId sets the ExternalId field's value.
func (s *AssumeRoleInput) SetExternalId(v string) *AssumeRoleInput {
	s.ExternalId = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *AssumeRoleInput) SetPolicy(v string) *AssumeRoleInput {
	s.Policy = &v
	return s
}

// SetRoleArn sets the RoleArn field's value.
func (s *AssumeRoleInput) SetRoleArn(v string) *AssumeRoleInput {
	s.RoleArn = &v
	return s
}

// SetRoleSessionName sets the RoleSessionName field's value.
func (s *AssumeRoleInput) SetRoleSessionName(v string) *AssumeRoleInput {
	s.RoleSessionName = &v
	return s
}

// SetSerialNumber sets the SerialNumber field's value.
func (s *AssumeRoleInput) SetSerialNumber(v string) *AssumeRoleInput {
	s.SerialNumber = &v
	return s
}

// SetTokenCode sets the TokenCode field's value.
func (s *AssumeRoleInput) SetTokenCode(v string) *AssumeRoleInput {
	s.TokenCode = &v
	return s
}

// Contains the response to a successful AssumeRole request, including temporary
// AWS credentials that can be used to make AWS requests.
type AssumeRoleOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) and the assumed role ID, which are identifiers
	// that you can use to refer to the resulting temporary security credentials.
	// For example, you can reference these credentials as a principal in a resource-based
	// policy by using the ARN or assumed role ID. The ARN and ID include the RoleSessionName
	// that you specified when you called AssumeRole.
	AssumedRoleUser *AssumedRoleUser `type:"structure"`

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token.
	//
	// Note: The size of the security token that STS APIs return is not fixed. We
	// strongly recommend that you make no assumptions about the maximum size. As
	// of this writing, the typical size is less than 4096 bytes, but that can vary.
	// Also, future updates to AWS might require larger sizes.
	Credentials *Credentials `type:"structure"`

	// A percentage value that indicates the size of the policy in packed form.
	// The service rejects any policy with a packed size greater than 100 percent,
	// which means the policy exceeded the allowed space.
	PackedPolicySize *int64 `type:"integer"`
}

// String returns the string representation
func (s AssumeRoleOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AssumeRoleOutput) GoString() string {
	return s.String()
}

// SetAssumedRoleUser sets the AssumedRoleUser field's value.
func (s *AssumeRoleOutput) SetAssumedRoleUser(v *AssumedRoleUser) *AssumeRoleOutput {
	s.AssumedRoleUser = v
	return s
}

// SetCredentials sets the Credentials field's value.
func (s *AssumeRoleOutput) SetCredentials(v *Credentials) *AssumeRoleOutput {
	s.Credentials = v
	return s
}

// SetPackedPolicySize sets the PackedPolicySize field's value.
func (s *AssumeRoleOutput) SetPackedPolicySize(v int64) *AssumeRoleOutput {
	s.PackedPolicySize = &v
	return s
}

type AssumeRoleWithSAMLInput struct {
	_ struct{} `type:"structure"`

	// The duration, in seconds, of the role session. The value can range from 900
	// seconds (15 minutes) to 3600 seconds (1 hour). By default, the value is set
	// to 3600 seconds. An expiration can also be specified in the SAML authentication
	// response's SessionNotOnOrAfter value. The actual expiration time is whichever
	// value is shorter.
	//
	// This is separate from the duration of a console session that you might request
	// using the returned credentials. The request to the federation endpoint for
	// a console sign-in token takes a SessionDuration parameter that specifies
	// the maximum length of the console session, separately from the DurationSeconds
	// parameter on this API. For more information, see Enabling SAML 2.0 Federated
	// Users to Access the AWS Management Console (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-saml.html)
	// in the IAM User Guide.
	DurationSeconds *int64 `min:"900" type:"integer"`

	// An IAM policy in JSON format.
	//
	// The policy parameter is optional. If you pass a policy, the temporary security
	// credentials that are returned by the operation have the permissions that
	// are allowed by both the access policy of the role that is being assumed,
	// and the policy that you pass. This gives you a way to further restrict the
	// permissions for the resulting temporary security credentials. You cannot
	// use the passed policy to grant permissions that are in excess of those allowed
	// by the access policy of the role that is being assumed. For more information,
	// Permissions for AssumeRole, AssumeRoleWithSAML, and AssumeRoleWithWebIdentity
	// (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_assumerole.html)
	// in the IAM User Guide.
	//
	// The format for this parameter, as described by its regex pattern, is a string
	// of characters up to 2048 characters in length. The characters can be any
	// ASCII character from the space character to the end of the valid character
	// list (\u0020-\u00FF). It can also include the tab (\u0009), linefeed (\u000A),
	// and carriage return (\u000D) characters.
	//
	// The policy plain text must be 2048 bytes or shorter. However, an internal
	// conversion compresses it into a packed binary format with a separate limit.
	// The PackedPolicySize response element indicates by percentage how close to
	// the upper size limit the policy is, with 100% equaling the maximum allowed
	// size.
	Policy *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the SAML provider in IAM that describes
	// the IdP.
	//
	// PrincipalArn is a required field
	PrincipalArn *string `min:"20" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the role that the caller is assuming.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`

	// The base-64 encoded SAML authentication response provided by the IdP.
	//
	// For more information, see Configuring a Relying Party and Adding Claims (http://docs.aws.amazon.com/IAM/latest/UserGuide/create-role-saml-IdP-tasks.html)
	// in the Using IAM guide.
	//
	// SAMLAssertion is a required field
	SAMLAssertion *string `min:"4" type:"string" required:"true"`
}

// String returns the string representation
func (s AssumeRoleWithSAMLInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AssumeRoleWithSAMLInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssumeRoleWithSAMLInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssumeRoleWithSAMLInput"}
	if s.DurationSeconds != nil && *s.DurationSeconds < 900 {
		invalidParams.Add(request.NewErrParamMinValue("DurationSeconds", 900))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Policy", 1))
	}
	if s.PrincipalArn == nil {
		invalidParams.Add(request.NewErrParamRequired("PrincipalArn"))
	}
	if s.PrincipalArn != nil && len(*s.PrincipalArn) < 20 {
		invalidParams.Add(request.NewErrParamMinLen("PrincipalArn", 20))
	}
	if s.RoleArn == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(request.NewErrParamMinLen("RoleArn", 20))
	}
	if s.SAMLAssertion == nil {
		invalidParams.Add(request.NewErrParamRequired("SAMLAssertion"))
	}
	if s.SAMLAssertion != nil && len(*s.SAMLAssertion) < 4 {
		invalidParams.Add(request.NewErrParamMinLen("SAMLAssertion", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDurationSeconds sets the DurationSeconds field's value.
func (s *AssumeRoleWithSAMLInput) SetDurationSeconds(v int64) *AssumeRoleWithSAMLInput {
	s.DurationSeconds = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *AssumeRoleWithSAMLInput) SetPolicy(v string) *AssumeRoleWithSAMLInput {
	s.Policy = &v
	return s
}

// SetPrincipalArn sets the PrincipalArn field's value.
func (s *AssumeRoleWithSAMLInput) SetPrincipalArn(v string) *AssumeRoleWithSAMLInput {
	s.PrincipalArn = &v
	return s
}

// SetRoleArn sets the RoleArn field's value.
func (s *AssumeRoleWithSAMLInput) SetRoleArn(v string) *AssumeRoleWithSAMLInput {
	s.RoleArn = &v
	return s
}

// SetSAMLAssertion sets the SAMLAssertion field's value.
func (s *AssumeRoleWithSAMLInput) SetSAMLAssertion(v string) *AssumeRoleWithSAMLInput {
	s.SAMLAssertion = &v
	return s
}

// Contains the response to a successful AssumeRoleWithSAML request, including
// temporary AWS credentials that can be used to make AWS requests.
type AssumeRoleWithSAMLOutput struct {
	_ struct{} `type:"structure"`

	// The identifiers for the temporary security credentials that the operation
	// returns.
	AssumedRoleUser *AssumedRoleUser `type:"structure"`

	// The value of the Recipient attribute of the SubjectConfirmationData element
	// of the SAML assertion.
	Audience *string `type:"string"`

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token.
	//
	// Note: The size of the security token that STS APIs return is not fixed. We
	// strongly recommend that you make no assumptions about the maximum size. As
	// of this writing, the typical size is less than 4096 bytes, but that can vary.
	// Also, future updates to AWS might require larger sizes.
	Credentials *Credentials `type:"structure"`

	// The value of the Issuer element of the SAML assertion.
	Issuer *string `type:"string"`

	// A hash value based on the concatenation of the Issuer response value, the
	// AWS account ID, and the friendly name (the last part of the ARN) of the SAML
	// provider in IAM. The combination of NameQualifier and Subject can be used
	// to uniquely identify a federated user.
	//
	// The following pseudocode shows how the hash value is calculated:
	//
	// BASE64 ( SHA1 ( "https://example.com/saml" + "123456789012" + "/MySAMLIdP"
	// ) )
	NameQualifier *string `type:"string"`

	// A percentage value that indicates the size of the policy in packed form.
	// The service rejects any policy with a packed size greater than 100 percent,
	// which means the policy exceeded the allowed space.
	PackedPolicySize *int64 `type:"integer"`

	// The value of the NameID element in the Subject element of the SAML assertion.
	Subject *string `type:"string"`

	// The format of the name ID, as defined by the Format attribute in the NameID
	// element of the SAML assertion. Typical examples of the format are transient
	// or persistent.
	//
	// If the format includes the prefix urn:oasis:names:tc:SAML:2.0:nameid-format,
	// that prefix is removed. For example, urn:oasis:names:tc:SAML:2.0:nameid-format:transient
	// is returned as transient. If the format includes any other prefix, the format
	// is returned with no modifications.
	SubjectType *string `type:"string"`
}

// String returns the string representation
func (s AssumeRoleWithSAMLOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AssumeRoleWithSAMLOutput) GoString() string {
	return s.String()
}

// SetAssumedRoleUser sets the AssumedRoleUser field's value.
func (s *AssumeRoleWithSAMLOutput) SetAssumedRoleUser(v *AssumedRoleUser) *AssumeRoleWithSAMLOutput {
	s.AssumedRoleUser = v
	return s
}

// SetAudience sets the Audience field's value.
func (s *AssumeRoleWithSAMLOutput) SetAudience(v string) *AssumeRoleWithSAMLOutput {
	s.Audience = &v
	return s
}

// SetCredentials sets the Credentials field's value.
func (s *AssumeRoleWithSAMLOutput) SetCredentials(v *Credentials) *AssumeRoleWithSAMLOutput {
	s.Credentials = v
	return s
}

// SetIssuer sets the Issuer field's value.
func (s *AssumeRoleWithSAMLOutput) SetIssuer(v string) *AssumeRoleWithSAMLOutput {
	s.Issuer = &v
	return s
}

// SetNameQualifier sets the NameQualifier field's value.
func (s *AssumeRoleWithSAMLOutput) SetNameQualifier(v string) *AssumeRoleWithSAMLOutput {
	s.NameQualifier = &v
	return s
}

// SetPackedPolicySize sets the PackedPolicySize field's value.
func (s *AssumeRoleWithSAMLOutput) SetPackedPolicySize(v int64) *AssumeRoleWithSAMLOutput {
	s.PackedPolicySize = &v
	return s
}

// SetSubject sets the Subject field's value.
func (s *AssumeRoleWithSAMLOutput) SetSubject(v string) *AssumeRoleWithSAMLOutput {
	s.Subject = &v
	return s
}

// SetSubjectType sets the SubjectType field's value.
func (s *AssumeRoleWithSAMLOutput) SetSubjectType(v string) *AssumeRoleWithSAMLOutput {
	s.SubjectType = &v
	return s
}

type AssumeRoleWithWebIdentityInput struct {
	_ struct{} `type:"structure"`

	// The duration, in seconds, of the role session. The value can range from 900
	// seconds (15 minutes) to 3600 seconds (1 hour). By default, the value is set
	// to 3600 seconds.
	//
	// This is separate from the duration of a console session that you might request
	// using the returned credentials. The request to the federation endpoint for
	// a console sign-in token takes a SessionDuration parameter that specifies
	// the maximum length of the console session, separately from the DurationSeconds
	// parameter on this API. For more information, see Creating a URL that Enables
	// Federated Users to Access the AWS Management Console (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_enable-console-custom-url.html)
	// in the IAM User Guide.
	DurationSeconds *int64 `min:"900" type:"integer"`

	// An IAM policy in JSON format.
	//
	// The policy parameter is optional. If you pass a policy, the temporary security
	// credentials that are returned by the operation have the permissions that
	// are allowed by both the access policy of the role that is being assumed,
	// and the policy that you pass. This gives you a way to further restrict the
	// permissions for the resulting temporary security credentials. You cannot
	// use the passed policy to grant permissions that are in excess of those allowed
	// by the access policy of the role that is being assumed. For more information,
	// see Permissions for AssumeRoleWithWebIdentity (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_assumerole.html)
	// in the IAM User Guide.
	//
	// The format for this parameter, as described by its regex pattern, is a string
	// of characters up to 2048 characters in length. The characters can be any
	// ASCII character from the space character to the end of the valid character
	// list (\u0020-\u00FF). It can also include the tab (\u0009), linefeed (\u000A),
	// and carriage return (\u000D) characters.
	//
	// The policy plain text must be 2048 bytes or shorter. However, an internal
	// conversion compresses it into a packed binary format with a separate limit.
	// The PackedPolicySize response element indicates by percentage how close to
	// the upper size limit the policy is, with 100% equaling the maximum allowed
	// size.
	Policy *string `min:"1" type:"string"`

	// The fully qualified host component of the domain name of the identity provider.
	//
	// Specify this value only for OAuth 2.0 access tokens. Currently www.amazon.com
	// and graph.facebook.com are the only supported identity providers for OAuth
	// 2.0 access tokens. Do not include URL schemes and port numbers.
	//
	// Do not specify this value for OpenID Connect ID tokens.
	ProviderId *string `min:"4" type:"string"`

	// The Amazon Resource Name (ARN) of the role that the caller is assuming.
	//
	// RoleArn is a required field
	RoleArn *string `min:"20" type:"string" required:"true"`

	// An identifier for the assumed role session. Typically, you pass the name
	// or identifier that is associated with the user who is using your application.
	// That way, the temporary security credentials that your application will use
	// are associated with that user. This session name is included as part of the
	// ARN and assumed role ID in the AssumedRoleUser response element.
	//
	// The regex used to validate this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can
	// also include underscores or any of the following characters: =,.@-
	//
	// RoleSessionName is a required field
	RoleSessionName *string `min:"2" type:"string" required:"true"`

	// The OAuth 2.0 access token or OpenID Connect ID token that is provided by
	// the identity provider. Your application must get this token by authenticating
	// the user who is using your application with a web identity provider before
	// the application makes an AssumeRoleWithWebIdentity call.
	//
	// WebIdentityToken is a required field
	WebIdentityToken *string `min:"4" type:"string" required:"true"`
}

// String returns the string representation
func (s AssumeRoleWithWebIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AssumeRoleWithWebIdentityInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssumeRoleWithWebIdentityInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssumeRoleWithWebIdentityInput"}
	if s.DurationSeconds != nil && *s.DurationSeconds < 900 {
		invalidParams.Add(request.NewErrParamMinValue("DurationSeconds", 900))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Policy", 1))
	}
	if s.ProviderId != nil && len(*s.ProviderId) < 4 {
		invalidParams.Add(request.NewErrParamMinLen("ProviderId", 4))
	}
	if s.RoleArn == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleArn"))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(request.NewErrParamMinLen("RoleArn", 20))
	}
	if s.RoleSessionName == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleSessionName"))
	}
	if s.RoleSessionName != nil && len(*s.RoleSessionName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("RoleSessionName", 2))
	}
	if s.WebIdentityToken == nil {
		invalidParams.Add(request.NewErrParamRequired("WebIdentityToken"))
	}
	if s.WebIdentityToken != nil && len(*s.WebIdentityToken) < 4 {
		invalidParams.Add(request.NewErrParamMinLen("WebIdentityToken", 4))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDurationSeconds sets the DurationSeconds field's value.
func (s *AssumeRoleWithWebIdentityInput) SetDurationSeconds(v int64) *AssumeRoleWithWebIdentityInput {
	s.DurationSeconds = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *AssumeRoleWithWebIdentityInput) SetPolicy(v string) *AssumeRoleWithWebIdentityInput {
	s.Policy = &v
	return s
}

// SetProviderId sets the ProviderId field's value.
func (s *AssumeRoleWithWebIdentityInput) SetProviderId(v string) *AssumeRoleWithWebIdentityInput {
	s.ProviderId = &v
	return s
}

// SetRoleArn sets the RoleArn field's value.
func (s *AssumeRoleWithWebIdentityInput) SetRoleArn(v string) *AssumeRoleWithWebIdentityInput {
	s.RoleArn = &v
	return s
}

// SetRoleSessionName sets the RoleSessionName field's value.
func (s *AssumeRoleWithWebIdentityInput) SetRoleSessionName(v string) *AssumeRoleWithWebIdentityInput {
	s.RoleSessionName = &v
	return s
}

// SetWebIdentityToken sets the WebIdentityToken field's value.
func (s *AssumeRoleWithWebIdentityInput) SetWebIdentityToken(v string) *AssumeRoleWithWebIdentityInput {
	s.WebIdentityToken = &v
	return s
}

// Contains the response to a successful AssumeRoleWithWebIdentity request,
// including temporary AWS credentials that can be used to make AWS requests.
type AssumeRoleWithWebIdentityOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) and the assumed role ID, which are identifiers
	// that you can use to refer to the resulting temporary security credentials.
	// For example, you can reference these credentials as a principal in a resource-based
	// policy by using the ARN or assumed role ID. The ARN and ID include the RoleSessionName
	// that you specified when you called AssumeRole.
	AssumedRoleUser *AssumedRoleUser `type:"structure"`

	// The intended audience (also known as client ID) of the web identity token.
	// This is traditionally the client identifier issued to the application that
	// requested the web identity token.
	Audience *string `type:"string"`

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security token.
	//
	// Note: The size of the security token that STS APIs return is not fixed. We
	// strongly recommend that you make no assumptions about the maximum size. As
	// of this writing, the typical size is less than 4096 bytes, but that can vary.
	// Also, future updates to AWS might require larger sizes.
	Credentials *Credentials `type:"structure"`

	// A percentage value that indicates the size of the policy in packed form.
	// The service rejects any policy with a packed size greater than 100 percent,
	// which means the policy exceeded the allowed space.
	PackedPolicySize *int64 `type:"integer"`

	// The issuing authority of the web identity token presented. For OpenID Connect
	// ID Tokens this contains the value of the iss field. For OAuth 2.0 access
	// tokens, this contains the value of the ProviderId parameter that was passed
	// in the AssumeRoleWithWebIdentity request.
	Provider *string `type:"string"`

	// The unique user identifier that is returned by the identity provider. This
	// identifier is associated with the WebIdentityToken that was submitted with
	// the AssumeRoleWithWebIdentity call. The identifier is typically unique to
	// the user and the application that acquired the WebIdentityToken (pairwise
	// identifier). For OpenID Connect ID tokens, this field contains the value
	// returned by the identity provider as the token's sub (Subject) claim.
	SubjectFromWebIdentityToken *string `min:"6" type:"string"`
}

// String returns the string representation
func (s AssumeRoleWithWebIdentityOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AssumeRoleWithWebIdentityOutput) GoString() string {
	return s.String()
}

// SetAssumedRoleUser sets the AssumedRoleUser field's value.
func (s *AssumeRoleWithWebIdentityOutput) SetAssumedRoleUser(v *AssumedRoleUser) *AssumeRoleWithWebIdentityOutput {
	s.AssumedRoleUser = v
	return s
}

// SetAudience sets the Audience field's value.
func (s *AssumeRoleWithWebIdentityOutput) SetAudience(v string) *AssumeRoleWithWebIdentityOutput {
	s.Audience = &v
	return s
}

// SetCredentials sets the Credentials field's value.
func (s *AssumeRoleWithWebIdentityOutput) SetCredentials(v *Credentials) *AssumeRoleWithWebIdentityOutput {
	s.Credentials = v
	return s
}

// SetPackedPolicySize sets the PackedPolicySize field's value.
func (s *AssumeRoleWithWebIdentityOutput) SetPackedPolicySize(v int64) *AssumeRoleWithWebIdentityOutput {
	s.PackedPolicySize = &v
	return s
}

// SetProvider sets the Provider field's value.
func (s *AssumeRoleWithWebIdentityOutput) SetProvider(v string) *AssumeRoleWithWebIdentityOutput {
	s.Provider = &v
	return s
}

// SetSubjectFromWebIdentityToken sets the SubjectFromWebIdentityToken field's value.
func (s *AssumeRoleWithWebIdentityOutput) SetSubjectFromWebIdentityToken(v string) *AssumeRoleWithWebIdentityOutput {
	s.SubjectFromWebIdentityToken = &v
	return s
}

// The identifiers for the temporary security credentials that the operation
// returns.
type AssumedRoleUser struct {
	_ struct{} `type:"structure"`

	// The ARN of the temporary security credentials that are returned from the
	// AssumeRole action. For more information about ARNs and how to use them in
	// policies, see IAM Identifiers (http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html)
	// in Using IAM.
	//
	// Arn is a required field
	Arn *string `min:"20" type:"string" required:"true"`

	// A unique identifier that contains the role ID and the role session name of
	// the role that is being assumed. The role ID is generated by AWS when the
	// role is created.
	//
	// AssumedRoleId is a required field
	AssumedRoleId *string `min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s AssumedRoleUser) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AssumedRoleUser) GoString() string {
	return s.String()
}

// SetArn sets the Arn field's value.
func (s *AssumedRoleUser) SetArn(v string) *AssumedRoleUser {
	s.Arn = &v
	return s
}

// SetAssumedRoleId sets the AssumedRoleId field's value.
func (s *AssumedRoleUser) SetAssumedRoleId(v string) *AssumedRoleUser {
	s.AssumedRoleId = &v
	return s
}

// AWS credentials for API authentication.
type Credentials struct {
	_ struct{} `type:"structure"`

	// The access key ID that identifies the temporary security credentials.
	//
	// AccessKeyId is a required field
	AccessKeyId *string `min:"16" type:"string" required:"true"`

	// The date on which the current credentials expire.
	//
	// Expiration is a required field
	Expiration *time.Time `type:"timestamp" timestampFormat:"iso8601" required:"true"`

	// The secret access key that can be used to sign requests.
	//
	// SecretAccessKey is a required field
	SecretAccessKey *string `type:"string" required:"true"`

	// The token that users must pass to the service API to use the temporary credentials.
	//
	// SessionToken is a required field
	SessionToken *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Credentials) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Credentials) GoString() string {
	return s.String()
}

// SetAccessKeyId sets the AccessKeyId field's value.
func (s *Credentials) SetAccessKeyId(v string) *Credentials {
	s.AccessKeyId = &v
	return s
}

// SetExpiration sets the Expiration field's value.
func (s *Credentials) SetExpiration(v time.Time) *Credentials {
	s.Expiration = &v
	return s
}

// SetSecretAccessKey sets the SecretAccessKey field's value.
func (s *Credentials) SetSecretAccessKey(v string) *Credentials {
	s.SecretAccessKey = &v
	return s
}

// SetSessionToken sets the SessionToken field's value.
func (s *Credentials) SetSessionToken(v string) *Credentials {
	s.SessionToken = &v
	return s
}

type DecodeAuthorizationMessageInput struct {
	_ struct{} `type:"structure"`

	// The encoded message that was returned with the response.
	//
	// EncodedMessage is a required field
	EncodedMessage *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DecodeAuthorizationMessageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DecodeAuthorizationMessageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DecodeAuthorizationMessageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DecodeAuthorizationMessageInput"}
	if s.EncodedMessage == nil {
		invalidParams.Add(request.NewErrParamRequired("EncodedMessage"))
	}
	if s.EncodedMessage != nil && len(*s.EncodedMessage) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EncodedMessage", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEncodedMessage sets the EncodedMessage field's value.
func (s *DecodeAuthorizationMessageInput) SetEncodedMessage(v string) *DecodeAuthorizationMessageInput {
	s.EncodedMessage = &v
	return s
}

// A document that contains additional information about the authorization status
// of a request from an encoded message that is returned in response to an AWS
// request.
type DecodeAuthorizationMessageOutput struct {
	_ struct{} `type:"structure"`

	// An XML document that contains the decoded message.
	DecodedMessage *string `type:"string"`
}

// String returns the string representation
func (s DecodeAuthorizationMessageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DecodeAuthorizationMessageOutput) GoString() string {
	return s.String()
}

// SetDecodedMessage sets the DecodedMessage field's value.
func (s *DecodeAuthorizationMessageOutput) SetDecodedMessage(v string) *DecodeAuthorizationMessageOutput {
	s.DecodedMessage = &v
	return s
}

// Identifiers for the federated user that is associated with the credentials.
type FederatedUser struct {
	_ struct{} `type:"structure"`

	// The ARN that specifies the federated user that is associated with the credentials.
	// For more information about ARNs and how to use them in policies, see IAM
	// Identifiers (http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html)
	// in Using IAM.
	//
	// Arn is a required field
	Arn *string `min:"20" type:"string" required:"true"`

	// The string that identifies the federated user associated with the credentials,
	// similar to the unique ID of an IAM user.
	//
	// FederatedUserId is a required field
	FederatedUserId *string `min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s FederatedUser) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s FederatedUser) GoString() string {
	return s.String()
}

// SetArn sets the Arn field's value.
func (s *FederatedUser) SetArn(v string) *FederatedUser {
	s.Arn = &v
	return s
}

// SetFederatedUserId sets the FederatedUserId field's value.
func (s *FederatedUser) SetFederatedUserId(v string) *FederatedUser {
	s.FederatedUserId = &v
	return s
}

type GetCallerIdentityInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetCallerIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetCallerIdentityInput) GoString() string {
	return s.String()
}

// Contains the response to a successful GetCallerIdentity request, including
// information about the entity making the request.
type GetCallerIdentityOutput struct {
	RequestID string `json:"RequestId"`

	// The AWS account ID number of the account that owns or contains the calling
	// entity.
	Account *string `json:"AccountId"`

	// The AWS ARN associated with the calling entity.
	Arn *string `json:"Arn"`

	// The unique identifier of the calling entity. The exact value depends on the
	// type of entity making the call. The values returned are those listed in the
	// aws:userid column in the Principal table (http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_variables.html#principaltable)
	// found on the Policy Variables reference page in the IAM User Guide.
	UserId *string `json:"UserId"`
}

// String returns the string representation
func (s GetCallerIdentityOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetCallerIdentityOutput) GoString() string {
	return s.String()
}

// SetAccount sets the Account field's value.
func (s *GetCallerIdentityOutput) SetAccount(v string) *GetCallerIdentityOutput {
	s.Account = &v
	return s
}

// SetArn sets the Arn field's value.
func (s *GetCallerIdentityOutput) SetArn(v string) *GetCallerIdentityOutput {
	s.Arn = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *GetCallerIdentityOutput) SetUserId(v string) *GetCallerIdentityOutput {
	s.UserId = &v
	return s
}

type GetFederationTokenInput struct {
	_ struct{} `type:"structure"`

	// The duration, in seconds, that the session should last. Acceptable durations
	// for federation sessions range from 900 seconds (15 minutes) to 129600 seconds
	// (36 hours), with 43200 seconds (12 hours) as the default. Sessions obtained
	// using AWS account (root) credentials are restricted to a maximum of 3600
	// seconds (one hour). If the specified duration is longer than one hour, the
	// session obtained by using AWS account (root) credentials defaults to one
	// hour.
	DurationSeconds *int64 `min:"900" type:"integer"`

	// The name of the federated user. The name is used as an identifier for the
	// temporary security credentials (such as Bob). For example, you can reference
	// the federated user name in a resource-based policy, such as in an Amazon
	// S3 bucket policy.
	//
	// The regex used to validate this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can
	// also include underscores or any of the following characters: =,.@-
	//
	// Name is a required field
	Name *string `min:"2" type:"string" required:"true"`

	// An IAM policy in JSON format that is passed with the GetFederationToken call
	// and evaluated along with the policy or policies that are attached to the
	// IAM user whose credentials are used to call GetFederationToken. The passed
	// policy is used to scope down the permissions that are available to the IAM
	// user, by allowing only a subset of the permissions that are granted to the
	// IAM user. The passed policy cannot grant more permissions than those granted
	// to the IAM user. The final permissions for the federated user are the most
	// restrictive set based on the intersection of the passed policy and the IAM
	// user policy.
	//
	// If you do not pass a policy, the resulting temporary security credentials
	// have no effective permissions. The only exception is when the temporary security
	// credentials are used to access a resource that has a resource-based policy
	// that specifically allows the federated user to access the resource.
	//
	// The format for this parameter, as described by its regex pattern, is a string
	// of characters up to 2048 characters in length. The characters can be any
	// ASCII character from the space character to the end of the valid character
	// list (\u0020-\u00FF). It can also include the tab (\u0009), linefeed (\u000A),
	// and carriage return (\u000D) characters.
	//
	// The policy plain text must be 2048 bytes or shorter. However, an internal
	// conversion compresses it into a packed binary format with a separate limit.
	// The PackedPolicySize response element indicates by percentage how close to
	// the upper size limit the policy is, with 100% equaling the maximum allowed
	// size.
	//
	// For more information about how permissions work, see Permissions for GetFederationToken
	// (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_control-access_getfederationtoken.html).
	Policy *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetFederationTokenInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetFederationTokenInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFederationTokenInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetFederationTokenInput"}
	if s.DurationSeconds != nil && *s.DurationSeconds < 900 {
		invalidParams.Add(request.NewErrParamMinValue("DurationSeconds", 900))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("Name", 2))
	}
	if s.Policy != nil && len(*s.Policy) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Policy", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDurationSeconds sets the DurationSeconds field's value.
func (s *GetFederationTokenInput) SetDurationSeconds(v int64) *GetFederationTokenInput {
	s.DurationSeconds = &v
	return s
}

// SetName sets the Name field's value.
func (s *GetFederationTokenInput) SetName(v string) *GetFederationTokenInput {
	s.Name = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *GetFederationTokenInput) SetPolicy(v string) *GetFederationTokenInput {
	s.Policy = &v
	return s
}

// Contains the response to a successful GetFederationToken request, including
// temporary AWS credentials that can be used to make AWS requests.
type GetFederationTokenOutput struct {
	_ struct{} `type:"structure"`

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token.
	//
	// Note: The size of the security token that STS APIs return is not fixed. We
	// strongly recommend that you make no assumptions about the maximum size. As
	// of this writing, the typical size is less than 4096 bytes, but that can vary.
	// Also, future updates to AWS might require larger sizes.
	Credentials *Credentials `type:"structure"`

	// Identifiers for the federated user associated with the credentials (such
	// as arn:aws:sts::123456789012:federated-user/Bob or 123456789012:Bob). You
	// can use the federated user's ARN in your resource-based policies, such as
	// an Amazon S3 bucket policy.
	FederatedUser *FederatedUser `type:"structure"`

	// A percentage value indicating the size of the policy in packed form. The
	// service rejects policies for which the packed size is greater than 100 percent
	// of the allowed value.
	PackedPolicySize *int64 `type:"integer"`
}

// String returns the string representation
func (s GetFederationTokenOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetFederationTokenOutput) GoString() string {
	return s.String()
}

// SetCredentials sets the Credentials field's value.
func (s *GetFederationTokenOutput) SetCredentials(v *Credentials) *GetFederationTokenOutput {
	s.Credentials = v
	return s
}

// SetFederatedUser sets the FederatedUser field's value.
func (s *GetFederationTokenOutput) SetFederatedUser(v *FederatedUser) *GetFederationTokenOutput {
	s.FederatedUser = v
	return s
}

// SetPackedPolicySize sets the PackedPolicySize field's value.
func (s *GetFederationTokenOutput) SetPackedPolicySize(v int64) *GetFederationTokenOutput {
	s.PackedPolicySize = &v
	return s
}

type GetSessionTokenInput struct {
	_ struct{} `type:"structure"`

	// The duration, in seconds, that the credentials should remain valid. Acceptable
	// durations for IAM user sessions range from 900 seconds (15 minutes) to 129600
	// seconds (36 hours), with 43200 seconds (12 hours) as the default. Sessions
	// for AWS account owners are restricted to a maximum of 3600 seconds (one hour).
	// If the duration is longer than one hour, the session for AWS account owners
	// defaults to one hour.
	DurationSeconds *int64 `min:"900" type:"integer"`

	// The identification number of the MFA device that is associated with the IAM
	// user who is making the GetSessionToken call. Specify this value if the IAM
	// user has a policy that requires MFA authentication. The value is either the
	// serial number for a hardware device (such as GAHT12345678) or an Amazon Resource
	// Name (ARN) for a virtual device (such as arn:aws:iam::123456789012:mfa/user).
	// You can find the device for an IAM user by going to the AWS Management Console
	// and viewing the user's security credentials.
	//
	// The regex used to validated this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can
	// also include underscores or any of the following characters: =,.@:/-
	SerialNumber *string `min:"9" type:"string"`

	// The value provided by the MFA device, if MFA is required. If any policy requires
	// the IAM user to submit an MFA code, specify this value. If MFA authentication
	// is required, and the user does not provide a code when requesting a set of
	// temporary security credentials, the user will receive an "access denied"
	// response when requesting resources that require MFA authentication.
	//
	// The format for this parameter, as described by its regex pattern, is a sequence
	// of six numeric digits.
	TokenCode *string `min:"6" type:"string"`
}

// String returns the string representation
func (s GetSessionTokenInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetSessionTokenInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSessionTokenInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetSessionTokenInput"}
	if s.DurationSeconds != nil && *s.DurationSeconds < 900 {
		invalidParams.Add(request.NewErrParamMinValue("DurationSeconds", 900))
	}
	if s.SerialNumber != nil && len(*s.SerialNumber) < 9 {
		invalidParams.Add(request.NewErrParamMinLen("SerialNumber", 9))
	}
	if s.TokenCode != nil && len(*s.TokenCode) < 6 {
		invalidParams.Add(request.NewErrParamMinLen("TokenCode", 6))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDurationSeconds sets the DurationSeconds field's value.
func (s *GetSessionTokenInput) SetDurationSeconds(v int64) *GetSessionTokenInput {
	s.DurationSeconds = &v
	return s
}

// SetSerialNumber sets the SerialNumber field's value.
func (s *GetSessionTokenInput) SetSerialNumber(v string) *GetSessionTokenInput {
	s.SerialNumber = &v
	return s
}

// SetTokenCode sets the TokenCode field's value.
func (s *GetSessionTokenInput) SetTokenCode(v string) *GetSessionTokenInput {
	s.TokenCode = &v
	return s
}

// Contains the response to a successful GetSessionToken request, including
// temporary AWS credentials that can be used to make AWS requests.
type GetSessionTokenOutput struct {
	_ struct{} `type:"structure"`

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token.
	//
	// Note: The size of the security token that STS APIs return is not fixed. We
	// strongly recommend that you make no assumptions about the maximum size. As
	// of this writing, the typical size is less than 4096 bytes, but that can vary.
	// Also, future updates to AWS might require larger sizes.
	Credentials *Credentials `type:"structure"`
}

// String returns the string representation
func (s GetSessionTokenOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetSessionTokenOutput) GoString() string {
	return s.String()
}

// SetCredentials sets the Credentials field's value.
func (s *GetSessionTokenOutput) SetCredentials(v *Credentials) *GetSessionTokenOutput {
	s.Credentials = v
	return s
}
