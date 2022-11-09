# go 生成拼图验证码

## 安装
```
go get github.com/wangzmgit/jigsaw
```

## 基本用法
首先需要指定背景图和遮罩图位置，背景图需要使用尺寸统一的png图片，遮罩图片需使用正方形的半透明png图片。
```go
j := jigsaw.New()
j.SetBgDir("./images/bg/")
j.SetMaskPath("./images/mask.png")
```

在需要生成图片的地方调用
```go
img, bg, x, y, err := jigsaw.Create()
// img为生成小图的base64字符串
// bg为生成背景图的base64字符串
// x为小图左边在背景图的位置
// y为小图上边在背景图的位置
// err为错误信息
```

## 修改图片尺寸
默认背景图大小为310*160，遮罩大小为50\*50，可通过以下代码修改背景图和遮罩大小
```go
j.SetBgSize(背景图宽度,背景图高度)
j.SetMaskSize(遮罩尺寸)
```


## 修改最大背景图数量
默认情况下最多读取背景图路径下的10张图片，可通过以下代码修改最大背景图数量
```go
j.SetMaxBgNums(最大背景图数量)
```

## 生成图片示例
![text][img_0]
![text][img_1]


[img_0]:data:image/jpg;base64,/9j/2wCEAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxNDQ0Hyc5PTgyPC4zNDIBCQkJDAsMGA0NGDIhHCEyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMv/AABEIADIAMgMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/APSKKKK7DzwooooAKKKKACiiigDpBZWy9IE/EZpTaWxGPIj/AO+RU1Fct2UVH021f/lntP8Ask1Wk0ZTzFKR7MM1qUU1NoRzs1jcQZLJlf7y8iq1dXVK602KfLIBG/qOhrSNTuFjA3qP4hSb1/vCoJND1UyufKByT0cU3+wtV/55f+Pj/GuV4yd/gZ3LCU7fxEdtRRRWhxBRRRQAUUUUAFFFFABVC9kdCdrsPoav1nX/AFNa0fiEyvps8sl4FeV2G08Fia2awtK/4/R/umt2rxKtPQEFFFFc4wooooA//9k=

[img_1]:data:image/jpg;base64,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

