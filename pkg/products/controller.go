package products

import (
	"github.com/gin-gonic/gin"

	"github.com/j-kap/myretail/pkg/firestore"
	"github.com/j-kap/myretail/pkg/redsky"
)

type handler struct {
	DB     firestore.Client
	client redsky.Client
}

func New(db firestore.Client, client redsky.Client) *handler {
	return &handler{
		DB:     db,
		client: client,
	}
}

func RegisterRoutes(r *gin.Engine, h *handler) {
	routes := r.Group("/products")

	routes.GET("/:id", h.Get)
	routes.PUT("/:id", h.UpdatePrice)
}
