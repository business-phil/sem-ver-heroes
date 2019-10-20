package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Character struct {
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

func main() {
	fmt.Println("Hello SemVer Hero!")

	// TODO: Pass in filename as an argument
	jsonFile, err := os.Open("sample-input.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	var character Character
	json.Unmarshal(byteValue, &character)

	fmt.Printf("Successfully parsed input for %s:\n", character.Name)
	fmt.Println(character)
}
