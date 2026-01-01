package handler

import (
	"net/http"
	"workflow-approval-service/model"
	"workflow-approval-service/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WorkflowStepHandler struct {
	stepService service.WorkflowStepService
}

func NewWorkflowStepHandler(ss service.WorkflowStepService) *WorkflowStepHandler {
	return &WorkflowStepHandler{stepService: ss}
}

// POST /workflows/:id/steps
func (h *WorkflowStepHandler) CreateStep(c *gin.Context) {
	workflowID := c.Param("id")
	var req model.WorkflowStep
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	wfID, err := uuid.Parse(workflowID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid workflow id"})
		return
	}
	req.WorkflowID = wfID

	if err := h.stepService.Create(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": req})
}

// GET /workflows/:id/steps
func (h *WorkflowStepHandler) GetSteps(c *gin.Context) {
	workflowID := c.Param("id")
	steps, err := h.stepService.GetByWorkflowID(workflowID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": steps})
}
