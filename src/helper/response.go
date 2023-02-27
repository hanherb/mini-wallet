package helper

import "github.com/gin-gonic/gin"

func ResponseSuccess(data interface{}) gin.H {
	return gin.H{
		"status": "success",
		"data":   data,
	}
}

func ResponseBadRequest(err string) gin.H {
	return gin.H{
		"status": "bad request",
		"data": gin.H{
			"error": err,
		},
	}
}

func ResponseUnauthorized(err string) gin.H {
	return gin.H{
		"status": "unauthorized",
		"data": gin.H{
			"error": err,
		},
	}
}

func ResponseForbidden(err string) gin.H {
	return gin.H{
		"status": "forbidden",
		"data": gin.H{
			"error": err,
		},
	}
}

func ResponseNotFound(err string) gin.H {
	return gin.H{
		"status": "not found",
		"data": gin.H{
			"error": err,
		},
	}
}

func ResponseISE(err string) gin.H {
	return gin.H{
		"status": "internal server error",
		"data": gin.H{
			"error": err,
		},
	}
}
