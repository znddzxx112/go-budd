package utils

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"io/ioutil"
	"testing"
	"time"
)

const (
	PRIVATEKEY = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAr1SJsILc1Bu6ezkB74ZpUjvB/XhHmxLlC7o/JelXPyWVFny9
71EsvNws9W9ol9K/zicVLOFvwJ1H/tasVizUoXTVulfSPV/RT0Ih0yFsy1cuPAAE
kY6NQE2B351C0lkKgFVs8WzkqCqTDEsq5nEActNL1wwnLRSTgWdiCp3jaWmZhrP2
OxCmKmOqlXZwvKXYCynN+lPFwYs82DMxlvbvhCMNNhUDN2pRMrDi2CrKaUIEJjlz
q7/9mGuKh8J+kyOCPA4+rA18E3Xvgwdr9jU9UBv3+rYLPmmQt3huTZt4wB+AvnGs
s+zFef0xPuLWOA4NeG69ACak+lNhsJ0F2B/ruwIDAQABAoIBAA50XsGhmEkYXCBq
i8FOiIJSEOUvtF+RiDaWTTx350x6cbcf45zGSXTshfxaCfpyUoPRbwp7L3ZmwRe+
ZQjZu1HwSuqI8PlEXAI3xogcelatQk+KBNZnNTf2690envKuipIX+NpSiuteTy81
Rz0mVc4ho2AATX6G8r45YrjTWXtIKYQ2zM8PacZHAPq/2FLs2hijsQnO/k0vDcai
DEO1tFq5FX0hg7YLLO5NY+vlBU+KtxqXgoKxw6o62PgLVj792b6gDD+Rf0NFD3sV
a+ovfabsbAIkWQXXna8eQYPDB8vEuxU4auegh06hK6x75j3fM8hld5rpFtHxTMZc
ZaB1GGkCgYEA1LEo22U6VmK+v0jyYvBNmyyjFXFGEL2My8amfs5tXPPoiSS993l/
jDhiND3djyPPOvN9gNUK8e0jFjvVfimqBD2DtxJ3yv/WIhBgGb2tEBig7DAuVmOF
R2AfXP0PDDUGApIrqnskw/27eNsbN16Y6C6O/5+XCSoFNw/t/GoII2UCgYEA0wfY
2yloKJAfxCBdb1HC8F910jE1H3j2An71++diSmG6GmXSAdwy7ppkj76A6iDSUq78
Vap81aeoZD1RY2xD9sskehEvyHJS3hJhPp+pTb91+sEt3Tb4ZgT0frd7VyL6dx2P
xmy839ZaWBSULXozWYWD0gR3dQHQGU1hMLdDMJ8CgYAUMIDtWo2QF619oOIJTEBB
o5oTdf4tpqCP01qInPW6phiDtC4oKKtX1x5TUTAP31lTgjt+mDqCvnzfJmzcnf3a
izBOl30xktlzwFQu/VdJON1FrkknwCitns07WXYtNXdNlSx58ViLFjHOhhXuSpd/
KrQ+RZrjrs5x7Jwtoq8FFQKBgQCC+KHMEzzDv/8XGAclKZdU28oh88xGhioYjndY
KdjE1TZNX9ggs2sgzs2OsOsNY+Vkz5loCXGEoJNP8wZqMC1WI+m1oKkJPHrVvI6B
8VGAIU21nOM8Ifu0RWsAoht4jwrnln8+5Qmr2jsj41G7j9YCv2x6ka8Z/VAMBAxH
0dVvfQKBgFSxDRxQK75AmUWMu+eZ/9OE7cVlsd2z54313Ca/2KtExmSwY+pOp9SK
7L6AI5jKRhhY5ExmyFlYwAU2XhG8aoKtlp/icObR1FDkQHb5rmvZv+7RBF/z2NsY
r1qweZYXZn0Swwfr/lLC4p81YvrQ7Dg89zEci84/lQtfKztMy4NX
-----END RSA PRIVATE KEY-----`
	PUBLICKEY = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAr1SJsILc1Bu6ezkB74Zp
UjvB/XhHmxLlC7o/JelXPyWVFny971EsvNws9W9ol9K/zicVLOFvwJ1H/tasVizU
oXTVulfSPV/RT0Ih0yFsy1cuPAAEkY6NQE2B351C0lkKgFVs8WzkqCqTDEsq5nEA
ctNL1wwnLRSTgWdiCp3jaWmZhrP2OxCmKmOqlXZwvKXYCynN+lPFwYs82DMxlvbv
hCMNNhUDN2pRMrDi2CrKaUIEJjlzq7/9mGuKh8J+kyOCPA4+rA18E3Xvgwdr9jU9
UBv3+rYLPmmQt3huTZt4wB+AvnGss+zFef0xPuLWOA4NeG69ACak+lNhsJ0F2B/r
uwIDAQAB
-----END PUBLIC KEY-----`
)

//jwt加密
func TestJwtEnDoce(t *testing.T) {
	uuid, _ := uuid.NewV4()
	claims := jwt.MapClaims{
		"externalId":    uuid.String(),
		"udAccountUuid": "test123",
		"name":          "张**",
		"mobile":        "132****8613",
		"email":         "710***334@QQ,COM",
		"exp":           time.Now().Unix() + 120,
	}
	/*	priKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(PRIVATEKEY))
		if err != nil {
			fmt.Println("解析私钥错误:", err)
			return
		}*/
	token, err := Signed(claims, []byte("123456"))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(token)
}

