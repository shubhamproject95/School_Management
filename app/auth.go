package app

import (
	"context"
	"net/http"
	"strings"

	"School_gql/service"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
		loginUser, err := service.NewUserService().Auth(token)

		if err != nil {
			// controller.ApiResponse(w, &controller.Res{
			// 	Code:    901,
			// 	Message: "error",
			// 	Data:    "User not authorized " + err.Error(),
			// })
			return
		}

		ctx := context.WithValue(r.Context(), "user", loginUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
