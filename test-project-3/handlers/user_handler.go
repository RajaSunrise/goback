package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/test/test-project-3/models"
	"github.com/test/test-project-3/services"
	"github.com/google/uuid"

	// --- Import dinamis berdasarkan framework ---
	"github.com/gin-gonic/gin"
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







// CreateUser menangani request POST untuk membuat user baru.
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	user, err := h.userService.CreateUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// GetUsers menangani request GET untuk mendapatkan daftar user dengan paginasi.
func (h *UserHandler) GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

	response, err := h.userService.GetAllUsers(c.Request.Context(), page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetUser menangani request GET untuk mendapatkan satu user berdasarkan ID.
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := parseUserID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser menangani request PUT untuk memperbarui data user.
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := parseUserID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	user, err := h.userService.UpdateUser(c.Request.Context(), id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser menangani request DELETE untuk menghapus user.
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := parseUserID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}








// parseUserID adalah helper untuk mengurai ID dari string, beradaptasi dengan tipe database.
func parseUserID(idStr string) (UserID, error) {

	id, err := uuid.Parse(idStr)
	return id, err

}