package router

import (
	"workflow-approval-service/handler"

	"github.com/gin-gonic/gin"
)

func RegisterWorkflowRoutes(r *gin.Engine, workflowHandler *handler.WorkflowHandler) {
	r.POST("/workflows", workflowHandler.Create)
	r.GET("/workflows/:id", workflowHandler.GetByID)
	r.GET("/workflows", workflowHandler.GetAll)
}

func RegisterWorkflowStepRoutes(r *gin.Engine, workflowStepHandler *handler.WorkflowStepHandler) {
	r.POST("/workflows/:id/steps", workflowStepHandler.CreateStep)
	r.GET("/workflows/:id/steps", workflowStepHandler.GetSteps)
}

func RegisterRequestRoutes(r *gin.Engine, requestHandler *handler.RequestHandler) {
	r.POST("/requests", requestHandler.Create)
	r.GET("/requests/:id", requestHandler.GetByID)
	r.POST("/requests/:id/approve", requestHandler.Approve)
	r.POST("/requests/:id/reject", requestHandler.Reject)
}
