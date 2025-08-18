package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// MarketContent creates the UI for the content market.
func MarketContent(contentContainer *fyne.Container) *fyne.Container {
	marketContent := container.NewVBox(
		widget.NewEntry(), // 搜索栏
		container.NewGridWithColumns(3, // 示例内容卡片
			widget.NewCard("内容标题1", "创作者A", widget.NewLabel("价格: 100 Sui")),
			widget.NewCard("内容标题2", "创作者B", widget.NewLabel("价格: 150 Sui")),
			widget.NewCard("内容标题3", "创作者C", widget.NewLabel("价格: 200 Sui")),
		),
		widget.NewButton("加载更多", func() {
			// 加载更多内容的逻辑
		}),
	)
	return marketContent
}

// DetailContent creates the UI for content details.
func DetailContent(contentContainer *fyne.Container) *fyne.Container {
	var detailContent *fyne.Container // 声明为指针，以便在后面赋值
	
	purchaseTokenContent := PurchaseTokenContent(contentContainer, detailContent) // Pass detailContent here

	detailContent = container.NewVBox(
		widget.NewLabel("内容标题详情"),
		widget.NewLabel("创作者: 创作者X"),
		widget.NewLabel("这是一个内容的详细描述。"),
		widget.NewLabel("价格: 250 Sui"),
		widget.NewButton("购买 AccessToken", func() {
			contentContainer.Objects = []fyne.CanvasObject{purchaseTokenContent}
			contentContainer.Refresh()
		}),
		widget.NewButton("返回", func() {
			contentContainer.Objects = []fyne.CanvasObject{MarketContent(contentContainer)} // Assuming MarketContent returns a new instance
			contentContainer.Refresh()
		}),
	)
	return detailContent
}

// PurchaseTokenContent creates the UI for the AccessToken purchase process.
func PurchaseTokenContent(contentContainer *fyne.Container, detailContent *fyne.Container) *fyne.Container {
	purchaseTokenContent := container.NewVBox(
		widget.NewLabel("购买确认信息: 购买内容X, 价格: 250 Sui"),
		widget.NewLabel("钱包连接状态: 已连接 (0x...abcd)"),
		widget.NewLabel("交易详情预览: 扣除 250 Sui, Gas 费预估: 0.001 Sui"),
		container.NewHBox(
			widget.NewButton("确认购买", func() {
				// 触发 Sui 钱包签名交易
			}),
			widget.NewButton("取消", func() {
				contentContainer.Objects = []fyne.CanvasObject{detailContent} // 返回内容详情
				contentContainer.Refresh()
			}),
		),
		widget.NewLabel("交易状态提示: 等待用户确认..."),
	)
	return purchaseTokenContent
}