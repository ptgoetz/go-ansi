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
	// ColorBlack represents a black foreground
	ColorBlack = normalForeground + black
	// ColorRed represents a red foreground
	ColorRed = normalForeground + red
	// ColorGreen represents a green foreground
	ColorGreen = normalForeground + green
	// ColorYellow represents a yellow foreground
	ColorYellow = normalForeground + yellow
	// ColorBlue represents a blue foreground
	ColorBlue = normalForeground + blue
	// ColorMagenta represents a magenta foreground
	ColorMagenta = normalForeground + magenta
	// ColorCyan represents a cyan foreground
	ColorCyan = normalForeground + cyan
	// ColorWhite represents a white foreground
	ColorWhite = normalForeground + white
	// ColorRGB is an ANSI code to render a color in `r;g;b;` format.
	ColorRGB = 38
	// ColorDefault represents the default foreground
	ColorDefault = normalForeground + none

	// ColorBrightBlack represents a bright black foreground
	ColorBrightBlack = brightForeground + black
	// ColorBrightRed represents a bright red foreground
	ColorBrightRed = brightForeground + red
	// ColorBrightGreen represents a bright green foreground
	ColorBrightGreen = brightForeground + green
	// ColorBrightYellow represents a bright yellow foreground
	ColorBrightYellow = brightForeground + yellow
	// ColorBrightBlue represents a bright blue foreground
	ColorBrightBlue = brightForeground + blue
	// ColorBrightMagenta represents a bright magenta foreground
	ColorBrightMagenta = brightForeground + magenta
	// ColorBrightCyan represents a bright cyan foreground
	ColorBrightCyan = brightForeground + cyan
	// ColorBrightWhite represents a bright white foreground
	ColorBrightWhite = brightForeground + white
	
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
