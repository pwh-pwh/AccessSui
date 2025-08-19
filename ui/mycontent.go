package ui

import (
	"image/color" // 导入 color 包
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas" // 导入 canvas 包
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme" // 导入 theme 包
	"fyne.io/fyne/v2/widget"
)

// MyContentContent creates the UI for managing user's AccessTokens.
func MyContentContent(contentContainer *fyne.Container) *fyne.Container {
	myContentContent := container.NewVBox(
		widget.NewLabel("我的 AccessToken 列表"),
		container.NewGridWithColumns(1, // 示例 AccessToken 列表
			widget.NewCard("内容: 内容标题1", "ID: 0x...123",
				container.NewVBox(
					widget.NewLabel("状态: 有效"),
					container.NewHBox(
						widget.NewButtonWithIcon("解锁/查看", theme.ContentCopyIcon(), func() { // 使用图标
							contentContainer.Objects = []fyne.CanvasObject{UnlockContent(contentContainer)}
							contentContainer.Refresh()
						}),
						widget.NewButtonWithIcon("转让", theme.MailSendIcon(), func() { /* 转让逻辑 */ }), // 使用图标
						widget.NewButtonWithIcon("撤销", theme.DeleteIcon(), func() { /* 撤销逻辑 */ }), // 使用图标
					),
				),
			),
			widget.NewCard("内容: 内容标题2", "ID: 0x...456",
				container.NewVBox(
					widget.NewLabel("状态: 已过期"),
					container.NewHBox(
						widget.NewButtonWithIcon("解锁/查看", theme.ContentCopyIcon(), func() { // 使用图标
							contentContainer.Objects = []fyne.CanvasObject{UnlockContent(contentContainer)}
							contentContainer.Refresh()
						}),
						widget.NewButtonWithIcon("转让", theme.MailSendIcon(), func() { /* 转让逻辑 */ }), // 使用图标
						widget.NewButtonWithIcon("撤销", theme.DeleteIcon(), func() { /* 撤销逻辑 */ }), // 使用图标
					),
				),
			),
		),
		widget.NewButton("筛选/排序", func() { /* 筛选/排序逻辑 */ }),
	)
	return myContentContent
}

// UnlockContent creates the UI for unlocking and displaying content.
func UnlockContent(contentContainer *fyne.Container) *fyne.Container {
	unlockContent := container.NewVBox(
		widget.NewLabel("内容标题: 解锁内容X"),
		widget.NewLabel("解锁状态: 正在请求随机数..."),
		widget.NewLabel("签名提示: 请在您的钱包中确认签名。"),
		container.NewMax( // 内容显示区域
			container.NewVBox(
				widget.NewLabel("解密内容显示区域:"),
				widget.NewLabel("这是一个文本内容的示例。"),
				canvas.NewRectangle(color.RGBA{R: 100, G: 100, B: 100, A: 255}), // 图片占位符
				widget.NewLabel("此处可显示视频或PDF播放器"),
			),
		),
		widget.NewButton("关闭", func() {
			contentContainer.Objects = []fyne.CanvasObject{MyContentContent(contentContainer)} // 返回我的内容
			contentContainer.Refresh()
		}),
	)
	return unlockContent
}