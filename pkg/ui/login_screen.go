package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateLoginScreen builds the login/welcome screen
func CreateLoginScreen(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("Welcome to SkillKonnect")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := widget.NewLabel("Connect skills, build networks")
	subtitle.Alignment = fyne.TextAlignCenter

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email or Username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	loginBtn := widget.NewButton("Login", func() {
		email := emailEntry.Text
		password := passwordEntry.Text
		if email == "" || password == "" {
			fmt.Println("Please fill in all fields")
		} else {
			fmt.Println("Logged in as:", email)
			// Navigate to main screen
			state.ShowScreen("main")
		}
	})

	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		subtitle,
		emailEntry,
		passwordEntry,
		loginBtn,
		layout.NewSpacer(),
	)

	return container.NewPadded(container.NewCenter(content))
}
