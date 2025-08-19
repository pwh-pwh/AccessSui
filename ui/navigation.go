package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// CreateSidebar creates a sidebar navigation using AppTabs
func CreateSidebar(window fyne.Window, contentContainer *fyne.Container) *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItem("市场", MarketContent(contentContainer)),
		container.NewTabItem("我的内容", MyContentContent(contentContainer)),
		container.NewTabItem("上传", UploadContent(contentContainer)),
		container.NewTabItem("历史与收藏", HistoryAndFavoritesContent()),
		container.NewTabItem("设置", SettingsContent()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	return tabs
}
