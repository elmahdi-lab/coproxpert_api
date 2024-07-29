package middleware_test

//
//import (
//	"errors"
//	"github.com/joho/godotenv"
//	"net/http"
//	"testing"
//	"time"
//
//	"github.com/gofiber/fiber/v2"
//	"github.com/google/uuid"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"ithumans.com/coproxpert/middleware"
//	"ithumans.com/coproxpert/models"
//)
//
//type UserRepositoryMock struct {
//	mock.Mock
//}
//
//func (m *UserRepositoryMock) FindByToken(token string) (*models.User, error) {
//	args := m.Called(token)
//	return args.Get(0).(*models.User), args.Error(1)
//}
//
//func TestAuthMiddlewareWithValidToken(t *testing.T) {
//	err := godotenv.Load("../.env")
//	if err != nil {
//		t.Fatalf("Failed to load .env file: %s\n", err)
//	}
//	userRepoMock := new(UserRepositoryMock)
//	validToken := uuid.New()
//	expiryTime := time.Now().Add(time.Hour)
//	user := models.User{Token: &validToken, TokenExpiresAt: &expiryTime}
//	userRepoMock.On("FindByToken", validToken).Return(&user, nil)
//
//	app := fiber.New()
//	app.Use(middleware.AuthMiddleware())
//	req, _ := http.NewRequest("GET", "/", nil)
//	req.Header.Set("Authorization", validToken.String())
//	resp, _ := app.Test(req)
//
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//	userRepoMock.AssertExpectations(t)
//}
//
//func TestAuthMiddlewareWithInvalidToken(t *testing.T) {
//	userRepoMock := new(UserRepositoryMock)
//	userRepoMock.On("FindByToken", "invalidToken").Return(models.User{}, errors.New("invalid token"))
//
//	app := fiber.New()
//	app.Use(middleware.AuthMiddleware())
//	req, _ := http.NewRequest("GET", "/", nil)
//	req.Header.Set("Authorization", "invalidToken")
//	resp, _ := app.Test(req)
//
//	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
//	userRepoMock.AssertExpectations(t)
//}
//
//func TestAuthMiddlewareWithExpiredToken(t *testing.T) {
//	userRepoMock := new(UserRepositoryMock)
//	expiredToken := uuid.New()
//	expiredTime := time.Now().Add(-time.Hour)
//
//	user := models.User{Token: &expiredToken, TokenExpiresAt: &expiredTime}
//	userRepoMock.On("FindByToken", "expiredToken").Return(user, nil)
//
//	app := fiber.New()
//	app.Use(middleware.AuthMiddleware())
//	req, _ := http.NewRequest("GET", "/", nil)
//	req.Header.Set("Authorization", "expiredToken")
//	resp, _ := app.Test(req)
//
//	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
//	userRepoMock.AssertExpectations(t)
//}
//
//func TestAuthMiddlewareWithoutToken(t *testing.T) {
//	userRepoMock := new(UserRepositoryMock)
//
//	app := fiber.New()
//	app.Use(middleware.AuthMiddleware())
//	req, _ := http.NewRequest("GET", "/", nil)
//	resp, _ := app.Test(req)
//
//	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
//	userRepoMock.AssertExpectations(t)
//}
