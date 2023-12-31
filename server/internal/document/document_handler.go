package document

import (
	"fmt"
	"net/http"
	"strconv"

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

func (h *Handler) GetDocumentByID(c *gin.Context) {
	documentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("ID Must be an integer")})
		return
	}
	userID := c.GetInt64("userId")
	req := GetDocumentByIDReq{
		UserID:     userID,
		DocumentID: int64(documentID),
	}
	documents, err := h.Service.Getdocument(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, documents)
}

func (h *Handler) UpdateDocument(c *gin.Context) {
	var u UpdateDocumentReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.UserID = c.GetInt64("userId")

	documentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("ID Must be an integer")})
		return
	}
	u.DocumentID = int64(documentID)

	err = h.Service.UpdateDocument(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
func (h *Handler) ShareDocument(c *gin.Context) {
	var u ShareDocumentReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.CurrentUserID = c.GetInt64("userId")

	documentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("ID Must be an integer")})
		return
	}
	u.DocumentID = int64(documentID)

	err = h.Service.ShareDocument(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
