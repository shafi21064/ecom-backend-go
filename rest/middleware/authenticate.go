package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
	"strings"

	"github.com/shafi21064/ecom-go/util"
)

func (m *Middlewares) AuthenticateJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		headerArr := strings.Split(header, " ")

		if len(headerArr) < 2 {
			util.SendError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")

		if len(tokenParts) != 3 {
			util.SendError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenHeader := tokenParts[0]
		tokenClaim := tokenParts[1]
		tokenSignature := tokenParts[2]

		message := tokenHeader + "." + tokenClaim

		byteArrSecrate := []byte(m.cofig.JwtSecrateKey)
		byteArrMessage := []byte(message)

		hash := hmac.New(sha256.New, byteArrSecrate)

		hash.Write(byteArrMessage)

		signature := hash.Sum(nil)

		newSignatureB64 := util.Base64UrlEncoding(signature)

		if newSignatureB64 != tokenSignature {
			util.SendError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
