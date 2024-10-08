package model

import (
	"fmt"
	"math"
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
	// Red > Green > Blue > Red (2)
	// Red < Blue < Green < Red (1/2)
	// Red == Cyan == Red (1)
	
	// Red to green curve :: sin(3x/4)
	// Gree to cyan curve (from reds perspective) :: sin(3x/2)

	sourceHue := source.hue()
	targetHue := target.hue()		
	
	fmt.Printf("%v vs %v = %.2f", source, target ,sourceHue - targetHue)
	fmt.Println()
	if (math.Abs(sourceHue - targetHue) <= 0.01) { 
		return 0 
	}

	dir := 0.0
	if (targetHue > sourceHue) { 
		dir = 1
	} else { 
		dir = -1
	}

	var po2 = 0.0
	overclock := math.Abs((2*math.Pi/3)/(sourceHue-targetHue)) // how far away from a 3rd turn 
	if (overclock > 1) {
		fmt.Println("overclocked")
		po2 = (dir * math.Sin( (3/2)*sourceHue))
	} else { 
		fmt.Println("not overclocked")
		po2 = (dir * math.Sin( (3/2)*sourceHue))
	}
	return po2
}

func (source *Color) hue() float64 {
	var hue = 0.0
	var min = min(source.Red, source.Green, source.Blue) 
	var max = max(source.Red, source.Green, source.Blue)

	if (source.Red >= source.Green && source.Red >= source.Blue) {
		hue = float64(source.Green - source.Blue) / float64(max - min)
	} else if (source.Green >= source.Red && source.Green >= source.Blue) { 
		hue = 2.0 + float64(source.Blue - source.Red) / float64(max - min)
	} else if (source.Blue >= source.Red && source.Blue >= source.Green) { 
		hue = 4.0 + float64(source.Red - source.Green) / float64(max - min)
	} else  {
		fmt.Println("At least one color should be the maximum")
		fmt.Println(source)
	}

	hueRads := hue * (math.Pi/3.0)  
	if (hueRads < 0) { 
		hueRads = hueRads + (2*math.Pi)
	}
	return hueRads
}