package controllers

import (
	"chat-app/internal/api"
	"chat-app/pkg/config"
	"chat-app/pkg/jwt"
	"chat-app/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type (
	LoginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	SignupReq struct {
		api.User
		ConfirmPassword string `json:"confirmPassword"`
	}
)

func Login(c *gin.Context) {
	var req LoginReq
	if err := c.Bind(&req); err != nil {
		response.Error(c, 500, err)
		return
	}

	user, err := api.FindByUsername(req.Username)
	if err != nil {
		response.Error(c, 500, err)
		return
	}

	if user == nil {
		response.JSON(c, 400, "username or password valid")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		response.JSON(c, 400, "password error")
		return
	}

	generateTokenAndSetCookie(user.ID.Hex(), c)
	user.Password = ""
	response.JSON(c, 200, user)
}

func Signup(c *gin.Context) {
	var req SignupReq
	if err := c.Bind(&req); err != nil {
		response.Error(c, 500, err)
		return
	}

	if req.Password != req.ConfirmPassword {
		response.JSON(c, 400, "password do not match")
		return
	}
	user, _ := api.FindByUsername(req.Username)

	if user != nil {
		response.JSON(c, 401, "username already exists")
		return
	}

	user = &req.User

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	profilePic := "https://avatar.iran.liara.run/public"
	if req.Gender == "male" {
		profilePic += "/boy?username=" + user.Username
	} else {
		profilePic += "/girl?username=" + user.Username
	}
	user.ProfilePic = profilePic

	err := api.InsertOneUser(user)
	if err != nil {
		response.Error(c, 500, err)
		return
	}

	generateTokenAndSetCookie(req.User.ID.Hex(), c)
	user.Password = ""
	response.JSON(c, 200, user)
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", 0, "", "", true, true)
	response.JSON(c, 200, "logout successfully")
}

func generateTokenAndSetCookie(userId string, c *gin.Context) {
	token := jwt.GenerateAccessToken(map[string]interface{}{"id": userId})
	c.SetCookie("jwt",
		token,
		jwt.AccessTokenExpiredTime,
		"",
		"",
		config.GetConfig().Environment == config.ProductionEnv,
		true)
}
