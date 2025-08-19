package ui

import (
	"image/color" // 导入 color 包
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas" // 导入 canvas 包
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// MarketContent creates the UI for the content market.
func MarketContent(contentContainer *fyne.Container) *fyne.Container {
	searchEntry := widget.NewEntry() // 搜索栏
	
	contentGrid := container.NewGridWithColumns(2, // 示例内容卡片，改为2列
		widget.NewCard("内容标题1", "创作者A", container.NewVBox(
			canvas.NewRectangle(color.Gray{}), // 占位符图片
			widget.NewLabel("价格: 100 Sui"),
		)),
		widget.NewCard("内容标题2", "创作者B", container.NewVBox(
			canvas.NewRectangle(color.Gray{}), // 占位符图片
			widget.NewLabel("价格: 150 Sui"),
		)),
		widget.NewCard("内容标题3", "创作者C", container.NewVBox(
			canvas.NewRectangle(color.Gray{}), // 占位符图片
			widget.NewLabel("价格: 200 Sui"),
		)),
	)

	loadingIndicator := widget.NewProgressBarInfinite()
	loadingIndicator.Hide() // 默认隐藏

	loadMoreButton := widget.NewButton("加载更多", func() {
		loadingIndicator.Show() // 显示加载指示器
		// 模拟加载
		go func() {
			// 实际加载逻辑
			// time.Sleep(2 * time.Second) // 模拟网络延迟
			loadingIndicator.Hide() // 隐藏加载指示器
		}()
	})

	marketContent := container.NewVBox(
		searchEntry,
		contentGrid,
		loadMoreButton,
		loadingIndicator,
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
	transactionStatusLabel := widget.NewLabel("交易状态提示: 等待用户确认...")

	purchaseTokenContent := container.NewVBox(
		widget.NewLabel("购买确认信息: 购买内容X, 价格: 250 Sui"),
		widget.NewLabel("钱包连接状态: 已连接 (0x...abcd)"),
		widget.NewLabel("交易详情预览: 扣除 250 Sui, Gas 费预估: 0.001 Sui"),
		container.NewHBox(
			widget.NewButton("确认购买", func() {
				transactionStatusLabel.SetText("交易状态提示: 正在处理交易...")
				// 模拟交易处理
				go func() {
					// time.Sleep(2 * time.Second) // 模拟网络延迟
					transactionStatusLabel.SetText("交易状态提示: 交易成功！")
					// 触发 Sui 钱包签名交易
				}()
			}),
			widget.NewButton("取消", func() {
				contentContainer.Objects = []fyne.CanvasObject{detailContent} // 返回内容详情
				contentContainer.Refresh()
			}),
		),
		transactionStatusLabel, // 使用动态标签
	)
	return purchaseTokenContent
}