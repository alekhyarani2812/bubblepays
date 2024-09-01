package api

import (
	"encoding/json"
	"net/http"
	"time"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/repository/auth"
	"github.com/bubblepays/utils"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

var jwtkey = []byte("secret_key")
// LoadRoutes for routing to get the token
func LoadRoutes(router *mux.Router) {
	router.HandleFunc("/api/getToken", Login).Methods("POST")
}
// Login handler for authenticating the user
func Login(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	// data, err := token.GenerateToken(&ctx)
	var userLogin models.User
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userData, err := auth.Login(ctx, userLogin)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
		}
		utils.Render(w, response)
		return
	}
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &models.Jwt{
		Username: userData.Name,
		Role:     "bubblepays.admin", //@TODO need to create a role and get the data from it
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		log.Errorft(ctx, "got an error while doing signing error %v", err.Error())
		response := models.Response{
			Error:   err.Error(),
			Status:  http.StatusInternalServerError,
			Message: utils.Fail,
		}
		utils.Render(w, response)
		return
	}
	data := &models.AuthResponse{
		AccessToken: tokenString,
	}
	response := models.Response{
		Message:     utils.Success,
		Description: "You are singned in successfully",
		Error:       "",
		Status:      http.StatusOK,
		Data:        data,
	}
	utils.Render(w, response)
}
