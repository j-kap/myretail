package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/j-kap/myretail/pkg/products"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"cloud.google.com/go/firestore"

	db "github.com/j-kap/myretail/pkg/firestore"
)

func main() {
	ctx := context.Background()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	r := gin.Default()

	projectID := "silken-physics-355018"

	fsClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create client")
	}

	defer fsClient.Close()

	h := products.New(db.New(fsClient))

	products.RegisterRoutes(r, h)

	r.Run("0.0.0.0:8080")
}
