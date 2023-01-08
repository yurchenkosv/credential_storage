package interceptors

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// JWTInterceptor is a gRPC interceptor that validates a JWT and returns the user.
func JWTInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Extract the JWT from the request metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("failed to get metadata from context")
	}
	tokens := md.Get("authorization")
	if len(tokens) == 0 {
		return nil, errors.New("no authorization token found in metadata")
	}
	token := tokens[0]

	// Validate the JWT
	parser := jwt.Parser{
		ValidMethods: []string{jwt.SigningMethodHS256.Name},
	}
	claims, err := parser.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Replace with your own key and algorithm
		key := []byte("my-secret-key")
		if token.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	// Extract the user from the JWT
	// Replace with your own code to extract the user from the claims
	user, err := extractUserFromClaims(claims.Claims)
	if err != nil {
		return nil, err
	}

	// Set the user in the context and call the next handler
	ctx = context.WithValue(ctx, "user", user)
	return handler(ctx, req)
}

func extractUserFromClaims(claims jwt.Claims) (string, error) {
	// Check that the claims are of the expected type
	if claims, ok := claims.(jwt.MapClaims); ok {
		// Extract the user from the claims
		if user, ok := claims["user"].(string); ok {
			return user, nil
		}
		return "", errors.New("failed to extract user from claims")
	}
	return "", errors.New("invalid claims type")
}
