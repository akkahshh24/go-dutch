package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/akkahshh24/go-dutch/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "AuthorizationPayload"
)

var (
	ErrNoAuthorizationHeader        = errors.New("authorization header is required")
	ErrInvalidToken                 = errors.New("invalid token")
	ErrUnsupportedAuthorizationType = errors.New("unsupported authorization type")
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(authorizationHeaderKey)
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(ErrNoAuthorizationHeader))
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(ErrInvalidToken))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(ErrUnsupportedAuthorizationType))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}
