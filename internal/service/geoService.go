package service

import (
	"arvore/internal/model"
	"arvore/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

func NewGeoJsonTreeervice(
	treeRepo repository.TreeRepositoryInterface,
	riskRepo repository.RiskAssessmentRepositoryInterface,
) *TreeService {
	return &TreeService{
		treeRepo: treeRepo,
		riskRepo: riskRepo,
	}
}
func (s *TreeService) ImportGeoJSON(ctx context.Context) error {
	data, err := os.ReadFile("internal/data/arvores.geojson")
	if err != nil {
		return err
	}

	var geo model.GeoJSON
	if err := json.Unmarshal(data, &geo); err != nil {
		return err
	}

	for _, f := range geo.Features {

		// ⚠️ longitude vem primeiro!
		lng := f.Geometry.Coordinates[0]
		lat := f.Geometry.Coordinates[1]

		tree := model.Tree{
			Latitude:  lat,
			Longitude: lng,
			Especie:   f.Properties.Especie,
		}

		_, err := s.treeRepo.Create(ctx, tree)
		if err != nil {
			fmt.Println("erro ao inserir:", err)
			continue
		}
	}

	return nil
}
