package usecase

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	usecase2 "go_online_course_v2/internal/admin/usecase"
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
	Refresh(dtoRefreshToken dto.RefreshTokenRequestBody) (*dto.LoginResponse, *response.Errors)
	Logout(accessToken string) *response.Errors
}

type oauthUseCase struct {
	oauthClientRepository       repository.OauthClientRepository
	oauthAccessTokenRepository  repository.OauthAccessTokenRepository
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository
	userUseCase                 usecase.UserUseCase
	adminUseCase                usecase2.AdminUseCase
}

func (useCase *oauthUseCase) Logout(accessToken string) *response.Errors {
	//	find by accessToken on oauth_access_token table
	oauthAccessToken, err := useCase.oauthAccessTokenRepository.FindOneByAccessToken(accessToken)
	if err != nil {
		return err
	}

	//	find by oauth_refresh_token_id
	oauthRefreshToken, err := useCase.oauthRefreshTokenRepository.FindOneByOauthAccessTokenID(int(oauthAccessToken.ID))
	if err != nil {
		return err
	}

	//	delete data
	useCase.oauthRefreshTokenRepository.Delete(*oauthRefreshToken)
	useCase.oauthAccessTokenRepository.Delete(*oauthAccessToken)
	return nil
}

func (useCase *oauthUseCase) Refresh(dtoRefreshToken dto.RefreshTokenRequestBody) (*dto.LoginResponse, *response.Errors) {
	//	check oauth refresh token based from refresh token data
	oauthRefreshToken, err := useCase.oauthRefreshTokenRepository.FindOneByToken(dtoRefreshToken.RefreshToken)
	if err != nil {
		return nil, err
	}

	if oauthRefreshToken.ExpiredAt.Before(time.Now()) {
		return nil, &response.Errors{
			Code: 400,
			Err:  errors.New("oauth refresh token has expired"),
		}
	}

	var user dto.UserResponse

	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	if *oauthRefreshToken.OauthAccessToken.OauthClientID == 2 {
		dataAdmin, _ := useCase.adminUseCase.FindByID(int(oauthRefreshToken.UserID))

		user.ID = dataAdmin.ID
		user.Name = dataAdmin.Name
		user.Email = dataAdmin.Email
	} else {
		dataUser, _ := useCase.userUseCase.FindOneByID(int(oauthRefreshToken.UserID))

		user.ID = dataUser.ID
		user.Name = dataUser.Name
		user.Email = dataUser.Email
	}

	claims := &dto.ClaimResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	if *oauthRefreshToken.OauthAccessToken.OauthClientID == 2 {
		claims.IsAdmin = true
	}
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errSignedString := token.SignedString(jwtKey)

	if errSignedString != nil {
		return nil, &response.Errors{
			Code: 500,
			Err:  errSignedString,
		}
	}

	//	insert to oauth access token table
	dataOauthAccessToken := entity.OauthAccessToken{
		OauthClientID: oauthRefreshToken.OauthAccessToken.OauthClientID,
		UserID:        oauthRefreshToken.UserID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt:     &expirationTime,
	}

	saveOauthAccessToken, err := useCase.oauthAccessTokenRepository.Create(dataOauthAccessToken)

	if err != nil {
		return nil, err
	}

	expirationTimeOauthRefreshToken := time.Now().Add(24 * 366 * time.Hour)

	//	insert to oauth refresh token table
	dataOauthRefreshToken := entity.OauthRefreshToken{
		OauthAccessTokenID: &saveOauthAccessToken.ID,
		UserID:             oauthRefreshToken.UserID,
		Token:              utils.RandString(128),
		ExpiredAt:          &expirationTimeOauthRefreshToken,
	}

	//TODO check to make sure access & refresh token is unique

	saveOauthRefreshToken, err := useCase.oauthRefreshTokenRepository.Create(dataOauthRefreshToken)
	if err != nil {
		return nil, err
	}

	//	delete old oauth refresh token
	err = useCase.oauthRefreshTokenRepository.Delete(*oauthRefreshToken)
	if err != nil {
		return nil, err
	}

	//	delete old oauth access token
	err = useCase.oauthAccessTokenRepository.Delete(*oauthRefreshToken.OauthAccessToken)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  tokenString,
		RefreshToken: saveOauthRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
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

	//check middleware role
	if oauthClient.Name == "web-admin" {
		dataAdmin, err := useCase.adminUseCase.FindByEmail(dtoLoginRequestBody.Email)
		if err != nil {
			return nil, &response.Errors{
				Code: 400,
				Err:  errors.New("username or password is invalid"),
			}
		}
		user.ID = dataAdmin.ID
		user.Email = dataAdmin.Email
		user.Name = dataAdmin.Name
		user.Password = dataAdmin.Password

	} else {
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
	}

	//	define JWT
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	//	compare password using jwt
	errorBcrypt := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dtoLoginRequestBody.Password))

	if errorBcrypt != nil {
		return nil, &response.Errors{
			Code: 400,
			Err:  errors.New("username or password is invalid"),
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

	if oauthClient.Name == "web-admin" {
		claims.IsAdmin = true
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
	adminUseCase usecase2.AdminUseCase,
) OauthUseCase {
	return &oauthUseCase{
		oauthClientRepository:       oauthClientRepository,
		oauthAccessTokenRepository:  oauthAccessTokenRepository,
		oauthRefreshTokenRepository: oauthRefreshTokenRepository,
		userUseCase:                 userUseCase,
		adminUseCase:                adminUseCase,
	}
}
