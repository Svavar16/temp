// with
/**
To add
go get -u github.com/labstack/echo/v4
go get -u github.com/markbates/goth
go get -u github.com/dgrijalva/jwt-go

*/

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	// "github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

var (
	secretKey       = []byte("your_secret_key")
	googleClientKey = "your_google_client_key"
	googleSecretKey = "your_google_secret_key"
)

type UserClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Name   string   `json:"name"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}

func init() {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8080/auth/google/callback"),
	)
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the home page!")
	})

	// Middleware for handling access token
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenStr := c.Request().Header.Get("Authorization")

			// Extract the token from the "Authorization" header
			if len(tokenStr) < 8 || tokenStr[:7] != "Bearer " {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}
			tokenStr = tokenStr[7:]

			// Parse the token
			token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			if !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is invalid")
			}

			// Store the user object in the context
			c.Set("user", token)

			return next(c)
		}
	})

	// here are assuming the we get the provider with a query param
	e.GET("/auth", handleAuth)
	e.GET("/auth/callback", handleAuthCallback)
	e.GET("/protected", fetchProtectedData)

	e.Logger.Fatal(e.Start(":8080"))
}

func handleAuth(c echo.Context) error {
	gothic.BeginAuthHandler(c.Response().Writer, c.Request())
	return nil
}

func handleAuthCallback(c echo.Context) error {
	provider := c.QueryParam("provider")
	user, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error completing OAuth2 authentication")
	}

	// Here you can use the user information to create or authenticate the user in your system.
	fmt.Printf("User logged in with provider %s\n", provider)
	fmt.Printf("User ID: %s\n", user.UserID)
	fmt.Printf("User Email: %s\n", user.Email)
	fmt.Printf("User Name: %s\n", user.Name)

	// Generate JWTs for the user and send them in the response
	accessToken, refreshToken, err := generateTokens(user.UserID, user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating tokens")
	}

	response := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return c.JSON(http.StatusOK, response)
}

func generateTokens(userID, username string) (string, string, error) {
	// Generate the access token with a short expiration time (e.g., 15 minutes)
	accessExp := time.Now().Add(time.Minute * 15).Unix()
	accessClaims := &UserClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExp,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenStr, err := accessToken.SignedString(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return "", "", err
	}

	// Generate the refresh token with a longer expiration time (e.g., 30 days)
	refreshExp := time.Now().Add(time.Hour * 24 * 30).Unix()
	refreshClaims := &UserClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExp,
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenStr, err := refreshToken.SignedString(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

func fetchProtectedData(c echo.Context) error {
	// Get the user information from the context
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*UserClaims)

	// In a real application, you would use the userID or username to fetch data from the API or database.
	// For simplicity, we're just printing the user information here.
	fmt.Println("Fetching protected data for user:", claims.Username)

	return c.String(http.StatusOK, "Protected data fetched successfully!")
}

// this should refresh the access token, based on that he has one
func refreshAccessToken(c echo.Context) error {
	// Get the refresh token from the request
	refreshToken := c.FormValue("refresh_token")
	if refreshToken == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Refresh token not provided")
	}

	// Parse the refresh token
	token, err := jwt.ParseWithClaims(refreshToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid refresh token")
	}

	// Check if the token is valid
	if !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid refresh token")
	}

	// Extract the user claims from the refresh token
	claims := token.Claims.(*UserClaims)

	// Check if the refresh token is still valid (not expired)
	if time.Now().Unix() > claims.ExpiresAt {
		return echo.NewHTTPError(http.StatusUnauthorized, "Refresh token has expired")
	}

	// Generate a new access token with a short expiration time (e.g., 15 minutes)
	accessExp := time.Now().Add(time.Minute * 15).Unix()
	accessClaims := &UserClaims{
		UserID:   claims.UserID,
		Username: claims.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExp,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenStr, err := accessToken.SignedString(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating new access token")
	}

	response := map[string]string{
		"access_token": accessTokenStr,
	}

	return c.JSON(http.StatusOK, response)
}

// this would be a middleware that chould check if the user is valid or not.
// this is used in the main function, so e.use(checkUser), something like this.
func checkUser (next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("Authorization")

		// Extract the token from the "Authorization" header
		if len(tokenStr) < 8 || tokenStr[:7] != "Bearer " {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}
		tokenStr = tokenStr[7:]

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is invalid")
		}

		// Store the user object in the context
		c.Set("user", token)

		return next(c)
	}
}

// this is a middleware that checks users roles
func checkUserRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Retrieve the JWT token from the request header
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing JWT token")
			}

			// Parse the JWT token and extract the claims
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Provide the secret key used to sign the token
				// Replace "your-secret-key" with your actual secret key
				return []byte("your-secret-key"), nil
			})

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT token")
			}

			// Check if the token is valid
			if !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT token")
			}

			// Extract the user information from the token's claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT claims")
			}

			// Check if the user has the required role
			roles, ok := claims["roles"].([]interface{})
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "User roles not found in JWT claims")
			}

			hasRole := false
			for _, r := range roles {
				if r == role {
					hasRole = true
					break
				}
			}

			if !hasRole {
				return echo.NewHTTPError(http.StatusForbidden, "You don't have the required role")
			}

			return next(c)
		}
	}
}
