package handler

import (
	"encoding/json"
	"fmt"
	"fmt/service"
	"log"
	"net/http"
)

/**
* GET /api/menus
* 모든 메뉴에 대한 정보를 반환
 */
func GetMenusHandler(w http.ResponseWriter, r *http.Request) {
	menus, err := service.FindMenus()
	if err != nil {
		// 에러 처리
	}
	json.NewEncoder(w).Encode(menus)

}

/**
* POST /api/menus/score
* 요청 바디에 담긴 메뉴 이름과 사용자 ID를 통해 메뉴에 대한 평점을 삽입, 갱신
* 요청 body 예시
*	{
*		"menuName": "짜장면",
*		"userId": 1,
*		"score": 5
*	}
 */
func PostMenuScoreHandler(w http.ResponseWriter, r *http.Request) {
	// 콘솔에 로그를 출력
	log.Println("POST /api/menus/score")

	// 요청 바디를 Menu table 뿐만 아니라 User table과도 연관된 Score를 저장할 수 있는 구조체로 파싱
	var request struct {
		MenuName string `json:"menuName"`
		UserID   int64  `json:"userId"`
		Score    int    `json:"score"`
	}

	// error handling
	// 요청 바디를 파싱하는 과정에서 에러가 발생하면, 400 Bad Request 에러를 반환
	parseErr := json.NewDecoder(r.Body).Decode(&request)
	if parseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body: %v", parseErr)
		return
	}

	// findMenu 함수를 통해 요청 바디에 담긴 메뉴 이름을 가진 Menu를 찾음
	// 요청 바디에 담긴 메뉴 이름을 가진 Menu가 없으면 insertMenu 함수를 통해 Menu table에 새로운 메뉴를 삽입
	// 요청 바디에 담긴 메뉴 이름을 가진 Menu가 있으면 updateMenu 함수를 통해 Menu table에 메뉴의 평점을 갱신
	_, err := service.FindMenu(request.MenuName)
	if err != nil {
		// 값이 없으면 insertMenu
		service.InsertMenu(request.MenuName, request.Score)
	} else {
		// 값이 있으면 updateMenu
		service.UpdateMenu(request.MenuName, request.Score)
	}

	// findUser 함수를 통해 요청 바디에 담긴 사용자 ID를 가진 User를 찾음
	// 요청 바디에 담긴 사용자 ID를 가진 User가 없으면 값이 없다는 에러를 반환
	// 요청 바디에 담긴 사용자 ID를 가진 User가 있으면 updateMenuScore 함수를 통해 User table에 메뉴의 평점을 갱신
	user, err := service.FindUser(request.UserID)
	if err != nil {
		// 요청 바디에 담긴 사용자 ID를 가진 User가 없으면 에러를 반환
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user id: %v", err)
		return
	} else {
		// 요청 바디에 담긴 사용자 ID를 가진 User가 있으면 updateMenuScore
		service.UpdateMenuScore(user, request.MenuName, request.Score)
	}
}
