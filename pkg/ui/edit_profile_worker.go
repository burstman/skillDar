package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateEditProfileWorkerScreen builds the worker profile edit screen
func CreateEditProfileWorkerScreen(state AppState) fyne.CanvasObject {
	// Header
	title := widget.NewLabel("Edit Worker Profile")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	// Profile picture section
	profilePicBtn := widget.NewButton("Change Profile Picture", func() {
		fmt.Println("Change profile picture clicked")
		// TODO: Implement image picker
	})

	// Basic Information
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Full Name")
	nameEntry.SetText("Mohamed Hassan")

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Email:")
	emailEntry.SetText("mohamed@example.com")

	phoneEntry := widget.NewEntry()
	phoneEntry.SetPlaceHolder("Phone Number:")
	phoneEntry.SetText("+2164567890")

	locationEntry := widget.NewEntry()
	locationEntry.SetPlaceHolder("Location/City:")
	locationEntry.SetText("Sousse, Tunisia")

	// Professional Information
	professionEntry := widget.NewEntry()
	professionEntry.SetPlaceHolder("Profession: (e.g., Plumber, Electrician)")
	professionEntry.SetText("Plumber")

	experienceEntry := widget.NewEntry()
	experienceEntry.SetPlaceHolder("Years of Experience:")
	experienceEntry.SetText("12")

	hourlyRateEntry := widget.NewEntry()
	hourlyRateEntry.SetPlaceHolder("Hourly Rate ($)")
	hourlyRateEntry.SetText("80")

	minHoursEntry := widget.NewEntry()
	minHoursEntry.SetPlaceHolder("Minimum Hours")
	minHoursEntry.SetText("2")

	// Skills section
	skillsEntry := widget.NewMultiLineEntry()
	skillsEntry.SetPlaceHolder("List your skills (one per line)")
	skillsEntry.SetMinRowsVisible(3)
	skillsEntry.SetText("Pipe Installation\nLeak Repair\nBathroom Renovation")

	// About/Bio
	bioEntry := widget.NewMultiLineEntry()
	bioEntry.SetPlaceHolder("Tell clients about your experience and expertise...")
	bioEntry.SetMinRowsVisible(4)
	bioEntry.SetText("Professional plumber with 12 years of experience in all plumbing work...")

	// Availability toggle
	availableCheck := widget.NewCheck("Available for new jobs", func(checked bool) {
		fmt.Println("Availability:", checked)
	})
	availableCheck.SetChecked(true)

	// Certificates section
	certificatesLabel := widget.NewLabel("Certificates: 340")
	addCertBtn := widget.NewButton("Add Certificate", func() {
		fmt.Println("Add certificate clicked")
		// TODO: Implement certificate upload
	})

	// Save button
	saveBtn := widget.NewButton("Save Changes", func() {
		fmt.Println("Saving worker profile...")
		fmt.Println("Name:", nameEntry.Text)
		fmt.Println("Profession:", professionEntry.Text)
		fmt.Println("Experience:", experienceEntry.Text)
		fmt.Println("Rate:", hourlyRateEntry.Text)
		fmt.Println("Skills:", skillsEntry.Text)
		// TODO: Send to API
		state.ShowScreen("main")
	})
	saveBtn.Importance = widget.HighImportance

	// Cancel button
	cancelBtn := widget.NewButton("Cancel", func() {
		state.ShowScreen("main")
	})

	// Layout
	content := container.NewVBox(
		title,
		layout.NewSpacer(),
		profilePicBtn,

		widget.NewLabel("Personal Information"),
		nameEntry,
		emailEntry,
		phoneEntry,
		locationEntry,

		widget.NewLabel("Professional Information"),
		professionEntry,
		experienceEntry,

		widget.NewLabel("Pricing"),
		hourlyRateEntry,
		minHoursEntry,

		widget.NewLabel("Skills"),
		skillsEntry,

		widget.NewLabel("About"),
		bioEntry,

		widget.NewLabel("Certificates"),
		certificatesLabel,
		addCertBtn,

		availableCheck,
		layout.NewSpacer(),
		saveBtn,
		cancelBtn,
	)

	scroll := container.NewVScroll(content)
	return scroll
}
