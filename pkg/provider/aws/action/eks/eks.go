// Just copied from AKS

package eks

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	// "github.com/pulumi/pulumi-eks/sdk/v3/go/eks" // This version allows to define Auto Mode
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/redhat-developer/mapt/pkg/manager"
	maptContext "github.com/redhat-developer/mapt/pkg/manager/context"
	"github.com/redhat-developer/mapt/pkg/provider/aws"
	"github.com/redhat-developer/mapt/pkg/provider/util/output"
	"github.com/redhat-developer/mapt/pkg/util/logging"
)

type EKSRequest struct {
	Prefix   string
	Location string
	VMSize   string
	// "1.26.3"
	KubernetesVersion string
	// OnlySystemPool    bool
	// Spot              bool
	// SpotTolerance     spotAzure.EvictionRate
}

const autoMode bool = true

func Create(ctx *maptContext.ContextArgs, r *EKSRequest) (err error) {
	logging.Debug("Creating EKS")
	if err := maptContext.Init(ctx, aws.Provider()); err != nil {
		return err
	}
	cs := manager.Stack{
		StackName:           maptContext.StackNameByProject(stackName),
		ProjectName:         maptContext.ProjectName(),
		BackedURL:           maptContext.BackedURL(),
		ProviderCredentials: aws.DefaultCredentials,
		DeployFunc:          r.deployer,
	}
	sr, _ := manager.UpStack(cs)
	return r.manageResults(sr)
}

func Destroy(ctx *maptContext.ContextArgs) error {
	// Create mapt Context
	logging.Debug("Destroy EKS")
	if err := maptContext.Init(ctx, aws.Provider()); err != nil {
		return err
	}
	return aws.DestroyStack(
		aws.DestroyStackRequest{
			BackedURL: maptContext.BackedURL(),
			Stackname: stackName,
	})
}

