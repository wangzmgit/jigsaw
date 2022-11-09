package jigsaw

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"

	"github.com/disintegration/imaging"
)

var jigsaw *Jigsaw

type Jigsaw struct {
	maxBgNums  int // 最大背景图数量
	realBgNums int // 实际背景图数量

	maskSize int //遮罩图尺寸
	bgWidth  int // 背景图宽度
	bgHeight int //背景图高度

	maskPath string   // 遮罩文件
	bgList   []string // 背景图文件数组
}

func New() *Jigsaw {
	j := new(Jigsaw)
	j.maxBgNums = 10 //最大数量
	j.bgWidth = 310  //背景图宽度
	j.bgHeight = 160 //背景图高度
	j.maskSize = 50

	jigsaw = j
	return j
}

// 设置背景图所在目录
func (j *Jigsaw) SetBgDir(dir string) error {
	bgList := make([]string, j.maxBgNums+1) //背景图列表
	if files, err := ioutil.ReadDir(dir); err == nil {
		nums := 0 //背景图数量
		for _, f := range files {
			if nums >= j.maxBgNums {
				break
			}
			//仅记录png文件
			if !f.IsDir() && path.Ext(f.Name()) == ".png" {
				nums++
				bgList[nums] = dir + f.Name()
			}
		}

		if nums == 0 {
			return errors.New("no valid image found")
		}

		j.realBgNums = nums
		j.bgList = bgList[1:]
		return nil
	} else {
		return err
	}
}

// 设置遮罩图路径
func (j *Jigsaw) SetMaskPath(path string) error {
	if exist, _ := pathExists(path); exist {
		j.maskPath = path
		return nil
	}

	return errors.New("mask file does not exist")
}

//设置最大背景图数量
func (j *Jigsaw) SetMaxBgNums(n int) {
	if n > 0 {
		j.maxBgNums = n
	}
}

//设置背景图尺寸
func (j *Jigsaw) SetBgSize(width, height int) {
	if width > 0 && height > 0 {
		j.bgWidth = width
		j.bgHeight = height
	}
}

//设置遮罩尺寸
func (j *Jigsaw) SetMaskSize(size int) {
	if size > 0 {
		j.maskSize = size
	}
}

//生成图片和坐标
func Create() (string, string, int, int, error) {
	return jigsaw.Create()
}

func (j *Jigsaw) Create() (string, string, int, int, error) {
	if j.realBgNums == 0 {
		return "", "", 0, 0, errors.New("no valid image found")
	}
	nums := getRandInt(0, j.realBgNums)
	f, fileErr := os.Open(j.bgList[nums])
	if fileErr != nil {
		return "", "", 0, 0, fileErr
	}
	defer f.Close()

	//生成拼图位置坐标
	//生成拼图左边x坐标，取值范围拼图左右边界到两侧最小距离20px
	imageRandX := getRandInt(20, j.bgWidth-j.maskSize-20)
	//生成拼图上边y坐标，取值范围拼图上下边界到顶部或底部最小距离20px
	imageRandY := getRandInt(20, j.bgHeight-j.maskSize-20)

	//设置截取的最大坐标值和最小坐标值
	minPotion := image.Point{
		X: imageRandX,
		Y: imageRandY,
	}
	maxPotion := image.Point{
		X: imageRandX + j.maskSize,
		Y: imageRandY + j.maskSize,
	}
	//处理图像
	m, decodeErr := png.Decode(f) //转化为image对象
	if decodeErr != nil {
		return "", "", 0, 0, decodeErr
	}

	img := getSubImg(maxPotion, minPotion, m)
	bg, bgErr := getCodeImg(minPotion, m, j.maskPath)
	if bgErr != nil {
		return "", "", 0, 0, bgErr
	}

	return img, bg, imageRandX, imageRandY, nil
}

func getSubImg(maxPotion, minPotion image.Point, m image.Image) string {
	subimg := image.Rectangle{
		Max: maxPotion,
		Min: minPotion,
	}
	data := imaging.Crop(m, subimg)
	//转成base64
	emptyBuff := bytes.NewBuffer(nil)                           //开辟一个新的空buff
	jpeg.Encode(emptyBuff, data, nil)                           //img写入到buff
	return base64.StdEncoding.EncodeToString(emptyBuff.Bytes()) //buff转成base64
}

//获取背景图
func getCodeImg(minPotion image.Point, m image.Image, maskPath string) (string, error) {
	maskFile, fileErr := os.Open(maskPath)
	if fileErr != nil {
		return "", fileErr
	}

	maskimg, decodeErr := png.Decode(maskFile)
	if decodeErr != nil {
		return "", decodeErr
	}

	data := imaging.Overlay(m, maskimg, minPotion, 1.0)
	//转成base64
	emptyBuff := bytes.NewBuffer(nil)
	jpeg.Encode(emptyBuff, data, nil)
	return base64.StdEncoding.EncodeToString(emptyBuff.Bytes()), nil
}
