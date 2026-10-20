package httpapi

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/fisdemire/pickup_point_service/internal/auth"
	"github.com/fisdemire/pickup_point_service/internal/domain"
)

type TokenParser interface {
	Parse(token string) (auth.Identity, error)
}

const identityKey = "identity"

func AuthRequired(tp TokenParser) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if len(h) < 8 || !strings.EqualFold(h[:7], "Bearer ") {
			abort(c, http.StatusUnauthorized, domain.ErrUnauthorized)
			return
		}
		id, err := tp.Parse(strings.TrimSpace(h[7:]))
		if err != nil {
			abort(c, http.StatusUnauthorized, domain.ErrUnauthorized)
			return
		}
		c.Set(identityKey, id)
		c.Next()
	}
}

func RequireRole(roles ...domain.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.MustGet(identityKey).(auth.Identity)
		if !slices.Contains(roles, id.Role) {
			abort(c, http.StatusForbidden, domain.ErrForbidden)
			return
		}
		c.Next()
	}
}

func abort(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, errorDTO{Message: err.Error()})
}