// Main function to deploy all requried resources to AWS
func (r *EKSRequest) deployer(ctx *pulumi.Context) error {
	// Read back the default VPC and public subnets, which we will use.
	t := true
	vpc, err := ec2.LookupVpc(ctx, &ec2.LookupVpcArgs{Default: &t})
	if err != nil {
		return err
	}

	subnet, err := ec2.GetSubnets(ctx, &ec2.GetSubnetsArgs{
		Filters: []ec2.GetSubnetsFilter{
			{Name: "vpc-id", Values: []string{vpc.Id}},
		},
	})
	if err != nil {
		return err
	}
	eksRole, err := iam.NewRole(ctx, "eks-iam-eksRole", &iam.RoleArgs{
		AssumeRolePolicy: pulumi.String(`{
			"Version": "2008-10-17",
			"Statement": [{
					"Sid": "",
					"Effect": "Allow",
					"Principal": {
							"Service": "eks.amazonaws.com"
					},
					"Action": "sts:AssumeRole"
			}]
	}`),
	})
	if err != nil {
		return err
	}
	eksPolicies := []string{
		"arn:aws:iam::aws:policy/AmazonEKSServicePolicy",
		"arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
	}
	for i, eksPolicy := range eksPolicies {
		_, err := iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("rpa-%d", i), &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String(eksPolicy),
			Role:      eksRole.Name,
		})
		if err != nil {
			return err
		}
	}
	// Create the EC2 NodeGroup Role
	nodeGroupRole, err := iam.NewRole(ctx, "nodegroup-iam-role", &iam.RoleArgs{
		AssumeRolePolicy: pulumi.String(`{
			"Version": "2012-10-17",
			"Statement": [{
					"Sid": "",
					"Effect": "Allow",
					"Principal": {
							"Service": "ec2.amazonaws.com"
					},
					"Action": "sts:AssumeRole"
			}]
	}`),
	})
	if err != nil {
		return err
	}
	nodeGroupPolicies := []string{
		"arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
		"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
		"arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
	}
	for i, nodeGroupPolicy := range nodeGroupPolicies {
		_, err := iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("ngpa-%d", i), &iam.RolePolicyAttachmentArgs{
			Role:      nodeGroupRole.Name,
			PolicyArn: pulumi.String(nodeGroupPolicy),
		})
		if err != nil {
			return err
		}
	}
	// Create a Security Group that we can use to actually connect to our cluster
	clusterSg, err := ec2.NewSecurityGroup(ctx, "cluster-sg", &ec2.SecurityGroupArgs{
		VpcId: pulumi.String(vpc.Id),
		Egress: ec2.SecurityGroupEgressArray{
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("-1"),
				FromPort:   pulumi.Int(0),
				ToPort:     pulumi.Int(0),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
		},
		Ingress: ec2.SecurityGroupIngressArray{
			ec2.SecurityGroupIngressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(80),
				ToPort:     pulumi.Int(80),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
		},
	})
	if err != nil {
		return err
	}
	// Create EKS Cluster
	eksCluster, err := eks.NewCluster(ctx, "eks-cluster", &eks.ClusterArgs{
		RoleArn: pulumi.StringInput(eksRole.Arn),
		VpcConfig: &eks.ClusterVpcConfigArgs{
			PublicAccessCidrs: pulumi.StringArray{
				pulumi.String("0.0.0.0/0"),
			},
			SecurityGroupIds: pulumi.StringArray{
				clusterSg.ID().ToStringOutput(),
			},
			SubnetIds: toPulumiStringArray(subnet.Ids),
		},
		Version: pulumi.String(r.KubernetesVersion),
		// Set the Auto mode
		ComputeConfig: &eks.ClusterComputeConfigArgs{
			Enabled: pulumi.Bool(autoMode),
		},
		KubernetesNetworkConfig: &eks.ClusterKubernetesNetworkConfigArgs{
			ElasticLoadBalancing: &eks.ClusterKubernetesNetworkConfigElasticLoadBalancingArgs{
				Enabled: pulumi.Bool(autoMode),
			},
		},
		StorageConfig: &eks.ClusterStorageConfigArgs{
			BlockStorage: &eks.ClusterStorageConfigBlockStorageArgs{
				Enabled: pulumi.Bool(autoMode),
			},
		},
		BootstrapSelfManagedAddons: pulumi.BoolPtr(!autoMode),
	})
	if err != nil {
		return err
	}

	if(autoMode==false) {
		_, err = eks.NewNodeGroup(ctx, "node-group-2", &eks.NodeGroupArgs{
			ClusterName:   eksCluster.Name,
			NodeGroupName: pulumi.String("demo-eks-nodegroup-2"),
			NodeRoleArn:   pulumi.StringInput(nodeGroupRole.Arn),
			SubnetIds:     toPulumiStringArray(subnet.Ids),
			InstanceTypes: pulumi.StringArray{pulumi.String(r.VMSize)},
			ScalingConfig: &eks.NodeGroupScalingConfigArgs{
				// TODO: Hardcoded, but should be configurable
				DesiredSize: pulumi.Int(2),
				MaxSize:     pulumi.Int(5),
				MinSize:     pulumi.Int(0),
			},
		})
		if err != nil {
			return err
		}
	}

	kubeconfig := generateKubeconfig(eksCluster.Endpoint, eksCluster.CertificateAuthority.Data().Elem(), eksCluster.Name)

	ctx.Export("kubeconfig", kubeconfig)
	return nil
}

// Create the KubeConfig Structure as per https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html
func generateKubeconfig(clusterEndpoint pulumi.StringOutput, certData pulumi.StringOutput, clusterName pulumi.StringOutput) pulumi.StringOutput {
	return pulumi.Sprintf(`{
        "apiVersion": "v1",
        "clusters": [{
            "cluster": {
                "server": "%s",
                "certificate-authority-data": "%s"
            },
            "name": "kubernetes",
        }],
        "contexts": [{
            "context": {
                "cluster": "kubernetes",
                "user": "aws",
            },
            "name": "aws",
        }],
        "current-context": "aws",
        "kind": "Config",
        "users": [{
            "name": "aws",
            "user": {
                "exec": {
                    "apiVersion": "client.authentication.k8s.io/v1beta1",
                    "command": "aws",
                    "args": [
                        "eks",
                        "get-token",
                        "--cluster-name",
                        "%s",
                    ],
                },
            },
        }],
    }`, clusterEndpoint, certData, clusterName)
}

func toPulumiStringArray(a []string) pulumi.StringArrayInput {
	var res []pulumi.StringInput
	for _, s := range a {
		res = append(res, pulumi.String(s))
	}
	return pulumi.StringArray(res)
}

// MAPT AKS Kubeconfig
// Write exported values in context to files o a selected target folder
func (r *EKSRequest) manageResults(stackResult auto.UpResult) error {
	return output.Write(stackResult, maptContext.GetResultsOutputPath(), map[string]string{
		fmt.Sprintf("%s-%s", r.Prefix, outputKubeconfig): "kubeconfig",
	})
}
