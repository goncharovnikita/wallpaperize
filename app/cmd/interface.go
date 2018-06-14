package cmd

// App is application interface which will be used
// in all commands
type App interface {
	Daily()
	Info(format string)
	Random(loadOnly bool)
	Restore()
	Set(path string)
}
