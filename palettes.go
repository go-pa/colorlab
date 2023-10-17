package colorlab

var (
	AllPalettes = []Palette{
		SolarizedDarkPalette,
		SolarizedLightPalette,
		SolarizedDarkHighContrastPalette,
		SolarizedLightHighContrastPalette,
		SelenizedBlackPalette,
		SelenizedDarkPalette,
		SelenizedLightPalette,
		SelenizedWhitePalette,
		GruvboxDarkPalette,
		GruvboxLightPalette,
		ZenburnPalette,
		MonokaiPalette,
	}

	SolarizedDarkPalette = Palette{
		Name:      "solarized-dark",
		Solarized: solarized,
		Accent1Pair: AccentPairGenerator{
			BlendBackgroundAmout:       default1BgFg.BlendBackgroundAmout,
			BlendForegroundAmout:       default1BgFg.BlendForegroundAmout,
			Gamma:                      default1BgFg.Gamma,
			ForegroundBlendFinder:      NamedColorFinder("base1"),
			BackgroundBlendFinder:      default1BgFg.BackgroundBlendFinder,
			MinimumLightnessDifference: 0.4,
		},
		Accent2Pair: AccentPairGenerator{
			BlendBackgroundAmout:       default2BgFg.BlendBackgroundAmout,
			BlendForegroundAmout:       default2BgFg.BlendForegroundAmout,
			Gamma:                      default2BgFg.Gamma,
			ForegroundBlendFinder:      NamedColorFinder("base1"),
			BackgroundBlendFinder:      default2BgFg.BackgroundBlendFinder,
			MinimumLightnessDifference: 0.4,
		},
	}
	SolarizedLightPalette = Palette{
		Name:        "solarized-light",
		Solarized:   solarized,
		Inverse:     true,
		Accent1Pair: default1BgFg,
		Accent2Pair: default2BgFg,
	}

	SolarizedDarkHighContrastPalette = Palette{
		Name:      "solarized-dark-high-contrast",
		Solarized: solarizedDarkHighContrast,
		Accent1Pair: AccentPairGenerator{
			BlendBackgroundAmout:       default1BgFg.BlendBackgroundAmout,
			BlendForegroundAmout:       default1BgFg.BlendForegroundAmout,
			Gamma:                      default1BgFg.Gamma,
			ForegroundBlendFinder:      NamedColorFinder("base1"),
			BackgroundBlendFinder:      default1BgFg.BackgroundBlendFinder,
			MinimumLightnessDifference: 0.4,
		},
		Accent2Pair: AccentPairGenerator{
			BlendBackgroundAmout:       default2BgFg.BlendBackgroundAmout,
			BlendForegroundAmout:       default2BgFg.BlendForegroundAmout,
			Gamma:                      default2BgFg.Gamma,
			ForegroundBlendFinder:      NamedColorFinder("base1"),
			BackgroundBlendFinder:      default2BgFg.BackgroundBlendFinder,
			MinimumLightnessDifference: 0.4,
		},
	}
	SolarizedLightHighContrastPalette = Palette{
		Name:        "solarized-light-high-contrast",
		Solarized:   solarizedLightHighContrast,
		Inverse:     true,
		Accent1Pair: default1BgFg,
		Accent2Pair: default2BgFg,
	}
	GruvboxDarkPalette = Palette{
		Name:        "gruvbox-dark",
		Solarized:   gruvboxDark,
		Accent1Pair: default1BgFg,
		Accent2Pair: default2BgFg,
	}
	GruvboxLightPalette = Palette{
		Name:        "gruvbox-light",
		Solarized:   gruvboxLight,
		Inverse:     true,
		Accent1Pair: default1BgFg,
		Accent2Pair: default2BgFg,
	}
	ZenburnPalette = Palette{
		Name:        "zenburn",
		Solarized:   zenburn,
		Accent1Pair: default1BgFg,
		Accent2Pair: default2BgFg,
	}
	MonokaiPalette = Palette{
		Name:        "monokai",
		Solarized:   monokai,
		Accent1Pair: default1BgFg,
		Accent2Pair: default2BgFg,
	}
	SelenizedDarkPalette = Palette{
		Name:        "selenized-dark",
		Solarized:   selenizedDark,
		Accent1Pair: default1BgFg,
		Accent2Pair: default2BgFg,
	}
	SelenizedBlackPalette = Palette{
		Name:        "selenized-black",
		Solarized:   selenizedBlack,
		Accent1Pair: default1BgFg,
		Accent2Pair: default2BgFg,
	}
	SelenizedLightPalette = Palette{
		Name:        "selenized-light",
		Solarized:   selenizedLight,
		Accent1Pair: default1BgFg,
		Accent2Pair: default2BgFg,
	}
	SelenizedWhitePalette = Palette{
		Name:        "selenized-white",
		Solarized:   selenizedWhite,
		Accent1Pair: default1BgFg,
		Accent2Pair: default2BgFg,
	}
)

