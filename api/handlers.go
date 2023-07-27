package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// handler is the GET HandlerFunc
func (a *API) handler(c *gin.Context) {
	source := c.Query("source")
	target := c.Query("target")
	amountStr := c.Query("amount")

	amountStr = strings.ReplaceAll(amountStr, "$", "")
	amountStr = strings.ReplaceAll(amountStr, ",", "")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if source == "" || target == "" || amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "No source/target/amount",
		})
		return
	}

	amount, err = a.er.ConvertCurrency(source, target, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    "success",
		"amount": FormatAmount(amount),
	})

}
