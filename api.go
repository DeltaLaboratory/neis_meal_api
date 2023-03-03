package neis_meal_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	client *http.Client
}

func (c *Client) GetMeal(OFCCode string, SchoolCode string) (*Info, error) {
	res, err := c.client.Get(fmt.Sprintf("https://open.neis.go.kr/hub/mealServiceDietInfo?Type=json&pIndex=1&pSize=5&ATPT_OFCDC_SC_CODE=%s&SD_SCHUL_CODE=%s&MLSV_YMD=%s", OFCCode, SchoolCode, time.Now().Format("20060102")))
	if err != nil {
		return nil, fmt.Errorf("failed to get NEIS api: %w", err)
	}
	result := new(Info)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}
	if err := json.Unmarshal(body, result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	for i := range result.MealServiceDietInfo {
		for j := range result.MealServiceDietInfo[i].Row {
			result.MealServiceDietInfo[i].Row[j].FoodName = formatList(result.MealServiceDietInfo[i].Row[j].FoodName)
			result.MealServiceDietInfo[i].Row[j].FoodOrigin = formatList(result.MealServiceDietInfo[i].Row[j].FoodOrigin)
			result.MealServiceDietInfo[i].Row[j].NeutronInfo = formatList(result.MealServiceDietInfo[i].Row[j].NeutronInfo)
		}
	}
	return result, nil
}

func formatList(list string) string {
	var r []string
	k := strings.Split(list, "<br/>")
	for _, v := range k {
		r = append(r, strings.TrimSpace(v))
	}
	return strings.Join(r, "\n")
}
