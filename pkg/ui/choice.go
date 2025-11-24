package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateChoiceScreen builds the choice screen
func CreateChoiceScreen(state AppState) fyne.CanvasObject {
	// Header with icon and title
	iconLabel := widget.NewLabel("💼")
	iconLabel.Alignment = fyne.TextAlignCenter
	iconLabel.TextStyle = fyne.TextStyle{Bold: true}

	headerTitle1 := canvas.NewText("Your Home, Our Expertise.", color.Black)
	headerTitle1.Alignment = fyne.TextAlignCenter
	headerTitle1.TextSize = 18
	headerTitle1.TextStyle = fyne.TextStyle{Bold: true}

	headerBox := container.NewVBox(
		iconLabel,
		headerTitle1,
	)

	// Skilled Worker section
	skilledWorkerIcon := widget.NewLabel("🔧")
	skilledWorkerTitle := widget.NewLabel(" Skilled Worker")
	skilledWorkerTitle.TextStyle = fyne.TextStyle{Bold: true}

	skilledWorkerHeader := container.NewHBox(
		skilledWorkerIcon,
		skilledWorkerTitle,
	)

	skilledWorkerDesc := widget.NewLabel("Find job opportunities and connect\nwith clients.")

	skilledWorkerBtn := widget.NewButton("I am a Skilled Worker", func() {
		state.ShowScreen("main")
	})

	skilledWorkerBox := container.NewVBox(
		container.NewCenter(skilledWorkerHeader),
		container.NewCenter(skilledWorkerDesc),
		skilledWorkerBtn,
	)

	// Decorative image
	plumberImage := canvas.NewImageFromResource(state.GetImage("plumberFix"))
	plumberImage.FillMode = canvas.ImageFillContain
	plumberImage.SetMinSize(fyne.NewSize(390, 220))

	// Client section
	clientIcon := widget.NewLabel("🏠")
	clientTitle := widget.NewLabel(" Client")
	clientTitle.TextStyle = fyne.TextStyle{Bold: true}

	clientHeader := container.NewHBox(
		clientIcon,
		clientTitle,
	)

	clientDesc := widget.NewLabel("Find reliable professionals for your\nhome projects.")

	clientBtn := widget.NewButton("I am a Client", func() {
		state.ShowScreen("main")
	})

	clientBox := container.NewVBox(
		clientHeader,
		clientDesc,
		clientBtn,
	)

	// Main scrollable layout
	content := container.NewVBox(
		headerBox,
		//layout.NewSpacer(),
		skilledWorkerBox,
		//layout.NewSpacer(),
		plumberImage,
		layout.NewSpacer(),
		clientBox,
	)

	// Wrap in scroll container for mobile-like experience
	scroll := container.NewScroll(content)

	// Return padded content
	return container.NewPadded(scroll)
}
