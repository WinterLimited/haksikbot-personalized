# 음식점 추천 시스템: 학식 메뉴 버전

이 프로젝트는 사용자의 취향에 맞는 학식 메뉴를 추천하는 애플리케이션입니다. 사용자는 제공되는 학식 메뉴를 평가하고, 이 평가를 기반으로 새로운 학식 메뉴의 적합성을 계산하여 사용자에게 추천합니다. 

## 사용 기술

# Spring
- **Spring Boot**: 애플리케이션 서버를 구현하고 REST API를 제공합니다.
- **JPA**: 데이터베이스와의 통신을 담당합니다. 메뉴별 사용자 점수를 저장하고, 필요한 경우 사용자의 평가 기록을 조회합니다.
- **React.js**: 사용자 인터페이스를 구현합니다. 아직 연동은 진행하지 않았습니다.

## 앱 소개

사용자는 제공되는 학식 메뉴를 '좋아요', '보통이에요', '싫어요' 세 가지 카테고리로 평가할 수 있습니다. 이 평가는 각 메뉴에 점수로 변환되어 데이터베이스에 저장됩니다. 

학식 메뉴가 제공될 때, 사용자가 이전에 평가한 메뉴 점수를 기반으로 새로운 메뉴의 적합성을 계산합니다.
정보가 없는 메뉴의 적합성 점수는 사용자 전원의 메뉴 평가 점수의 평균으로 계산됩니다. 이 점수는 5점 만점으로, 사용자에게 얼마나 적합한 메뉴인지를 알려줍니다.

## 사용 방법 - Spring

### 사용자의 메뉴 점수 적합도를 계산하여 반환하는 API

```java
@GetMapping("/api/{userId}/{menuName}")
public int getUserMenuScore(@RequestParam Long userId, @RequestParam String menuName) {
    // ...
}
```

### 사용자의 메뉴 점수를 업데이트하는 API

```java
@PostMapping("/api/menuscore")
public void updateUserMenuScore(@RequestParam Long userId, @RequestParam String menuName, @RequestParam int score) {
    // ...
}
```

## 개발 진행 상황 - Spring

현재까지 스프링 부트를 사용하여 백엔드 서버를 구현하였고, JPA를 통한 데이터베이스 연동, 메뉴 점수 계산 로직 등이 구현되어 있습니다. 다음 단계로는 React.js를 이용해 프론트엔드를 구현하고, 백엔드 서버와의 연동을 진행할 예정입니다.


# Golang
- **Go**: 애플리케이션 서버를 http 통신으로 이용하여 REST API를 제공합니다.

## 사용 방법 - Go

### 메뉴 점수를 반환하는 API

```Go
// GET /api/menus
http.HandleFunc("/api/menus", getMenusHandler)
```

### 메뉴 이름과 사용자 ID를 통해 메뉴에 대한 평점을 삽입, 갱신

```Go
// POST /api/menus/score
http.HandleFunc("/api/menus/score", postMenuScoreHandler)
```

### 사용자 ID를 통해 사용자의 이름과 메뉴별 평점을 반환

```Go
// GET /api/users/{userId}
http.HandleFunc("/api/users/", getUserHandler)
```

## 개발 진행 상황 - Go

코드의 리팩토링, ScyllaDB와의 연동을 통한 데이터 관리 예정입니다. =>
[caucat_haksikbot repository에서 진행](https://github.com/WinterLimited/caucat_haksikbot)


