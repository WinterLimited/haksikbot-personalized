package handler

import (
	"encoding/json"
	"fmt"
	"fmt/service"
	"net/http"
	"strconv"
	"strings"
)

/**
* GET /api/users/{userId}
* 요청 URL에 담긴 사용자 ID를 통해 사용자의 이름과 메뉴별 평점을 반환
* TODO: {userId}를 파싱하는 과정이 너무 복잡해서 개선방법을 알아보기
 */
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
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
	user, err := service.FindUser(userId)
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
