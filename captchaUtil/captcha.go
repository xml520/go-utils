package captchaUtil

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type Captcha struct {
	captcha *base64Captcha.Captcha
	store   base64Captcha.Store
}

// NewCaptcha 生成验证码实例 ，参数可空
func NewCaptcha(d base64Captcha.Driver, s base64Captcha.Store) *Captcha {
	if d == nil {
		d = (&base64Captcha.DriverString{
			Height:          40,
			Width:           110,
			NoiseCount:      0,
			ShowLineOptions: 2 | 4,
			Length:          4,
			Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
			BgColor: &color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 0,
			},
			Fonts: []string{"wqy-microhei.ttc"}}).ConvertFonts()
	}
	if s == nil {
		s = base64Captcha.DefaultMemStore
	}

	c := &Captcha{
		store:   s,
		captcha: base64Captcha.NewCaptcha(d, s),
	}
	return c
}

// Generate 生成新的验证码
func (c *Captcha) Generate() (id, base64, v string, err error) {
	return c.captcha.Generate()
}

// Verify 验证
func (c *Captcha) Verify(id, value string, clear bool) bool {
	return base64Captcha.DefaultMemStore.Verify(id, value, clear)
}
