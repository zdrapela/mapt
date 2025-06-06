// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets properties of a DNS resolver policy.
//
// Uses Azure REST API version 2023-07-01-preview.
func LookupDnsResolverPolicy(ctx *pulumi.Context, args *LookupDnsResolverPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDnsResolverPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDnsResolverPolicyResult
	err := ctx.Invoke("azure-native:network:getDnsResolverPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDnsResolverPolicyArgs struct {
	// The name of the DNS resolver policy.
	DnsResolverPolicyName string `pulumi:"dnsResolverPolicyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes a DNS resolver policy.
type LookupDnsResolverPolicyResult struct {
	// ETag of the DNS resolver policy.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The current provisioning state of the DNS resolver policy. This is a read-only property and any attempt to set this value will be ignored.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resourceGuid property of the DNS resolver policy resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDnsResolverPolicyOutput(ctx *pulumi.Context, args LookupDnsResolverPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDnsResolverPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDnsResolverPolicyResultOutput, error) {
			args := v.(LookupDnsResolverPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getDnsResolverPolicy", args, LookupDnsResolverPolicyResultOutput{}, options).(LookupDnsResolverPolicyResultOutput), nil
		}).(LookupDnsResolverPolicyResultOutput)
}

type LookupDnsResolverPolicyOutputArgs struct {
	// The name of the DNS resolver policy.
	DnsResolverPolicyName pulumi.StringInput `pulumi:"dnsResolverPolicyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDnsResolverPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsResolverPolicyArgs)(nil)).Elem()
}

// Describes a DNS resolver policy.
type LookupDnsResolverPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDnsResolverPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsResolverPolicyResult)(nil)).Elem()
}

func (o LookupDnsResolverPolicyResultOutput) ToLookupDnsResolverPolicyResultOutput() LookupDnsResolverPolicyResultOutput {
	return o
}

func (o LookupDnsResolverPolicyResultOutput) ToLookupDnsResolverPolicyResultOutputWithContext(ctx context.Context) LookupDnsResolverPolicyResultOutput {
	return o
}

// ETag of the DNS resolver policy.
func (o LookupDnsResolverPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupDnsResolverPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupDnsResolverPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDnsResolverPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of the DNS resolver policy. This is a read-only property and any attempt to set this value will be ignored.
func (o LookupDnsResolverPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resourceGuid property of the DNS resolver policy resource.
func (o LookupDnsResolverPolicyResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverPolicyResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDnsResolverPolicyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDnsResolverPolicyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDnsResolverPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDnsResolverPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDnsResolverPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDnsResolverPolicyResultOutput{})
}
