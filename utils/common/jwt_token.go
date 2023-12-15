package common

import (
	"errors"
	"final-project-kelompok-1/config"
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	modelutil "final-project-kelompok-1/utils/common/model_util"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken interface {
	GenerateToken(payload model.Users) (dto.AuthResponseDto, error)
	GenerateTokenStudent(payload model.Student) (dto.AuthResponseDto, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
	RefreshToken(oldTokenString string) (dto.AuthResponseDto, error)
}

type jwtToken struct {
	cfg config.TokenConfig
}

func (j *jwtToken) GenerateToken(payload model.Users) (dto.AuthResponseDto, error) {
	claims := modelutil.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.JwtLifeTime)),
		},
		UserId: payload.UserID,
		Role:   payload.Role,
	}

	jwtNewClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtNewClaims.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, errors.New("failed to generate token")
	}

	return dto.AuthResponseDto{Token: token}, nil
}

func (j *jwtToken) GenerateTokenStudent(payload model.Student) (dto.AuthResponseDto, error) {
	claims := modelutil.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.JwtLifeTime)),
		},
		UserId: payload.StudentID,
		Role:   payload.Role,
	}

	jwtNewClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtNewClaims.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, errors.New("failed to generate token")
	}

	return dto.AuthResponseDto{Token: token}, nil
}

func (j *jwtToken) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, errors.New("failed to verify token")
	}

	// kita convert dari token.Claims ke jwt.MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok || claims["iss"] != j.cfg.IssuerName {
		return nil, errors.New("invalid claim token")
	}

	return claims, nil
}

func (j *jwtToken) RefreshToken(oldTokenString string) (dto.AuthResponseDto, error) {
	// IMPLEMENT ME
	token, err := jwt.Parse(oldTokenString, func(token *jwt.Token) (any, error) {
		return j.cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok || claims["iss"] != j.cfg.IssuerName {
		return dto.AuthResponseDto{}, errors.New("invalid claim token")
	}

	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString(j.cfg.JwtSignatureKey)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return dto.AuthResponseDto{Token: newTokenString}, nil
}

func NewJwtToken(cfg config.TokenConfig) JwtToken {
	return &jwtToken{cfg: cfg}
}
