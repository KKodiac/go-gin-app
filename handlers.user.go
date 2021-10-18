// handlers.user.go

package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func showLoginPage(c *gin.Context) {
	render(c, gin.H{
		"Title": "Login Page",
	}, "login.html")
}

func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if isUserValid(username, password) {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"Title": "Successful Login",
			"Name":  getAccountName(username, password),
		}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided",
		})
	}
}

func showRegistrationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Registration Page",
	}, "registration.html")
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")

	if _, err := registerNewUser(username, password, name); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"Title": "Successful registration & Login",
			"Name":  name}, "login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest, "registration.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error(),
		})
	}
}
