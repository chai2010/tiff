package fax

import (
	"bytes"
	"image"
	"image/gif"
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

var testImageNames = []string{"red"}
var benchmarkImageNames = []string{"red"}

func sample(name string) []byte {
	data, err := ioutil.ReadFile("../../testdata/compress/" + name + ".tiff")
	if err != nil {
		panic(err)
	}
	return data[8:] // strip TIFF header
}

func prototype(name string) (result image.Image) {
	file, err := os.Open("../../testdata/compress/" + name + ".gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	result, err = gif.Decode(file)
	if err != nil {
		panic(err)
	}
	return result
}

func TestFullDecode(t *testing.T) {
	for _, name := range testImageNames {
		failFile := name + "-fail.png"
		os.Remove(failFile)

		expected := prototype(name)
		bounds := expected.Bounds()

		reader := bytes.NewBuffer(sample(name))
		result, err := DecodeG4(reader, bounds.Dx(), bounds.Dy())
		if err != nil {
			t.Fatalf("decode %s gave %s", name, err)
		}
		if !bounds.Eq(result.Bounds()) {
			t.Fatalf("%s bound to %#v", name, result.Bounds())
		}

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				ra, ga, ba, aa := result.At(x, y).RGBA()
				rb, gb, bb, ab := expected.At(x, y).RGBA()
				if ra != rb || ga != gb || ba != bb || aa != ab {
					file, err := os.Create(failFile)
					if err != nil {
						t.Fatalf("create %s gave %s", failFile, err)
					}
					defer file.Close()

					err = png.Encode(file, result)
					if err != nil {
						t.Fatalf("encode %s gave %s", failFile, err)
					}

					t.Error("recorded mismatch:", failFile)
				}
			}
		}
	}
}

func BenchmarkFullDecode(b *testing.B) {
	b.StopTimer()
	var pixelCount int64
	for _, name := range benchmarkImageNames {
		data := sample(name)
		bounds := prototype(name).Bounds()
		width := bounds.Dx()
		height := bounds.Dy()
		pixelCount += int64(width * height)

		for remaining := b.N; remaining > 0; remaining-- {
			reader := bytes.NewBuffer(data)
			b.StartTimer()
			_, err := DecodeG4(reader, width, height)
			b.StopTimer()
			if err != nil {
				b.Fatal(err)
			}
		}
	}
	b.SetBytes(pixelCount)
}
