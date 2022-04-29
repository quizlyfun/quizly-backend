package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/answersuck/vault/internal/domain"

	"github.com/answersuck/vault/pkg/logging"
)

type tagService interface {
	GetAll(ctx context.Context) ([]*domain.Tag, error)
}

type tagHandler struct {
	log     logging.Logger
	service tagService
}

func newTagHandler(handler *gin.RouterGroup, d *Deps) {
	h := &tagHandler{
		log:     d.Logger,
		service: d.TagService,
	}

	g := handler.Group("tags")
	{
		g.GET("", h.getAll)
	}
}

func (h *tagHandler) getAll(c *gin.Context) {
	t, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		h.log.Error(fmt.Errorf("http - v1 - tag - getAll - h.service.GetAll: %w", err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, t)
}