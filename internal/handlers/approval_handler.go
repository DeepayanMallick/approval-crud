package handlers

import (
	"net/http"

	"github.com/deepayanMallick/approval-crud/internal/models"
	"github.com/deepayanMallick/approval-crud/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ApprovalHandler struct {
	repo repository.ApprovalRepository
}

func NewApprovalHandler(repo repository.ApprovalRepository) *ApprovalHandler {
	return &ApprovalHandler{repo: repo}
}

func (h *ApprovalHandler) CreateApproval(c *gin.Context) {
	var approval models.Approval
	if err := c.ShouldBindJSON(&approval); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.CreateApproval(c.Request.Context(), &approval); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, approval)
}

func (h *ApprovalHandler) GetApproval(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid approval ID"})
		return
	}

	approval, err := h.repo.GetApprovalByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "approval not found"})
		return
	}

	c.JSON(http.StatusOK, approval)
}

func (h *ApprovalHandler) GetApprovalByFlowID(c *gin.Context) {
	flowID, err := uuid.Parse(c.Param("flow_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flow ID"})
		return
	}

	approval, err := h.repo.GetApprovalByFlowID(c.Request.Context(), flowID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "approval not found"})
		return
	}

	c.JSON(http.StatusOK, approval)
}

func (h *ApprovalHandler) UpdateApproval(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid approval ID"})
		return
	}

	var approval models.Approval
	if err := c.ShouldBindJSON(&approval); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	approval.ID = id

	if err := h.repo.UpdateApproval(c.Request.Context(), &approval); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, approval)
}

func (h *ApprovalHandler) DeleteApproval(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid approval ID"})
		return
	}

	if err := h.repo.DeleteApproval(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "approval deleted successfully"})
}

func (h *ApprovalHandler) ListApprovals(c *gin.Context) {
	approvals, err := h.repo.ListApprovals(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, approvals)
}
