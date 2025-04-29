package services

import (
	awsparams "github.com/redhat-developer/mapt/cmd/mapt/cmd/aws/constants"
	params "github.com/redhat-developer/mapt/cmd/mapt/cmd/constants"
	maptContext "github.com/redhat-developer/mapt/pkg/manager/context"
	awsEKS "github.com/redhat-developer/mapt/pkg/provider/aws/action/eks"
	"github.com/redhat-developer/mapt/pkg/util/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	cmdEKS     = "eks"
	cmdEKSDesc = "eks operations"

	paramVersion                = "version"
	paramVersionDesc            = "EKS K8s cluster version"
	defaultParamVMSize          = "t3.medium"
	paramVMSizeDesc             = "VMSize to be used on the user pool. Typically this is used to provision spot node pools"
	defaultVersion              = "1.31"
	paramScalingDesiredSize     = "workers-desired"
	paramScalingDesiredSizeDesc = "Worker nodes scaling desired size"
	defaultScalingDesiredSize   = "1"
	paramScalingMaxSize         = "workers-max"
	paramScalingMaxSizeDesc     = "Worker nodes scaling maximum size"
	defaultScalingMaxSize       = "3"
	paramScalingMinSize         = "workers-min"
	paramScalingMinSizeDesc     = "Worker nodes scaling minimum size"
	defaultScalingMinSize       = "1"
	paramAddons                 = "addons"
	paramAddonsDesc             = "List of EKS addons to be installed, separated by commas."
	defaultAddons               = ""
)

func GetEKSCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   cmdEKS,
		Short: cmdEKSDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			return nil
		},
	}
	c.AddCommand(getCreateEKS(), getDestroyEKS())
	return c
}

func getCreateEKS() *cobra.Command {
	c := &cobra.Command{
		Use:   params.CreateCmdName,
		Short: params.CreateCmdName,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			if err := awsEKS.Create(
				&maptContext.ContextArgs{
					ProjectName:   viper.GetString(params.ProjectName),
					BackedURL:     viper.GetString(params.BackedURL),
					ResultsOutput: viper.GetString(params.ConnectionDetailsOutput),
					Debug:         viper.IsSet(params.Debug),
					DebugLevel:    viper.GetUint(params.DebugLevel),
					Tags:          viper.GetStringMapString(params.Tags),
				},
				&awsEKS.EKSRequest{
					Prefix:             viper.GetString(params.ProjectName),
					Location:           viper.GetString(awsparams.ParamLocation),
					VMSize:             viper.GetString(awsparams.ParamVMSize),
					KubernetesVersion:  viper.GetString(paramVersion),
					ScalingDesiredSize: viper.GetInt(paramScalingDesiredSize),
					ScalingMaxSize:     viper.GetInt(paramScalingMaxSize),
					ScalingMinSize:     viper.GetInt(paramScalingMinSize),
					Addons:             viper.GetStringSlice(paramAddons),
				}); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
	flagSet := pflag.NewFlagSet(params.CreateCmdName, pflag.ExitOnError)
	flagSet.StringP(params.ConnectionDetailsOutput, "", "", params.ConnectionDetailsOutputDesc)
	flagSet.StringToStringP(params.Tags, "", nil, params.TagsDesc)
	flagSet.StringP(awsparams.ParamLocation, "", "", awsparams.ParamLocationDesc)
	flagSet.StringP(awsparams.ParamVMSize, "", defaultParamVMSize, paramVMSizeDesc)
	flagSet.StringP(paramVersion, "", defaultVersion, paramVersionDesc)
	flagSet.StringP(paramScalingDesiredSize, "", defaultScalingDesiredSize, paramScalingDesiredSizeDesc)
	flagSet.StringP(paramScalingMaxSize, "", defaultScalingMaxSize, paramScalingMaxSizeDesc)
	flagSet.StringP(paramScalingMinSize, "", defaultScalingMinSize, paramScalingMinSizeDesc)
	flagSet.StringSliceP(paramAddons, "", []string{}, paramAddonsDesc)
	c.PersistentFlags().AddFlagSet(flagSet)
	return c
}

func getDestroyEKS() *cobra.Command {
	return &cobra.Command{
		Use:   params.DestroyCmdName,
		Short: params.DestroyCmdName,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}
			if err := awsEKS.Destroy(
				&maptContext.ContextArgs{
					ProjectName: viper.GetString(params.ProjectName),
					BackedURL:   viper.GetString(params.BackedURL),
					Debug:       viper.IsSet(params.Debug),
					DebugLevel:  viper.GetUint(params.DebugLevel),
				}); err != nil {
				logging.Error(err)
			}
			return nil
		},
	}
}
