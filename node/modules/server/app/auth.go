package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/syahrul12345/Blocknalytics/node/modules/server/models"
	"github.com/syahrul12345/Blocknalytics/node/modules/server/utils"
)

//JwtAuthentication authenticates the received JWT token
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		notAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := request.URL.Path

		//check if response does not require authentication
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(writer, request)
				return
			}
		}
		//other wise it requires authentication
		response := make(map[string]interface{})
		tokenHeader := request.Header.Get("Authorization")

		if tokenHeader == "" {
			response = utils.Message(false, "Missing auth token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
		}
		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
			return
		}
		tokenPart := splitted[1] // the information that we're interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		//malformed token, return 403
		if err != nil {
			fmt.Println(err)
			response = utils.Message(false, "Malformed auth token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
			return
		}
		//token is invalid
		if !token.Valid {
			fmt.Println(token.Valid)
			response = utils.Message(false, "Token is invalid")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
			return
		}

		//everything went well
		fmt.Sprintf("User ", tk.UserName)
		ctx := context.WithValue(request.Context(), "user", tk.UserID)
		request = request.WithContext(ctx)
		next.ServeHTTP(writer, request)
	})
}
