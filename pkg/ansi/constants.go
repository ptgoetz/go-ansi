package ansi

// CursorDirection represents the direction a cursor can travel (Up, Down, Left, Right)
type CursorDirection string

const (
	// CursorUp moves the cursor up.
	CursorUp CursorDirection = "A"
	// CursorDown moves the cursor down.
	CursorDown CursorDirection = "B"
	// CursorRight moves the cursor right.
	CursorRight CursorDirection = "C"
	// CursorLeft moves the cursor left.
	CursorLeft CursorDirection = "D"
)

// Style constants
const (
	StyleNone = iota
	StyleBold
	StyleDim
	StyleItalic
	StyleUnderline
	StyleBlink
	StyleInverse
	StyleStrikethrough
)

// Basic color constants
const (
	black = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	white
	none = 9
)

// ANSI foreground/background constants
const (
	normalForeground = 30
	normalBackground = 40
	brightForeground = 90
	brightBackground = 100
)

const (
	// Start is the ANSI code to start an instruction
	Start = "\033["
	// Reset is an ANSI instruction that will restore all ANSI formatting defaults
	Reset = Start + "0m"
)

const (
	ColorBlack   = normalForeground + black
	ColorRed     = normalForeground + red
	ColorGreen   = normalForeground + green
	ColorYellow  = normalForeground + yellow
	ColorBlue    = normalForeground + blue
	ColorMagenta = normalForeground + magenta
	ColorCyan    = normalForeground + cyan
	ColorWhite   = normalForeground + white
	ColorRGB     = 38
	ColorDefault = normalForeground + none

	ColorBrightBlack   = brightForeground + black
	ColorBrightRed     = brightForeground + red
	ColorBrightGreen   = brightForeground + green
	ColorBrightYellow  = brightForeground + yellow
	ColorBrightBlue    = brightForeground + blue
	ColorBrightMagenta = brightForeground + magenta
	ColorBrightCyan    = brightForeground + cyan
	ColorBrightWhite   = brightForeground + white

	ColorBgBlack   = normalBackground + black
	ColorBgRed     = normalBackground + red
	ColorBgGreen   = normalBackground + green
	ColorBgYellow  = normalBackground + yellow
	ColorBgBlue    = normalBackground + blue
	ColorBgMagenta = normalBackground + magenta
	ColorBgCyan    = normalBackground + cyan
	ColorBgWhite   = normalBackground + white
	ColorBgRGB     = 48
	ColorBgDefault = normalBackground + none

	ColorBgBrightBlack   = brightBackground + black
	ColorBgBrightRed     = brightBackground + red
	ColorBgBrightGreen   = brightBackground + green
	ColorBgBrightYellow  = brightBackground + yellow
	ColorBgBrightBlue    = brightBackground + blue
	ColorBgBrightMagenta = brightBackground + magenta
	ColorBgBrightCyan    = brightBackground + cyan
	ColorBgBrightWhite   = brightBackground + white
)
