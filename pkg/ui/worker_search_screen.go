package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateWorkerSearchScreen creates the worker search screen for clients
func CreateWorkerSearchScreen(state AppState) fyne.CanvasObject {
	// Header with location
	locationIcon := widget.NewLabel("📍")
	locationText := widget.NewLabel("Your Location\nCairo, Egypt")
	locationText.TextStyle = fyne.TextStyle{Bold: true}

	notificationBtn := widget.NewButton("🔔", func() {
		// Handle notifications
	})

	header := container.NewBorder(
		nil, nil,
		container.NewHBox(locationIcon, locationText),
		notificationBtn,
	)

	// Search bar
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search for craftsman near you...")

	searchBar := container.NewBorder(
		nil, nil, nil,
		widget.NewButton("🔍", func() {
			// Handle search
		}),
		searchEntry,
	)

	// Map placeholder (simple colored rectangle for now)
	mapPlaceholder := canvas.NewRectangle(color.RGBA{200, 220, 240, 255})
	mapPlaceholder.SetMinSize(fyne.NewSize(0, 250))

	// Add some worker markers on the map
	worker1 := canvas.NewCircle(color.RGBA{40, 125, 247, 255})
	worker1.Resize(fyne.NewSize(30, 30))
	worker1.Move(fyne.NewPos(100, 80))

	worker2 := canvas.NewCircle(color.RGBA{76, 175, 80, 255})
	worker2.Resize(fyne.NewSize(30, 30))
	worker2.Move(fyne.NewPos(200, 150))

	worker3 := canvas.NewCircle(color.RGBA{255, 152, 0, 255})
	worker3.Resize(fyne.NewSize(30, 30))
	worker3.Move(fyne.NewPos(150, 180))

	mapContainer := container.NewWithoutLayout(
		mapPlaceholder,
		worker1,
		worker2,
		worker3,
	)

	// Action buttons
	nearbyBtn := widget.NewButton("🏠 Nearby apartments before neighboring", func() {
		// Handle nearby apartments
	})
	nearbyBtn.Importance = widget.HighImportance

	needHelpBtn := widget.NewButton("⏰ Need help now? Add us only", func() {
		// Handle urgent help
	})
	needHelpBtn.Importance = widget.DangerImportance

	// Professional categories section
	categoriesLabel := widget.NewLabel("Professional Categories")
	categoriesLabel.TextStyle = fyne.TextStyle{Bold: true}

	// Category cards
	painterCard := createCategoryCard("🎨", "Painter")
	plumberCard := createCategoryCard("🔧", "Plumber")
	carpenterCard := createCategoryCard("🪛", "Carpenter")
	electricianCard := createCategoryCard("⚡", "Electrician")
	masonCard := createCategoryCard("🧱", "Mason")
	acCard := createCategoryCard("❄️", "AC Technician")

	categoriesGrid := container.NewGridWithColumns(3,
		painterCard,
		plumberCard,
		carpenterCard,
		electricianCard,
		masonCard,
		acCard,
	)

	// Nearby workers section
	nearbyWorkersLabel := widget.NewLabel("Nearby Workers (2)")
	nearbyWorkersLabel.TextStyle = fyne.TextStyle{Bold: true}

	worker1Card := createWorkerCard("Mohamed Hassan", "Plumber", "180 EGP/hour", "4.9", "127")
	worker2Card := createWorkerCard("Karim Fathy", "Plumber", "160 EGP/hour", "4.5", "76")

	// Main content
	content := container.NewVBox(
		header,
		searchBar,
		mapContainer,
		nearbyBtn,
		needHelpBtn,
		categoriesLabel,
		categoriesGrid,
		nearbyWorkersLabel,
		worker1Card,
		worker2Card,
	)

	scroll := container.NewScroll(content)

	// Bottom navigation bar
	bottomNav := createBottomNavForClient(state)

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

// createBottomNavForClient creates bottom navigation for client screens
func createBottomNavForClient(state AppState) fyne.CanvasObject {
	// Navigation bar background
	navBg := canvas.NewRectangle(color.RGBA{245, 245, 245, 255})

	// Create navigation buttons
	homeBtn := widget.NewButton("🏠\nHome", func() {
		state.ShowScreen("worker_search_screen")
	})
	homeBtn.Importance = widget.HighImportance

	ordersBtn := widget.NewButton("📋\nOrders", func() {
		// TODO: Navigate to orders screen
	})

	chatBtn := widget.NewButton("💬\nChat", func() {
		// TODO: Navigate to chat screen
	})

	profileBtn := widget.NewButton("👤\nProfile", func() {
		state.ShowScreen("edit_profile_client")
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

// createCategoryCard creates a card for a professional category
func createCategoryCard(icon, name string) fyne.CanvasObject {
	iconLabel := widget.NewLabel(icon)
	iconLabel.Alignment = fyne.TextAlignCenter
	iconLabel.TextStyle = fyne.TextStyle{Bold: true}

	nameLabel := widget.NewLabel(name)
	nameLabel.Alignment = fyne.TextAlignCenter

	card := container.NewVBox(
		iconLabel,
		nameLabel,
	)

	return container.NewPadded(card)
}

// createWorkerCard creates a card for displaying worker information
func createWorkerCard(name, profession, price, rating, reviews string) fyne.CanvasObject {
	// Avatar placeholder
	avatar := canvas.NewCircle(color.RGBA{200, 200, 200, 255})
	avatar.Resize(fyne.NewSize(50, 50))

	// Worker info
	nameLabel := widget.NewLabel(name)
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}

	professionLabel := widget.NewLabel(profession)

	priceLabel := widget.NewLabel(price)
	priceLabel.TextStyle = fyne.TextStyle{Bold: true}
	priceLabel.Alignment = fyne.TextAlignTrailing

	ratingLabel := widget.NewLabel("⭐ " + rating)
	reviewsLabel := widget.NewLabel("🚩 " + reviews + " reviews")

	statusLabel := widget.NewLabel("✅ Available")
	statusLabel.Alignment = fyne.TextAlignCenter

	availableBtn := widget.NewButton("Available", func() {
		// Contact worker
	})
	availableBtn.Importance = widget.SuccessImportance

	// Layout
	leftInfo := container.NewVBox(
		nameLabel,
		professionLabel,
		container.NewHBox(ratingLabel, reviewsLabel),
	)

	rightInfo := container.NewVBox(
		priceLabel,
		layout.NewSpacer(),
		availableBtn,
	)

	cardContent := container.NewBorder(
		nil, nil,
		container.NewHBox(avatar, leftInfo),
		rightInfo,
	)

	return container.NewPadded(cardContent)
}
