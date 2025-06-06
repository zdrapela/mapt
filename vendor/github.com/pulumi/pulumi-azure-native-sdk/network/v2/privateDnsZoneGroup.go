// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Private dns zone group resource.
//
// Uses Azure REST API version 2023-02-01. In version 1.x of the Azure Native provider, it used API version 2020-11-01.
//
// Other available API versions: 2021-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
type PrivateDnsZoneGroup struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A collection of private dns zone configurations of the private dns zone group.
	PrivateDnsZoneConfigs PrivateDnsZoneConfigResponseArrayOutput `pulumi:"privateDnsZoneConfigs"`
	// The provisioning state of the private dns zone group resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
}

// NewPrivateDnsZoneGroup registers a new resource with the given unique name, arguments, and options.
func NewPrivateDnsZoneGroup(ctx *pulumi.Context,
	name string, args *PrivateDnsZoneGroupArgs, opts ...pulumi.ResourceOption) (*PrivateDnsZoneGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateEndpointName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateEndpointName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20200301:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230901:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20231101:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240101:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240301:PrivateDnsZoneGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240501:PrivateDnsZoneGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateDnsZoneGroup
	err := ctx.RegisterResource("azure-native:network:PrivateDnsZoneGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateDnsZoneGroup gets an existing PrivateDnsZoneGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateDnsZoneGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateDnsZoneGroupState, opts ...pulumi.ResourceOption) (*PrivateDnsZoneGroup, error) {
	var resource PrivateDnsZoneGroup
	err := ctx.ReadResource("azure-native:network:PrivateDnsZoneGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateDnsZoneGroup resources.
type privateDnsZoneGroupState struct {
}

type PrivateDnsZoneGroupState struct {
}

func (PrivateDnsZoneGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDnsZoneGroupState)(nil)).Elem()
}

type privateDnsZoneGroupArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// A collection of private dns zone configurations of the private dns zone group.
	PrivateDnsZoneConfigs []PrivateDnsZoneConfig `pulumi:"privateDnsZoneConfigs"`
	// The name of the private dns zone group.
	PrivateDnsZoneGroupName *string `pulumi:"privateDnsZoneGroupName"`
	// The name of the private endpoint.
	PrivateEndpointName string `pulumi:"privateEndpointName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a PrivateDnsZoneGroup resource.
type PrivateDnsZoneGroupArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// A collection of private dns zone configurations of the private dns zone group.
	PrivateDnsZoneConfigs PrivateDnsZoneConfigArrayInput
	// The name of the private dns zone group.
	PrivateDnsZoneGroupName pulumi.StringPtrInput
	// The name of the private endpoint.
	PrivateEndpointName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (PrivateDnsZoneGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDnsZoneGroupArgs)(nil)).Elem()
}

type PrivateDnsZoneGroupInput interface {
	pulumi.Input

	ToPrivateDnsZoneGroupOutput() PrivateDnsZoneGroupOutput
	ToPrivateDnsZoneGroupOutputWithContext(ctx context.Context) PrivateDnsZoneGroupOutput
}

func (*PrivateDnsZoneGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDnsZoneGroup)(nil)).Elem()
}

func (i *PrivateDnsZoneGroup) ToPrivateDnsZoneGroupOutput() PrivateDnsZoneGroupOutput {
	return i.ToPrivateDnsZoneGroupOutputWithContext(context.Background())
}

func (i *PrivateDnsZoneGroup) ToPrivateDnsZoneGroupOutputWithContext(ctx context.Context) PrivateDnsZoneGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDnsZoneGroupOutput)
}

type PrivateDnsZoneGroupOutput struct{ *pulumi.OutputState }

func (PrivateDnsZoneGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDnsZoneGroup)(nil)).Elem()
}

func (o PrivateDnsZoneGroupOutput) ToPrivateDnsZoneGroupOutput() PrivateDnsZoneGroupOutput {
	return o
}

func (o PrivateDnsZoneGroupOutput) ToPrivateDnsZoneGroupOutputWithContext(ctx context.Context) PrivateDnsZoneGroupOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o PrivateDnsZoneGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDnsZoneGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o PrivateDnsZoneGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateDnsZoneGroup) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A collection of private dns zone configurations of the private dns zone group.
func (o PrivateDnsZoneGroupOutput) PrivateDnsZoneConfigs() PrivateDnsZoneConfigResponseArrayOutput {
	return o.ApplyT(func(v *PrivateDnsZoneGroup) PrivateDnsZoneConfigResponseArrayOutput { return v.PrivateDnsZoneConfigs }).(PrivateDnsZoneConfigResponseArrayOutput)
}

// The provisioning state of the private dns zone group resource.
func (o PrivateDnsZoneGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDnsZoneGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateDnsZoneGroupOutput{})
}
