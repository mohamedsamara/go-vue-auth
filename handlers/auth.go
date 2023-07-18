package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mohamedsamara/golang-vue/auth"
	"github.com/mohamedsamara/golang-vue/constants"
	"github.com/mohamedsamara/golang-vue/models"
	"github.com/mohamedsamara/golang-vue/utils"
)

func (h *BaseHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload models.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteFailureResponse(w, constants.API_PROCESSING_ERROR)
		return
	}

	if payload.Email == "" || payload.Name == "" || payload.Password == "" {
		utils.WriteFailureResponse(w, "Email, password or name is missing.")
		return
	}

	hashedPassword, hashPasswordError := utils.HashPassword(payload.Password)

	if hashPasswordError != nil {
		utils.WriteFailureResponse(w, "Wrong password.")
		return
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
	}

	dbResult := h.db.Create(&newUser)

	if dbResult.Error != nil {
		msg := "Wrong password."
		if strings.Contains(dbResult.Error.Error(), "duplicate key value violates unique") {
			msg = "User with that email already exists."
		}
		utils.WriteFailureResponse(w, msg)
		return
	}

	token, tokenError := auth.CreateAccessToken(15, newUser.ID)

	if tokenError != nil {
		utils.WriteFailureResponse(w, "Generating JWT token failed.")
		return
	}

	response := formatUserTokenResponse(newUser, token)

	utils.WriteSuccessResponse(w, "User registered.", response)
}

func (h *BaseHandler) Login(w http.ResponseWriter, r *http.Request) {

	var payload models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteFailureResponse(w, constants.API_PROCESSING_ERROR)
		return
	}

	if payload.Email == "" || payload.Password == "" {
		utils.WriteFailureResponse(w, "Email or password is missing.")
		return
	}

	var user models.User

	dbResult := h.db.First(&user, "email = ?", strings.ToLower(payload.Email))

	if dbResult.Error != nil {
		utils.WriteFailureResponse(w, "Invalid email or Password.")
		return
	}

	verifyError := utils.VerifyPassword(user.Password, payload.Password)

	if verifyError != nil {
		utils.WriteFailureResponse(w, "Invalid email or Password.")
		return
	}

	token, tokenError := auth.CreateAccessToken(1, user.ID)

	if tokenError != nil {
		utils.WriteFailureResponse(w, "Generating JWT token failed.")
		return
	}

	response := formatUserTokenResponse(user, token)

	utils.WriteSuccessResponse(w, "User logged in.", response)
}

func formatUserTokenResponse(user models.User, token string) map[string]interface{} {
	response := make(map[string]interface{})

	response["user"] = utils.FormatUserResponse(user)
	response["token"] = token

	return response
}
