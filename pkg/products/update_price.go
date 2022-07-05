package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h handler) UpdatePrice(c *gin.Context) {
	id := c.Param("id")

	var body Product
	if err := c.BindJSON(&body); err != nil {
		log.Error().Err(err).Str("id", id).Msg("Error parsing product")
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := h.DB.SetProductPrice(c, id, body.Price.Value, body.Price.Currency)
	if err != nil {
		log.Error().Err(err).Str("id", id).Msg("Error setting product price")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	c.JSON(http.StatusOK, body)
}
