//go:generate fyne bundle -o bundle.go assets

package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	skilltheme "skillDar/pkg/theme"
	uiscreen "skillDar/pkg/ui"
)

// AppState manages navigation and theme across screens
type AppState struct {
	app         fyne.App
	window      fyne.Window
	isDarkTheme bool
	screens     map[string]fyne.CanvasObject
	icons       map[string]fyne.Resource // Map of all app icons
	userRole    string                   // "client" or "worker"
}

// ShowScreen displays a screen by name with the top bar
func (as *AppState) ShowScreen(screenName string) {
	if screen, exists := as.screens[screenName]; exists {
		// Wrap screen with top bar
		layout := container.NewBorder(
			as.createTopBar(), // Top (theme toggle)
			nil,               // Bottom
			nil,               // Left
			nil,               // Right
			screen,            // Center (screen content)
		)
		as.window.SetContent(layout)
	}
}

// createTopBar builds the top navigation bar with theme toggle
func (as *AppState) createTopBar() *fyne.Container {
	var themeBtn *widget.Button
	themeBtn = widget.NewButtonWithIcon("", as.icons["darkTheme"], func() {
		as.toggleTheme()
		themeBtn.SetIcon(as.getThemeIcon())
	})
	return container.NewHBox(
		themeBtn,
		layout.NewSpacer(),
	)
}

// toggleTheme switches between light and dark theme
func (as *AppState) toggleTheme() {
	as.isDarkTheme = !as.isDarkTheme
	fmt.Println("Theme toggled. isDarkTheme:", as.isDarkTheme)

	variant := theme.VariantLight
	if as.isDarkTheme {
		variant = theme.VariantDark
	}
	as.app.Settings().SetTheme(skilltheme.NewSkillKonnectTheme(variant))
	as.window.Content().Refresh()
}

// getThemeIcon returns the appropriate icon for current theme
func (as *AppState) getThemeIcon() fyne.Resource {
	if as.isDarkTheme {
		return as.icons["lightTheme"]
	}
	return as.icons["darkTheme"]
}

// GetImage returns an image resource by name
func (as *AppState) GetImage(name string) fyne.Resource {
	return as.icons[name]
}

// SetUserRole sets the user role (client or worker)
func (as *AppState) SetUserRole(role string) {
	as.userRole = role
	fmt.Println("User role set to:", role)
}

// GetUserRole returns the current user role
func (as *AppState) GetUserRole() string {
	return as.userRole
}

func main() {
	// Create the app
	a := app.New()
	w := a.NewWindow("SkillKonnect")
	w.SetMaster()
	w.Resize(fyne.NewSize(390, 844)) // iPhone 12/13 size

	// Initialize app state
	state := &AppState{
		app:         a,
		window:      w,
		isDarkTheme: false, // Start with LIGHT theme
		screens:     make(map[string]fyne.CanvasObject),
		icons: map[string]fyne.Resource{
			"lightTheme": resourceThemeLightlPng,
			"darkTheme":  resourceDarckThemePng,
			"plumberFix": resourcePlumberFixJpg,
			"client":     resourceClientJpg,
			"logoImage":  resourceSkilldarPng,
			// Add more icons here as needed:
			// "home":     resourceHomePng,
			// "settings": resourceSettingsPng,
			// "profile":  resourceProfilePng,
		},
	}

	// Set initial theme
	a.Settings().SetTheme(skilltheme.NewSkillKonnectTheme(theme.VariantLight))

	// Register screens
	state.screens["welcome"] = uiscreen.CreateWelcomeScreen(state)
	state.screens["choice"] = uiscreen.CreateChoiceScreen(state)
	state.screens["login"] = uiscreen.CreateLoginScreen(state)
	state.screens["main"] = uiscreen.CreateMainScreen(state)
	state.screens["profile"] = uiscreen.CreateProfileScreen(state)
	state.screens["edit_profile_client"] = uiscreen.CreateEditProfileClientScreen(state)
	state.screens["edit_profile_worker"] = uiscreen.CreateEditProfileWorkerScreen(state)

	// Show welcome screen first
	state.ShowScreen("welcome")

	// Make sure window is visible
	w.Show()
	w.CenterOnScreen()

	// Show and run
	w.ShowAndRun()
}
