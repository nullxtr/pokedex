package pokeapi

type NamedURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
	ID                     int         `json:"id"`
	Name                   string      `json:"name"`
	BaseExperience         int         `json:"base_experience"`
	Height                 int         `json:"height"`
	IsDefault              bool        `json:"is_default"`
	Order                  int         `json:"order"`
	Weight                 int         `json:"weight"`
	Abilities              []Ability   `json:"abilities"`
	Forms                  []NamedURL  `json:"forms"`
	GameIndices            []GameIndex `json:"game_indices"`
	HeldItems              []HeldItem  `json:"held_items"`
	LocationAreaEncounters string      `json:"location_area_encounters"`
	Moves                  []Move      `json:"moves"`
	Species                NamedURL    `json:"species"`
	Sprites                Sprites     `json:"sprites"`
	Cries                  Cries       `json:"cries"`
	Stats                  []Stat      `json:"stats"`
	Types                  []TypeSlot  `json:"types"`
	PastTypes              []PastType  `json:"past_types"`
}

type Ability struct {
	IsHidden bool     `json:"is_hidden"`
	Slot     int      `json:"slot"`
	Ability  NamedURL `json:"ability"`
}

type GameIndex struct {
	GameIndex int      `json:"game_index"`
	Version   NamedURL `json:"version"`
}

type HeldItem struct {
	Item           NamedURL         `json:"item"`
	VersionDetails []VersionDetails `json:"version_details"`
}

type VersionDetails struct {
	Rarity  int      `json:"rarity"`
	Version NamedURL `json:"version"`
}

type Move struct {
	Move                NamedURL            `json:"move"`
	VersionGroupDetails []MoveVersionDetail `json:"version_group_details"`
}

type MoveVersionDetail struct {
	LevelLearnedAt  int      `json:"level_learned_at"`
	VersionGroup    NamedURL `json:"version_group"`
	MoveLearnMethod NamedURL `json:"move_learn_method"`
}

type Sprites struct {
	BackDefault      string       `json:"back_default"`
	BackFemale       interface{}  `json:"back_female"`
	BackShiny        string       `json:"back_shiny"`
	BackShinyFemale  interface{}  `json:"back_shiny_female"`
	FrontDefault     string       `json:"front_default"`
	FrontFemale      interface{}  `json:"front_female"`
	FrontShiny       string       `json:"front_shiny"`
	FrontShinyFemale interface{}  `json:"front_shiny_female"`
	Other            OtherSprites `json:"other"`
	Versions         Versions     `json:"versions"`
}

type OtherSprites struct {
	DreamWorld struct {
		FrontDefault string      `json:"front_default"`
		FrontFemale  interface{} `json:"front_female"`
	} `json:"dream_world"`
	Home struct {
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"home"`
	OfficialArtwork struct {
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"official-artwork"`
	Showdown struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"showdown"`
}

type Versions struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationII   GenerationII   `json:"generation-ii"`
	GenerationIII  GenerationIII  `json:"generation-iii"`
	GenerationIV   GenerationIV   `json:"generation-iv"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVI   GenerationVI   `json:"generation-vi"`
	GenerationVII  GenerationVII  `json:"generation-vii"`
	GenerationVIII GenerationVIII `json:"generation-viii"`
}

type GenerationI struct {
	RedBlue struct {
		BackDefault  string `json:"back_default"`
		BackGray     string `json:"back_gray"`
		FrontDefault string `json:"front_default"`
		FrontGray    string `json:"front_gray"`
	} `json:"red-blue"`
	Yellow struct {
		BackDefault  string `json:"back_default"`
		BackGray     string `json:"back_gray"`
		FrontDefault string `json:"front_default"`
		FrontGray    string `json:"front_gray"`
	} `json:"yellow"`
}

type GenerationII struct {
	Crystal struct {
		BackDefault  string `json:"back_default"`
		BackShiny    string `json:"back_shiny"`
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"crystal"`
	Gold struct {
		BackDefault  string `json:"back_default"`
		BackShiny    string `json:"back_shiny"`
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"gold"`
	Silver struct {
		BackDefault  string `json:"back_default"`
		BackShiny    string `json:"back_shiny"`
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"silver"`
}

type GenerationIII struct {
	Emerald struct {
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"emerald"`
	FireredLeafgreen struct {
		BackDefault  string `json:"back_default"`
		BackShiny    string `json:"back_shiny"`
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"firered-leafgreen"`
	RubySapphire struct {
		BackDefault  string `json:"back_default"`
		BackShiny    string `json:"back_shiny"`
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"ruby-sapphire"`
}

type GenerationIV struct {
	DiamondPearl struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"diamond-pearl"`
	HeartgoldSoulsilver struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"heartgold-soulsilver"`
	Platinum struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"platinum"`
}

type GenerationV struct {
	BlackWhite struct {
		Animated struct {
			BackDefault      string      `json:"back_default"`
			BackFemale       interface{} `json:"back_female"`
			BackShiny        string      `json:"back_shiny"`
			BackShinyFemale  interface{} `json:"back_shiny_female"`
			FrontDefault     string      `json:"front_default"`
			FrontFemale      interface{} `json:"front_female"`
			FrontShiny       string      `json:"front_shiny"`
			FrontShinyFemale interface{} `json:"front_shiny_female"`
		} `json:"animated"`
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"black-white"`
}

type GenerationVI struct {
	OmegarubyAlphasapphire struct {
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"omegaruby-alphasapphire"`
	XY struct {
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"x-y"`
}

type GenerationVII struct {
	Icons struct {
		FrontDefault string      `json:"front_default"`
		FrontFemale  interface{} `json:"front_female"`
	} `json:"icons"`
	UltraSunUltraMoon struct {
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"ultra-sun-ultra-moon"`
}

type GenerationVIII struct {
	Icons struct {
		FrontDefault string      `json:"front_default"`
		FrontFemale  interface{} `json:"front_female"`
	} `json:"icons"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Stat struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	Stat     NamedURL `json:"stat"`
}

type TypeSlot struct {
	Slot int      `json:"slot"`
	Type NamedURL `json:"type"`
}

type PastType struct {
	Generation NamedURL   `json:"generation"`
	Types      []TypeSlot `json:"types"`
}
