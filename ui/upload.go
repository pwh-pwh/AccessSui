package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// UploadContent creates the UI for content uploading by creators.
func UploadContent(contentContainer *fyne.Container) *fyne.Container {
	uploadContent := container.NewVBox(
		widget.NewLabel("上传新内容"),
		widget.NewEntry(),          // 内容标题输入框
		widget.NewMultiLineEntry(), // 内容描述输入框
		widget.NewButton("选择内容文件", func() { /* 文件选择逻辑 */ }),
		widget.NewButton("上传封面图/缩略图", func() { /* 封面图上传逻辑 */ }),
		container.NewHBox(
			widget.NewLabel("价格:"),
			widget.NewEntry(), // 价格输入框
			widget.NewCheck("订阅制", func(b bool) { /* 订阅制选项 */ }),
			widget.NewLabel("版税比例:"),
			widget.NewEntry(), // 版税比例设置
		),
		widget.NewButton("上传并铸造", func() { /* 上传并铸造逻辑 */ }),
		widget.NewLabel("上传进度: 0%"), // 上传进度条/状态提示
	)
	return uploadContent
}