package ui

// AppState defines the interface for app state management
// This allows screens to access navigation and app-level state
type AppState interface {
	ShowScreen(screenName string)
}
