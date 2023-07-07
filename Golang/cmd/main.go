//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//// User data structure
//type User struct {
//	ID         int64       `json:"id"`
//	Name       string      `json:"name"`
//	MenuScores []MenuScore `json:"menuScores"`
//}
//
//// Menu data structure
//// Menu에 대한 모든 사용자들의 평점을 저장하고, Count 필드를 통해 평균을 구할 수 있도록 함
//type Menu struct {
//	Name  string `json:"name"`
//	Score int    `json:"score"`
//	Count int    `json:"count"`
//}
//
//// MenuScore data structure
//// User별로 Menu에 대한 평점을 매길 수 있도록 함
//type MenuScore struct {
//	MenuName string `json:"menuName"`
//	Score    int    `json:"score"`
//	Count    int    `json:"count"`
//}
//
//// menus data
//var menus = []Menu{
//	{
//		Name:  "짜장면",
//		Score: 0,
//		Count: 0,
//	},
//	{
//		Name:  "짬뽕",
//		Score: 0,
//		Count: 0,
//	},
//	{
//		Name:  "탕수육",
//		Score: 0,
//		Count: 0,
//	},
//	{
//		Name:  "볶음밥",
//		Score: 0,
//		Count: 0,
//	},
//}
//
//// users data
//var users = []User{
//	{
//		ID:   1,
//		Name: "Winter",
//		MenuScores: []MenuScore{
//			{
//				MenuName: "짜장면",
//				Score:    0,
//				Count:    0,
//			},
//			{
//				MenuName: "짬뽕",
//				Score:    0,
//				Count:    0,
//			},
//			{
//				MenuName: "탕수육",
//				Score:    0,
//				Count:    0,
//			},
//			{
//				MenuName: "볶음밥",
//				Score:    0,
//				Count:    0,
//			},
//		},
//	},
//}
//
//func main() {
//	// rest api를 통해 CRUD 메서드를 구현
//
//	// GET /api/menus
//	// 모든 메뉴에 대한 정보를 반환
//	http.HandleFunc("/api/menus", getMenusHandler)
//
//	// POST /api/menus/score
//	// 요청 바디에 담긴 메뉴 이름과 사용자 ID를 통해 메뉴에 대한 평점을 삽입, 갱신
//	http.HandleFunc("/api/menus/score", postMenuScoreHandler)
//
//	// GET /api/users/{userId}
//	// 요청 URL에 담긴 사용자 ID를 통해 사용자의 이름과 메뉴별 평점을 반환
//	http.HandleFunc("/api/users/", getUserHandler)
//
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
//
//// rest api를 통해 CRUD 메서드를 구현
//
//// Menu table에 값 존재여부 확인
//// error handling 추가
//func findMenu(menuName string) (*Menu, error) {
//	for _, menu := range menus {
//		if menu.Name == menuName {
//			return &menu, nil
//		}
//	}
//	return nil, fmt.Errorf("Menu not found")
//}
//
//// Menu table에 새로운 메뉴 삽입
//func insertMenu(menuName string, score int) (*Menu, error) {
//	menu := Menu{
//		Name:  menuName,
//		Score: score,
//		Count: 1,
//	}
//	menus = append(menus, menu)
//	return &menu, nil
//}
//
//// Menu table에 메뉴의 평점 갱신
//// *Menu로 값을 직접 받아서 갱신
//func updateMenu(menuName string, score int) (*Menu, error) {
//	// 해당 값을 찾아서 갱신
//	for i := range menus {
//		if menus[i].Name == menuName {
//			menus[i].Score = (menus[i].Score*menus[i].Count + score) / (menus[i].Count + 1)
//			menus[i].Count++
//			return &menus[i], nil
//		}
//	}
//	return nil, fmt.Errorf("Menu not found")
//}
//
//// User table에 값 존재여부 확인
//// error handling 추가
//func findUser(userID int64) (*User, error) {
//	for _, user := range users {
//		if user.ID == userID {
//			return &user, nil
//		}
//	}
//	return nil, fmt.Errorf("User not found")
//}
//
//// User table에 메뉴의 평점 갱신
//func updateMenuScore(user *User, name string, score int) {
//	// 해당 값을 찾아서 갱신
//	for i := range user.MenuScores {
//		if user.MenuScores[i].MenuName == name {
//			user.MenuScores[i].Score = (user.MenuScores[i].Score*user.MenuScores[i].Count + score) / (user.MenuScores[i].Count + 1)
//			user.MenuScores[i].Count++
//			return
//		}
//	}
//}

package main

import (
	"fmt"
	"fmt/api"
	"log"
	"net/http"
)

func main() {
	// API 라우트 등록
	api.RegisterRoutes()

	// 서버 실행
	port := 8080
	log.Printf("Server listening on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
