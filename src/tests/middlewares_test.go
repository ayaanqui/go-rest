package routes

import (
	"testing"

	"github.com/ayaanqui/go-rest-server/src/routes"
)

func TestGetBearerToken(t *testing.T) {
	token1, err := routes.GetBearerToken("Bearer my_token")
    if token1 != "my_token" && err == nil {
        t.Errorf("did not parse bearer token properly")
    }

    _, err = routes.GetBearerToken("Bearer")
    if err == nil {
        t.Errorf("did not parse bearer token properly")
    }

    _, err = routes.GetBearerToken("")
    if err == nil {
        t.Errorf("did not parse bearer token properly")
    }
}