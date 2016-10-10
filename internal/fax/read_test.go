package fax

// Test the decoding.

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"strings"
	"testing"
)

type goldenDecode struct {
	base2  string
	pixels string
}

func TestDecode(t *testing.T) {
	var tests = [...]goldenDecode{
		{verticalCodes[-1] + verticalCodes[0], "wwb"},
		{passCode + verticalCodes[-2] + passCode, "www.wbb"},
		{verticalCodes[-3] + passCode + verticalCodes[1] + verticalCodes[0], "bbb.wbb"},
		{verticalCodes[-2] + verticalCodes[-1] + passCode + verticalCodes[-1] + passCode, "wbw.bbb"},
		{horizontalCode + whiteCodes[2] + blackCodes[2] + verticalCodes[-1] + horizontalCode + blackCodes[2] + whiteCodes[1], "wwbb.wbbw"},
		{verticalCodes[-3] + verticalCodes[-1] + passCode + passCode + horizontalCode + whiteCodes[0] + blackCodes[1], "wbbw.wwwb"},
		{horizontalCode + whiteCodes[0] + blackCodes[3] + passCode + passCode + horizontalCode + whiteCodes[0] + blackCodes[1], "bbbw.wwwb"},

		// out of bounds:
		{verticalCodes[3] + verticalCodes[-1] + verticalCodes[2], "ww.wb"},
		{verticalCodes[-3] + verticalCodes[1] + verticalCodes[-2] + verticalCodes[3], "bb.bb"},
		{verticalCodes[-1] + verticalCodes[2] + verticalCodes[-2] + verticalCodes[2], "wb.bb"},
		{horizontalCode + whiteCodes[1] + blackCodes[2], "wb"},
		{horizontalCode + sharedMakeUpCodes[6] + whiteCodes[1] + blackCodes[1], "ww"},
	}
	for _, golden := range tests {
		verifyImage(t, golden.base2, golden.pixels)
	}
}

// TestUncompressedDecode verifies uncompressed mode support.
func TestUncompressedDecode(t *testing.T) {
	base2 := "0000001111"
	_, e := DecodeG4(packImage(base2), 4, 1)
	if e != uncompressedMode {
		t.Fatalf("Wanted error %s, got %s", uncompressedMode, e)
	}
}

// TestExtensionDecode verifies additional mode support.
func TestExtensionDecode(t *testing.T) {
	base2 := "0000001011"
	wanted := "fax: unknown extension 0b011"
	_, e := DecodeG4(packImage(base2), 4, 1)
	if e == nil || e.Error() != wanted {
		t.Fatalf("Wanted error %s, got %s", wanted, e)
	}
}

// TestRunLengthDecode verifies the number of pixels for all
// horizontal Huffman codes.
func TestRunLengthDecode(t *testing.T) {
	for white, whiteCode := range whiteCodes {
		for black, blackCode := range blackCodes {
			base2 := horizontalCode + whiteCode + blackCode
			expected := strings.Repeat("w", white) + strings.Repeat("b", black)
			verifyImage(t, base2, expected)
		}
	}
	for makeup, makeupCode := range whiteMakeUpCodes {
		for white, whiteCode := range whiteCodes {
			base2 := horizontalCode + makeupCode + whiteCode + blackCodes[2]
			whiteCount := white + (makeup+1)*64
			expected := strings.Repeat("w", whiteCount) + "bb"
			verifyImage(t, base2, expected)
		}
	}
	for makeup, makeupCode := range sharedMakeUpCodes {
		for white, whiteCode := range whiteCodes {
			base2 := horizontalCode + makeupCode + whiteCode + blackCodes[2]
			whiteCount := white + (makeup+28)*64
			expected := strings.Repeat("w", whiteCount) + "bb"
			verifyImage(t, base2, expected)
		}
	}
	for makeup, makeupCode := range blackMakeUpCodes {
		for black, blackCode := range blackCodes {
			base2 := horizontalCode + whiteCodes[2] + makeupCode + blackCode
			blackCount := black + (makeup+1)*64
			expected := "ww" + strings.Repeat("b", blackCount)
			verifyImage(t, base2, expected)
		}
	}
	for makeup, makeupCode := range sharedMakeUpCodes {
		for black, blackCode := range blackCodes {
			base2 := horizontalCode + whiteCodes[2] + makeupCode + blackCode
			blackCount := black + (makeup+28)*64
			expected := "ww" + strings.Repeat("b", blackCount)
			verifyImage(t, base2, expected)
		}
	}
}

func verifyImage(t *testing.T, base2 string, pixels string) {
	rows := strings.Split(pixels, ".")
	height := len(rows)
	width := len(rows[0])
	if width == 0 {
		height = 0
	}

	result, error := DecodeG4(packImage(base2), width, 1)
	if error != nil {
		t.Fatal(base2, "resulted in error:", error)
	}

	dimensions := image.Rect(0, 0, width, height)
	if !dimensions.Eq(result.Bounds()) {
		t.Fatal(base2, "bounds", result.Bounds(), "instead of", dimensions)
	}

	for y, row := range rows {
		for x, c := range row {
			color := result.At(x, y)
			r, g, b, a := color.RGBA()
			if c == 'w' {
				if r != 0xFFFF || g != 0xFFFF || b != 0xFFFF || a != 0xFFFF {
					t.Fatalf("%s [%d,%d] = %#v, want %s",
						base2, x+1, y+1, color, pixels)
				}
			} else {
				if r != 0x0000 || g != 0x0000 || b != 0x0000 || a != 0xFFFF {
					t.Fatalf("%s [%d,%d] = %#v, want %s.",
						base2, x+1, y+1, color, pixels)
				}
			}
		}
	}
}

func packImage(base2 string) io.ByteReader {
	base2 += terminatorCode
	if tail := len(base2) % 8; tail != 0 {
		padBitCount := 8 - tail
		base2 += strings.Repeat("0", padBitCount)
	}

	var data bytes.Buffer
	for i := 0; i < len(base2); i += 8 {
		var c byte
		fmt.Sscanf(base2[i:i+8], "%b", &c)
		data.WriteByte(c)
	}
	return &data
}
