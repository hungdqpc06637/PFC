package services

import (
	"errors"
	"fmt"

	"web-api/internal/pkg/database"
	"web-api/internal/pkg/models/request"
	"web-api/internal/pkg/models/types"

	"github.com/dgrijalva/jwt-go"

	"time"
)

type LoginService struct {
	*BaseService
}

var LG = &LoginService{}

// Hàm Login
func (s *LoginService) Login(requestParams *request.LoginRequest) (types.Login, error) {
	// db, err := database.ERPConnection()
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return types.Login{}, err
	}
	dbInstance, _ := db.DB()
	defer dbInstance.Close()

	// query := `
	// SELECT USERID, USERNAME, PWD, Role
	// FROM Busers
	// WHERE USERID = ? AND PWD = ?
	// `
	query := `    
	SELECT USERID, USERNAME, PWD
	FROM Busers
	WHERE USERNAME = ? AND PWD = ?
	`
	var result types.Login
	err = db.Raw(query, requestParams.Userid, requestParams.Password).Scan(&result).Error
	if err != nil {
		fmt.Println("Query error:", err)
		return types.Login{}, err
	}

	if result.UserId == "" {
		err = errors.New("user not found")
		fmt.Println("User not found error:", err)
		return types.Login{}, err
	}

	// Tạo JWT token
	tokenString, err := s.generateToken(result.UserId)
	if err != nil {
		fmt.Println("Token generation error:", err)
		return types.Login{}, err
	}

	result.Token = tokenString
	return result, nil
}

// Hàm tạo JWT token
func (s *LoginService) generateToken(userId string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	// Tạo claims cho token
	claims := &jwt.StandardClaims{
		Subject:   userId,
		ExpiresAt: expirationTime.Unix(),
	}

	// Tạo token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("TYXUAN@123"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
