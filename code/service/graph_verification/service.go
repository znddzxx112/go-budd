package graph_verification

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

type GraphVerification interface {
	// 仅仅生成图形验证码,通过流输出到前端
	ProduceGraphCode(w http.ResponseWriter) (string, error)
	// 生成图形验证码,并保存至数据库中,通过流输出到前端
	ProduceGraphCodeAndSave(w http.ResponseWriter) (imageToken string, hash error)
	// 验证图形验证码
	CheckCode(imageToken string, imageCode string) error
}

type graphVerification struct {
	db *gorm.DB
}

// 图片框长宽大小
const (
	dx1 = 150
	dy2 = 50
)

func NewGraphVerification(db *gorm.DB) GraphVerification {
	gv := new(graphVerification)
	gv.db = db
	return gv
}

func (gv *graphVerification) CheckCode(imageToken string, imageCode string) error {
	g := NewGraphValidWithToken(gv.db, imageToken)
	if g == nil {
		return fmt.Errorf("%s", "未找到图形验证码")
	}
	//if g.Code != imageCode {
	//	return fmt.Errorf("%s", "图片验证码输入不正确")
	//}
	// 图形验证码
	g.Delete(gv.db, map[string]interface{}{"token": imageToken})

	if strings.ToLower(g.Code) != strings.ToLower(imageCode) {
		return fmt.Errorf("%s", "图片验证码输入不正确")
	}

	return nil
}

// 生成图形验证码,通过流输出到前端
func (gv *graphVerification) ProduceGraphCodeAndSave(w http.ResponseWriter) (token string, error error) {
	token = GenerateTicket("graph", 10)
	w.Header().Add("imageToken", token)
	code, err := gv.ProduceGraphCode(w)
	if err != nil {
		return "", err
	}

	graphValid := GraphValid{
		CreatedAt: time.Now(),
		Code:      code,
		Token:     token,
	}
	err = graphValid.Create(gv.db)
	if err != nil {
		return "", err
	}
	return token, nil
}

// 生成图形验证码,通过流输出到前端
func (gv *graphVerification) ProduceGraphCode(w http.ResponseWriter) (string, error) {

	//所有的相对路径都是相对于该项目开始
	err := ReadFonts("./resources/fonts", ".ttf")
	if err != nil {
		log.Fatal(err)
	}

	captchaImage, err := NewCaptchaImage(dx1, dy2, RandLightColor())

	captchaImage.DrawNoise(CaptchaComplexLower)

	//captchaImage.DrawTextNoise(CaptchaComplexLower)

	//生成验证码的位数
	code := RandText(4)
	captchaImage.DrawText(code)

	captchaImage.DrawBorder(ColorToRGB(0x17A7A7A))

	//captchaImage.DrawSineLine()

	if err != nil {
		return "", err
	}

	captchaImage.SaveImage(w, ImageFormatJpeg)
	return code, nil
}
