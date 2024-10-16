package model

import (
	"math"
	// "fmt"
)

type CombatStats struct { 
	Strength		uint8
	Fortitude 		uint8
	Constitution 	uint8
	Intelligence	uint8
	Wisdom			uint8 
	Willpower 		uint8
	Agility			uint8
	Dexterity		uint8
	Endurance 		uint8 
	Color			Color	
}

type Color struct { 
	Red		uint8
	Green   uint8
	Blue	uint8
}

func (source *CombatStats) calculateStrengthDamage(target CombatStats) float64 {	
	return float64(source.Strength) / float64(target.Fortitude) 
}
func (source *CombatStats) calculateIntelligenceDamage(target CombatStats) float64  {
	return float64(source.Intelligence) / float64(target.Wisdom)
}
func (source *CombatStats) calculateAgilityDamage(target CombatStats) float64 { 
	return float64(source.Agility) / float64(target.Dexterity)
}

var PRIMARY_SOURCE_STAT_MODIFIER uint8 = 10
var SECONDARY_SOURCE_STAT_MODIFIER uint8 = 10
var TERTIARY_SOURCE_STAT_MODIFIER uint8 = 10

func calculateResourceTotal(primary uint8, secondary uint8, tertiary uint8) uint64 { 
	return uint64(math.Round(
			math.Pow(float64(primary*PRIMARY_SOURCE_STAT_MODIFIER), 2.0) + 
		   	math.Pow(float64(secondary*SECONDARY_SOURCE_STAT_MODIFIER), 1) + 
		   	math.Pow(float64(tertiary*TERTIARY_SOURCE_STAT_MODIFIER), 0.5)))
}

func (source *CombatStats) MaxHp() uint64 { 
	return calculateResourceTotal(source.Constitution, source.Wisdom, source.Agility) 
}

func (source *CombatStats) MaxMp() uint64 { 
	return calculateResourceTotal(source.Willpower, source.Dexterity, source.Strength) 
}

func (source *CombatStats) MaxSp() uint64 { 
	return calculateResourceTotal(source.Endurance, source.Fortitude, source.Intelligence)
}

func (source *Color) ColorDamageModifier(target *Color) float64 {
	var mod float64 = 1.0
	diff := target.hue() - source.hue()
	if (diff < 0) {
		mod = -1.0
	}

	diff = math.Abs(diff)
	if (diff >= 0 && diff <= (2*math.Pi/3.0)) {
		return mod * math.Sin((3/4.0)*diff)
	}

	diff -= (2*math.Pi/3.0)
	if (diff >= 0 && diff <= (2*math.Pi/3.0)) {
		return mod * math.Cos((3/2.0) * diff)
	}
	
	return mod * -math.Sin((3/4.0) * diff)
}

func (source *Color) hue() float64 {
	var hue = 0.0
	var min = min(source.Red, source.Green, source.Blue) 
	var max = max(source.Red, source.Green, source.Blue)

	if (source.Red >= source.Green && source.Red >= source.Blue) {
		hue = (float64(source.Green) - float64(source.Blue)) / float64(max - min)
	} else if (source.Green >= source.Red && source.Green >= source.Blue) { 
		hue = 2.0 + (float64(source.Blue) - float64(source.Red)) / float64(max - min)
	} else if (source.Blue >= source.Red && source.Blue >= source.Green) { 
		hue = 4.0 + (float64(source.Red) - float64(source.Green)) / float64(max - min)
	}
	hueRads := hue * (math.Pi/3.0)  
	if (hueRads < 0) { 
		hueRads += 2*math.Pi
	}
	return hueRads
}