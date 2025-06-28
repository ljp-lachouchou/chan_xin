package ctxdata

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

const Identify = "github.com/ljp-lachouchou/chan_xin"
const PageIdentify = "github.com/ljp-lachouchou/chan_xin/page_token"

func GetToken(secretKey string, iat, seconds int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[Identify] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func GenPageToken(secretKey string, iat, seconds int64, pageLocation int32) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[PageIdentify] = pageLocation
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func ParseByTokenString(secretKey []byte, token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
}
