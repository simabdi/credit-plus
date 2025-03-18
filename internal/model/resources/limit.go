package resource

import (
	"credit-plus/internal/model/entity"
	"credit-plus/internal/model/formatter"
)

func LimitResource(limit entity.Limit) formatter.LimitFormatter {
	Resource := formatter.LimitFormatter{
		Uuid:   limit.Uuid,
		Tenor:  limit.Tenor,
		Amount: limit.Amount,
	}

	return Resource
}

func LimitCollectionResource(limits []entity.Limit) []formatter.LimitFormatter {
	resourceCollection := []formatter.LimitFormatter{}

	for _, limit := range limits {
		format := LimitResource(limit)
		resourceCollection = append(resourceCollection, format)
	}

	return resourceCollection
}
