// JWT (Access Token):
// JWT is a compact, self-contained token format that consists of three parts: header, payload, and signature.
// The payload typically contains information about the user (such as user ID, username, and roles) and any other relevant data.
// After the user successfully authenticates, the server generates a JWT and sends it back to the client.
// The client includes the JWT in subsequent requests, usually in the Authorization header as a bearer token.
// Since JWTs are signed, servers can verify their integrity and authenticity without needing to store them. This makes JWTs stateless and scalable.

// Refresh Token:
// Refresh tokens are long-lived tokens that are issued alongside the JWT.
// Unlike JWTs, refresh tokens are stored securely on the server-side.
// When the JWT expires (which typically happens after a short duration to enhance security), the client sends the refresh token to the server to obtain a new JWT.
// The server validates the refresh token and issues a new JWT if the refresh token is valid.
// Refresh tokens help mitigate some security risks associated with JWTs. Since they are not sent with every request, the exposure window is smaller, reducing the impact of token theft.
// If a refresh token is compromised or expires, the user is required to re-authenticate, enhancing security.

// Here's a typical flow for using JWT and refresh tokens:

// User Authentication:
// User provides credentials (e.g., username and password) to the server for authentication.
// The server verifies the credentials and issues a JWT and a refresh token if authentication is successful. The JWT contains a short expiration time (e.g., 15 minutes), while the refresh token has a longer expiration time (e.g., 7 days).
// Both tokens are sent back to the client.

// Accessing Protected Resources:
// The client includes the JWT in the Authorization header when accessing protected resources.
// The server verifies the JWT's signature and expiration time to ensure its validity.

// Refreshing the JWT:
// When the JWT expires, the client sends the refresh token to the server.
// The server verifies the refresh token and issues a new JWT if the refresh token is valid.
// If the refresh token is expired or invalid, the server requires the user to re-authenticate.

// Logout and Token Revocation:
// To log out or invalidate tokens (e.g., in case of a security breach), the server can maintain a blacklist of invalidated tokens or use other mechanisms to revoke tokens.

// This combination of JWTs and refresh tokens provides a balance between security, scalability, and usability in web application authentication and authorization.

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// In a real application, you may authenticate the user with username and password
	// and issue a JWT accordingly. For simplicity, we're just hardcoding a username here.
	username := "exampleUser"

	// Create the JWT claims, which includes the username and expiry time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create the token with the claims and sign it using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error generating JWT token: %v", err)
		return
	}

	// Send the token as JSON response
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"token": "%s"}`, tokenString)
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		// Parse and validate the token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized: %v", err)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid token")
			return
		}

		// Call the next handler if token is valid
		next.ServeHTTP(w, r)
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Protected resource accessed successfully")
}

func main() {
	// Define HTTP routes
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/protected", authMiddleware(protectedHandler))

	// Start the HTTP server
	fmt.Println("Server running on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
