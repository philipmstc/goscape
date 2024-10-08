package model

import (
	"fmt"
	"testing"
	"math"
)

func TestCombatStats(t *testing.T) { 
	stats := CombatStats { 1,3,5,1,3,5,1,3,5, Color{255,0,0}}
	hp := stats.MaxHp()
	mp := stats.MaxMp()
	sp := stats.MaxSp()
	fmt.Printf("HP: %d, MP: %d, SP: %d", hp, mp, sp)
	fmt.Println()
	if (! (hp == mp && hp == sp)) {
		t.Error("All points totals should be the same")
	}	 
}

func TestRedVsBlue(t *testing.T) {
	red := Color{255,0,0}
	if (math.Abs(red.hue()) > 0.001) {
		t.Errorf("Expected red to be 0pi, was %.2fpi", red.hue())
	}
	blue := Color{0,0,255}
	if (math.Abs(blue.hue() - 4*math.Pi/3) > 0.001) {
		t.Errorf("Expected blue to be 2/3 pi, was %.2fpi", blue.hue())
	}
	green := Color{0,255,0}
	if (math.Abs(green.hue() - 2*math.Pi/3) > 0.001) { 
		t.Errorf("Expected green to be 1/3 pi, was %.2fpi", green.hue())
	}
	cyan := Color{0,255,255}
	if (math.Abs(cyan.hue() - math.Pi) > 0.001) { 
		t.Errorf ("Expected cyan to be pi, was %.2fpi", cyan.hue())
	}

	rvb := red.ColorDamageModifier(&blue)
	if (rvb != -1.0) { 
		t.Errorf("Expected -1, got %.2f", rvb)
	}
	rvg := red.ColorDamageModifier(&green)
	if (rvg != 1.0) { 
		t.Errorf("Expected 1, got %.2f", rvg)
	}
	rvc := red.ColorDamageModifier(&cyan)
	if (rvc != 0) { 
		t.Errorf("Expected 0, got %.2f", rvc)
	}
}
