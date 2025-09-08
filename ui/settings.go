package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	myconfig "github.com/pwh-pwh/AccessSui/config" // Corrected full import path
)

// SettingsContent creates the UI for application settings.
func SettingsContent(w fyne.Window) *fyne.Container {
	// Wallet status label
	walletStatusLabel := widget.NewLabel("钱包状态: 未连接")
	walletStatusLabel.TextStyle.Bold = true

	walletStatusColor := canvas.NewRectangle(color.RGBA{A: 0x00}) // Transparent by default
	walletStatusColor.SetMinSize(fyne.NewSize(10, 10))

	// Function to update wallet status
	updateWalletStatus := func() {
		mnemonic, _ := myconfig.LoadMnemonic()
		if mnemonic != "" {
			walletStatusLabel.SetText("钱包状态: 已连接")
			walletStatusColor.FillColor = color.RGBA{G: 0xff, A: 0xff} // Green
		} else {
			walletStatusLabel.SetText("钱包状态: 未连接")
			walletStatusColor.FillColor = color.RGBA{R: 0xff, A: 0xff} // Red
		}
		canvas.Refresh(walletStatusColor) // Refresh the canvas object to show color change
	}

	// Initial status update
	updateWalletStatus()

	// Connect Wallet button
	connectWalletButton := widget.NewButton("连接钱包", func() {
		mnemonicEntry := widget.NewPasswordEntry() // Use password entry for mnemonic for security
		mnemonicEntry.SetPlaceHolder("输入您的助记词...")

		items := []*widget.FormItem{
			widget.NewFormItem("助记词", mnemonicEntry),
		}

		dialog.ShowForm("连接钱包", "连接", "取消", items, func(confirm bool) {
			if confirm {
				mnemonic := mnemonicEntry.Text
				if mnemonic == "" {
					dialog.ShowError(fmt.Errorf("助记词不能为空，请输入助记词"), w) // Use fmt.Errorf
					return
				}
				err := myconfig.SaveMnemonic(mnemonic)
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				dialog.ShowInformation("成功", "钱包连接成功！", w)
				updateWalletStatus() // Update status after connection
			}
		}, w)
	})

	// Disconnect Wallet button
	disconnectWalletButton := widget.NewButton("断开钱包", func() {
		dialog.ShowConfirm("断开钱包", "您确定要断开钱包并清除助记词吗？", func(confirm bool) {
			if confirm {
				err := myconfig.ClearMnemonic()
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				dialog.ShowInformation("成功", "钱包已断开！", w)
				updateWalletStatus() // Update status after disconnection
			}
		}, w)
	})

	settingsContent := container.NewVBox(
		widget.NewLabel("设置"),
		container.NewHBox(
			walletStatusLabel,
			walletStatusColor, // Add color indicator
		),
		container.NewHBox(
			widget.NewLabel("语言:"),
			widget.NewSelect([]string{"中文", "English"}, func(s string) { /* 语言选择逻辑 */ }),
		),
		widget.NewButton("清除本地缓存", func() { /* 清除缓存逻辑 */ }),
		layout.NewSpacer(),
		connectWalletButton,
		disconnectWalletButton,
	)
	return settingsContent
}
