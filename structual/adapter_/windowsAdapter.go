package main

type windowsAdapter struct {
	windowMachine *windows
}
//adapter mapping interface to another
func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}