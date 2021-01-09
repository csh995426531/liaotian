package handler

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"liaotian/domain/auth/proto"
	"time"
)

type Handler struct {

}

func (h *Handler) Generated(ctx context.Context, request *proto.GeneratedRequest, response *proto.GeneratedResponse) error {

	if request.UserId == 0 || request.Name == "" {
		return ErrorBadRequest
	}

	secret := "k0xdv7apeo21sfjo"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": request.UserId,
		"name": request.Name,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Unix() + 86400,
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return ErrorInternalServerError(err)
	}

	response.Data = tokenString
	response.Message = "success"
	return nil
}

func (h *Handler) Parse(ctx context.Context, request *proto.ParseRequest, response *proto.ParseResponse) error {
	if request.Token == "" {
		return ErrorBadRequest
	}

	secret := "k0xdv7apeo21sfjo";
	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil {
		return ErrorInternalServerError(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		response.Data = &proto.User{
			UserId: int64(claims["user_id"].(float64)),
			Name: claims["name"].(string),
		}
		response.Message = "success"
		return nil
	} else {
		return ErrorInternalServerError(err)
	}
}