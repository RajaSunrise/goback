package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/test/test-project-fiber/models"
	"github.com/test/test-project-fiber/services"
	"github.com/google/uuid"

	// --- Import dinamis berdasarkan framework ---
	"github.com/gofiber/fiber/v2"
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
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body: " + err.Error()})
	}

	user, err := h.userService.CreateUser(c.Context(), &req)
	if err != nil {
		// Di sini kita bisa memeriksa tipe error untuk memberikan status code yang lebih baik
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUsers menangani request GET untuk mendapatkan daftar user dengan paginasi.
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))

	response, err := h.userService.GetAllUsers(c.Context(), page, perPage)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(response)
}

// GetUser menangani request GET untuk mendapatkan satu user berdasarkan ID.
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := parseUserID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}

	user, err := h.userService.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

// UpdateUser menangani request PUT untuk memperbarui data user.
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := parseUserID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body: " + err.Error()})
	}

	user, err := h.userService.UpdateUser(c.Context(), id, &req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

// DeleteUser menangani request DELETE untuk menghapus user.
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := parseUserID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format"})
	}

	if err := h.userService.DeleteUser(c.Context(), id); err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}











// parseUserID adalah helper untuk mengurai ID dari string, beradaptasi dengan tipe database.
func parseUserID(idStr string) (UserID, error) {

	id, err := uuid.Parse(idStr)
	return id, err

}