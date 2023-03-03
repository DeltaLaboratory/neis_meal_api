package neis_meal_api

/*
1	    ATPT_OFCDC_SC_CODE	    시도교육청코드
2	    ATPT_OFCDC_SC_NM	    시도교육청명
3	    SD_SCHUL_CODE	    표준학교코드
4	    SCHUL_NM	    학교명
5	    MMEAL_SC_CODE	    식사코드
6	    MMEAL_SC_NM	    식사명
7	    MLSV_YMD	    급식일자
8	    MLSV_FGR	    급식인원수
9	    DDISH_NM	    요리명
10	    ORPLC_INFO	    원산지정보
11	    CAL_INFO	    칼로리정보
12	    NTR_INFO	    영양정보
13	    MLSV_FROM_YMD	    급식시작일자
14	    MLSV_TO_YMD	    급식종료일자
*/

type Info struct {
	MealServiceDietInfo []struct {
		Head []struct {
			ListTotalCount int `json:"list_total_count,omitempty"`
			Result         struct {
				Code    string `json:"CODE"`
				Message string `json:"MESSAGE"`
			} `json:"RESULT,omitempty"`
		} `json:"head,omitempty"`
		Row []struct {
			SchoolCode  string `json:"SD_SCHUL_CODE"`
			SchoolName  string `json:"SCHUL_NM"`
			MealCode    string `json:"MMEAL_SC_CODE"`
			MealName    string `json:"MMEAL_SC_NM"`
			MealDate    string `json:"MLSV_YMD"`
			FoodName    string `json:"DDISH_NM"`
			FoodOrigin  string `json:"ORPLC_INFO"`
			CalorieInfo string `json:"CAL_INFO"`
			NeutronInfo string `json:"NTR_INFO"`
		} `json:"row,omitempty"`
	} `json:"mealServiceDietInfo"`
}
