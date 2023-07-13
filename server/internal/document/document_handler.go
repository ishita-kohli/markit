package document

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateDocument(c *gin.Context) {
	var u CreateDocumentReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.UserID = c.GetInt64("userId")

	res, err := h.Service.CreateDocument(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Listdocuments(c *gin.Context) {
	role, ok := c.GetQuery("role")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf(" invalid role query")})
		return
	}
	userID := c.GetInt64("userId")
	req := DocumentlistReq{
		UserID: userID,
		Role:   role,
	}
	documents, err := h.Service.Listdocuments(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, documents)
}
