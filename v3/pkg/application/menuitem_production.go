//go:build darwin && production && !devtools

package application

func newShowDevToolsMenuItem() *MenuItem {
	return nil
}
