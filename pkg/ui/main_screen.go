package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateMainScreen builds the main app screen with bottom navigation
func CreateMainScreen(state AppState) fyne.CanvasObject {
	// Content container that will change based on selected tab
	var currentContent *fyne.Container

	userRole := state.GetUserRole()

	// Initial content based on role
	if userRole == "worker" {
		currentContent = container.NewVBox(createWorkerHomeContent(state))
	} else {
		currentContent = container.NewVBox(createClientHomeContent(state))
	}

	// Bottom navigation bar
	bottomNav := createBottomNavigationBar(state, currentContent)

	// Main layout with bottom navigation
	mainLayout := container.NewBorder(
		nil,                                 // top
		bottomNav,                           // bottom
		nil,                                 // left
		nil,                                 // right
		container.NewScroll(currentContent), // center
	)

	return mainLayout
}

// createBottomNavigationBar creates the bottom navigation menu
func createBottomNavigationBar(state AppState, contentContainer *fyne.Container) fyne.CanvasObject {
	userRole := state.GetUserRole()

	// Create navigation buttons
	homeBtn := createNavButton("🏠", "Home", true)
	ordersBtn := createNavButton("📋", "Orders", false)
	chatBtn := createNavButton("💬", "Chat", false)
	profileBtn := createNavButton("👤", "Profile", false)

	// Navigation bar background
	navBg := canvas.NewRectangle(color.RGBA{245, 245, 245, 255})

	// Button handlers
	homeBtn.OnTapped = func() {
		updateNavButtons(homeBtn, ordersBtn, chatBtn, profileBtn)
		if userRole == "worker" {
			contentContainer.Objects = []fyne.CanvasObject{createWorkerHomeContent(state)}
		} else {
			contentContainer.Objects = []fyne.CanvasObject{createClientHomeContent(state)}
		}
		contentContainer.Refresh()
	}

	ordersBtn.OnTapped = func() {
		updateNavButtons(ordersBtn, homeBtn, chatBtn, profileBtn)
		contentContainer.Objects = []fyne.CanvasObject{createOrdersContent(state)}
		contentContainer.Refresh()
	}

	chatBtn.OnTapped = func() {
		updateNavButtons(chatBtn, homeBtn, ordersBtn, profileBtn)
		contentContainer.Objects = []fyne.CanvasObject{createChatContent(state)}
		contentContainer.Refresh()
	}

	profileBtn.OnTapped = func() {
		updateNavButtons(profileBtn, homeBtn, ordersBtn, chatBtn)
		if userRole == "worker" {
			state.ShowScreen("edit_profile_worker")
		} else {
			state.ShowScreen("edit_profile_client")
		}
	}

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

// createNavButton creates a navigation button
func createNavButton(icon, label string, active bool) *widget.Button {
	btn := widget.NewButton(icon+"\n"+label, nil)
	if active {
		btn.Importance = widget.HighImportance
	}
	return btn
}

// updateNavButtons updates the active state of navigation buttons
func updateNavButtons(active *widget.Button, others ...*widget.Button) {
	active.Importance = widget.HighImportance
	active.Refresh()
	for _, btn := range others {
		btn.Importance = widget.MediumImportance
		btn.Refresh()
	}
}

// createClientHomeContent creates the home content for clients
func createClientHomeContent(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("Available Workers")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	// Search bar
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search for workers...")

	// Professional categories
	categoriesLabel := widget.NewLabel("Professional Categories")
	categoriesLabel.TextStyle = fyne.TextStyle{Bold: true}

	painterCard := createSimpleCategoryCard("🎨", "Painter")
	plumberCard := createSimpleCategoryCard("🔧", "Plumber")
	carpenterCard := createSimpleCategoryCard("🪛", "Carpenter")
	electricianCard := createSimpleCategoryCard("⚡", "Electrician")
	masonCard := createSimpleCategoryCard("🧱", "Mason")
	acCard := createSimpleCategoryCard("❄️", "AC Tech")

	categoriesGrid := container.NewGridWithColumns(3,
		painterCard, plumberCard, carpenterCard,
		electricianCard, masonCard, acCard,
	)

	// Available workers
	workersLabel := widget.NewLabel("Available Workers")
	workersLabel.TextStyle = fyne.TextStyle{Bold: true}

	worker1 := createSimpleWorkerCard("Mohamed Hassan", "Plumber", "4.9", "0.8 km", "80 EGP/hr", true)
	worker2 := createSimpleWorkerCard("Karim Fathy", "Electrician", "4.5", "1.2 km", "75 EGP/hr", true)
	worker3 := createSimpleWorkerCard("Ahmed Ali", "Carpenter", "4.8", "2.1 km", "90 EGP/hr", false)

	return container.NewVBox(
		title,
		searchEntry,
		categoriesLabel,
		categoriesGrid,
		workersLabel,
		worker1,
		worker2,
		worker3,
	)
}

// createWorkerHomeContent creates the home content for workers
func createWorkerHomeContent(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("Available Projects")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	// Search/filter bar
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search for projects...")

	// Available projects
	projectsLabel := widget.NewLabel("Nearby Projects")
	projectsLabel.TextStyle = fyne.TextStyle{Bold: true}

	project1 := createProjectCard("Fix Kitchen Sink", "Plumbing needed", "150 EGP", "1.5 km", "2 hours ago")
	project2 := createProjectCard("Install AC Unit", "AC installation", "300 EGP", "0.8 km", "5 hours ago")
	project3 := createProjectCard("Repair Door", "Carpentry work", "120 EGP", "3.2 km", "1 day ago")

	return container.NewVBox(
		title,
		searchEntry,
		projectsLabel,
		project1,
		project2,
		project3,
	)
}

// createOrdersContent creates the orders/bookings content
func createOrdersContent(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("My Orders")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	noOrders := widget.NewLabel("No orders yet")
	noOrders.Alignment = fyne.TextAlignCenter

	return container.NewVBox(
		layout.NewSpacer(),
		title,
		noOrders,
		layout.NewSpacer(),
	)
}

// createChatContent creates the chat/messages content
func createChatContent(state AppState) fyne.CanvasObject {
	title := widget.NewLabel("Messages")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	noMessages := widget.NewLabel("No messages yet")
	noMessages.Alignment = fyne.TextAlignCenter

	return container.NewVBox(
		layout.NewSpacer(),
		title,
		noMessages,
		layout.NewSpacer(),
	)
}

// createProjectCard creates a project card for workers
func createProjectCard(title, description, price, distance, posted string) fyne.CanvasObject {
	titleLabel := widget.NewLabel(title)
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	descLabel := widget.NewLabel(description)

	priceLabel := widget.NewLabel("💰 " + price)
	distanceLabel := widget.NewLabel("📍 " + distance)
	postedLabel := widget.NewLabel("🕐 " + posted)

	details := container.NewHBox(distanceLabel, postedLabel)

	applyBtn := widget.NewButton("Apply", func() {
		// Handle apply to project
	})
	applyBtn.Importance = widget.HighImportance

	cardContent := container.NewVBox(
		titleLabel,
		descLabel,
		priceLabel,
		details,
		applyBtn,
	)

	return container.NewPadded(cardContent)
}

// createSimpleWorkerCard creates a worker card for clients (simplified version)
func createSimpleWorkerCard(name, profession, rating, distance, price string, available bool) fyne.CanvasObject {
	nameLabel := widget.NewLabel(name)
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}

	professionLabel := widget.NewLabel(profession)

	ratingLabel := widget.NewLabel("⭐ " + rating)
	distanceLabel := widget.NewLabel("📍 " + distance)

	priceLabel := widget.NewLabel(price)

	statusLabel := widget.NewLabel("✅ Available")
	if !available {
		statusLabel.Text = "⏰ Busy"
	}

	contactBtn := widget.NewButton("Contact", func() {
		// Handle contact worker
	})
	if available {
		contactBtn.Importance = widget.SuccessImportance
	}

	info := container.NewVBox(
		nameLabel,
		professionLabel,
		container.NewHBox(ratingLabel, distanceLabel),
	)

	rightSide := container.NewVBox(
		priceLabel,
		statusLabel,
		contactBtn,
	)

	cardContent := container.NewBorder(
		nil, nil,
		info,
		rightSide,
	)

	return container.NewPadded(cardContent)
}

// createSimpleCategoryCard creates a category card
func createSimpleCategoryCard(icon, name string) fyne.CanvasObject {
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
