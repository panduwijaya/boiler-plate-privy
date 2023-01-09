package dto

import (
	"cake-store/cake-store/internal/consts"
	"cake-store/cake-store/internal/entity"
	"cake-store/cake-store/internal/presentations"
)

func CakeToResponse(src entity.Cake) presentations.CakeDetail {
	x := presentations.CakeDetail{
            ID: src.ID,
            Title: src.Title,
            Description: src.Description,
            Rating: src.Rating,
            Image: src.Image,
            CreatedAt: src.CreatedAt.String(),
            UpdatedAt: src.UpdatedAt.String(),
	}

	if !src.CreatedAt.IsZero() {
		x.CreatedAt = src.CreatedAt.Format(consts.LayoutDateTimeFormat)
	}
	
	//if !src.UpdatedAt.IsZero() {
	//	x.UpdatedAt = src.UpdatedAt.Format(consts.LayoutDateTimeFormat)
	//}

	//if !src.DeletedAt.IsZero() {
	//	x.DeletedAt = src.DeletedAt.Format(consts.LayoutDateTimeFormat)
	//}

	return x
}

func CakesToResponse(inputs []entity.Cake) []presentations.CakeDetail {
	var (
		result = []presentations.CakeDetail{}
	)

	for i := 0; i < len(inputs); i++ {
		result = append(result, CakeToResponse(inputs[i]))
	}

	return result
}