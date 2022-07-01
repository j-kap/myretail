package products

import (
	"github.com/gin-gonic/gin"

	"github.com/j-kap/myretail/pkg/firestore"
)

type handler struct {
	DB firestore.Client
}

func RegisterRoutes(r *gin.Engine, db firestore.Client) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/products")

	routes.GET("/:id", h.Get)
	routes.PUT("/:id", h.UpdatePrice)
}
