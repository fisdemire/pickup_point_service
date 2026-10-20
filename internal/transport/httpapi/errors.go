package httpapi

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fisdemire/pickup_point_service/internal/domain"
)

var statusByErr = []struct {
	err  error
	code int
}{
	{domain.ErrUnauthorized, http.StatusUnauthorized},
	{domain.ErrInvalidCredentials, http.StatusUnauthorized},
	{domain.ErrForbidden, http.StatusForbidden},
	{domain.ErrEmailTaken, http.StatusBadRequest},
	{domain.ErrPVZNotFound, http.StatusBadRequest},
	{domain.ErrOpenReceptionExists, http.StatusBadRequest},
	{domain.ErrNoOpenReception, http.StatusBadRequest},
	{domain.ErrNoProducts, http.StatusBadRequest},
}

func writeError(c *gin.Context, err error) {
	var ve *domain.ValidationError
	if errors.As(err, &ve) {
		c.JSON(http.StatusBadRequest, errorDTO{Message: ve.Msg})
		return
	}
	for _, m := range statusByErr {
		if errors.Is(err, m.err) {
			c.JSON(m.code, errorDTO{Message: m.err.Error()})
			return
		}
	}
	slog.ErrorContext(c.Request.Context(), "internal error", "err", err, "route", c.FullPath())
	c.JSON(http.StatusInternalServerError, errorDTO{Message: "internal error"})
}

func badBody(c *gin.Context) {
	c.JSON(http.StatusBadRequest, errorDTO{Message: "invalid request body"})
}
