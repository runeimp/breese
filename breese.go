package breese

const Version = "0.1.0-alpha"

type CodePoint struct {
	Character    string
	ITUR         bool
	Prosign      bool
	QCode        bool
	UTFMorse     bool
	ValueASCII   string
	ValueUnicode string
}

var (
	MorseCodeInternational = make(map[string]CodePoint)
	MorseCodeContinental   = make(map[string]CodePoint)
)

func init() {
	/*
	 * International (ITU)
	 */

	// Letters
	MorseCodeInternational["a"] = CodePoint{
		Character:    "a",
		ITUR:         true,
		ValueASCII:   ".-", // Period + Hyphen
		ValueUnicode: "·-", // Bullet + Endash
	}
	MorseCodeInternational["b"] = CodePoint{
		Character:    "b",
		ITUR:         true,
		ValueASCII:   "-...",
		ValueUnicode: "-···",
	}
	MorseCodeInternational["c"] = CodePoint{
		Character:    "c",
		ITUR:         true,
		ValueASCII:   "-.-.",
		ValueUnicode: "-·-·",
	}
	MorseCodeInternational["d"] = CodePoint{
		Character:    "d",
		ITUR:         true,
		ValueASCII:   "-..",
		ValueUnicode: "-··",
	}
	MorseCodeInternational["e"] = CodePoint{
		Character:    "e",
		ITUR:         true,
		ValueASCII:   ".",
		ValueUnicode: "·",
	}
	MorseCodeInternational["f"] = CodePoint{
		Character:    "f",
		ITUR:         true,
		ValueASCII:   "..-.",
		ValueUnicode: "··-·",
	}
	MorseCodeInternational["g"] = CodePoint{
		Character:    "g",
		ITUR:         true,
		ValueASCII:   "--.",
		ValueUnicode: "--·",
	}
	MorseCodeInternational["h"] = CodePoint{
		Character:    "h",
		ITUR:         true,
		ValueASCII:   "....",
		ValueUnicode: "····",
	}
	MorseCodeInternational["i"] = CodePoint{
		Character:    "i",
		ITUR:         true,
		ValueASCII:   "..",
		ValueUnicode: "··",
	}
	MorseCodeInternational["j"] = CodePoint{
		Character:    "j",
		ITUR:         true,
		ValueASCII:   ".---",
		ValueUnicode: "·---",
	}
	MorseCodeInternational["k"] = CodePoint{
		Character:    "k",
		ITUR:         true,
		ValueASCII:   "-.-",
		ValueUnicode: "-·-",
	}
	MorseCodeInternational["l"] = CodePoint{
		Character:    "l",
		ITUR:         true,
		ValueASCII:   ".-..",
		ValueUnicode: "·-··",
	}
	MorseCodeInternational["m"] = CodePoint{
		Character:    "m",
		ITUR:         true,
		ValueASCII:   "--",
		ValueUnicode: "--",
	}
	MorseCodeInternational["n"] = CodePoint{
		Character:    "n",
		ITUR:         true,
		ValueASCII:   "-.",
		ValueUnicode: "-·",
	}
	MorseCodeInternational["o"] = CodePoint{
		Character:    "o",
		ITUR:         true,
		ValueASCII:   "---",
		ValueUnicode: "---",
	}
	MorseCodeInternational["p"] = CodePoint{
		Character:    "p",
		ITUR:         true,
		ValueASCII:   ".--.",
		ValueUnicode: "·--·",
	}
	MorseCodeInternational["q"] = CodePoint{
		Character:    "q",
		ITUR:         true,
		ValueASCII:   "--.-",
		ValueUnicode: "--·-",
	}
	MorseCodeInternational["r"] = CodePoint{
		Character:    "r",
		ITUR:         true,
		ValueASCII:   ".-.",
		ValueUnicode: "·-·",
	}
	MorseCodeInternational["s"] = CodePoint{
		Character:    "s",
		ITUR:         true,
		ValueASCII:   "...",
		ValueUnicode: "···",
	}
	MorseCodeInternational["t"] = CodePoint{
		Character:    "t",
		ITUR:         true,
		ValueASCII:   "-",
		ValueUnicode: "-",
	}
	MorseCodeInternational["u"] = CodePoint{
		Character:    "u",
		ITUR:         true,
		ValueASCII:   "..-",
		ValueUnicode: "··-",
	}
	MorseCodeInternational["v"] = CodePoint{
		Character:    "v",
		ITUR:         true,
		ValueASCII:   "...-",
		ValueUnicode: "···-",
	}
	MorseCodeInternational["w"] = CodePoint{
		Character:    "w",
		ITUR:         true,
		ValueASCII:   ".--",
		ValueUnicode: "·--",
	}
	MorseCodeInternational["x"] = CodePoint{
		Character:    "x",
		ITUR:         true,
		ValueASCII:   "-..-",
		ValueUnicode: "-··-",
	}
	MorseCodeInternational["y"] = CodePoint{
		Character:    "y",
		ITUR:         true,
		ValueASCII:   "-.--",
		ValueUnicode: "-·--",
	}
	MorseCodeInternational["z"] = CodePoint{
		Character:    "z",
		ITUR:         true,
		ValueASCII:   "--..",
		ValueUnicode: "--··",
	}

	// Digits
	MorseCodeInternational["0"] = CodePoint{
		Character:    "0",
		ITUR:         true,
		ValueASCII:   "-----",
		ValueUnicode: "-----",
	}
	MorseCodeInternational["1"] = CodePoint{
		Character:    "1",
		ITUR:         true,
		ValueASCII:   ".----",
		ValueUnicode: "·----",
	}
	MorseCodeInternational["2"] = CodePoint{
		Character:    "2",
		ITUR:         true,
		ValueASCII:   "..---",
		ValueUnicode: "··---",
	}
	MorseCodeInternational["3"] = CodePoint{
		Character:    "3",
		ITUR:         true,
		ValueASCII:   "...--",
		ValueUnicode: "···--",
	}
	MorseCodeInternational["4"] = CodePoint{
		Character:    "4",
		ITUR:         true,
		ValueASCII:   "....-",
		ValueUnicode: "····-",
	}
	MorseCodeInternational["5"] = CodePoint{
		Character:    "5",
		ITUR:         true,
		ValueASCII:   ".....",
		ValueUnicode: "·····",
	}
	MorseCodeInternational["6"] = CodePoint{
		Character:    "6",
		ITUR:         true,
		ValueASCII:   "-....",
		ValueUnicode: "-····",
	}
	MorseCodeInternational["7"] = CodePoint{
		Character:    "7",
		ITUR:         true,
		ValueASCII:   "--...",
		ValueUnicode: "--···",
	}
	MorseCodeInternational["8"] = CodePoint{
		Character:    "8",
		ITUR:         true,
		ValueASCII:   "---..",
		ValueUnicode: "---··",
	}
	MorseCodeInternational["9"] = CodePoint{
		Character:    "9",
		ITUR:         true,
		ValueASCII:   "----.",
		ValueUnicode: "----·",
	}

	// Extended Letters
	MorseCodeInternational["á"] = CodePoint{
		Character:    "á",
		ITUR:         true,
		ValueASCII:   ".--.-",
		ValueUnicode: "·--·-",
	}
	MorseCodeInternational["å"] = CodePoint{
		Character:    "å",
		ITUR:         true,
		ValueASCII:   ".--.-", // NOTE: Same as á? Need to look into this.
		ValueUnicode: "·--·-", // NOTE: Same as á? Need to look into this.
	}
	MorseCodeInternational["é"] = CodePoint{
		Character:    "é",
		ITUR:         true,
		ValueASCII:   "..-..",
		ValueUnicode: "··-··",
	}
	MorseCodeInternational["ñ"] = CodePoint{
		Character:    "ñ",
		ITUR:         true,
		ValueASCII:   "--.--",
		ValueUnicode: "--·--",
	}
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }
	// MorseCodeInternational["___"] = CodePoint{
	// 	Character:    "___",
	// 	ITUR:         true,
	// 	ValueASCII:   "- . .",
	// 	ValueUnicode: "- · ·",
	// }

	/*
	 * Continental (Friedrich Clemens Gerke "Hamburg alphabet" in 1848)
	 */
	MorseCodeContinental["ä"] = CodePoint{
		Character:    "ä",
		ITUR:         true,
		ValueASCII:   ". - . -",
		ValueUnicode: "· - · -",
	}
	MorseCodeContinental["Ch"] = CodePoint{
		Character:    "Ch",
		ITUR:         true,
		ValueASCII:   "- - - -",
		ValueUnicode: "- - - -",
	}
	MorseCodeContinental["ö"] = CodePoint{
		Character:    "ö",
		ITUR:         true,
		ValueASCII:   "- - - .",
		ValueUnicode: "- - - ·",
	}
	MorseCodeContinental["ü"] = CodePoint{
		Character:    "ü",
		ITUR:         true,
		ValueASCII:   ". . - -",
		ValueUnicode: "· · - -",
	}

	/*
	 * American (Original Morse Code)
	 */
}
