package ui

import (
	"errors" // 导入 errors 包
	"strconv" // 导入 strconv 包
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout" // 导入 layout 包
	"fyne.io/fyne/v2/widget"
)

// UploadContent creates the UI for content uploading by creators.
func UploadContent(contentContainer *fyne.Container) *fyne.Container {
	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("请输入内容标题")
	titleEntry.OnChanged = func(s string) {
		if len(s) == 0 {
			titleEntry.SetValidationError(errors.New("标题不能为空"))
		} else {
			titleEntry.SetValidationError(nil)
		}
	}

	priceEntry := widget.NewEntry()
	priceEntry.SetPlaceHolder("请输入价格 (Sui)")
	priceEntry.OnChanged = func(s string) {
		if _, err := strconv.ParseFloat(s, 64); err != nil && len(s) > 0 {
			priceEntry.SetValidationError(errors.New("价格必须是有效数字"))
		} else {
			priceEntry.SetValidationError(nil)
		}
	}

	form := container.New(layout.NewFormLayout(),
		widget.NewLabel("内容标题"), titleEntry,
		widget.NewLabel("内容描述"), widget.NewMultiLineEntry(),
		widget.NewLabel("内容文件"), widget.NewButton("选择内容文件", func() { /* 文件选择逻辑 */ }),
		widget.NewLabel("封面图"), widget.NewButton("上传封面图/缩略图", func() { /* 封面图上传逻辑 */ }),
		widget.NewLabel("价格"), priceEntry,
		widget.NewLabel("订阅制"), widget.NewCheck("", func(b bool) { /* 订阅制选项 */ }),
		widget.NewLabel("版税比例"), widget.NewEntry(),
	)

	bottomContent := container.NewVBox(
		widget.NewButton("上传并铸造", func() { /* 上传并铸造逻辑 */ }),
		widget.NewLabel("上传进度: 0%"),
	)

	return container.NewVBox(form, bottomContent)
}