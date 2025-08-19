package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/mbaklor/fyne-catppuccin"
	"github.com/pwh-pwh/AccessSui/ui" // 导入 ui 包
)

func main() {
	a := app.New()
	ctp := catppuccin.New()
	ctp.SetFlavor(catppuccin.Frappe)
	a.Settings().SetTheme(ctp)
	w := a.NewWindow("去中心化知识加密分享平台")
	w.Resize(fyne.NewSize(800, 600))

	// 创建一个可切换内容的容器
	contentContainer := container.NewStack()

	// 侧边导航栏
	marketBtn := widget.NewButton("内容市场", func() {
		contentContainer.Objects = []fyne.CanvasObject{ui.MarketContent(contentContainer)}
		contentContainer.Refresh()
	})
	myContentBtn := widget.NewButton("我的内容", func() {
		contentContainer.Objects = []fyne.CanvasObject{ui.MyContentContent(contentContainer)}
		contentContainer.Refresh()
	})
	uploadBtn := widget.NewButton("上传内容", func() {
		contentContainer.Objects = []fyne.CanvasObject{ui.UploadContent(contentContainer)}
		contentContainer.Refresh()
	})
	historyBtn := widget.NewButton("历史与收藏", func() {
		contentContainer.Objects = []fyne.CanvasObject{ui.HistoryAndFavoritesContent()}
		contentContainer.Refresh()
	})
	settingsBtn := widget.NewButton("设置", func() {
		contentContainer.Objects = []fyne.CanvasObject{ui.SettingsContent()}
		contentContainer.Refresh()
	})

	sidebar := container.NewVBox(
		marketBtn,
		myContentBtn,
		uploadBtn,
		historyBtn,
		layout.NewSpacer(), // 将设置按钮推到底部
		settingsBtn,
	)

	// 初始显示内容市场界面
	contentContainer.Objects = []fyne.CanvasObject{ui.MarketContent(contentContainer)}

	// 主布局：侧边栏 + 内容区域
	mainLayout := container.NewHSplit(sidebar, contentContainer)
	mainLayout.SetOffset(0.2) // 侧边栏占据20%宽度

	w.SetContent(mainLayout)
	w.ShowAndRun()
}
