package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateWelcomeScreen creates the welcome screen UI
// onGetStarted is a callback function that runs when "Get Started" button is clicked
func CreateWelcomeScreen(state AppState) fyne.CanvasObject {
	// UI elements
	title := widget.NewLabel("Welcome to SkillDar")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := widget.NewLabel("Your Home, Our Expertise")
	subtitle.Alignment = fyne.TextAlignCenter

	getStartedBtn := widget.NewButton("Get Started", func() {
		state.ShowScreen("login")
	})

	// Layout - vertically stacked with spacers for centering
	content := container.NewVBox(
		title,
		subtitle,
		layout.NewSpacer(),
		getStartedBtn,
		layout.NewSpacer(),
		layout.NewSpacer(),
	)

	// Return padded content
	return container.NewPadded(content)
}
