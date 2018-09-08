package gcpcln

import (
	"context"
	"os"
	"testing"
	"time"
)

func Test_Client_Fetch(t *testing.T) {
	ctx := context.Background()
	cln, err := NewClient(ctx, makeConfig())
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	bktName := os.Getenv("APP_PROJECT_BUCKET")
	prefix := "test"
	dt := time.Second * 10
	// NOTE: tester should add account with id 21 to db
	accountID := 21

	// run method fetch
	cln.Fetch(bktName, prefix, accountID, dt)
}
