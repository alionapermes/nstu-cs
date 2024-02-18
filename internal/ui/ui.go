package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type RectState int

// enum-impl in go
const (
	RectStateBlue  RectState = 1
	RectStateGreen RectState = 2
)

// var syntax:
// var <name> <type> = <expr>
var (
	blue  color.Color = color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	green color.Color = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
)

type UI struct {
	Container *fyne.Container

	rectState RectState

	btn  *widget.Button
	rect *canvas.Rectangle
}

func New() *UI {
	ui := &UI{rectState: RectStateBlue}

	return ui.init()
}

// struct's method syntax:
// func (<this> <struct-owner-type>) <name>([<args>]) <return-type>
func (self *UI) init() *UI {
	return self.initRect().
		initBtn().
		initContainer()
}

func (self *UI) initBtn() *UI {
	btn := widget.NewButton("click", func() {
		switch self.rectState {
		// case in go does "break" by default
		case RectStateBlue:
			self.rect.FillColor = green
			self.rectState = RectStateGreen
		case RectStateGreen:
			self.rect.FillColor = blue
			self.rectState = RectStateBlue
		}

		self.rect.Refresh()
	})

	btn.Resize(fyne.NewSize(60, 60))

	self.btn = btn

	return self
}

func (self *UI) initRect() *UI {
	rect := canvas.NewRectangle(blue)
	rect.Resize(fyne.NewSize(240, 320))
	rect.Move(fyne.NewPos(100, 50))

	self.rect = rect

	return self
}

func (self *UI) initContainer() *UI {
	self.Container = container.NewWithoutLayout(
		self.btn,
		self.rect,
	)

	return self
}
