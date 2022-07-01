package products

import (
	"github.com/gin-gonic/gin"

	"github.com/j-kap/myretail/pkg/firestore"
)

type handler struct {
	DB firestore.Client
}

func New(db firestore.Client) *handler {
	return &handler{
		DB: db,
	}
}

func RegisterRoutes(r *gin.Engine, h *handler) {
	routes := r.Group("/products")

	routes.GET("/:id", h.Get)
	routes.PUT("/:id", h.UpdatePrice)
}
