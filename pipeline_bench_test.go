package main

import (
	"testing"
)

// BenchmarkPipeline measures the throughput of the entire image processing pipeline.
func BenchmarkPipeline(b *testing.B) {
	imagePaths := []string{"images/image1.jpeg", "images/image2.jpeg"}

	// Ensure the images exist and are accessible before running the benchmark
	// ...

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		channel1 := loadImage(imagePaths)
		channel2 := resize(channel1)
		channel3 := convertToGrayscale(channel2)
		writeResults := saveImage(channel3)

		for success := range writeResults {
			if !success {
				b.Error("Failed to process an image during benchmark.")
				return
			}
		}
	}
}

// BenchmarkLoadImage measures the performance of loading images.
func BenchmarkLoadImage(b *testing.B) {
	imagePaths := []string{"images/image1.jpeg", "images/image2.jpeg"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		out := loadImage(imagePaths)

		// Drain the channel to ensure all images are loaded
		for range out {
		}
	}
}

// BenchmarkResize measures the performance of resizing images.
func BenchmarkResize(b *testing.B) {
	imagePaths := []string{"images/image1.jpeg", "images/image2.jpeg"}
	channel1 := loadImage(imagePaths)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		out := resize(channel1)

		// Drain the channel to ensure all images are resized
		for range out {
		}
	}
}

// BenchmarkConvertToGrayscale measures the performance of converting images to grayscale.
func BenchmarkConvertToGrayscale(b *testing.B) {
	imagePaths := []string{"images/image1.jpeg", "images/image2.jpeg"}
	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		out := convertToGrayscale(channel2)

		// Drain the channel to ensure all images are converted
		for range out {
		}
	}
}

// BenchmarkSaveImage measures the performance of saving images.
func BenchmarkSaveImage(b *testing.B) {
	imagePaths := []string{"images/image1.jpeg", "images/image2.jpeg"}
	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		out := saveImage(channel3)

		// Drain the channel to ensure all images are saved
		for success := range out {
			if !success {
				b.Error("Failed to save an image during benchmark.")
				return
			}
		}
	}
}
