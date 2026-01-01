package handler

import (
	"net/http"
	"workflow-approval-service/model"
	"workflow-approval-service/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestHandler struct {
	service service.RequestService
}

func NewRequestHandler(s service.RequestService) *RequestHandler {
	return &RequestHandler{service: s}
}

// Request body untuk POST /requests
type CreateRequestBody struct {
	WorkflowID string  `json:"workflow_id"`
	Amount     float64 `json:"amount"`
}

// POST /requests
func (h *RequestHandler) Create(c *gin.Context) {
	var body CreateRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	workflowUUID, err := uuid.Parse(body.WorkflowID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid workflow id"})
		return
	}

	req := model.Request{
		WorkflowID: workflowUUID,
		Amount:     body.Amount,
	}

	if err := h.service.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": req})
}

// GET /requests/:id
func (h *RequestHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	req, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": req})
}

// POST /requests/:id/approve
func (h *RequestHandler) Approve(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Approve(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

// POST /requests/:id/reject
func (h *RequestHandler) Reject(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Reject(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