// base 16 color arrangements
var (
	// The original solarized color palette
	solarized = Solarized{
		Base: Base{
			Base03: SolarizedBase03,
			Base02: SolarizedBase02,
			Base01: SolarizedBase01,
			Base00: SolarizedBase00,
			Base0:  SolarizedBase0,
			Base1:  SolarizedBase1,
			Base2:  SolarizedBase2,
			Base3:  SolarizedBase3,
		},
		Accents: Accents{
			Blue:    SolarizedBlue,
			Cyan:    SolarizedCyan,
			Green:   SolarizedGreen,
			Magenta: SolarizedMagenta,
			Orange:  SolarizedOrange,
			Red:     SolarizedRed,
			Violet:  SolarizedViolet,
			Yellow:  SolarizedYellow,
		},
	}
	solarizedDarkHighContrast = Solarized{
		Base:    solarized.Base.Clone().ChangeLightness(0.04, -0.02),
		Accents: solarized.Accents.Clone().ChangeLightness(0.05),
	}
	solarizedLightHighContrast = Solarized{
		Base:    solarized.Base.Clone().ChangeLightness(0.02, -0.05),
		Accents: solarized.Accents.Clone().ChangeLightness(-0.05),
	}

	// the current -d colors in the emacs theme
	oldDarkAccents = Accents{
		Blue:    "#00629D",
		Cyan:    "#00736F",
		Green:   "#546E00",
		Magenta: "#93115C",
		Orange:  "#8B2C02",
		Red:     "#990A1B",
		Violet:  "#3F4D91",
		Yellow:  "#7B6000",
	}

	// the current -l colors in the emacs theme
	oldLightAccents = Accents{
		Blue:    "#69B7F0",
		Cyan:    "#69CABF",
		Green:   "#B4C342",
		Magenta: "#F771AC",
		Orange:  "#F2804F",
		Red:     "#FF6E64",
		Violet:  "#9EA0E5",
		Yellow:  "#DEB542",
	}

	gruvboxDark = Solarized{
		Base: Base{
			Base03: GruvboxDark0,
			Base02: GruvboxDark0Soft,
			Base01: GruvboxDark4,
			Base00: GruvboxDark0,
			Base0:  GruvboxLight4,
			Base1:  GruvboxLight3,
			Base2:  GruvboxLight4,
			Base3:  GruvboxLight0,
		},
		Accents: Accents{
			Blue:    GruvboxBlue,
			Cyan:    GruvboxAqua,
			Green:   GruvboxGreen,
			Magenta: GruvboxBrightPurple,
			Orange:  GruvboxOrange,
			Red:     GruvboxBrightRed,
			Violet:  GruvboxPurple,
			Yellow:  GruvboxYellow,
		},
	}

	// note: not inversed here
	gruvboxLight = Solarized{
		Base: Base{
			Base03: GruvboxDark0,
			Base02: GruvboxDark0Soft,
			Base01: GruvboxDark3,
			Base00: GruvboxDark4,
			Base0:  GruvboxDark1,
			Base1:  GruvboxLight4,
			Base2:  GruvboxLight1,
			Base3:  GruvboxLight0,
		},
		Accents: Accents{
			Blue:    GruvboxDarkBlue,
			Cyan:    GruvboxAqua,
			Green:   GruvboxGreen,
			Magenta: GruvboxBrightPurple,
			Orange:  GruvboxDarkOrange,
			Red:     GruvboxDarkRed,
			Violet:  GruvboxDarkPurple,
			Yellow:  GruvboxDarkYellow,
		},
	}

	zenburn = Solarized{
		Base: Base{
			Base03: ZenburnBg,
			Base02: ZenburnBgP1,
			Base01: HexColor(ZenburnFgM1).Blend(ZenburnFg, 0.3),
			Base00: ZenburnBgP3,
			Base0:  ZenburnFg,
			Base1:  ZenburnFgP1,
			Base2:  HexColor(ZenburnFgP1).Blend(ZenburnFgP2, 0.5),
			Base3:  ZenburnFgP2,
		},
		Accents: Accents{
			Blue:    ZenburnBlue,
			Cyan:    ZenburnCyan,
			Green:   ZenburnGreen,
			Magenta: ZenburnMagenta,
			Orange:  ZenburnOrange,
			Red:     ZenburnRed,
			Violet:  HexColor(ZenburnBlue).Blend(ZenburnMagenta, 0.5),
			Yellow:  ZenburnYellow,
		},
	}

	monokai = Solarized{
		Base: Base{
			Base03: Monokai03,
			Base02: Monokai02,
			Base01: Monokai01,
			Base00: Monokai00,
			Base0:  Monokai0,
			Base1:  Monokai0,
			Base2:  Monokai0,
			Base3:  Monokai0,
		},
		Accents: Accents{
			Blue:    MonokaiBlue,
			Cyan:    MonokaiCyan,
			Green:   MonokaiGreen,
			Magenta: MonokaiMagenta,
			Orange:  MonokaiOrange,
			Red:     MonokaiRed,
			Violet:  MonokaiViolet,
			Yellow:  MonokaiYellow,
		},
	}
	selenizedDark = Solarized{
		Base: Base{
			Base03: "#103c48",
			Base02: "#184956",
			Base01: "#72898f",
			Base00: "#103c48",
			Base0:  "#adbcbc",
			Base1:  "#cad8d9",
			Base2:  "#ece3cc",
			Base3:  "#fbf3db",
		},
		Accents: Accents{
			Blue:    "#4695f7",
			Cyan:    "#41c7b9",
			Green:   "#75b938",
			Magenta: "#f275be",
			Orange:  "#ed8649",
			Red:     "#fa5750",
			Violet:  "#af88eb",
			Yellow:  "#dbb32d",
		},
	}
	selenizedBlack = Solarized{
		Base: Base{
			Base03: "#181818",
			Base02: "#252525",
			Base01: "#777777",
			Base00: "#181818",
			Base0:  "#b9b9b9",
			Base1:  "#dedede",
			Base2:  "#53676d",
			Base3:  "#3a4d53",
		},
		Accents: Accents{
			Blue:    "#368aeb",
			Cyan:    "#3fc5b7",
			Green:   "#70b433",
			Magenta: "#eb6eb7",
			Orange:  "#e67f43",
			Red:     "#ed4a46",
			Violet:  "#a580e2",
			Yellow:  "#dbb32d",
		},
	}
	selenizedLight = Solarized{
		Base: Base{
			Base03: "#fbf3db",
			Base02: "#ece3cc",
			Base01: "#909995",
			Base00: "#fbf3db",
			Base0:  "#53676d",
			Base1:  "#3a4d53",
			Base2:  "#adbcbc",
			Base3:  "#cad8d9",
		},
		Accents: Accents{
			Blue:    "#0072d4",
			Cyan:    "#009c8f",
			Green:   "#489100",
			Magenta: "#ca4898",
			Orange:  "#c25d1e",
			Red:     "#d2212d",
			Violet:  "#8762c6",
			Yellow:  "#ad8900",
		},
	}
	selenizedWhite = Solarized{
		Base: Base{
			Base03: "#ffffff",
			Base02: "#ebebeb",
			Base01: "#878787",
			Base00: "#ffffff",
			Base0:  "#474747",
			Base1:  "#282828",
			Base2:  "#b9b9b9",
			Base3:  "#dedede",
		},
		Accents: Accents{
			Blue:    "#0064e4",
			Cyan:    "#00ad9c",
			Green:   "#1d9700",
			Magenta: "#dd0f9d",
			Orange:  "#d04a00",
			Red:     "#d6000c",
			Violet:  "#7f51d6",
			Yellow:  "#c49700",
		},
	}
)

var (
	default1BgFg = AccentPairGenerator{
		BlendBackgroundAmout:       0.85,
		BlendForegroundAmout:       0.3,
		Gamma:                      0.01,
		MinimumLightnessDifference: 0.35,
		ForegroundBlendFinder:      ExtremeColorFgFinder,
		BackgroundBlendFinder:      ExtremeColorBgFinder,
	}
	default2BgFg = AccentPairGenerator{
		BlendBackgroundAmout:       0.6,
		BlendForegroundAmout:       0.45,
		Gamma:                      0.04,
		MinimumLightnessDifference: 0.35,
		ForegroundBlendFinder:      ExtremeColorFgFinder,
		BackgroundBlendFinder:      ExtremeColorBgFinder,
	}
)
