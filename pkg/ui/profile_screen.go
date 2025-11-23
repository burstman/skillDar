package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateProfileScreen builds the profile screen
func CreateProfileScreen(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("Profile")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	userInfo := widget.NewLabel("User: test@example.com")

	backBtn := widget.NewButton("Back to Main", func() {
		state.ShowScreen("main")
	})

	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		userInfo,
		backBtn,
		layout.NewSpacer(),
	)

	return container.NewPadded(container.NewCenter(content))
}
