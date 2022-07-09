package color

// HTML color names
// https://htmlcolorcodes.com/color-names/
var (
	// Red Colors

	IndianRed   = New(205, 92, 92)
	LightCoral  = New(240, 128, 128)
	Salmon      = New(250, 128, 114)
	DarkSalmon  = New(233, 150, 122)
	LightSalmon = New(255, 160, 122)
	Crimson     = New(220, 20, 60)
	Red         = New(255, 0, 0)
	FireBrick   = New(178, 34, 34)
	DarkRed     = New(139, 0, 0)

	// Pink Colors

	Pink            = New(255, 192, 203)
	LightPink       = New(255, 182, 193)
	HotPink         = New(255, 105, 180)
	DeepPink        = New(255, 20, 147)
	MediumVioletRed = New(199, 21, 133)
	PaleVioletRed   = New(219, 112, 147)

	// Orange Colors

	//LightSalmon = New(0, 0, 0)
	Coral      = New(0, 0, 0)
	Tomato     = New(0, 0, 0)
	OrangeRed  = New(0, 0, 0)
	DarkOrange = New(0, 0, 0)
	Orange     = New(0, 0, 0)

	// Yellow Colors

	Gold                 = New(0, 0, 0)
	Yellow               = New(0, 0, 0)
	LightYellow          = New(0, 0, 0)
	LemonChiffon         = New(0, 0, 0)
	LightGoldenrodYellow = New(0, 0, 0)
	PapayaWhip           = New(0, 0, 0)
	Moccasin             = New(0, 0, 0)
	PeachPuff            = New(0, 0, 0)
	PaleGoldenrod        = New(0, 0, 0)
	Khaki                = New(0, 0, 0)
	DarkKhaki            = New(0, 0, 0)

	// Purple Colors

	Lavender        = New(230, 230, 250)
	Thistle         = New(216, 191, 216)
	Plum            = New(221, 160, 221)
	Violet          = New(238, 130, 238)
	Orchid          = New(218, 112, 214)
	Fuchsia         = New(255, 0, 255)
	Magenta         = New(255, 0, 255)
	MediumOrchid    = New(186, 85, 211)
	MediumPurple    = New(147, 112, 219)
	RebeccaPurple   = New(102, 51, 153)
	BlueViolet      = New(102, 51, 153)
	DarkViolet      = New(148, 0, 211)
	DarkOrchid      = New(153, 50, 204)
	DarkMagenta     = New(139, 0, 139)
	Purple          = New(128, 0, 128)
	Indigo          = New(75, 0, 130)
	SlateBlue       = New(106, 90, 205)
	DarkSlateBlue   = New(72, 61, 139)
	MediumSlateBlue = New(123, 104, 238)

	// Green Colors

	GreenYellow       = New(173, 255, 47)
	Chartreuse        = New(127, 255, 0)
	LawnGreen         = New(124, 252, 0)
	Lime              = New(0, 255, 0)
	LimeGreen         = New(50, 205, 50)
	PaleGreen         = New(152, 251, 152)
	LightGreen        = New(144, 238, 144)
	MediumSpringGreen = New(0, 250, 154)
	SpringGreen       = New(0, 255, 127)
	MediumSeaGreen    = New(60, 179, 113)
	SeaGreen          = New(46, 139, 87)
	ForestGreen       = New(34, 139, 34)
	Green             = New(34, 139, 34)
	DarkGreen         = New(0, 100, 0)
	YellowGreen       = New(154, 205, 50)
	OliveDrab         = New(107, 142, 35)
	Olive             = New(128, 128, 0)
	DarkOliveGreen    = New(85, 107, 47)
	MediumAquamarine  = New(102, 205, 170)
	DarkSeaGreen      = New(143, 188, 139)
	LightSeaGreen     = New(32, 178, 170)
	DarkCyan          = New(0, 139, 139)
	Teal              = New(0, 128, 128)

	// Blue Colors

	Aqua            = New(0, 255, 255)
	Cyan            = New(0, 255, 255)
	LightCyan       = New(224, 255, 255)
	PaleTurquoise   = New(175, 238, 238)
	Aquamarine      = New(127, 255, 212)
	Turquoise       = New(64, 224, 208)
	MediumTurquoise = New(72, 209, 204)
	DarkTurquoise   = New(0, 206, 209)
	CadetBlue       = New(95, 158, 160)
	SteelBlue       = New(70, 130, 180)
	LightSteelBlue  = New(176, 196, 222)
	PowderBlue      = New(176, 224, 230)
	LightBlue       = New(173, 216, 230)
	SkyBlue         = New(135, 206, 235)
	LightSkyBlue    = New(135, 206, 250)
	DeepSkyBlue     = New(0, 191, 255)
	DodgerBlue      = New(30, 144, 255)
	CornflowerBlue  = New(100, 149, 237)
	//MediumSlateBlue = New(255, 255, 255)
	RoyalBlue    = New(65, 105, 225)
	Blue         = New(0, 0, 255)
	MediumBlue   = New(0, 0, 205)
	DarkBlue     = New(0, 0, 139)
	Navy         = New(0, 0, 128)
	MidnightBlue = New(25, 25, 112)

	// Brown colors

	Cornsilk       = New(255, 248, 220)
	BlanchedAlmond = New(255, 235, 205)
	Bisque         = New(255, 228, 196)
	NavajoWhite    = New(255, 222, 173)
	Wheat          = New(245, 222, 179)
	BurlyWood      = New(222, 184, 135)
	Tan            = New(210, 180, 140)
	RosyBrown      = New(188, 143, 143)
	SandyBrown     = New(244, 164, 96)
	Goldenrod      = New(218, 165, 32)
	DarkGoldenrod  = New(184, 134, 11)
	Peru           = New(205, 133, 63)
	Chocolate      = New(210, 105, 30)
	SaddleBrown    = New(139, 69, 19)
	Sienna         = New(160, 82, 45)
	Brown          = New(165, 42, 42)
	Maroon         = New(128, 0, 0)

	// White colors

	White         = New(255, 255, 255)
	Snow          = New(255, 250, 250)
	HoneyDew      = New(240, 255, 240)
	MintCream     = New(245, 255, 250)
	Azure         = New(240, 255, 255)
	AliceBlue     = New(240, 248, 255)
	GhostWhite    = New(248, 248, 255)
	WhiteSmoke    = New(245, 245, 245)
	SeaShell      = New(255, 245, 238)
	Beige         = New(245, 245, 220)
	FloralWhite   = New(253, 245, 230)
	Ivory         = New(255, 250, 240)
	AntiqueWhite  = New(250, 235, 215)
	Linen         = New(250, 240, 230)
	LavenderBlush = New(255, 240, 245)
	MistyRose     = New(255, 228, 225)

	// Black colors

	Gainsboro      = New(220, 220, 220)
	LightGray      = New(211, 211, 211)
	Silver         = New(192, 192, 192)
	DarkGray       = New(169, 169, 169)
	Gray           = New(128, 128, 128)
	DimGray        = New(105, 105, 105)
	LightSlateGray = New(119, 136, 153)
	SlateGray      = New(112, 128, 144)
	DarkSlateGray  = New(47, 79, 79)
	Black          = New(0, 0, 0)
)
