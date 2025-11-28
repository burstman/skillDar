package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateJobSearchScreen builds the job search screen for workers
func CreateJobSearchScreen(state AppState) fyne.CanvasObject {
	// Header
	title := widget.NewLabel("Available Jobs")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	// Search bar
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search jobs...")

	// Filter options
	categorySelect := widget.NewSelect([]string{"All Categories", "Plumbing", "Electrical", "Carpentry", "Painting", "Other"}, func(value string) {
		fmt.Println("Category selected:", value)
	})
	categorySelect.SetSelected("All Categories")

	sortSelect := widget.NewSelect([]string{"Newest", "Highest Budget", "Closest", "Urgent"}, func(value string) {
		fmt.Println("Sort by:", value)
	})
	sortSelect.SetSelected("Newest")

	filterRow := container.NewGridWithColumns(2, categorySelect, sortSelect)

	// Job listings (example jobs)
	job1 := createJobCard(
		"Fix Kitchen Sink Leak",
		"Urgent plumbing needed. Kitchen sink has a major leak...",
		"$50 - $100",
		"New York, NY",
		"2 hours ago",
		func() {
			showBidDialog(state, "Fix Kitchen Sink Leak")
		},
	)

	job2 := createJobCard(
		"Install Bathroom Fixtures",
		"Need professional to install new toilet and sink...",
		"$150 - $250",
		"Brooklyn, NY",
		"5 hours ago",
		func() {
			showBidDialog(state, "Install Bathroom Fixtures")
		},
	)

	job3 := createJobCard(
		"Water Heater Repair",
		"Water heater not working properly, needs inspection...",
		"$100 - $200",
		"Queens, NY",
		"1 day ago",
		func() {
			showBidDialog(state, "Water Heater Repair")
		},
	)

	// Jobs list
	jobsList := container.NewVBox(
		job1,
		job2,
		job3,
	)

	// Main content
	content := container.NewVBox(
		title,
		searchEntry,
		filterRow,
		widget.NewSeparator(),
		jobsList,
	)

	scroll := container.NewVScroll(content)

	// Bottom navigation bar
	bottomNav := createBottomNavForWorker(state)

	// Main layout with bottom navigation
	mainLayout := container.NewBorder(
		nil,       // top
		bottomNav, // bottom
		nil,       // left
		nil,       // right
		scroll,    // center
	)

	return mainLayout
}

// createBottomNavForWorker creates bottom navigation for worker screens
func createBottomNavForWorker(state AppState) fyne.CanvasObject {
	// Navigation bar background
	navBg := canvas.NewRectangle(color.RGBA{245, 245, 245, 255})

	// Create navigation buttons
	homeBtn := widget.NewButton("🏠\nHome", func() {
		state.ShowScreen("client_post_job")
	})
	homeBtn.Importance = widget.HighImportance

	ordersBtn := widget.NewButton("📋\nOrders", func() {
		// TODO: Navigate to orders screen
	})

	chatBtn := widget.NewButton("💬\nChat", func() {
		// TODO: Navigate to chat screen
	})

	profileBtn := widget.NewButton("👤\nProfile", func() {
		state.ShowScreen("edit_profile_worker")
	})

	navButtons := container.NewHBox(
		layout.NewSpacer(),
		homeBtn,
		layout.NewSpacer(),
		ordersBtn,
		layout.NewSpacer(),
		chatBtn,
		layout.NewSpacer(),
		profileBtn,
		layout.NewSpacer(),
	)

	return container.NewStack(
		navBg,
		container.NewPadded(navButtons),
	)
}

// createJobCard creates a card for a job listing
func createJobCard(title, description, budget, location, timeAgo string, onBid func()) *fyne.Container {
	// Job title
	titleLabel := widget.NewLabel(title)
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Description
	descLabel := widget.NewLabel(description)
	descLabel.Wrapping = fyne.TextWrapWord

	// Budget with icon
	budgetLabel := widget.NewLabel("💰 " + budget)

	// Location and time
	locationLabel := widget.NewLabel("📍 " + location)
	timeLabel := widget.NewLabel("🕐 " + timeAgo)
	timeLabel.Importance = widget.LowImportance

	metaRow := container.NewHBox(locationLabel, layout.NewSpacer(), timeLabel)

	// Bid button
	bidBtn := widget.NewButton("Place Bid", onBid)
	bidBtn.Importance = widget.HighImportance

	// Card content
	cardContent := container.NewVBox(
		titleLabel,
		descLabel,
		budgetLabel,
		metaRow,
		bidBtn,
	)

	// Card background
	background := canvas.NewRectangle(color.RGBA{R: 245, G: 245, B: 245, A: 255})
	card := container.NewStack(background, container.NewPadded(cardContent))

	return container.NewVBox(card, widget.NewSeparator())
}

// showBidDialog shows a dialog for placing a bid
func showBidDialog(state AppState, jobTitle string) {
	// Create bid entry
	bidEntry := widget.NewEntry()
	bidEntry.SetPlaceHolder("Enter your bid amount ($)")

	messageEntry := widget.NewMultiLineEntry()
	messageEntry.SetPlaceHolder("Message to client (optional)...")
	messageEntry.SetMinRowsVisible(3)

	estimatedTimeEntry := widget.NewEntry()
	estimatedTimeEntry.SetPlaceHolder("Estimated completion time (e.g., 2 hours)")

	// Submit button
	submitBtn := widget.NewButton("Submit Bid", func() {
		fmt.Println("Bid submitted for:", jobTitle)
		fmt.Println("Amount:", bidEntry.Text)
		fmt.Println("Time:", estimatedTimeEntry.Text)
		fmt.Println("Message:", messageEntry.Text)
		// TODO: Send bid to API
	})
	submitBtn.Importance = widget.SuccessImportance

	// For now, just print the bid info (you can create a modal dialog or dedicated screen later)
	fmt.Println("Opening bid dialog for:", jobTitle)
	fmt.Println("Bid Entry:", bidEntry.Text)
	fmt.Println("Time Entry:", estimatedTimeEntry.Text)
	fmt.Println("Message Entry:", messageEntry.Text)

	// TODO: You could navigate to a dedicated bid screen or show a popup
	// state.ShowScreen("bid_details")
}
