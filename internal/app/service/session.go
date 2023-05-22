package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-finance/internal/app/model"
	"github.com/go-finance/internal/app/schema"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserSession interface {
	GetUserByUsername(username string) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	GetUserByID(ID int) (model.User, error)
}

type AuthRepo interface {
	GetByID(userID int, refreshToken string) (model.Auth, error)
	CreateSession(auth model.Auth) error
	DeleteByUserID(userID int) error
}

type TokenGenerator interface {
	GenerateAccessToken(userID int) (string, time.Time, error)
	GenerateRefreshToken(userID int) (string, time.Time, error)
}

type SessionService struct {
	userRepo      UserSession
	authRepo      AuthRepo
	generateToken TokenGenerator
}

func NewSessionService(
	userRepo UserSession,
	authRepo AuthRepo,
	generateToken TokenGenerator,
) *SessionService {
	return &SessionService{
		userRepo:      userRepo,
		authRepo:      authRepo,
		generateToken: generateToken,
	}
}

func (ss *SessionService) Login(req *schema.LoginReq) (schema.LoginResp, error) {
	var resp schema.LoginResp

	// find existing user by userID
	existingUser, _ := ss.userRepo.GetUserByEmail(req.Email)
	if existingUser.ID <= 0 {
		log.Error("unable to get user by email")
		return resp, errors.New("reason.FailedLogin")
	}

	// verify password
	isVerified := ss.verifyPassword(existingUser.Password, req.Password)
	if !isVerified {
		log.Error("invalid password")
		return resp, errors.New("reason.FailedLogin")
	}

	// generate access token
	accessToken, _, err := ss.generateToken.GenerateAccessToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("access token creation : %w", err))
		return resp, errors.New("reason.FailedLogin")
	}

	// generate refresh token
	refreshToken, expiredAt, err := ss.generateToken.GenerateRefreshToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("refresh token creation : %w", err))
		return resp, errors.New("reason.FailedLogin")
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	// save refresh token
	authPayload := model.Auth{
		UserID:    existingUser.ID,
		Token:     refreshToken,
		AuthType:  "refresh_token",
		ExpiredAt: expiredAt,
	}
	err = ss.authRepo.CreateSession(authPayload)
	if err != nil {
		log.Error(fmt.Errorf("refresh token saving : %w", err))
		return resp, errors.New("reason.FailedLogin")
	}

	return resp, nil
}

func (ss *SessionService) Logout(UserID int) error {
	err := ss.authRepo.DeleteByUserID(UserID)
	if err != nil {
		log.Error(fmt.Errorf("refresh token saving : %w", err))
		return errors.New("reason.FailedLogout")
	}
	return nil
}

func (ss *SessionService) Refresh(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error) {
	var resp schema.RefreshTokenResp

	existingUser, _ := ss.userRepo.GetUserByID(req.UserID)
	if existingUser.ID <= 0 {
		return resp, errors.New("reason.FailedRefreshToken")
	}

	auth, err := ss.authRepo.GetByID(existingUser.ID, req.RefreshToken)
	if err != nil || auth.ID < 0 {
		log.Error(fmt.Errorf("error SessionService - refresh : %w", err))
		return resp, errors.New("reason.FailedRefreshToken")
	}

	accessToken, _, _ := ss.generateToken.GenerateAccessToken(existingUser.ID)

	resp.AccessToken = accessToken
	return resp, nil
}

func (ss *SessionService) verifyPassword(hashPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainPwd))
	return err == nil
}
