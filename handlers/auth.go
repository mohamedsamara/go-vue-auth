package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohamedsamara/go-vue/auth"
	"github.com/mohamedsamara/go-vue/constants"
	"github.com/mohamedsamara/go-vue/models"
	"github.com/mohamedsamara/go-vue/utils"
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

	accessToken, tokenError := auth.CreateAccessToken(constants.ACCESS_TOKEN_EXPIRE, newUser.ID)
	if tokenError != nil {
		utils.WriteFailureResponse(w, "Generating JWT token failed.")
		return
	}

	refreshToken, refreshTokenError := auth.CreateAccessToken(constants.REFRESH_TOKEN_EXPIRE, newUser.ID)
	if refreshTokenError != nil {
		utils.WriteFailureResponse(w, "Generating JWT token failed.")
		return
	}

	response := formatAuthResponse(accessToken, refreshToken)

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

	accessToken, tokenError := auth.CreateAccessToken(constants.ACCESS_TOKEN_EXPIRE, user.ID)
	if tokenError != nil {
		utils.WriteFailureResponse(w, "Generating JWT token failed.")
		return
	}

	refreshToken, refreshTokenError := auth.CreateAccessToken(constants.REFRESH_TOKEN_EXPIRE, user.ID)
	if refreshTokenError != nil {
		utils.WriteFailureResponse(w, "Generating JWT token failed.")
		return
	}

	response := formatAuthResponse(accessToken, refreshToken)

	utils.WriteSuccessResponse(w, "User logged in.", response)
}

func (h *BaseHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var payload models.RefreshTokenRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WritePermissionDenied(w)
		return
	}

	if payload.RefreshToken == "" {
		utils.WritePermissionDenied(w)
		return
	}

	refreshToken := payload.RefreshToken

	token, err := auth.VerifyJWT(refreshToken)

	if err != nil {
		utils.WritePermissionDenied(w)
		return
	}

	if !token.Valid {
		utils.WritePermissionDenied(w)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	userId := claims["id"]

	if !ok {
		utils.WritePermissionDenied(w)
		return
	}

	var user models.User

	h.db.First(&user, "id = ?", userId)

	if user.ID.String() != userId {
		utils.WritePermissionDenied(w)
		return
	}

	accessToken, tokenError := auth.CreateAccessToken(constants.ACCESS_TOKEN_EXPIRE, user.ID)
	if tokenError != nil {
		utils.WriteFailureResponse(w, "Generating JWT token failed.")
		return
	}

	response := formatAuthResponse(accessToken, refreshToken)

	utils.WriteSuccessResponse(w, "User logged in.", response)
}

func formatAuthResponse(accessToken string, refreshToken string) map[string]interface{} {
	response := make(map[string]interface{})
	response["accessToken"] = accessToken
	response["refreshToken"] = refreshToken
	return response
}
