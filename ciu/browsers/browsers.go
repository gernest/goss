package browsers

type Key string

const (
	A Key = "A"
	B Key = "B"
	C Key = "C"
	D Key = "D"
	E Key = "E"
	F Key = "F"
	G Key = "G"
	H Key = "H"
	I Key = "I"
	J Key = "J"
	K Key = "K"
	L Key = "L"
	M Key = "M"
	N Key = "N"
	O Key = "O"
	P Key = "P"
	Q Key = "Q"
	R Key = "R"
)

func (k Key) String() string {
	switch k {
	case A:
		return "and_chr"
	case B:
		return "and_ff"
	case C:
		return "and_qq"
	case D:
		return "and_uc"
	case E:
		return "android"
	case F:
		return "baidu"
	case G:
		return "bb"
	case H:
		return "chrome"
	case I:
		return "edge"
	case J:
		return "firefox"
	case K:
		return "ie"
	case L:
		return "ie_mob"
	case M:
		return "ios_saf"
	case N:
		return "op_mini"
	case O:
		return "op_mob"
	case P:
		return "opera"
	case Q:
		return "safari"
	case R:
		return "samsung"
	default:
		return ""
	}
}

// New retuns a map of encoded browser keys to their values.
func New() map[string]string {
	return map[string]string{
		"A": "and_chr",
		"B": "and_ff",
		"C": "and_qq",
		"D": "and_uc",
		"E": "android",
		"F": "baidu",
		"G": "bb",
		"H": "chrome",
		"I": "edge",
		"J": "firefox",
		"K": "ie",
		"L": "ie_mob",
		"M": "ios_saf",
		"N": "op_mini",
		"O": "op_mob",
		"P": "opera",
		"Q": "safari",
		"R": "samsung",
	}
}