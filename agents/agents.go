// DO NOT EDIT!
// Code generated by mage agents command
// source caniuse/data.json

// Package agents exposes details about all common web browsers. This uses data
// from caniuse-db.
package agents

// Agent contains details about a web browser.
type Agent struct {
	Name        string
	Browser     string
	Abbr        string
	Prefix      string
	Type        string
	UsageGlobal map[string]float64
	Versions    []string
}

// AgentsMap is mapping of browser name to browser details.
var AgentsMap = map[string]Agent{
	"and_chr": ChromeForAndroid,
	"and_ff":  FirefoxForAndroid,
	"and_qq":  QQForAndroid,
	"and_uc":  UCForAndroid,
	"android": Android,
	"baidu":   Baidu,
	"bb":      BlackBerry,
	"chrome":  Chrome,
	"edge":    Edge,
	"firefox": Firefox,
	"ie":      InternetExplorer,
	"ie_mob":  InternetExplorerMobile,
	"ios_saf": IOSSafari,
	"op_mini": OperaMini,
	"op_mob":  OperaMobile,
	"opera":   Opera,
	"safari":  Safari,
	"samsung": Samsung,
}
var (

	// ChromeForAndroid is Chrome for Android browser
	ChromeForAndroid = Agent{
		Name:        "and_chr",
		Browser:     "Chrome for Android",
		Abbr:        "Chr/And.",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"64": 30.4129},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "64", "", "", ""},
	}

	// FirefoxForAndroid is Firefox for Android browser
	FirefoxForAndroid = Agent{
		Name:        "and_ff",
		Browser:     "Firefox for Android",
		Abbr:        "FF/And.",
		Prefix:      "moz",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"57": 0.25047},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "57", "", "", ""},
	}

	// QQForAndroid is QQ Browser browser
	QQForAndroid = Agent{
		Name:        "and_qq",
		Browser:     "QQ Browser",
		Abbr:        "QQ",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"1.2": 0},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "1.2", "", "", ""},
	}

	// UCForAndroid is UC Browser for Android browser
	UCForAndroid = Agent{
		Name:        "and_uc",
		Browser:     "UC Browser for Android",
		Abbr:        "UC",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"11.8": 7.50853},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "11.8", "", "", ""},
	}

	// Android is Android Browser browser
	Android = Agent{
		Name:        "android",
		Browser:     "Android Browser",
		Abbr:        "And.",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"2.1": 0, "2.2": 0, "2.3": 0, "3": 0, "4": 0, "4.2-4.3": 0.205212, "62": 0, "4.1": 0.067427, "4.4": 0.578992, "4.4.3-4.4.4": 0.35619},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "2.1", "2.2", "2.3", "3", "4", "4.1", "4.2-4.3", "4.4", "4.4.3-4.4.4", "62", "", "", ""},
	}

	// Baidu is Baidu Browser browser
	Baidu = Agent{
		Name:        "baidu",
		Browser:     "Baidu Browser",
		Abbr:        "baidu",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"7.12": 0},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "7.12", "", "", ""},
	}

	// BlackBerry is Blackberry Browser browser
	BlackBerry = Agent{
		Name:        "bb",
		Browser:     "Blackberry Browser",
		Abbr:        "BB",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"7": 0.0133584, "10": 0.0534336},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "7", "10", "", "", ""},
	}

	// Chrome is Chrome browser
	Chrome = Agent{
		Name:        "chrome",
		Browser:     "Chrome",
		Abbr:        "Chr.",
		Prefix:      "webkit",
		Type:        "desktop",
		UsageGlobal: map[string]float64{"38": 0.039897, "42": 0.022165, "68": 0, "24": 0.013299, "28": 0.008866, "67": 0.022165, "5": 0.004879, "10": 0.004534, "55": 0.359073, "64": 15.2229, "13": 0.004879, "21": 0.004433, "32": 0.008866, "45": 0.035464, "48": 0.124124, "49": 0.820105, "53": 0.026598, "54": 0.084227, "63": 1.29887, "25": 0.008866, "27": 0.008866, "41": 0.022165, "57": 0.097526, "66": 0.053196, "7": 0.005591, "22": 0.017732, "35": 0.013299, "58": 0.164021, "59": 0.115258, "6": 0.004879, "47": 0.04433, "61": 0.172887, "11": 0.008866, "23": 0.008786, "50": 0.026598, "17": 0.004393, "20": 0.004393, "33": 0.013299, "14": 0.004706, "34": 0.022165, "39": 0.017732, "56": 0.115258, "62": 0.203918, "16": 0.004393, "30": 0.017732, "36": 0.017732, "44": 0.013299, "65": 8.28528, "26": 0.008866, "43": 0.062062, "37": 0.017732, "51": 0.053196, "52": 0.053196, "40": 0.017732, "60": 0.159588, "4": 0.004706, "29": 0.363506, "8": 0.005591, "9": 0.005591, "12": 0.004283, "15": 0.009154, "18": 0.017732, "19": 0.004418, "31": 0.026598, "46": 0.022165},
		Versions:    []string{"4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68"},
	}

	// Edge is Edge browser
	Edge = Agent{
		Name:        "edge",
		Browser:     "Edge",
		Abbr:        "Edge",
		Prefix:      "ms",
		Type:        "desktop",
		UsageGlobal: map[string]float64{"17": 0.022165, "18": 0, "12": 0.026598, "13": 0.035464, "14": 0.097526, "15": 0.128557, "16": 1.57371},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "12", "13", "14", "15", "16", "17", "18", ""},
	}

	// Firefox is Firefox browser
	Firefox = Agent{
		Name:        "firefox",
		Browser:     "Firefox",
		Abbr:        "FF",
		Prefix:      "moz",
		Type:        "desktop",
		UsageGlobal: map[string]float64{"19": 0.004443, "25": 0.008836, "30": 0.004433, "51": 0.062062, "3.5": 0.008786, "17": 0.004349, "48": 0.137423, "7": 0.005725, "12": 0.004418, "46": 0.008866, "58": 1.98155, "11": 0.004433, "31": 0.008866, "42": 0.008866, "14": 0.00453, "15": 0.004433, "13": 0.004486, "36": 0.013299, "37": 0.008866, "49": 0.022165, "56": 0.119691, "60": 0.048763, "5": 0.004879, "6": 0.020136, "23": 0.004433, "3": 0.004433, "24": 0.008786, "20": 0.004283, "4": 0.013299, "9": 0.00533, "21": 0.004418, "34": 0.008866, "47": 0.053196, "52": 0.4433, "59": 1.57371, "2": 0.004433, "26": 0.004393, "27": 0.004393, "32": 0.004433, "53": 0.026598, "8": 0.008866, "40": 0.013299, "41": 0.013299, "45": 0.035464, "57": 0.106392, "18": 0.004393, "50": 0.035464, "54": 0.035464, "3.6": 0.008866, "22": 0.004393, "33": 0.004433, "55": 0.04433, "61": 0, "10": 0.004283, "29": 0.008866, "43": 0.031031, "16": 0.004433, "35": 0.008866, "38": 0.039897, "39": 0.008866, "44": 0.075361, "28": 0.004418},
		Versions:    []string{"", "", "2", "3", "3.5", "3.6", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", ""},
	}

	// InternetExplorer is IE browser
	InternetExplorer = Agent{
		Name:        "ie",
		Browser:     "IE",
		Abbr:        "IE",
		Prefix:      "ms",
		Type:        "desktop",
		UsageGlobal: map[string]float64{"6": 0.00911931, "7": 0.013679, "8": 0.173267, "9": 0.118551, "10": 0.113991, "11": 2.76315, "5.5": 0.009298},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "5.5", "6", "7", "8", "9", "10", "11", "", "", ""},
	}

	// InternetExplorerMobile is IE Mobile browser
	InternetExplorerMobile = Agent{
		Name:        "ie_mob",
		Browser:     "IE Mobile",
		Abbr:        "IE.Mob",
		Prefix:      "ms",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"10": 0.033396, "11": 0.178112},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "10", "11", "", "", ""},
	}

	// IOSSafari is iOS Safari browser
	IOSSafari = Agent{
		Name:        "ios_saf",
		Browser:     "iOS Safari",
		Abbr:        "iOS",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"11.0-11.2": 8.61391, "4.2-4.3": 0.00107165, "10.0-10.2": 0.428659, "11.3": 0, "8": 0.0150031, "7.0-7.1": 0.0525108, "9.3": 0.429731, "5.0-5.1": 0.0192897, "4.0-4.1": 0, "6.0-6.1": 0.0150031, "8.1-8.4": 0.079302, "9.0-9.2": 0.0600123, "10.3": 0.993418, "3.2": 0.00107165},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "3.2", "4.0-4.1", "4.2-4.3", "5.0-5.1", "6.0-6.1", "7.0-7.1", "8", "8.1-8.4", "9.0-9.2", "9.3", "10.0-10.2", "10.3", "11.0-11.2", "11.3", "", "", ""},
	}

	// OperaMini is Opera Mini browser
	OperaMini = Agent{
		Name:        "op_mini",
		Browser:     "Opera Mini",
		Abbr:        "O.Mini",
		Prefix:      "o",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"all": 2.66111},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "all", "", "", ""},
	}

	// OperaMobile is Opera Mobile browser
	OperaMobile = Agent{
		Name:        "op_mob",
		Browser:     "Opera Mobile",
		Abbr:        "O.Mob",
		Prefix:      "o",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"11.5": 0, "12": 0, "12.1": 0, "37": 0.0111391, "10": 0, "11": 0, "11.1": 0},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "10", "11", "11.1", "11.5", "12", "12.1", "37", "", "", ""},
	}

	// Opera is Opera browser
	Opera = Agent{
		Name:        "opera",
		Browser:     "Opera",
		Abbr:        "Op.",
		Prefix:      "webkit",
		Type:        "desktop",
		UsageGlobal: map[string]float64{"46": 0.004433, "9.5-9.6": 0.00685, "20": 0.004879, "21": 0.006597, "42": 0.004418, "35": 0.004879, "38": 0.004367, "49": 0.008866, "9": 0.0082, "12.1": 0.04433, "28": 0.008698, "32": 0.005152, "16": 0.00685, "18": 0.005014, "30": 0.004879, "31": 0.004433, "10.0-10.1": 0, "10.6": 0.004706, "11.6": 0.008786, "15": 0.00685, "41": 0.004227, "47": 0.004433, "17": 0.00685, "25": 0.006015, "34": 0.009758, "36": 0.039897, "11.1": 0.006229, "45": 0.008866, "51": 0.740311, "52": 0.026598, "11.5": 0.004879, "27": 0.004393, "39": 0.004534, "43": 0.008668, "40": 0.004367, "44": 0.004227, "24": 0.006702, "26": 0.005595, "29": 0.004879, "37": 0.004283, "22": 0.006597, "23": 0.013434, "33": 0.005014, "48": 0.004433, "10.5": 0.008392, "11": 0.016581, "12": 0.008786, "19": 0.006015, "50": 0.048763},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "9", "9.5-9.6", "10.0-10.1", "10.5", "10.6", "11", "11.1", "11.5", "11.6", "12", "12.1", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", ""},
	}

	// Safari is Safari browser
	Safari = Agent{
		Name:        "safari",
		Browser:     "Safari",
		Abbr:        "Saf.",
		Prefix:      "webkit",
		Type:        "desktop",
		UsageGlobal: map[string]float64{"6": 0.004349, "8": 0.031031, "TP": 0, "4": 0, "5": 0.013299, "7": 0.008866, "9": 0.031031, "11.1": 0.022165, "3.2": 0.008692, "5.1": 0.070928, "7.1": 0.004283, "9.1": 0.164021, "10": 0.097526, "3.1": 0, "6.1": 0.013299, "10.1": 0.319176, "11": 1.56928},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "3.1", "3.2", "4", "5", "5.1", "6", "6.1", "7", "7.1", "8", "9", "9.1", "10", "10.1", "11", "11.1", "TP", "", ""},
	}

	// Samsung is Samsung Internet browser
	Samsung = Agent{
		Name:        "samsung",
		Browser:     "Samsung Internet",
		Abbr:        "SS",
		Prefix:      "webkit",
		Type:        "mobile",
		UsageGlobal: map[string]float64{"4": 0.943211, "5": 0.228029, "6.2": 1.71022},
		Versions:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "4", "5", "6.2", "", "", ""},
	}
)

// All is a helper function that returns a list of all agents.
func All() []Agent {
	return []Agent{ChromeForAndroid, FirefoxForAndroid, QQForAndroid, UCForAndroid, Android, Baidu, BlackBerry, Chrome, Edge, Firefox, InternetExplorer, InternetExplorerMobile, IOSSafari, OperaMini, OperaMobile, Opera, Safari, Samsung}
}