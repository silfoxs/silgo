package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silfoxs/silgo/internal/pkg/code"
)

func response(context *gin.Context, code int, res map[string]any, contextType string) {
	response := gin.H{
		"code": 0,
		"msg":  "",
		"data": map[string]any{},
	}

	for index, value := range res {
		response[index] = value
	}
	write(context, contextType, code, response)
}

func write(context *gin.Context, contextType string, code int, res gin.H) {
	switch contextType {
	case "IndentedJSON":
		context.IndentedJSON(code, res)
	case "SecureJSON":
		context.SecureJSON(code, res)
	case "JSON":
		context.JSON(code, res)
	case "AsciiJSON":
		context.AsciiJSON(code, res)
	case "PureJSON":
		context.PureJSON(code, res)
	case "XML":
		context.XML(code, res)
	case "YAML":
		context.YAML(code, res)
	case "ProtoBuf":
		context.ProtoBuf(code, res)
	}
}

func Success(context *gin.Context, data map[string]any) {
	if _, ok := data["code"]; !ok {
		data["code"] = code.Success
	}
	response(context, http.StatusOK, data, "JSON")
}

func Fail(context *gin.Context, data map[string]any) {
	if _, ok := data["code"]; !ok {
		data["code"] = code.Fail
	}
	response(context, http.StatusOK, data, "JSON")
}
