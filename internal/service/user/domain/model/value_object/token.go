package value_object

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserToken struct {
	ID   string
	Name string
	//Phone  string
	Claims jwt.StandardClaims
}

// DefaultClaims -
func (t *UserToken) DefaultClaims() {
	u := time.Now().Unix()
	t.Claims = jwt.StandardClaims{
		Audience: WOAH,
		Issuer:   TYR,
		IssuedAt: u,
		Subject:  fmt.Sprintf("%s-%s", MEMBERTOKEN, md5(MEMBERTOKEN, strconv.FormatInt(u, 10))),
	}
}

func md5(str1, str2 string) string {
	h := sha256.New()
	x := fmt.Sprintf("%s-%s-%s", str1, "ax", str2)
	h.Write([]byte(x))
	return hex.EncodeToString(h.Sum(nil))
}
