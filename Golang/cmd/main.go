package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// User data structure
type User struct {
	ID         int64       `json:"id"`
	Name       string      `json:"name"`
	MenuScores []MenuScore `json:"menuScores"`
}

// Menu data structure
// Menu에 대한 모든 사용자들의 평점을 저장하고, Count 필드를 통해 평균을 구할 수 있도록 함
type Menu struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Count int    `json:"count"`
}

// MenuScore data structure
// User별로 Menu에 대한 평점을 매길 수 있도록 함
type MenuScore struct {
	MenuName string `json:"menuName"`
	Score    int    `json:"score"`
	Count    int    `json:"count"`
}

// menus data
var menus = []Menu{
	{
		Name:  "짜장면",
		Score: 0,
		Count: 0,
	},
	{
		Name:  "짬뽕",
		Score: 0,
		Count: 0,
	},
	{
		Name:  "탕수육",
		Score: 0,
		Count: 0,
	},
	{
		Name:  "볶음밥",
		Score: 0,
		Count: 0,
	},
}

// users data
var users = []User{
	{
		ID:   1,
		Name: "Winter",
		MenuScores: []MenuScore{
			{
				MenuName: "짜장면",
				Score:    0,
				Count:    0,
			},
			{
				MenuName: "짬뽕",
				Score:    0,
				Count:    0,
			},
			{
				MenuName: "탕수육",
				Score:    0,
				Count:    0,
			},
			{
				MenuName: "볶음밥",
				Score:    0,
				Count:    0,
			},
		},
	},
}

func main() {
	// rest api를 통해 CRUD 메서드를 구현

	// GET /api/menus
	// 모든 메뉴에 대한 정보를 반환
	http.HandleFunc("/api/menus", getMenusHandler)

	// POST /api/menus/score
	// 요청 바디에 담긴 메뉴 이름과 사용자 ID를 통해 메뉴에 대한 평점을 삽입, 갱신
	http.HandleFunc("/api/menus/score", postMenuScoreHandler)

	// GET /api/users/{userId}
	// 요청 URL에 담긴 사용자 ID를 통해 사용자의 이름과 메뉴별 평점을 반환
	http.HandleFunc("/api/users/", getUserHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// rest api를 통해 CRUD 메서드를 구현

/**
* GET /api/menus
* 모든 메뉴에 대한 정보를 반환
 */
func getMenusHandler(w http.ResponseWriter, r *http.Request) {
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
func postMenuScoreHandler(w http.ResponseWriter, r *http.Request) {
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
	_, err := findMenu(request.MenuName)
	if err != nil {
		// 값이 없으면 insertMenu
		insertMenu(request.MenuName, request.Score)
	} else {
		// 값이 있으면 updateMenu
		updateMenu(request.MenuName, request.Score)
	}

	// findUser 함수를 통해 요청 바디에 담긴 사용자 ID를 가진 User를 찾음
	// 요청 바디에 담긴 사용자 ID를 가진 User가 없으면 값이 없다는 에러를 반환
	// 요청 바디에 담긴 사용자 ID를 가진 User가 있으면 updateMenuScore 함수를 통해 User table에 메뉴의 평점을 갱신
	user, err := findUser(request.UserID)
	if err != nil {
		// 요청 바디에 담긴 사용자 ID를 가진 User가 없으면 에러를 반환
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid user id: %v", err)
		return
	} else {
		// 요청 바디에 담긴 사용자 ID를 가진 User가 있으면 updateMenuScore
		updateMenuScore(user, request.MenuName, request.Score)
	}
}

/**
* GET /api/users/{userId}
* 요청 URL에 담긴 사용자 ID를 통해 사용자의 이름과 메뉴별 평점을 반환
* TODO: {userId}를 파싱하는 과정이 너무 복잡해서 개선방법을 알아보기
 */
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// URL path에서 userId를 추출합니다.
	// 예를 들어, "/api/users/1"에서 "1"을 추출하게 됩니다.
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) != 4 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid URL")
		return
	}

	userIdStr := pathSegments[3]
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid userId: %v", err)
		return
	}

	// userId에 대한 사용자를 찾습니다.
	user, err := findUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User not found: %v", err)
		return
	}

	// 사용자 정보를 JSON으로 변환하여 응답에 씁니다.
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to encode user data: %v", err)
		return
	}
}

// Menu table에 값 존재여부 확인
// error handling 추가
func findMenu(menuName string) (*Menu, error) {
	for _, menu := range menus {
		if menu.Name == menuName {
			return &menu, nil
		}
	}
	return nil, fmt.Errorf("Menu not found")
}

// Menu table에 새로운 메뉴 삽입
func insertMenu(menuName string, score int) (*Menu, error) {
	menu := Menu{
		Name:  menuName,
		Score: score,
		Count: 1,
	}
	menus = append(menus, menu)
	return &menu, nil
}

// Menu table에 메뉴의 평점 갱신
// *Menu로 값을 직접 받아서 갱신
func updateMenu(menuName string, score int) (*Menu, error) {
	// 해당 값을 찾아서 갱신
	for i := range menus {
		if menus[i].Name == menuName {
			menus[i].Score = (menus[i].Score*menus[i].Count + score) / (menus[i].Count + 1)
			menus[i].Count++
			return &menus[i], nil
		}
	}
	return nil, fmt.Errorf("Menu not found")
}

// User table에 값 존재여부 확인
// error handling 추가
func findUser(userID int64) (*User, error) {
	for _, user := range users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}

// User table에 메뉴의 평점 갱신
func updateMenuScore(user *User, name string, score int) {
	// 해당 값을 찾아서 갱신
	for i := range user.MenuScores {
		if user.MenuScores[i].MenuName == name {
			user.MenuScores[i].Score = (user.MenuScores[i].Score*user.MenuScores[i].Count + score) / (user.MenuScores[i].Count + 1)
			user.MenuScores[i].Count++
			return
		}
	}
}
