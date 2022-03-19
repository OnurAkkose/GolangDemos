package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKeyy = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAxu0/1a97ooW7hDcosq85ueTg4N4/ph6zwvLDoJnxqCJRrNgb
6IajnECHgdYQkSdm3Vqj/bAMBToUpScAMUkJzHceQpQEUll1K4LJ+2Y0YzP3ICnO
FdXYCLrPhyKKa1YNzJbdenrf5m+Msf0Jni/7UulSnf+qtGVvagoTJWzinESJJhD5
jYUoIoILD/KOskkA87b9JVcd+RybRgs7JfBkJCe9hnzuEDmLNQnbWag+HUogXRes
d0HTkrF4svaPjfhzYmE/eDKyoUTwZgykgmB3ePHsYaN96HQA4gHEGSUI1YFAKPu0
wVWI/u433BePeVxTpk5pVdfThsvjhrCHP2QekwIDAQABAoIBAA0VETumXMUlcl2R
mxWVPICjMr7XHux33G6XtJHdTe02xJRPahZytFPUUpURVOmW48bu7RYD++ZkGXN7
CyIyoXhW5SCPx3D6/R+tbEciQ5O6mSf+V9VLyPdaAcfrV5sTf7jSuyIQ7qSi9WfH
Mli64xZizdIiPEG84gThQL0XWhfYtCMz2u7ZIitRVVpWCyH+YYsdGfbm+H4vmn3v
8Ecb7fjwc8rcpPLGfnm7WWzVnilnebqR2nMUT+ZUooDzn8RAQxTT/OVYFQKm/zoY
wDq6iU2qlufBfWRN/1HkiDTn0HP2VHbglFcFFRNyf+CEVgSvdErAEZyurTAwcwuF
LKh6l9ECgYEAyRoikcJ8ZiXJj6QgZsicm27xhC6qtWBqZsRCwAcG77GlpXE3b/9Z
rnWpACT42jQD9ZCMDN/kpLDsKNgvmkOErwSGBUni9WrsdsZGwEuY1etrRn9Hw9GG
S8quL+M0nvjaCeTLCm7wzicNkgR1TqEWKOKxvt7KJQAYnTGoHNAfnjUCgYEA/TsX
vM62fgKwmhbrXyZJeBSXjd76X3/cz8wfTmRk0SxlEKPMNy4/5FZBqeVE5GKVdDXw
p/hfgDegb7dAFbqXcdvzmKlSzoM0XnlnbmbYf3JyibWrIbCQmHfqFZw4FYgk/Sx8
+izFLzq9Ky/ifSGZ28EYi8WPDjjpfXRDNTxZgqcCgYEAin98VWRbJkJZ+ZowUnlR
Gd8jaER3fujC/rmluvhb95IiIbnCU1jKG9Oeq6HK3QQ7wdBmE6vSnPXX/x08U4Ky
i5KS9mt3akvURMyzB1ZJEPLMc8XO1/aiBeq1YfeZUu7Rw0SV7T6Qi3nr56c4Xwmj
6E6P2YM4NplFWmVLgWR4kg0CgYEAwze6DgEh6LT0JmZC36BphRwC1gku5U5yEPPg
spNssWDTLOfJeES8VrA0gOBRoutpIiSvOR6dqP+5PEZ+LgIh3FHfUjI+txuo1Kgt
F4xLnLzDFeyqWBeA8TmIiU5cYiUJtu+EDW1UOhvDV7bbmPG9Zg9Pd/k+Vo2DWwa1
BSZYLwMCgYBZpYyAS8LSl9APA0/lBCsXBHoFoCWNmrPEK25y/ZA8/RThTzfFA2lq
umiRWhKw0uFot18s/5RC9hTtj8f55b80erLs2hQ9nA3V7kdplYXaQ+738tS3Ytlk
YhZb7dGW2NSKMX5LE+jv2gm0wkq7fO7swWVIofQI0mrHXjxY0T9Ppw==
-----END RSA PRIVATE KEY-----`

func GetJWT() (string, error) {

	jwtParsedKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(jwtKeyy))
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "ISSUER"
	claims["sub"] = "SUBJECT"
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(jwtParsedKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	fmt.Fprintf(w, string(validToken))
}

func handleRequests() {
	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
