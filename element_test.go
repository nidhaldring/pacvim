package main

import "testing"

func TestGetEelementType(t *testing.T) {
	if getElementType(' ') != SPACE {
		t.Error("type for space is not correct")
	}

	if getElementType('#') != BLOCKING {
		t.Error("type for blocking is not correct")
	}

	if getElementType('~') != DEADLY {
		t.Error("type for deadly is not correct")
	}

	if getElementType('a') != EATABLE {
		t.Error("type for deadly is not correct")
	}
}

func TestGetElementStyle(t *testing.T) {
	if getElementStyle(SPACE) != DefTheme {
		t.Error("space theme is not correct")
	}

	if getElementStyle(BLOCKING) != DefTheme {
		t.Error("blocking theme is not correct")
	}

	if getElementStyle(DEADLY) != DeadlyTheme {
		t.Error("deadly theme is not correct")
	}

	if getElementStyle(EATABLE) != DefTheme {
		t.Error("eatable theme is not correct")
	}

	if getElementStyle(EATEN) != EatenTheme {
		t.Error("deadly theme is not correct")
	}
}
