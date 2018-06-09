package cmd

// App is application interface which will be used
// in all commands
type App interface {
	Daily()
	Info()
}
