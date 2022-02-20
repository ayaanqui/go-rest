package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ayaanqui/go-rest/src/types"
)

func JsonResponse(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		json.NewEncoder(w).Encode(
			types.Response{
				Message: "Could not parse response",
			},
		)
	}
}