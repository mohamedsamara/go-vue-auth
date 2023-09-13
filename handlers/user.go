package handlers

import (
	"net/http"

	"github.com/mohamedsamara/go-vue-auth/constants"
	"github.com/mohamedsamara/go-vue-auth/utils"
)

func (h *BaseHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value(constants.AUTH_CONTEXT)
	utils.WriteSuccessResponse(w, "User info.", currentUser)
}
