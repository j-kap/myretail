package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/j-kap/myretail/pkg/products"

	"cloud.google.com/go/firestore"

	db "github.com/j-kap/myretail/pkg/firestore"
)

func main() {
	ctx := context.Background()

	r := gin.Default()

	projectID := "silken-physics-355018"

	fsClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	defer fsClient.Close()

	products.RegisterRoutes(r, db.New(fsClient))

	r.Run("localhost:8080")
}
