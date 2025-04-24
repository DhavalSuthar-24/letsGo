package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/DhavalSuthar-24/letsGo/internal/utils"
	"github.com/DhavalSuthar-24/letsGo/internal/models"
	"github.com/DhavalSuthar-24/letsGo/internal/services"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(AuthService *services.AuthService) *AuthController {
	return &AuthController{AuthService: AuthService}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.AuthService.Register(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
func (c *AuthController) Login(ctx *gin.Context) {
    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    
    if err := ctx.ShouldBindJSON(&credentials); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    user, err := c.AuthService.Login(credentials.Email, credentials.Password)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
    
    // Generate JWT token
    token, err := utils.GenerateJWT(user)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
        "user":    user,
    })
}

func (c *AuthController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.AuthService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *AuthController) GetAllUsers(ctx *gin.Context) {
	users, err := c.AuthService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
