// view.go: Contains definitions of: 
//	- Ray
//	- Camera
//	- Output Image
package main

import (
	"image"
	"image/color"
)

// A Ray is a directed line segment, with a start point.
type Ray struct {
	start, direction Vec3
}

// A Camera is a view window into the scene.
type Camera struct {
	pos, lookAt, up Vec3
	width, height   int   // size of the view-window, in pixels.
	fovY            entry // the field-of-view angle, in degrees, along Y-axis.
}

// for creating an image
func NewOutputImage(width, height int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

// scale a color in [0,1] to the [0,255] range.
func toRGB(e entry) uint8 {
	if e >= ONE {
		return uint8(255)
	}
	if e <= ZERO {
		return uint8(0)
	}
	return uint8(e * entry(256))
}

// for setting pixels
func Set(o *image.RGBA, x, y int, col *Vec3) {
	o.Set(x, y, color.RGBA{toRGB(col[cX]), toRGB(col[cY]), toRGB(col[cZ]), uint8(255)})
}
