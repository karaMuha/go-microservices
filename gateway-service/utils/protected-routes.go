package utils

import "github.com/vodkaslime/wildcard"

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
		if result, err := matcher.Match(i, endpoint); err == nil && result {
			return v, nil
		} else if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
