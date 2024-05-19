package learning

import (
	"testing"

	"github.com/stretchr/testify/assert"
) 

func TestDuplicateString(t *testing.T) {
	strArray := []string{"alekhya", "loukya", "sony", "loukya", "alekhya", "satya"}
	files := DuplicateFollderName(strArray)
	assert.Equal(t, []string{"alekhya", "alekhya1", "loukya", "loukya1", "satya", "sony"}, files, "Expcted result is []string. But got false.")
} 

func TestDuplicateString_negative(t *testing.T) {
	strArray := []string{"alekhya", "loukya", "sony", "advith", "bhavani", "satya"}
	files := DuplicateFollderName(strArray)
	assert.Equal(t, []string{"advith", "alekhya", "bhavani", "loukya", "satya", "sony"}, files, "Expcted result is []string. But got false.")
} 
     