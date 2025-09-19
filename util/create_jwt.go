package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"` //user id
	Name        string `json:"name"`
	Email       string `json:"Email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	hearderB64 := Base64UrlEncoding(byteArrHeader)

	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	payloadB64 := Base64UrlEncoding(byteArrData)

	message := hearderB64 + "." + payloadB64

	byteArrSecrate := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecrate)

	h.Write(byteArrMessage)

	signature := h.Sum(nil)

	signatureB64 := Base64UrlEncoding(signature)

	jwt := message + "." + signatureB64

	return jwt, nil
}

func Base64UrlEncoding(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
