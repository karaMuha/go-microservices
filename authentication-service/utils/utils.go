package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// used to sign jwt token
var privateKey *rsa.PrivateKey

// used to validate reqeust body
var Validator *validator.Validate

func ReadPrivateKeyFromFile(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	data, _ := pem.Decode(buffer)
	key, err := x509.ParsePKCS8PrivateKey(data.Bytes)

	if err != nil {
		return err
	}

	if key, ok := key.(*rsa.PrivateKey); ok {
		privateKey = key
		return nil
	}

	return errors.New("error while reading private key")
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func GenerateJwt(userId string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user_id":   userId,
		"user_role": role,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString(privateKey)
}

func VerifyJwt(jwtToken string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return privateKey.Public(), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return parsedToken, nil
}
