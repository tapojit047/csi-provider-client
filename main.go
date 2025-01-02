package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/grpc"
	"sigs.k8s.io/secrets-store-csi-driver/provider/v1alpha1"
)

func main() {
	conn, err := grpc.Dial(
		"unix:///var/lib/csi/test-csi.sock", // Path to the Unix socket
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := v1alpha1.NewCSIDriverProviderClient(conn)

	// Call the Version method
	versionResp, err := client.Version(context.Background(), &v1alpha1.VersionRequest{})
	if err != nil {
		log.Fatalf("Version call failed: %v", err)
	}
	log.Printf("Version Response: %v", versionResp)

	// Serialize Attributes and Secrets to JSON strings
	attributes := map[string]string{
		"key": "value", // Example key-value pair for attributes
	}
	attributesJSON, err := json.Marshal(attributes)
	if err != nil {
		log.Fatalf("Failed to serialize Attributes: %v", err)
	}

	secrets := map[string]string{
		"secretKey": "secretValue", // Example key-value pair for secrets
	}
	secretsJSON, err := json.Marshal(secrets)
	if err != nil {
		log.Fatalf("Failed to serialize Secrets: %v", err)
	}
	// Call the Mount method (example with mock data)
	mountResp, err := client.Mount(context.Background(), &v1alpha1.MountRequest{
		Attributes: string(attributesJSON),
		Secrets:    string(secretsJSON),
		Permission: "644",
		TargetPath: "/tmp/target",
		//MountOptions: []string{"ro"},
	})
	if err != nil {
		log.Fatalf("Mount call failed: %v", err)
	} else {
		if mountResp.ObjectVersion != nil {

			log.Printf("Mount response: ObjectVersion: %v", mountResp.ObjectVersion)
		}
		if mountResp.Files != nil {
			log.Printf("Mount response: Files: %v", mountResp.Files)
		}
		if mountResp.Error != nil {
			log.Printf("Mount response: Error: %v", mountResp.Error)
		}
	}
	time.Sleep(10 * time.Hour)
}
