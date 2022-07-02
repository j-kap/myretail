package main

import (
	"context"
	"fmt"

	db "github.com/j-kap/myretail/pkg/firestore"
	"github.com/j-kap/myretail/pkg/products"
	"github.com/j-kap/myretail/pkg/redsky"

	"cloud.google.com/go/firestore"
	"github.com/JeremyLoy/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Port        int    `config:"PORT"`
	ProductsURL string `config:"PRODUCTS_URL"`
}

func main() {
	ctx := context.Background()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var c Config
	err := config.FromEnv().To(&c)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config")
	}

	if c.Port == 0 {
		c.Port = 8080
	}

	if c.ProductsURL == "" {
		log.Fatal().Msg("PRODUCTS_URL must be set")
	}

	r := gin.Default()

	projectID := "silken-physics-355018"

	fsClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create client")
	}

	defer fsClient.Close()

	h := products.New(db.New(fsClient), redsky.New(c.ProductsURL))

	products.RegisterRoutes(r, h)

	r.Run(fmt.Sprintf("0.0.0.0:%d", c.Port))
}
