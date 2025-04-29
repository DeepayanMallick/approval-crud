package routes

import (
	"github.com/deepayanMallick/approval-crud/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupApprovalRoutes(router *gin.Engine, approvalHandler *handlers.ApprovalHandler) {
	api := router.Group("/api/v1/approvals")
	{
		api.POST("/", approvalHandler.CreateApproval)
		api.GET("/", approvalHandler.ListApprovals)
		api.GET("/:id", approvalHandler.GetApproval)
		api.GET("/flow/:flow_id", approvalHandler.GetApprovalByFlowID)
		api.PUT("/:id", approvalHandler.UpdateApproval)
		api.DELETE("/:id", approvalHandler.DeleteApproval)
	}
}
