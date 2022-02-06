package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/Go-React-Fullstack/inputs"
)

func Body(input interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {

		switch input.(type) {
		case inputs.Login:
			input = &inputs.Login{}
		case inputs.Register:
			input = &inputs.Register{}
		}

		if err := context.BindJSON(&input); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}

		context.Set("input", input)
	}
}
