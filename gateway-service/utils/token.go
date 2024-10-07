package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

var privateKey *rsa.PrivateKey

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
