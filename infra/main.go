package main

import (
	"io/ioutil"
	"log"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	machineType := "t2d-standard-4"
	sshPubKey := "user:ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIII6AeuI/N5EUIL4+HqK8YmpoYyTwQvSFUSl6ev6Cg+Q user@work-vm-big"
	bootImage := "debian-cloud/debian-11"
	// bootImage := "projects/debian-cloud/global/images/debian-11-bullseye-arm64-v20230509" // Use for ARM instances

	pulumi.Run(func(ctx *pulumi.Context) error {
		// Prepare the startup script
		startupScript, err := ioutil.ReadFile("startup.sh")
		if err != nil {
			log.Fatalf("Failed to read startup script: %v", err)
		}

		// Configure network interface with ephemeral external IP
		networkInterface := compute.InstanceNetworkInterfaceArgs{
			Network: pulumi.String("default"),
			AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
				&compute.InstanceNetworkInterfaceAccessConfigArgs{},
			},
		}

		// Create a new instance
		instance, err := compute.NewInstance(ctx, "redisbench-"+machineType, &compute.InstanceArgs{
			Zone:        pulumi.String("us-central1-a"),
			MachineType: pulumi.String(machineType),
			BootDisk: &compute.InstanceBootDiskArgs{
				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
					Image: pulumi.String(bootImage),
				},
			},
			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{&networkInterface},
			Metadata: pulumi.StringMap{
				"ssh-keys": pulumi.String(sshPubKey),
			},
			MetadataStartupScript: pulumi.String(startupScript),
		})
		if err != nil {
			return err
		}

		// Export the instance's ephemeral external IP address
		ctx.Export("instanceExternalIp", instance.NetworkInterfaces.Index(pulumi.Int(0)).AccessConfigs().ApplyT(func(accessConfigs []compute.InstanceNetworkInterfaceAccessConfig) string {
			// Find the first access config that has a NAT IP address assigned
			for _, accessConfig := range accessConfigs {
				if accessConfig.NatIp != nil {
					return *accessConfig.NatIp
				}
			}
			return ""
		}))

		return nil
	})
}
