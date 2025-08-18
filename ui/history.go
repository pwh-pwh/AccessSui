package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// HistoryAndFavoritesContent creates the UI for browsing history and favorites.
func HistoryAndFavoritesContent() *fyne.Container {
	historyAndFavoritesContent := container.NewVBox(
		container.NewAppTabs(
			container.NewTabItem("访问历史", container.NewGridWithColumns(1,
				container.NewHBox(
					widget.NewLabel("内容: 历史内容1"),
					widget.NewLabel("创作者: 创作者Y"),
					widget.NewButton("查看内容", func() { /* 查看内容逻辑 */ }),
					widget.NewButton("从历史中移除", func() { /* 移除逻辑 */ }),
				),
			)),
			container.NewTabItem("我的收藏", container.NewGridWithColumns(1,
				container.NewHBox(
					widget.NewLabel("内容: 收藏内容1"),
					widget.NewLabel("创作者: 创作者Z"),
					widget.NewButton("查看内容", func() { /* 查看内容逻辑 */ }),
					widget.NewButton("取消收藏", func() { /* 取消收藏逻辑 */ }),
				),
			)),
		),
	)
	return historyAndFavoritesContent
}