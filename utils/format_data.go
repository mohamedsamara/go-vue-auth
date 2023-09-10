package utils

import "github.com/mohamedsamara/golang-vue/models"

func FormatUserResponse(u models.User) map[string]interface{} {
	user := make(map[string]interface{})

	user["id"] = u.ID
	user["name"] = u.Name
	user["email"] = u.Email
	user["createdAt"] = u.CreatedAt
	user["updatedAt"] = u.UpdatedAt

	return user
}
