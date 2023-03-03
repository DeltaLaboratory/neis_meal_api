package neis_meal_api

import (
	"net/http"
	"testing"
)

func TestClient_GetMeal(t *testing.T) {
	client := Client{&http.Client{}}
	res, err := client.GetMeal("B10", "7091435")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", res)
}
