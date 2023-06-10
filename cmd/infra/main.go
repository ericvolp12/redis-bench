package main

import (
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Prepare the startup script
		startupScript, err := ioutil.ReadFile("startup.sh")
		if err != nil {
			log.Fatalf("Failed to read startup script: %v", err)
		}

		// Create a new instance
		_, err = compute.NewInstance(ctx, "instance", &compute.InstanceArgs{
			MachineType: pulumi.String("n1-standard-1"),
			BootDisk: &compute.InstanceBootDiskArgs{
				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
					Image: pulumi.String("debian-cloud/debian-9"),
				},
			},
			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
				&compute.InstanceNetworkInterfaceArgs{
					Network: pulumi.String("default"),
				},
			},
			MetadataStartupScript: pulumi.String(startupScript),
		})
		if err != nil {
			return err
		}

		// SSH into the instance and run the command
		go func() {
			cmd := exec.Command("ssh", "user@instanceIP", "/usr/local/go/bin/go run /home/user/bsky-experiments/cmd/redis-bench/main.go")
			err := cmd.Run()
			if err != nil {
				log.Fatalf("Failed to run command: %v", err)
			}

			// Copy the result file from the instance
			cmd = exec.Command("scp", "user@instanceIP:/home/user/results.txt", "localpath/to/save/results.txt")
			err = cmd.Run()
			if err != nil {
				log.Fatalf("Failed to copy results file: %v", err)
			}
		}()

		return nil
	})
}
