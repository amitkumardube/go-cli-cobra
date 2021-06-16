package secret

import (
	"context"
	"fmt"
	"io"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"google.golang.org/api/iterator"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// listSecrets lists all secrets in the given project.
func ListSecrets(w io.Writer, parent string) (error) {
	// parent := "projects/my-project"

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
			return fmt.Errorf("failed to create secretmanager client: %v", err)
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.ListSecretsRequest{
			Parent: parent,
	}

	// Call the API.
	it := client.ListSecrets(ctx, req)
	for {
			resp, err := it.Next()
			if err == iterator.Done {
					break
			}

			if err != nil {
					return fmt.Errorf("failed to list secret versions: %v", err)
			}

			fmt.Fprintf(w, "%s\n", resp.Name)
	}

	return nil
}