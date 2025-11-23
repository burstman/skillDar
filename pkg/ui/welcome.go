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
	title := widget.NewLabel("Welcome to SkillKonnect")
	title.Alignment = fyne.TextAlignLeading
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := widget.NewLabel("Connect skills, build networks")
	subtitle.Alignment = fyne.TextAlignCenter

	getStartedBtn := widget.NewButton("Get Started", func() {
		state.ShowScreen("login")
	})

	// Layout - vertically stacked with spacers for centering
	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		subtitle,
		layout.NewSpacer(),
		getStartedBtn,
		layout.NewSpacer(),
	)

	// Return padded and centered content
	return container.NewPadded(container.NewCenter(content))
}
