package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}

func (pxCanvas *PxCanvas) MouseIn() {
}

func (pxCanvas *PxCanvas) MouseOut() {
}