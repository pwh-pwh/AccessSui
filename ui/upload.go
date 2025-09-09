package ui

import (
	"context"
	"errors" // 导入 errors 包
	"fmt"
	"io"
	"strconv" // 导入 strconv 包

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout" // 导入 layout 包
	"fyne.io/fyne/v2/widget"
	"github.com/pwh-pwh/AccessSui/client"
	"github.com/pwh-pwh/AccessSui/store"
)

// UploadContent creates the UI for content uploading by creators.
func UploadContent(w fyne.Window, contentContainer *fyne.Container) *fyne.Container {
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
	descEntry := widget.NewMultiLineEntry()
	form := container.New(layout.NewFormLayout(),
		widget.NewLabel("内容标题"), titleEntry,
		widget.NewLabel("内容描述"), descEntry,
		widget.NewLabel("内容文件"), widget.NewButton("选择内容文件", func() {
			dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil || reader == nil {
					return
				}
				defer reader.Close()
				data, err := io.ReadAll(reader)
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				descEntry.SetText(string(data))
			}, w)
		}),
		widget.NewLabel("价格"), priceEntry,
	)
	msgLabel := widget.NewLabel("上传进度: 0%")
	msgLabel.Selectable = true
	bottomContent := container.NewVBox(
		widget.NewButton("上传并铸造", func() {
			/* 上传并铸造逻辑 */
			blodId, err := store.StoreData([]byte(descEntry.Text))
			if err != nil {
				msgLabel.SetText(err.Error())
				return
			}

			suiClient, err := client.GetSuiClient()
			if err != nil {
				msgLabel.SetText(err.Error())
				return
			}
			resp, err := suiClient.PublishContent(context.Background(),
				suiClient.GetAddress(), blodId, titleEntry.Text, blodId, priceEntry.Text, "", "100000000")
			if err != nil {
				msgLabel.SetText(err.Error())
				return
			}
			msgLabel.SetText(fmt.Sprintf("交易tx: %s\nblodId: %s", resp.Digest, blodId))
		}),
		msgLabel,
	)

	return container.NewVBox(form, bottomContent)
}
