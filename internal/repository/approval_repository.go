package repository

import (
	"context"

	"github.com/deepayanMallick/approval-crud/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ApprovalRepository interface {
	CreateApproval(ctx context.Context, approval *models.Approval) error
	GetApprovalByID(ctx context.Context, id uuid.UUID) (*models.Approval, error)
	GetApprovalByFlowID(ctx context.Context, flowID uuid.UUID) (*models.Approval, error)
	UpdateApproval(ctx context.Context, approval *models.Approval) error
	DeleteApproval(ctx context.Context, id uuid.UUID) error
	ListApprovals(ctx context.Context) ([]models.Approval, error)
}

type approvalRepo struct {
	db *sqlx.DB
}

func NewApprovalRepository(db *sqlx.DB) ApprovalRepository {
	return &approvalRepo{db: db}
}

func (r *approvalRepo) CreateApproval(ctx context.Context, approval *models.Approval) error {
	query := `
		INSERT INTO approvals (
			flow_id, flow_name, status, created_by, comments
		) VALUES (
			:flow_id, :flow_name, :status, :created_by, :comments
		) RETURNING id, created_at, updated_at
	`

	rows, err := r.db.NamedQueryContext(ctx, query, approval)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(approval)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *approvalRepo) GetApprovalByID(ctx context.Context, id uuid.UUID) (*models.Approval, error) {
	var approval models.Approval
	query := `SELECT * FROM approvals WHERE id = $1`

	err := r.db.GetContext(ctx, &approval, query, id)
	if err != nil {
		return nil, err
	}

	return &approval, nil
}

func (r *approvalRepo) GetApprovalByFlowID(ctx context.Context, flowID uuid.UUID) (*models.Approval, error) {
	var approval models.Approval
	query := `SELECT * FROM approvals WHERE flow_id = $1`

	err := r.db.GetContext(ctx, &approval, query, flowID)
	if err != nil {
		return nil, err
	}

	return &approval, nil
}

func (r *approvalRepo) UpdateApproval(ctx context.Context, approval *models.Approval) error {
	query := `
		UPDATE approvals SET
			flow_name = :flow_name,
			status = :status,
			updated_by = :updated_by,
			comments = :comments,
			updated_at = NOW()
		WHERE id = :id
	`

	_, err := r.db.NamedExecContext(ctx, query, approval)
	return err
}

func (r *approvalRepo) DeleteApproval(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM approvals WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *approvalRepo) ListApprovals(ctx context.Context) ([]models.Approval, error) {
	var approvals []models.Approval
	query := `SELECT * FROM approvals ORDER BY created_at DESC`

	err := r.db.SelectContext(ctx, &approvals, query)
	if err != nil {
		return nil, err
	}

	return approvals, nil
}
