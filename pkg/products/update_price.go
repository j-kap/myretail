package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) UpdatePrice(c *gin.Context) {
	id := c.Param("id")

	var body Product
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := h.DB.SetProductPrice(c, id, body.Price.Value, body.Price.Currency)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, body)
}
