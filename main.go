package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type skillKonnectTheme struct{}

// Color lets you override specific named colors.
func (skillKonnectTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary:
		return color.RGBA{R: 0x28, G: 0x7D, B: 0xF7, A: 0xFF} // brand blue
	}
	// fall back to Fyne's default theme
	return theme.DefaultTheme().Color(name, variant)
}

func (skillKonnectTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (skillKonnectTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (skillKonnectTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func main() {
	// Create the app
	a := app.New()

	// Usage:
	a.Settings().SetTheme(skillKonnectTheme{})

	// Main window
	w := a.NewWindow("SkillKonnect")
	w.SetMaster() // Uncomment this for full mobile app behavior (full screen, no window borders)

	// Simulate phone size on desktop (ignored on real phone)
	// Common Android phone sizes:
	// - iPhone SE: 375x667
	// - iPhone 12/13: 390x844
	// - Samsung Galaxy S21: 360x800
	// - Pixel 5: 393x851
	// - Average Android phone: ~360-410 width, ~800-900 height
	w.Resize(fyne.NewSize(390, 844)) // iPhone 12/13 size - good Android approximation

	// UI elements
	title := widget.NewLabel("Welcome to SkillKonnect")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	subtitle := widget.NewLabel("Connect skills, build networks")
	subtitle.Alignment = fyne.TextAlignCenter

	getStartedBtn := widget.NewButton("Get Started", func() {
		pop := widget.NewPopUp(widget.NewLabel("Hello from mobile!"), w.Canvas())
		pop.Show()
	})

	// Layout
	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		subtitle,
		layout.NewSpacer(),
		getStartedBtn,
		layout.NewSpacer(),
	)

	// Center everything with padding (looks great on phones)
	w.SetContent(container.NewPadded(container.NewCenter(content)))

	// Make sure window is visible
	w.Show()

	// Center the window on screen
	w.CenterOnScreen()

	// Show and run
	w.ShowAndRun()
}
