package main

const (
	fgBlack   = "\x1b[38;05;0m"
	fgRed     = "\x1b[38;05;1m"
	fgGreen   = "\x1b[38;05;2m"
	fgYellow  = "\x1b[38;05;3m"
	fgBlue    = "\x1b[38;05;4m"
	fgMagenta = "\x1b[38;05;5m"
	fgCyan    = "\x1b[38;05;6m"
	fgGrey    = "\x1b[38;05;7m"
	fgWhite   = "\x1b[38;05;255m"

	bgBlack   = "\x1b[48;05;0m"
	bgRed     = "\x1b[48;05;1m"
	bgGreen   = "\x1b[48;05;2m"
	bgYellow  = "\x1b[48;05;3m"
	bgBlue    = "\x1b[48;05;4m"
	bgMagenta = "\x1b[48;05;5m"
	bgCyan    = "\x1b[48;05;6m"
	bgGrey    = "\x1b[48;05;7m"
	bgWhite   = "\x1b[48;05;255m"

	bold  = "\x1b[1m"
	reset = "\x1b[0m"
)

func sgrBoldRed(text string) string {
	return fgMagenta + bold + text + reset
}

func sgrBoldBlue(text string) string {
	return fgYellow + bold + text + reset
}
