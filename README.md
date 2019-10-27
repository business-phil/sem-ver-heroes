# SemVer Heroes

A library (or something?) for tracking and versioning character information.  The intention is to commit or otherwise save each change so that players can go back and reference previous character states.  I was inspired by Not Another D&D Podcast's NannerFly Effect tour where they went back and replayed past events.  How would I have known what what a character's stats were 5 levels ago?

SemVer rules:
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

### Roadmap / Goals

1. Create a JSON template for D&D 5E character sheet information
2. Make the template generic so it can be pulled in as a separate library
3. Add some examples

### What this is

* Generic library for tracking character information
* For Dungeons and Dragons 5th Edition character sheets and other metadata
* Probably JSON-based, but TBD for how this will work long-term

### What this is not (yet)

* Usable
* For non-D&D games
* A GUI for viewing/editing character information in the browser