type token struct {
	ExternalId string `json:"externalId"`
	UUID       string `json:"udAccountUuid"`
	Name       string `json:"name"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	Expired    int64  `json:"exp"`
}

const PB = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3jhqQzom5vqUeoWgwpojBW4iWi3G6zMw
QiDJPNKZg2uNZAmWmUDam22slOm1ibA/4vShfm/NH0Ry5ojp7KlXivOps+gY/brwpxS0WotUHAWj
T+FjbRftRDN55OKrYLHWrfaQhmBTji0TMeZkWlsYUCsXzAxdTW5SCQviJki6kdH2Mv0U5ZDKKE4U
NC+JQQFVld70cv65C93FZ/6eBY4rW7XUeJsTqpYl2ADo00kpPjxcUl0pJP4XBu/D0prMT1tNjnp+
CjTyKC1hji1Mrnl9TKZgcyFk8jM9T9MuryKQlBJfN90E1Vdok6xEVdeSYSmv99FJ/hoRnooU4NPz
qTQp2wIDAQAB
-----END PUBLIC KEY-----`

func TestJwtDeDoce(t *testing.T) {
	tokenString := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjQ4Njk0ODk4NDAxOTkyMjkyNTQifQ.eyJlbWFpbCI6IioqKioqQHRhbmppLmNvbS5jbiIsIm5hbWUiOiIqKirosK3orrAiLCJtb2JpbGUiOiIxNDUqKioqODkwOSIsImV4dGVybmFsSWQiOiIzODA5MDc2NzA5NDIyMjk0ODgiLCJ1ZEFjY291bnRVdWlkIjoiMDMxY2E5ZjUxMDliOTMxMDcxOWMxYzNjYzYyMDcyMDBNU1RUZ3ZXSTB6UiIsIm91SWQiOiI1MzgzOTQwMzA3Mjg4ODc2MTEwIiwib3VOYW1lIjoi5rWZ5pS_6ZKJIiwib3BlbklkIjpudWxsLCJpZHBVc2VybmFtZSI6IioqKmppIiwidXNlcm5hbWUiOiIqKioqKioqbGkiLCJhcHBsaWNhdGlvbk5hbWUiOiJKV1QiLCJleHRlbmRGaWVsZHMiOnsidGhlbWVDb2xvciI6ImdyZWVuIiwiYXBwTmFtZSI6IkpXVCIsInloVWlkIjoic2hhbmdzaGFuZGFsYW95dTEyMyIsImRmVXNlcklkIjpudWxsfSwiZXhwIjoxNTY1OTM4Mjc5LCJqdGkiOiI3ZFdoNWwyYTB4MThCZUdrbXlkR1FBIiwiaWF0IjoxNTY1OTM3Njc5LCJuYmYiOjE1NjU5Mzc2MTksInN1YiI6Im1laWxpbWVsaSJ9.ILnj0YyEIeBwKp4JbspE_Ix28NKoVHsS-DnKzLBTomGwIo-zu_G7MTMEjXQIDY_BUEZd0bITp8oFlgTwM5j5BhqyoW4f2ARafErO9nkdxYGMeLPabvEKeN6I8Pl-nq_TBe0uEE1N8-fEfWQaKz5Zd_OvEjxsCaQh-k5pi80qWyjqhOe1TQoSciE9n38_m5bMA05sYdV3vhXLLnY2jgtF6n_KfEssBKrMT59maFjqbSMhUIJg5hJHu_6jgPnfh_LnW5Ep9D1AhgaLrY3iszZzjuYyAvrE9xuW_bpnHZShUradDXKuxax_GAh4Nrv8DLO2CqsrG0KMeTG5S_gq8yEXYw"
	//pb, err := jwt.ParseRSAPublicKeyFromPEM([]byte(PB))
	//if err != nil {
	//	fmt.Println("公钥解析错误:", err)
	//	return
	//}
	pubkey, err := ioutil.ReadFile("../../resources/rsa/public_key_ud.pem")
	block, _ := pem.Decode([]byte(pubkey))
	pb, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("公钥解析错误:", err)
		return
	}
	info, result := Pares(tokenString, pb)
	t.Log(result)
	t.Log(info)
	infoBytes, err := json.Marshal(info)
	if err != nil {
		t.Fatal(err.Error())
	}
	record := make(map[string]interface{})
	err = json.Unmarshal(infoBytes, &record)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(record)
}

func TestJwtDeDoce2(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjcxMCoqKjMzNEBRUSxDT00iLCJleHAiOjE2MjYyMzI2MDcsImV4dGVybmFsSWQiOiIzYTQ0Yjk3MC05MGQ0LTQxYWMtOWMwNS1iOWQxOWRlN2Q3NmEiLCJtb2JpbGUiOiIxMzIqKioqODYxMyIsIm5hbWUiOiLlvKAqKiIsInVkQWNjb3VudFV1aWQiOiJ0ZXN0MTIzIn0.D6f3dBkvMYj69NqyumnrAB9RreJA0cxbIsycnHgQK3U"
	info, result := Pares(tokenString, []byte("123456"))
	t.Log(result)
	t.Log(info)
	infoBytes, err := json.Marshal(info)
	if err != nil {
		t.Fatal(err.Error())
	}
	record := make(map[string]interface{})
	err = json.Unmarshal(infoBytes, &record)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(record)
}
