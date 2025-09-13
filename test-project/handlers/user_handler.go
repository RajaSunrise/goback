package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/test/test-project/models"
	"github.com/test/test-project/services"
	"github.com/google/uuid"

	// --- Import dinamis berdasarkan framework ---
)

// Tipe ID dinamis berdasarkan pilihan database
type UserID = uuid.UUID

// UserHandler menangani request HTTP yang berhubungan dengan User.
type UserHandler struct {
	userService services.UserService
}

// NewUserHandler membuat instance baru dari UserHandler.
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}














// parseUserID adalah helper untuk mengurai ID dari string, beradaptasi dengan tipe database.
func parseUserID(idStr string) (UserID, error) {

	id, err := uuid.Parse(idStr)
	return id, err

}