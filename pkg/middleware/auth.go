package middleware

import (
	"context"
	"encoding/json"
	dto "gorm-imp/dto/result"
	jwtToken "gorm-imp/pkg/jwt"
	"net/http"
	"strings"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"string"`
}

// Auth func return http.Handler and pass request to the next handler with ctx and value inside it

// func Auth(next http.Handler) http.HandlerFunc {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")

// 		// retrieve token from header with following code
// 		token := r.Header.Get("Authorization")

// 		// lakukan pengecekan nil header
// 		if token == "" {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "unauthorized"}
// 			json.NewEncoder(w).Encode(response)
// 			return
// 		}

// 		// retrieve the 2nd string from "bearer <string>"
// 		token = strings.Split(token, " ")[1]
// 		claims, err := jwtToken.DecodeToken(token)

// 		fmt.Println(claims)

// 		if err != nil {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			response := Result{Code: http.StatusUnauthorized, Message: "unauthorized"}
// 			json.NewEncoder(w).Encode(response)
// 			return
// 		}

// 		// context is temp memory to carry value evaluated from middleware like auth, that can be carried to the next handler
// 		ctx := context.WithValue(r.Context(), "userInfo", claims)

// 		// NOTE: still doesnt understand logic of code block below
// 		r = r.WithContext(ctx)

// 		// serveHTTP() method do pass the ctx with value to the next handler
// 		next.ServeHTTP(w, r.WithContext(ctx))

// 	})

// }

// NOTE: did some typo at top function :()

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			response := Result{Code: http.StatusUnauthorized, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		ctx := context.WithValue(r.Context(), "userInfo", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
