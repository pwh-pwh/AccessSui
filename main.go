package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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
		contentContainer.Objects = []fyne.CanvasObject{widget.NewLabel("内容市场界面")}
		contentContainer.Refresh()
	})
	myContentBtn := widget.NewButton("我的内容", func() {
		contentContainer.Objects = []fyne.CanvasObject{widget.NewLabel("我的内容界面")}
		contentContainer.Refresh()
	})
	uploadBtn := widget.NewButton("上传内容", func() {
		contentContainer.Objects = []fyne.CanvasObject{widget.NewLabel("上传内容界面")}
		contentContainer.Refresh()
	})
	settingsBtn := widget.NewButton("设置", func() {
		contentContainer.Objects = []fyne.CanvasObject{widget.NewLabel("设置界面")}
		contentContainer.Refresh()
	})

	sidebar := container.NewVBox(
		marketBtn,
		myContentBtn,
		uploadBtn,
		layout.NewSpacer(), // 将设置按钮推到底部
		settingsBtn,
	)

	// 内容市场界面
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

	// 内容详情界面 (需要先定义，因为 purchaseTokenContent 会引用它)
	var detailContent *fyne.Container // 声明为指针，以便在后面赋值
	detailContent = container.NewVBox(
		widget.NewLabel("内容标题详情"),
		widget.NewLabel("创作者: 创作者X"),
		widget.NewLabel("这是一个内容的详细描述。"),
		widget.NewLabel("价格: 250 Sui"),
		widget.NewButton("购买 AccessToken", func() {
			// 购买逻辑，这里会跳转到购买流程界面
			// contentContainer.Objects = []fyne.CanvasObject{purchaseTokenContent}
			// contentContainer.Refresh()
		}),
		widget.NewButton("返回", func() {
			contentContainer.Objects = []fyne.CanvasObject{marketContent}
			contentContainer.Refresh()
		}),
	)

	// 购买 AccessToken 流程界面
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

	// 更新内容详情界面的购买按钮，使其指向 purchaseTokenContent
	detailContent.Objects[4] = widget.NewButton("购买 AccessToken", func() {
		contentContainer.Objects = []fyne.CanvasObject{purchaseTokenContent}
		contentContainer.Refresh()
	})

	// 我的内容 / AccessToken 管理界面
	myContentContent := container.NewVBox(
		widget.NewLabel("我的 AccessToken 列表"),
		container.NewGridWithColumns(1, // 示例 AccessToken 列表
			container.NewHBox(
				widget.NewLabel("内容: 内容标题1"),
				widget.NewLabel("ID: 0x...123"),
				widget.NewLabel("状态: 有效"),
				widget.NewButton("解锁/查看", func() { /* 解锁/查看逻辑 */ }),
				widget.NewButton("转让", func() { /* 转让逻辑 */ }),
				widget.NewButton("撤销", func() { /* 撤销逻辑 */ }),
			),
			container.NewHBox(
				widget.NewLabel("内容: 内容标题2"),
				widget.NewLabel("ID: 0x...456"),
				widget.NewLabel("状态: 已过期"),
				widget.NewButton("解锁/查看", func() { /* 解锁/查看逻辑 */ }),
				widget.NewButton("转让", func() { /* 转让逻辑 */ }),
				widget.NewButton("撤销", func() { /* 撤销逻辑 */ }),
			),
		),
		widget.NewButton("筛选/排序", func() { /* 筛选/排序逻辑 */ }),
	)

	// 更新侧边导航栏按钮的点击事件
	marketBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{marketContent}
		contentContainer.Refresh()
	}
	myContentBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{myContentContent}
		contentContainer.Refresh()
	}
	uploadBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{widget.NewLabel("上传内容界面")}
		contentContainer.Refresh()
	}
	settingsBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{widget.NewLabel("设置界面")}
		contentContainer.Refresh()
	}

	// 解锁内容界面
	unlockContent := container.NewVBox(
		widget.NewLabel("内容标题: 解锁内容X"),
		widget.NewLabel("解锁状态: 正在请求随机数..."),
		widget.NewLabel("签名提示: 请在您的钱包中确认签名。"),
		container.NewMax( // 内容显示区域
			widget.NewLabel("此处显示解密后的内容 (文本/图片/视频/PDF)"),
		),
		widget.NewButton("关闭", func() {
			contentContainer.Objects = []fyne.CanvasObject{myContentContent} // 返回我的内容
			contentContainer.Refresh()
		}),
	)

	// 更新我的内容界面的“解锁/查看”按钮，使其指向 unlockContent
	// 假设第一个 AccessToken 的解锁按钮
	if len(myContentContent.Objects) > 1 { // 确保有内容列表
		if hbox, ok := myContentContent.Objects[1].(*fyne.Container); ok {
			if len(hbox.Objects) > 0 {
				if innerHBox, ok := hbox.Objects[0].(*fyne.Container); ok {
					if len(innerHBox.Objects) > 3 {
						if unlockBtn, ok := innerHBox.Objects[3].(*widget.Button); ok {
							unlockBtn.OnTapped = func() {
								contentContainer.Objects = []fyne.CanvasObject{unlockContent}
								contentContainer.Refresh()
							}
						}
					}
				}
			}
		}
	}

	// 创作者上传内容界面
	uploadContent := container.NewVBox(
		widget.NewLabel("上传新内容"),
		widget.NewEntry(), // 内容标题输入框
		widget.NewMultiLineEntry(), // 内容描述输入框
		widget.NewButton("选择内容文件", func() { /* 文件选择逻辑 */ }),
		widget.NewButton("上传封面图/缩略图", func() { /* 封面图上传逻辑 */ }),
		container.NewHBox(
			widget.NewLabel("价格:"),
			widget.NewEntry(), // 价格输入框
			widget.NewCheck("订阅制", func(b bool) { /* 订阅制选项 */ }),
			widget.NewLabel("版税比例:"),
			widget.NewEntry(), // 版税比例设置
		),
		widget.NewButton("上传并铸造", func() { /* 上传并铸造逻辑 */ }),
		widget.NewLabel("上传进度: 0%"), // 上传进度条/状态提示
	)

	// 更新侧边导航栏按钮的点击事件
	marketBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{marketContent}
		contentContainer.Refresh()
	}
	myContentBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{myContentContent}
		contentContainer.Refresh()
	}
	uploadBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{uploadContent}
		contentContainer.Refresh()
	}
	settingsBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{widget.NewLabel("设置界面")}
		contentContainer.Refresh()
	}

	// 访问历史与收藏管理界面
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

	// 更新侧边导航栏按钮的点击事件
	marketBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{marketContent}
		contentContainer.Refresh()
	}
	myContentBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{myContentContent}
		contentContainer.Refresh()
	}
	uploadBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{uploadContent}
		contentContainer.Refresh()
	}
	settingsBtn.OnTapped = func() {
		contentContainer.Objects = []fyne.CanvasObject{historyAndFavoritesContent} // 暂时将设置按钮指向这里
		contentContainer.Refresh()
	}

	// 初始显示内容市场界面
	contentContainer.Objects = []fyne.CanvasObject{marketContent}

	// 主布局：侧边栏 + 内容区域
	mainLayout := container.NewHSplit(sidebar, contentContainer)
	mainLayout.SetOffset(0.2) // 侧边栏占据20%宽度

	w.SetContent(mainLayout)
	w.ShowAndRun()
}
