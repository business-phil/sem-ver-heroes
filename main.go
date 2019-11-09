package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type CharacterInput struct {
	Name          string        `json:"name"`
	Version       string        `json:"version"`
	Race          string        `json:"race"`
	Class         string        `json:"class"`
	Background    string        `json:"background"`
	Alignment     string        `json:"alignment"`
	HitPoints     int           `json:"hitPoints"`
	Speed         int           `json:"speed"`
	Abilities     Abilities     `json:"abilities"`
	Proficiencies Proficiencies `json:"proficiencies"`
	Items         Items         `json:"items"`
	Spells        []Spell       `json:"spells"`
}

// TODO: Consider removing HitPoints and Speed, which are only used
// to generate other stats
type CharacterOutput struct {
	Name              string        `json:"name"`
	Version           string        `json:"version"`
	Level             int           `json:"level"`
	Race              string        `json:"race"`
	Class             string        `json:"class"`
	Background        string        `json:"background"`
	Alignment         string        `json:"alignment"`
	HitPoints         int           `json:"hitPoints"`
	Speed             int           `json:"speed"`
	Proficiency       int           `json:"proficiency"`
	ArmorClass        int           `json:"armorClass"`
	Initiative        int           `json:"initiative"`
	PassivePerception int           `json:"passivePerception"`
	Magic             Magic         `json:"magic"`
	Abilities         Abilities     `json:"abilities"`
	AbilityModifiers  Abilities     `json:"abilityModifiers"`
	SavingThrows      Abilities     `json:"savingThrows"`
	Skills            Skills        `json:"skills"`
	Proficiencies     Proficiencies `json:"proficiencies"`
	Items             Items         `json:"items"`
	Spells            []Spell       `json:"spells"`
}

// TODO: consider int type alternatives for small values
type Abilities struct {
	Stength      int `json:"stength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
}

// TODO: replace strings with enums
type Proficiencies struct {
	Skills       []string `json:"skills"`
	SavingThrows []string `json:"savingThrows"`
	Resistances  []string `json:"resistances"`
	Armor        []string `json:"armor"`
	Weapons      []string `json:"weapons"`
	Languages    []string `json:"languages"`
}

type Items struct {
	Currency  Currency    `json:"currency"`
	Equipment []Equipment `json:"equipment"`
	Weapons   []Weapon    `json:"weapons"`
	Gear      []Gear      `json:"gear"`
}

type Currency struct {
	Gold int `json:"gold"`
}

type Equipment struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Rarity     string `json:"rarity"`
	IsEquipped bool   `json:"isEquipped"`
	ArmorClass int    `json:"armorClass"`
}

// TODO: add type for dice value (ex: "1d8")
type Weapon struct {
	Name       string   `json:"name"`
	AttackType string   `json:"attackType"`
	Rarity     string   `json:"rarity"`
	IsEquipped bool     `json:"isEquipped"`
	BaseDamage string   `json:"baseDamage"`
	DamageType string   `json:"damageType"`
	Properties []string `json:"properties"`
}

type Gear struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Spell struct {
	Name          string `json:"name"`
	Level         int    `json:"level"`
	Type          string `json:"type"`
	Range         int    `json:"range"`
	CastingTime   int    `json:"castingTime"`
	Concentration bool   `json:"concentration"`
}

// TODO: Map class to spell abilities to calculate Magic stats
type Magic struct {
	Modifier int `json:"modifier"`
	Attack   int `json:"attack"`
	SaveDC   int `json:"saveDC"`
}

type Attack struct {
	WeaponName string `json:"weaponName"`
	AttackType string `json:"attackType"`
	Range      int    `json:"range"`
	Hit        int    `json:"hit"`
	Damage     Damage `json:"damage"`
}

type Damage struct {
	Base     string `json:"base"`
	Modifier int    `json:"modifier"`
	Type     string `json:"type"`
}

type Skills struct {
	Acrobatics     int `json:"acrobatics"`
	AnimalHandling int `json:"animalHandling"`
	Arcana         int `json:"arcana"`
	Athletics      int `json:"athletics"`
	Deception      int `json:"deception"`
	History        int `json:"history"`
	Insight        int `json:"insight"`
	Intimidation   int `json:"intimidation"`
	Investigation  int `json:"investigation"`
	Medicine       int `json:"medicine"`
	Nature         int `json:"nature"`
	Perception     int `json:"perception"`
	Performance    int `json:"performance"`
	Persuasion     int `json:"persuasion"`
	Religion       int `json:"religion"`
	SleightOfHand  int `json:"sleightOfHand"`
	Stealth        int `json:"stealth"`
	Survival       int `json:"survival"`
}

func calculateAbilityModifier(ability int) int {
	normalizedAbility := ability - 10
	if normalizedAbility < 0 {
		normalizedAbility = normalizedAbility - 1
	}
	return int(normalizedAbility / 2)
}

func (character *CharacterOutput) GenerateAbilityModifiers() {
	character.AbilityModifiers.Stength = calculateAbilityModifier(character.Abilities.Stength)
	character.AbilityModifiers.Dexterity = calculateAbilityModifier(character.Abilities.Dexterity)
	character.AbilityModifiers.Constitution = calculateAbilityModifier(character.Abilities.Constitution)
	character.AbilityModifiers.Intelligence = calculateAbilityModifier(character.Abilities.Intelligence)
	character.AbilityModifiers.Wisdom = calculateAbilityModifier(character.Abilities.Wisdom)
	character.AbilityModifiers.Charisma = calculateAbilityModifier(character.Abilities.Charisma)
}

func (character *CharacterOutput) GenerateMagicStats() {
	switch character.Class {
	case "Bard", "Paladin", "Warlock", "Sorcerer":
		character.Magic.Modifier = character.AbilityModifiers.Charisma
	case "Wizard":
		character.Magic.Modifier = character.AbilityModifiers.Intelligence
	case "Cleric", "Druid", "Ranger":
		character.Magic.Modifier = character.AbilityModifiers.Wisdom
	}

	character.Magic.Attack = character.Proficiency + character.Magic.Modifier
	character.Magic.SaveDC = 8 + character.Magic.Attack
}

func main() {
	fmt.Println("Hello SemVer Hero!")

	// TODO: Pass in filename as an argument
	inputFile, err := os.Open("sample-input.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	inputBytes, err := ioutil.ReadAll(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: Either consolidate Input/Output types, or use Input to populate Output
	var character *CharacterOutput
	json.Unmarshal(inputBytes, &character)

	fmt.Printf("Successfully parsed input for %s:\n", character.Name)
	fmt.Println(character)

	// Calculate Level
	versionFloat, err := strconv.ParseFloat(character.Version, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	character.Level = int(versionFloat)

	// Calculate Proficiency
	proficiencyByLevel := [21]int{0, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6}
	character.Proficiency = proficiencyByLevel[character.Level]

	// Calculate Ability Modifiers
	character.GenerateAbilityModifiers()

	// Calculate Initiaitve
	character.Initiative = character.AbilityModifiers.Dexterity

	// Calculate Magic Stats
	character.GenerateMagicStats()

	fmt.Printf("Successfully generated stats for %s:\n", character.Name)
	fmt.Println(character)
}
