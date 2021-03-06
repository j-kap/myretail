package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/j-kap/myretail/pkg/redsky"
)

func (h handler) Get(c *gin.Context) {
	id := c.Param("id")

	var product Product

	product.ID = id

	prodResponse, err := h.client.GetProduct(id)
	if err != nil {
		if err == redsky.Err404NotFound {
			log.Info().Str("id", id).Msg("Product not found")
			c.AbortWithStatusJSON(http.StatusNotFound, errResponse(err))
		} else {
			log.Error().Err(err).Str("id", id).Msg("Error getting product")
			c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse(err))
		}

		return
	}

	product.Name = prodResponse.Data.Product.Item.Description.Title

	prodPrice, err := h.DB.GetProductPrice(c, id)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Error getting product price")
		prodPrice.Value = "UNKNOWN"
		prodPrice.Currency = "UNKNOWN"
	}

	product.Price.Value = prodPrice.Value
	product.Price.Currency = prodPrice.Currency

	c.JSON(http.StatusOK, &product)
}
