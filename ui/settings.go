package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// SettingsContent creates the UI for application settings.
func SettingsContent() *fyne.Container {
	settingsContent := container.NewVBox(
		widget.NewLabel("设置"),
		container.NewHBox(
			widget.NewLabel("语言:"),
			widget.NewSelect([]string{"中文", "English"}, func(s string) { /* 语言选择逻辑 */ }),
		),
		widget.NewButton("清除本地缓存", func() { /* 清除缓存逻辑 */ }),
		widget.NewButton("连接/断开钱包", func() { /* 钱包管理逻辑 */ }),
	)
	return settingsContent
}