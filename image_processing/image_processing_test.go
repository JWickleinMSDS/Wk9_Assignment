package imageprocessing

import (
	"image"
	_ "image/jpeg" // what you need to decobe jpeg images
	"testing"

	"golang.org/x/image/draw"
)

// TestGrayscale tests the grayscale conversion and resizing of an input image, comparing it to an expected output.
func TestGrayscale(t *testing.T) {
	// Grab the input image
	inputImage, err := ReadImage("testdata/inputImage.jpeg")
	if err != nil {
		t.Fatalf("Failed to read input image: %v", err)
	}

	// Convert the inputed image to grayscale
	grayImage := Grayscale(inputImage)

	// Resize the grayscale image to 500x500 pixels
	resizedGrayImage := resizeImageTo500x500(grayImage)

	// Grab the expected grayscale image
	expectedImage, err := ReadImage("testdata/expectedGrayImage.jpeg")
	if err != nil {
		t.Fatalf("Failed to read expected grayscale image: %v", err)
	}

	// Compare the resized grayscale image to the expected grayscale image
	if !imagesEqual(resizedGrayImage, expectedImage) {
		t.Errorf("The processed image does not match the expected output.")
	}
}

// resizeImageTo500x500 resizes an image to 500x500 pixels.
func resizeImageTo500x500(src image.Image) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Src, nil)
	return dst
}

// imagesEqual compares two images pixel by pixel.
func imagesEqual(img1, img2 image.Image) bool {
	if img1.Bounds() != img2.Bounds() {
		return false
	}
	for y := img1.Bounds().Min.Y; y < img1.Bounds().Max.Y; y++ {
		for x := img1.Bounds().Min.X; x < img1.Bounds().Max.X; x++ {
			if img1.At(x, y) != img2.At(x, y) {
				return false
			}
		}
	}
	return true
}
