package articles

import (
	"articles/core/models"
)

func validateArticleRequest(request models.RequestArticleModel) bool{
	var valid bool  = true

	if request.Title == "" || request.Content==""{
		valid = false
		}

	return valid
}
