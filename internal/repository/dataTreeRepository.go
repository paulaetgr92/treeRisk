package repository

import (
	db "arvore/db/sqlc"
	"context"
)

type DataTreeRepository struct {
	*BaseRepository
}

func NewDataTreeRepository(base *BaseRepository) *DataTreeRepository {
	return &DataTreeRepository{BaseRepository: base}
}

func (r *RiskAssesmentNewRepository) CreateDataTreeRepository(ctx context.Context, arg db.CreateTreeParams) (db.Tree, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.Tree{}, err
	}
	return r.Queries.CreateTree(ctx, arg)
}

func (r *RiskAssesmentNewRepository) GetDataRepository(ctx context.Context, id int64) (db.RiskAssessment, error) {
	err := r.GetConnection(ctx)
	if err != nil {
		return db.RiskAssessment{}, err
	}
	return r.Queries.GetRiskByID(ctx, id)
}
