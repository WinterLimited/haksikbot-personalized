package api

import (
	"fmt/api/handler"
	"net/http"
)

// RegisterRoutes registers the API routes
func RegisterRoutes() {

	// /api/menus 경로로 요청이 들어오면, 요청 메서드에 따라 다른 핸들러 함수를 실행
	http.HandleFunc("/api/menus", func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET /api/menus
		// handler.GetMenusHandler(wr, r)
		case http.MethodGet: // GET
			handler.GetMenusHandler(wr, r)
		case http.MethodPost: // POST
			handler.PostMenuScoreHandler(wr, r)
		}
	})

	// api/users 경로로 요청이 들어오면, 요청 메서드에 따라 다른 핸들러 함수를 실행
	http.HandleFunc("/api/users/{userId}", func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: // GET
			handler.GetUserHandler(wr, r)
		}
	})
}
