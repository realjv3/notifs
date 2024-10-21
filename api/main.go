package main

import (
	"fmt"
	"net/http"

	"github.com/realjv3/notifs/auth"
	"github.com/realjv3/notifs/notifications"
	"github.com/realjv3/notifs/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Initialize the database
	db := initDB()
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize the web framework
	r := gin.Default()
	err = r.SetTrustedProxies(nil)
	if err != nil {
		fmt.Println(err)
		return
	} // Hide trusted proxies warning
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	r.Use(cors.New(config))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	s := auth.NewService(db, r)

	// TODO move route definitions to their respective packages by entity

	// A simple login endpoint that takes an id and email
	// and returns a signed JWT
	r.POST("/login", func(c *gin.Context) {
		// validate
		email := c.PostForm("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
			return
		}

		// authenticate
		tokenString, err := s.Authenticate(email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	})

	r.GET("/notifications", func(c *gin.Context) {
		// authorize
		claims := c.MustGet("claims").(jwt.MapClaims)
		email := claims["email"].(string)
		if email == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// fetch user
		u := users.NewService(db)
		user, err := u.GetUserByEmail(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// fetch notifications by user and preferred type
		// TODO possibly accommodate filtering by type via HTTP query param
		n := notifications.NewService(db)
		notifs, err := n.GetNotificationsByPrefType(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"notifications": notifs})
	})

	r.POST("/preferences", func(c *gin.Context) {
		// validate
		preference := c.PostForm("preference")
		if preference != "EMAIL" && preference != "PUSH" && preference != "SMS" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid preference"})
			return
		}

		// authorize
		claims := c.MustGet("claims").(jwt.MapClaims)
		email := claims["email"].(string)
		if email == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// fetch user
		u := users.NewService(db)
		user, err := u.GetUserByEmail(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// update a user's notification preference
		user.Preference = preference
		err = u.SetUserNotificationPref(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	err = r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
