package usecase

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go_online_course_v2/internal/oauth/dto"
	"go_online_course_v2/internal/oauth/entity"
	"go_online_course_v2/internal/oauth/repository"
	"go_online_course_v2/internal/user/usecase"
	"go_online_course_v2/pkg/response"
	"go_online_course_v2/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type OauthUseCase interface {
	Login(dtoLoginRequestBody dto.LoginRequest) (*dto.LoginResponse, *response.Errors)
}

type oauthUseCase struct {
	oauthClientRepository       repository.OauthClientRepository
	oauthAccessTokenRepository  repository.OauthAccessTokenRepository
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository
	userUseCase                 usecase.UserUseCase
}

func (useCase *oauthUseCase) Login(dtoLoginRequestBody dto.LoginRequest) (*dto.LoginResponse, *response.Errors) {
	//	Check isValid clientID and clientSecret
	oauthClient, err := useCase.oauthClientRepository.FindByClientIDAndClientSecret(
		dtoLoginRequestBody.ClientID,
		dtoLoginRequestBody.ClientSecret,
	)
	if err != nil {
		return nil, err
	}

	var user dto.UserResponse

	dataUser, err := useCase.userUseCase.FindByEmail(dtoLoginRequestBody.Email)

	if err != nil {
		return nil, &response.Errors{
			Code: 400,
			Err:  errors.New("username or password is invalid"),
		}
	}

	//set data user for login
	user.ID = dataUser.ID
	user.Email = dataUser.Email
	user.Name = dataUser.Name
	user.Password = dataUser.Password

	//	define JWT
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	//	compare password using jwt
	errorBcrypt := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dtoLoginRequestBody.Password))

	if errorBcrypt != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  errorBcrypt,
		}
	}

	//set expirationTime for jwt
	expirationTime := time.Now().Add(24 * 365 * time.Hour)
	claims := dto.ClaimResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	//	insert data to oauth_access_tokens table
	dataOauthAccessToken := entity.OauthAccessToken{
		OauthClientID: &oauthClient.ID,
		UserID:        user.ID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt:     &expirationTime,
	}

	oauthAccessToken, err := useCase.oauthAccessTokenRepository.Create(dataOauthAccessToken)

	if err != nil {
		return nil, err
	}

	//set expiration_time oauthAccessToken
	expirationTimeOauthAccessToken := time.Now().Add(24 * 366 * time.Hour)

	//	insert data to oauth_refresh_tokens table
	dataOauthRefreshToken := entity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccessToken.ID,
		UserID:             user.ID,
		Token:              utils.RandString(128),
		ExpiredAt:          &expirationTimeOauthAccessToken,
	}

	oauthRefreshToken, err := useCase.oauthRefreshTokenRepository.Create(dataOauthRefreshToken)
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		AccessToken:  oauthAccessToken.Token,
		RefreshToken: oauthRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

func NewOauthUseCase(
	oauthClientRepository repository.OauthClientRepository,
	oauthAccessTokenRepository repository.OauthAccessTokenRepository,
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository,
	userUseCase usecase.UserUseCase,
) OauthUseCase {
	return &oauthUseCase{
		oauthClientRepository:       oauthClientRepository,
		oauthAccessTokenRepository:  oauthAccessTokenRepository,
		oauthRefreshTokenRepository: oauthRefreshTokenRepository,
		userUseCase:                 userUseCase,
	}
}
