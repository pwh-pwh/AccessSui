package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// MyContentContent creates the UI for managing user's AccessTokens.
func MyContentContent(contentContainer *fyne.Container) *fyne.Container {
	myContentContent := container.NewVBox(
		widget.NewLabel("我的 AccessToken 列表"),
		container.NewGridWithColumns(1, // 示例 AccessToken 列表
			container.NewHBox(
				widget.NewLabel("内容: 内容标题1"),
				widget.NewLabel("ID: 0x...123"),
				widget.NewLabel("状态: 有效"),
				widget.NewButton("解锁/查看", func() {
					contentContainer.Objects = []fyne.CanvasObject{UnlockContent(contentContainer)}
					contentContainer.Refresh()
				}),
				widget.NewButton("转让", func() { /* 转让逻辑 */ }),
				widget.NewButton("撤销", func() { /* 撤销逻辑 */ }),
			),
			container.NewHBox(
				widget.NewLabel("内容: 内容标题2"),
				widget.NewLabel("ID: 0x...456"),
				widget.NewLabel("状态: 已过期"),
				widget.NewButton("解锁/查看", func() {
					contentContainer.Objects = []fyne.CanvasObject{UnlockContent(contentContainer)}
					contentContainer.Refresh()
				}),
				widget.NewButton("转让", func() { /* 转让逻辑 */ }),
				widget.NewButton("撤销", func() { /* 撤销逻辑 */ }),
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
			widget.NewLabel("此处显示解密后的内容 (文本/图片/视频/PDF)"),
		),
		widget.NewButton("关闭", func() {
			contentContainer.Objects = []fyne.CanvasObject{MyContentContent(contentContainer)} // 返回我的内容
			contentContainer.Refresh()
		}),
	)
	return unlockContent
}