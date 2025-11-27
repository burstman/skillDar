package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateMainScreen builds the main app screen
func CreateMainScreen(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("Main Screen")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	profileBtn := widget.NewButton("Go to Profile", func() {
		if state.GetUserRole() == "worker" {
			state.ShowScreen("edit_profile_worker")
			return
		}
		state.ShowScreen("edit_profile_client")
	})

	logoutBtn := widget.NewButton("Logout", func() {
		state.ShowScreen("login")
	})

	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		profileBtn,
		logoutBtn,
		layout.NewSpacer(),
	)

	return container.NewPadded(container.NewCenter(content))
}
