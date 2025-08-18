package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/pwh-pwh/AccessSui/ui" // 导入 ui 包
)

// package id:0xbda58f110ce755a63c007d68cc53f7ac68c780dc8fb1fb16ad52d797143b4799
//tx: 6S5b62crgfbikAEZKLwphkQqtf5wuJwQqkkRW4ywnxug
//obj:0x7dc80959fdd7b4c68ba0caa2f0f1182fb297817742caf170dc1f787d58317f3d

const (
	PackageId  = "0xbda58f110ce755a63c007d68cc53f7ac68c780dc8fb1fb16ad52d797143b4799"
	CountObjId = "0x7dc80959fdd7b4c68ba0caa2f0f1182fb297817742caf170dc1f787d58317f3d"
)

/*func init() {
	err := godotenv.Load(".env") // 默认读取 .env 文件
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}*/

func main() {
	a := app.New()
	w := a.NewWindow("去中心化知识加密分享平台")
	w.Resize(fyne.NewSize(800, 600))

	// 创建一个可切换内容的容器
	contentContainer := container.NewMax()

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
