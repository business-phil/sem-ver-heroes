# SemVer Heroes

A library (or something?) for tracking and versioning character information.  The intention is to commit or otherwise save each change so that players can go back and reference previous character states.  I was inspired by Not Another D&D Podcast's NannerFly Effect tour where they went back and replayed past events.  How would I have known what what a character's stats were 5 levels ago?

**SemVer rules**:
* `{major}` for character levels
* `{minor}` for changes that would affect a character sheet since the last level-up
* `{patch}` for suckers, I guess?

**Example**: Otto Otto, level 4 Dwarf Cleric who has acquired extra equipment and/or had a stat change twice since leveling up
```json
{
  "name": "Otto Otto",
  "version": "4.2",
  "race": "Dwarf",
  "class": "Cleric",
  // ...
}
```
See also the [sample-input.json](https://github.com/business-phil/sem-ver-heroes/blob/master/sample-input.json) file for a full character sheet example.  You can compare this to [sample-output.json](https://github.com/business-phil/sem-ver-heroes/blob/master/sample-ouput.json) to see what values are generated.

### Dynamic Stat Calculations

In order to reduce the amount of data a user would need to track on their end, this library calculates as many stats as possible.  Documenting all of the abilities and dynamic rules available would be near-impossible - not to mention what that would do to the package size - so only some basic calculations are supported at this time.

**Supported**:
* Saving Throws and Skill modifiers based on ability stats and proficiencies
* Magic modifiers based on stats and class
* Attacks based on weapons, stats and specific weapon proficiencies (i.e., "Warhammer")

**In Development**:
* Weapon proficiencies for weapon categories (i.e., "simple weapons")
* Other attack options based on weapon properties (ex: two-handed option for "versatile")
* Armor Class based on equipped weapons, stats and armor proficiencies

**Not Supported (yet?)**:
* Spells and spell slots
* Other abilities and items that affect stats
* Skill overrides and/or additional modifiers to calculate after other stat generation

### Roadmap / Goals

:ballot_box_with_check: ~Create a JSON template for D&D 5E character sheet information~  
:black_square_button: Generate as many stat as possible in order to mitigate the amount of input data required  
:black_square_button: Make the template generic so it can be pulled in as a separate library  
:black_square_button: Add some examples  
:black_square_button: Document as many rules as possible (ex: case-sensitive proficiencies) and publish to [GoDoc](https://godoc.org/)  

### What this is

* Generic library for tracking character information
* For Dungeons and Dragons 5th Edition character sheets and other metadata
* Probably JSON-based, but TBD for how this will work long-term

### What this is not (yet)

* Usable
* For non-D&D games
* A GUI for viewing/editing character information in the browser
