package pokeapi

type RespLocationArea struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon NamedAPIResource       `json:"pokemon"`
	Version []PokemonVersionDetail `json:"version"`
}

type PokemonVersionDetail struct {
	Version          NamedAPIResource  `json:"version"`
	MaxChance        int               `json:"max_chance"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource `json:"encounter_method_id"`
	VersionDetails  []VersionDetail  `json:"version_details"`
}

type EncounterDetail struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	Chance          int                `json:"chance"`
	Method          NamedAPIResource   `json:"method"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type VersionDetail struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}
