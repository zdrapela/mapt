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

// Network watcher in a resource group.
//
// Uses Azure REST API version 2023-02-01. In version 1.x of the Azure Native provider, it used API version 2020-11-01.
//
// Other available API versions: 2022-05-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-05-01.
type NetworkWatcher struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the network watcher resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkWatcher registers a new resource with the given unique name, arguments, and options.
func NewNetworkWatcher(ctx *pulumi.Context,
	name string, args *NetworkWatcherArgs, opts ...pulumi.ResourceOption) (*NetworkWatcher, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20160901:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230901:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20231101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240101:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240301:NetworkWatcher"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240501:NetworkWatcher"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkWatcher
	err := ctx.RegisterResource("azure-native:network:NetworkWatcher", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkWatcher gets an existing NetworkWatcher resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkWatcher(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkWatcherState, opts ...pulumi.ResourceOption) (*NetworkWatcher, error) {
	var resource NetworkWatcher
	err := ctx.ReadResource("azure-native:network:NetworkWatcher", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkWatcher resources.
type networkWatcherState struct {
}

type NetworkWatcherState struct {
}

func (NetworkWatcherState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherState)(nil)).Elem()
}

type networkWatcherArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the network watcher.
	NetworkWatcherName *string `pulumi:"networkWatcherName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkWatcher resource.
type NetworkWatcherArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the network watcher.
	NetworkWatcherName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkWatcherArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherArgs)(nil)).Elem()
}

type NetworkWatcherInput interface {
	pulumi.Input

	ToNetworkWatcherOutput() NetworkWatcherOutput
	ToNetworkWatcherOutputWithContext(ctx context.Context) NetworkWatcherOutput
}

func (*NetworkWatcher) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkWatcher)(nil)).Elem()
}

func (i *NetworkWatcher) ToNetworkWatcherOutput() NetworkWatcherOutput {
	return i.ToNetworkWatcherOutputWithContext(context.Background())
}

func (i *NetworkWatcher) ToNetworkWatcherOutputWithContext(ctx context.Context) NetworkWatcherOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkWatcherOutput)
}

type NetworkWatcherOutput struct{ *pulumi.OutputState }

func (NetworkWatcherOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkWatcher)(nil)).Elem()
}

func (o NetworkWatcherOutput) ToNetworkWatcherOutput() NetworkWatcherOutput {
	return o
}

func (o NetworkWatcherOutput) ToNetworkWatcherOutputWithContext(ctx context.Context) NetworkWatcherOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o NetworkWatcherOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkWatcher) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location.
func (o NetworkWatcherOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkWatcher) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o NetworkWatcherOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkWatcher) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the network watcher resource.
func (o NetworkWatcherOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkWatcher) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o NetworkWatcherOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkWatcher) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o NetworkWatcherOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkWatcher) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkWatcherOutput{})
}
