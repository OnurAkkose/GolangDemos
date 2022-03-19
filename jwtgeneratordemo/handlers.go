package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAxXH05CS7qG9gZxMPBL2TemZLNp9Hn8Jyaklb7PfDs4rbKtkc
UWGRfHdqO1cOYsMuuRNp5iOyeuDxR9YgnngrNPxcynqY/wAuoZFLNtCjPT1SQnr/
8neSNs2Jm57yIgaUWlMj3Xf9T2orzVwX0bmo+R4EQHYLhNgZ6ETHWj8i4/CXme67
v+yyqGiPtKkZ10XNBmGb9QV71kfuH9B+M74xoaGH1EnXcCTGaob0URk7sr6nKZjf
16WWPb4DwkdaHmrt3B/JIHUfK5iQ8fRSWCKtpZ4FeDVaFHa+IGIQdbTmoh2tgH/1
eh9QLQLxfpysAYv3hg+Jclg25TQRpLmoYL1TVQIDAQABAoIBAHhrMfJavODTXLZA
l25KuMFz1fdwWVyEWQedyiTST1cXHugZFf5ERVjl87JRPALcC4jw4CtuJhJvUnh/
jdFYdPGR4H3VG66aS9ZaKMc7o/NbjOni3mrgphoqbPyuifpajOwxvLP71FA07pYG
zrgoRXf56Jnv5MKWkeZo3DScQQlGxlnL42cTQ9jKz+X73GDgyAdwQo4kJ5bWSPtM
BCBeZ4U9X6/aQmwdJtf41YeBfZH/bt04U6AsxfEr2wfZ4Ja9pAh4tFlre/cwA1eJ
OjAAxFC9NyW3LSGN5FkmMDdXo2j5Id4j2UxiT9j4R7mTCC91AH0+KU8fSHV5OoDu
vbXjS0ECgYEA7mL9o+k8dQLeqvHoitw1NKQw/yzGWWYEuR+BUp/TjzcpYS31KkgP
dIQKtq4GgrPk+K9vyN840d2osTZBePmHYnKZWMogjdqcpe/frFZTu5flX8GVGPci
hJP9KZQHHxQhYK1q24+12KLnbM+o0wi3o5T2LG13vjyLwKpauIKdSVkCgYEA1AiR
yC6sYRPgWqExSyYsrk/r6h/B3wCdY0j1fm2JGlDr12bSMitNHMjVSJZIWX2H6Y0I
xAMeS3Cvkw7avygABzdiNT7nEzbaQ4aAswOEzS2LGNpc24riDas5MxbDROsYWQYh
ajRtqsHDsPEITlfhw/kZ1XJenDakI/oReSGjXl0CgYArYjqU1QKarO4HIEVY2CCa
tvLvzZ8/b9CWPESV295tpvVR6UI/8qNVah5lBqDKsqCOHadzYCSAFR35Ok2KDad0
5DRCM27fQhTWIiSLwZ41erxUw+81fux6QlCFe5ocLtawOH9E/A91IJiLdfNcjK6X
B6oRhc6QAYbRhm9COwsS6QKBgGLC8xRpGQXGulO8jTdRurIeq+ZLkIQMx0J9s0uG
PbwyQQf97p5LqQllSmMbDOwSGoJgnNqgETZWcJFw89biDNFPrMDcYcmDXTripYO8
edkQA6A55dKk6BIx1NzFF4M8dgTRZDMR5JsK0dnTC2liadhcaPoQ4ZylnuLbEdR4
JJ5JAoGBAKu36KD6dXlkcnww8n1Y/V7HRfnNjFzPktzqvDlx7FQo+jHHnpU5L3vc
ZZPGOfO2hASwgmGHYJaWJ/Gx4B06b0NgKN/Xeet1GwkVC2QqWVO0KY3NItAVzBFM
uCDa+PyjJ32rfrK8s9VvPBDh3lx6DWYrgAuffhmMwcICmL9ex9Gz
-----END RSA PRIVATE KEY-----`

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

//
type Claims struct {
	Permission string `json:"permission"`
	Role       string `json:"role"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	expectedPassword, ok := users[creds.Username]

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		//Usrname: creds.Username,
		Role:       "user",
		Permission: "read",
		StandardClaims: jwt.StandardClaims{
			Audience:  "AUDIENCE",
			Issuer:    "ISSUER",
			Subject:   "SUBJECT",
			Id:        "h69UsZiHQpNQDhGGXTCPiQ",
			IssuedAt:  expirationTime.Unix(),
			NotBefore: expirationTime.Unix(),
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	jwtParsedKey, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(jwtKey))
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	//pkey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(jwtKey))
	// Create the JWT string
	tokenString, err := token.SignedString(jwtParsedKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s!", "claims.Username")))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	// (BEGIN) The code uptil this point is the same as the first part of the `Welcome` route
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// (END) The code uptil this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
