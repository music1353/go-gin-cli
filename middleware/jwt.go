package middleware

import (
    "log"
    "net/http"
    "time"
    "errors"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

// jwt secret key
var jwtSecret = []byte("secret")


type Claims struct {
    ID       int    `json:"id"`
    Account  string `json:"account"`
    Role     string `json:"role"`
    jwt.StandardClaims
}

// GenerateToken: generate tokens used for auth
func GenerateToken(id int, account string, role string) (token string, err error) {
    nowTime := time.Now()
    expireTime := nowTime.Add(3 * time.Hour) // token發放後多久過期

    claims := Claims{
        ID: id,
        Account: account,
        Role: role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expireTime.Unix(),
            IssuedAt:  nowTime.Unix(),
            Issuer:    "go-gin-cli",
        },
    }

    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err = tokenClaims.SignedString(jwtSecret)
    if err != nil {
        log.Println(err)
        return
    }

    return
}

// ParseToken: parsing token
func ParseToken(token string) (claims *Claims, err error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil {
        var errMsg string
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors & jwt.ValidationErrorMalformed != 0 {
                errMsg = "token is malformed"
            } else if ve.Errors & jwt.ValidationErrorUnverifiable != 0{
                errMsg = "token could not be verified because of signing problems"
            } else if ve.Errors & jwt.ValidationErrorSignatureInvalid != 0 {
                errMsg = "signature validation failed"
            } else if ve.Errors & jwt.ValidationErrorExpired != 0 {
                errMsg = "token is expired"
            } else if ve.Errors & jwt.ValidationErrorNotValidYet != 0 {
                errMsg = "token is not yet valid before sometime"
            } else {
                errMsg = "can not handle this token"
            }
        }
        err = errors.New(errMsg)
        return
    }

    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return
}


// jwt auth middleware
func AuthJwtMiddle() gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
	    token := strings.Split(auth, "Bearer ")[1]
        
        claims, err := ParseToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "result": "",
                "msg": err.Error(),
            })
            c.Abort()
            return
        }

        c.Set("id", claims.ID)
        c.Set("account", claims.Account)
        c.Set("role", claims.Role)

        c.Next()
        return
    }
}