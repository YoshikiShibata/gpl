// Copyright © 2015 Yoshiki Shibata

package main

import (
	"fmt"
	"image"
	"image/color"
)

func superSample(srcImg *image.RGBA) (*image.RGBA, error) {
	bounds := srcImg.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	if width%2 != 0 || height%2 != 0 {
		return nil, fmt.Errorf("illegal bounds : %v", bounds)
	}

	reducedImage := image.NewRGBA(image.Rect(0, 0, width/2, height/2))
	for x := 0; x < width/2; x++ {
		for y := 0; y < height/2; y++ {
			reducedImage.Set(x, y, superSampledBy4SubPixels(srcImg, x*2, y*2))
		}
	}
	return reducedImage, nil
}

func superSampledBy4SubPixels(srcImage *image.RGBA, x, y int) color.Color {
	pixels := []color.RGBA{srcImage.RGBAAt(x, y),
		srcImage.RGBAAt(x+1, y),
		srcImage.RGBAAt(x, y+1),
		srcImage.RGBAAt(x+1, y+1)}
	var sumR, sumG, sumB uint32
	for _, rgba := range pixels {
		r, g, b, _ := rgba.RGBA()
		sumR += r
		sumG += g
		sumB += b
	}
	return color.RGBA{uint8(sumR / 4), uint8(sumG / 4), uint8(sumB / 4), 0xff}
}
