package api

import (
	"fmt"
	"fmt/api/handler"
	"net/http"
	"strconv"
	"strings"
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
	http.HandleFunc("/api/users", func(wr http.ResponseWriter, r *http.Request) {

		// URL path에서 userId를 추출합니다.
		// 예를 들어, "/api/users/1"에서 "1"을 추출하게 됩니다.
		pathSegments := strings.Split(r.URL.Path, "/")
		if len(pathSegments) != 4 {
			wr.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(wr, "Invalid URL")
			return
		}

		userIdStr := pathSegments[3]
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(wr, "Invalid userId: %v", err)
			return
		}

		switch r.Method {
		case http.MethodGet: // GET
			handler.GetUserHandler(wr, r, userId)
		}
	})
}
