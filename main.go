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
	ProficiencyBonus  int           `json:"proficiencyBonus"`
	ArmorClass        int           `json:"armorClass"`
	Initiative        int           `json:"initiative"`
	PassivePerception int           `json:"passivePerception"`
	Magic             Magic         `json:"magic"`
	Abilities         Abilities     `json:"abilities"`
	AbilityModifiers  Abilities     `json:"abilityModifiers"`
	SavingThrows      Abilities     `json:"savingThrows"`
	SkillModifiers    Skills        `json:"skillModifiers"`
	Proficiencies     Proficiencies `json:"proficiencies"`
	Attacks           []Attack      `json:"attacks"`
	Spells            []Spell       `json:"spells"`
	Items             Items         `json:"items"`
}

type Magic struct {
	Modifier int `json:"modifier"`
	Attack   int `json:"attack"`
	SaveDC   int `json:"saveDC"`
}

type Abilities struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
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

// TODO: replace strings with enums
type Proficiencies struct {
	Skills       []string `json:"skills"`
	SavingThrows []string `json:"savingThrows"`
	Resistances  []string `json:"resistances"`
	Armor        []string `json:"armor"`
	Weapons      []string `json:"weapons"`
	Languages    []string `json:"languages"`
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

type Spell struct {
	Name          string `json:"name"`
	Level         int    `json:"level"`
	Type          string `json:"type"`
	Range         int    `json:"range"`
	CastingTime   int    `json:"castingTime"`
	Concentration bool   `json:"concentration"`
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
	Name         string `json:"name"`
	Type         string `json:"type"`
	Rarity       string `json:"rarity"`
	IsEquipped   bool   `json:"isEquipped"`
	AddDexterity bool   `json:"addDexterity"`
	ArmorClass   int    `json:"armorClass"`
}

// TODO: add type for dice value (ex: "1d8")
type Weapon struct {
	Name       string   `json:"name"`
	AttackType string   `json:"attackType"`
	Range      int      `json:"range"`
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

func calculateAbilityModifier(ability int) int {
	normalizedAbility := ability - 10
	if normalizedAbility < 0 {
		normalizedAbility = normalizedAbility - 1
	}
	return int(normalizedAbility / 2)
}

func (character *CharacterOutput) GenerateAbilityModifiers() {
	character.AbilityModifiers.Strength = calculateAbilityModifier(character.Abilities.Strength)
	character.AbilityModifiers.Dexterity = calculateAbilityModifier(character.Abilities.Dexterity)
	character.AbilityModifiers.Constitution = calculateAbilityModifier(character.Abilities.Constitution)
	character.AbilityModifiers.Intelligence = calculateAbilityModifier(character.Abilities.Intelligence)
	character.AbilityModifiers.Wisdom = calculateAbilityModifier(character.Abilities.Wisdom)
	character.AbilityModifiers.Charisma = calculateAbilityModifier(character.Abilities.Charisma)
}

func calculateModifierWithProficiency(skill string, baseModifier int, proficiencies []string, proficiencyBonus int) int {
	for _, p := range proficiencies {
		if p == skill {
			return baseModifier + proficiencyBonus
		}
	}
	return baseModifier
}

func (character *CharacterOutput) GenerateSkillModifiers() {
	character.SkillModifiers.Acrobatics = calculateModifierWithProficiency("acrobatics", character.AbilityModifiers.Dexterity, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.AnimalHandling = calculateModifierWithProficiency("animalHandling", character.AbilityModifiers.Wisdom, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Arcana = calculateModifierWithProficiency("arcana", character.AbilityModifiers.Intelligence, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Athletics = calculateModifierWithProficiency("athletics", character.AbilityModifiers.Strength, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Deception = calculateModifierWithProficiency("deception", character.AbilityModifiers.Charisma, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.History = calculateModifierWithProficiency("history", character.AbilityModifiers.Intelligence, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Insight = calculateModifierWithProficiency("insight", character.AbilityModifiers.Wisdom, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Intimidation = calculateModifierWithProficiency("intimidation", character.AbilityModifiers.Charisma, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Investigation = calculateModifierWithProficiency("investigation", character.AbilityModifiers.Intelligence, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Medicine = calculateModifierWithProficiency("medicine", character.AbilityModifiers.Wisdom, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Nature = calculateModifierWithProficiency("nature", character.AbilityModifiers.Intelligence, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Perception = calculateModifierWithProficiency("perception", character.AbilityModifiers.Wisdom, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Performance = calculateModifierWithProficiency("performance", character.AbilityModifiers.Charisma, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Persuasion = calculateModifierWithProficiency("persuasion", character.AbilityModifiers.Charisma, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Religion = calculateModifierWithProficiency("religion", character.AbilityModifiers.Intelligence, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.SleightOfHand = calculateModifierWithProficiency("sleightOfHand", character.AbilityModifiers.Dexterity, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Stealth = calculateModifierWithProficiency("stealth", character.AbilityModifiers.Dexterity, character.Proficiencies.Skills, character.ProficiencyBonus)
	character.SkillModifiers.Survival = calculateModifierWithProficiency("survival", character.AbilityModifiers.Wisdom, character.Proficiencies.Skills, character.ProficiencyBonus)
}

func (character *CharacterOutput) GenerateSavingThrows() {
	character.SavingThrows.Strength = calculateModifierWithProficiency("strength", character.AbilityModifiers.Strength, character.Proficiencies.SavingThrows, character.ProficiencyBonus)
	character.SavingThrows.Dexterity = calculateModifierWithProficiency("dexterity", character.AbilityModifiers.Dexterity, character.Proficiencies.SavingThrows, character.ProficiencyBonus)
	character.SavingThrows.Constitution = calculateModifierWithProficiency("constitution", character.AbilityModifiers.Constitution, character.Proficiencies.SavingThrows, character.ProficiencyBonus)
	character.SavingThrows.Intelligence = calculateModifierWithProficiency("intelligence", character.AbilityModifiers.Intelligence, character.Proficiencies.SavingThrows, character.ProficiencyBonus)
	character.SavingThrows.Wisdom = calculateModifierWithProficiency("wisdom", character.AbilityModifiers.Wisdom, character.Proficiencies.SavingThrows, character.ProficiencyBonus)
	character.SavingThrows.Charisma = calculateModifierWithProficiency("charisma", character.AbilityModifiers.Charisma, character.Proficiencies.SavingThrows, character.ProficiencyBonus)
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

	character.Magic.Attack = character.ProficiencyBonus + character.Magic.Modifier
	character.Magic.SaveDC = 8 + character.Magic.Attack
}

func (character *CharacterOutput) GenerateAttacks() {
	for _, weapon := range character.Items.Weapons {
		attack := Attack{
			WeaponName: weapon.Name,
			AttackType: weapon.AttackType,
			Range:      weapon.Range,
			Damage: Damage{
				Base: weapon.BaseDamage,
				Type: weapon.DamageType,
			},
		}

		var hitModifier int
		if weapon.AttackType == "melee" {
			hitModifier = character.AbilityModifiers.Strength

			for _, property := range weapon.Properties {
				if property == "finesse" && character.AbilityModifiers.Dexterity > character.AbilityModifiers.Strength {
					hitModifier = character.AbilityModifiers.Dexterity
				}
			}
		} else {
			hitModifier = character.AbilityModifiers.Dexterity
		}

		attack.Hit = calculateModifierWithProficiency(weapon.Name, hitModifier, character.Proficiencies.Weapons, character.ProficiencyBonus)
		attack.Damage.Modifier = hitModifier

		character.Attacks = append(character.Attacks, attack)
	}
}

func (character *CharacterOutput) GenerateArmorClass() {
	var baseArmorClass int
	for _, armor := range character.Items.Equipment {
		if armor.IsEquipped {
			if armor.AddDexterity {
				baseArmorClass = baseArmorClass + armor.ArmorClass + character.AbilityModifiers.Dexterity
			} else {
				baseArmorClass = baseArmorClass + armor.ArmorClass
			}
		}
	}

	character.ArmorClass = baseArmorClass
}

func main() {
	fmt.Println("Hello SemVer Hero!")

	inputFileName := "examples/levi-input.json"
	outputFileName := "examples/levi-output.json"

	// TODO: Pass in filename as an argument
	inputFile, err := os.Open(inputFileName)
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

	// Calculate ProficiencyBonus
	proficiencyByLevel := [...]int{0, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6}
	character.ProficiencyBonus = proficiencyByLevel[character.Level]

	// Calculate Ability Modifiers
	character.GenerateAbilityModifiers()

	// Calculate Skill Modifiers
	character.GenerateSkillModifiers()

	// Calculate Saving Throws
	character.GenerateSavingThrows()

	// Calculate Initiaitve and Passive Perception
	character.Initiative = character.AbilityModifiers.Dexterity
	character.PassivePerception = 10 + character.SkillModifiers.Perception

	// Calculate Magic Stats
	character.GenerateMagicStats()

	// Calculate Attacks
	character.GenerateAttacks()

	// Calculate Armor Class
	character.GenerateArmorClass()

	fmt.Printf("Successfully generated stats for %s:\n", character.Name)
	fmt.Println(character)

	data, err := json.MarshalIndent(character, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(outputFileName, data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully wrote %s's stats to  %s\n", character.Name, outputFileName)
}
