package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	auth AuthService
}

func (h *Handler) dummyLogin(c *gin.Context) {
	var req dummyLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badBody(c)
		return
	}
	token, err := h.auth.DummyLogin(req.Role)
	if err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func (h *Handler) register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badBody(c)
		return
	}
	u, err := h.auth.Register(c.Request.Context(), req.Email, req.Password, req.Role)
	if err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusCreated, toUserDTO(u))
}

func (h *Handler) login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		badBody(c)
		return
	}
	token, err := h.auth.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusOK, token)
}
