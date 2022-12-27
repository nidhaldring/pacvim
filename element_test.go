package main

import "testing"

func TestGetEelementType(t *testing.T) {
	if getEelementType(' ') != SPACE {
		t.Error("type for space is not correct")
	}

	if getEelementType('#') != BLOCKING {
		t.Error("type for blocking is not correct")
	}

	if getEelementType('~') != DEADLY {
		t.Error("type for deadly is not correct")
	}

	if getEelementType('a') != EATABLE {
		t.Error("type for deadly is not correct")
	}
}

func TestGetElementStyle(t *testing.T) {
	if getEelementStyle(SPACE) != DefTheme {
		t.Error("space theme is not correct")
	}

	if getEelementStyle(BLOCKING) != DefTheme {
		t.Error("blocking theme is not correct")
	}

	if getEelementStyle(DEADLY) != DeadlyTheme {
		t.Error("deadly theme is not correct")
	}

	if getEelementStyle(EATABLE) != DefTheme {
		t.Error("eatable theme is not correct")
	}

	if getEelementStyle(EATEN) != EatenTheme {
		t.Error("deadly theme is not correct")
	}
}
