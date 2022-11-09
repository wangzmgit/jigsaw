package jigsaw

import (
	"fmt"
	"testing"
)

func TestCaptcha(t *testing.T) {
	j := New()
	j.SetBgDir("./images/bg/")
	j.SetMaskPath("./images/mask.png")
	img, bg, x, y, err := j.Create()
	fmt.Printf("小图%s\n背景%s\n%d %d\n", img, bg, x, y)
	if err != nil {
		fmt.Println(err)
	}
}
