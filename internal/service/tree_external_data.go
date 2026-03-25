package service

import (
	"arvore/internal/model"
	"arvore/internal/repository"
)

type TreeExternalService struct {
	repos *repository.TreeExternalRepository
}

func NewTreeExternalService(r *repository.TreeExternalRepository) *TreeExternalService {
	return &TreeExternalService{repos: r}
}

func (s *TreeExternalService) GetTrees() ([]model.TreedData, error) {
	rawTrees, err := s.repos.GetTrees()
	if err != nil {
		return nil, err
	}

	var trees []model.TreedData

	for _, t := range rawTrees {
		tree := model.TreedData{
			Latitude:  t["lat"].(float64),
			Longitude: t["lng"].(float64),
			Species:   safeString(t["species"]),
			Health:    safeString(t["health"]),
		}
		trees = append(trees, tree)
	}

	return trees, nil
}

func safeString(v interface{}) string {
	if v == nil {
		return ""
	}
	return v.(string)
}

func (s *TreeExternalService) GetNearbyTrees(lat, lng, distance string) ([]model.TreedData, error) {
	return s.repos.GetNearbyTrees(lat, lng, distance)
}
