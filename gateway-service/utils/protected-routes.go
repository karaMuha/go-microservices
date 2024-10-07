package utils

import (
	"github.com/vodkaslime/wildcard"
)

var protectedRoutes map[string][]string
var matcher *wildcard.Matcher

func SetProtectedRoutes() {
	matcher = wildcard.NewMatcher()
	protectedRoutes = make(map[string][]string)
	protectedRoutes["GET /ping"] = []string{""}
	protectedRoutes["POST /users/signup"] = []string{""}
	protectedRoutes["POST /users/confirm/*/*"] = []string{""}
	protectedRoutes["POST /users/login"] = []string{""}
	protectedRoutes["POST /users/logout"] = []string{""}
	protectedRoutes["GET /users/*"] = []string{"user", "admin"}
	protectedRoutes["GET /users"] = []string{"admin"}
	protectedRoutes["GET /logs/{id}"] = []string{"admin"}
	protectedRoutes["GET /logs"] = []string{"admin"}
}

func IsProtectedRoute(endpoint string) ([]string, error) {
	for i, v := range protectedRoutes {
		result, err := matcher.Match(i, endpoint)

		if err != nil {
			return nil, err
		}

		if result {
			return v, nil
		}
	}

	return nil, nil
}
