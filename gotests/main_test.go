package main

import (
	"testing"
)

func TestNoBranch(t *testing.T) {
	NoBranch(1)
	NoBranch(2)
}

func TestOneBranch(t *testing.T) {
	OneBranch(1)
	OneBranch(2)
}

func TestTwoBranch(t *testing.T) {
	TwoBranch(6, 0)
	TwoBranch(7, 0)
	TwoBranch(13, 0)
	TwoBranch(14, 0)
}
