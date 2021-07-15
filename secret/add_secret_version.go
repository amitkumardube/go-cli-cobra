package secret

import (
	"context"
	"fmt"
	"io"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"github.com/amitkumardube/go-cli-cobra/common"
)

// addSecretVersion adds a new secret version to the given secret with the
// provided payload.
func AddSecretVersion(w io.Writer, parent string) error {
	// parent := "projects/my-project/secrets/my-secret"

	// Get the payload to store. We are generate a random number and store as payload.

	secret_value , err := common.GenerateRandomString(20)

	if err != nil {
		return err
	}

	payload := []byte(secret_value)

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
			return fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.AddSecretVersionRequest{
			Parent: parent,
			Payload: &secretmanagerpb.SecretPayload{
					Data: payload,
			},
	}

	// Call the API.
	result, err := client.AddSecretVersion(ctx, req)
	if err != nil {
			return fmt.Errorf("failed to add secret version: %v", err)
	}
	fmt.Fprintf(w, "Added secret version: %s\n", result.Name)
	return nil
}
