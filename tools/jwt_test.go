package tools_test

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
	"time"
	"weicai.zhao.io/tools"
)

func TestJwt(t *testing.T) {
	type claim struct {
		Key string
		jwt.RegisteredClaims
	}
	convey.Convey("test jwt sign and parse: ", t, func() {
		var (
			claims = &claim{
				Key: "value",
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer: "asd",
				},
			}
		)
		tokenStr, err := tools.JwtSign(claims)
		convey.So(err, convey.ShouldBeNil)

		var (
			newClaim = &claim{
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer: "asd",
				},
			}
		)

		tmpClaim, err := tools.JwtParse(tokenStr, newClaim)

		convey.So(err, convey.ShouldBeNil)
		convey.So(tmpClaim.(*claim).Key, convey.ShouldEqual, claims.Key)
	})

	convey.Convey("test jwt sign and parse: ", t, func() {
		var (
			claims = &claim{
				Key: "value",
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "asd",
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Second)),
				},
			}
		)
		tokenStr, err := tools.JwtSign(claims)
		convey.So(err, convey.ShouldBeNil)

		time.Sleep(time.Second * 1)
		var (
			newClaim = &claim{
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer: "asd",
				},
			}
		)

		_, err = tools.JwtParse(tokenStr, newClaim)
		log.Println("this should be err", err)
		convey.So(err, convey.ShouldNotBeNil)
		//convey.So(err, convey.ShouldEqual, jwt.ErrTokenExpired)
	})
}
