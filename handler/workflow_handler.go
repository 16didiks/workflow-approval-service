package handler

import (
	"net/http"
	"workflow-approval-service/model"
	"workflow-approval-service/service"

	"github.com/gin-gonic/gin"
)

type WorkflowHandler struct {
	service service.WorkflowService
}

func NewWorkflowHandler(s service.WorkflowService) *WorkflowHandler {
	return &WorkflowHandler{service: s}
}

// POST /workflows
func (h *WorkflowHandler) Create(c *gin.Context) {
	var req model.Workflow
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	if err := h.service.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": req})
}

// GET /workflows/:id
func (h *WorkflowHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	wf, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": wf})
}

// GET /workflows
func (h *WorkflowHandler) GetAll(c *gin.Context) {
	wfs, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": wfs})
}
