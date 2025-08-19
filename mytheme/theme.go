package mytheme

import (
	"fyne.io/fyne/v2"
	catppuccin "github.com/mbaklor/fyne-catppuccin"
)

type MyTheme struct {
	catppuccin.Theme
}

func (ctp MyTheme) Font(style fyne.TextStyle) fyne.Resource {
	return resourceOpposansTtf
}
