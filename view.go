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

// for setting pixels
func Set(o *image.RGBA, x, y int, col *Vec3) {
	o.Set(x, y, color.RGBA{uint8(col[cX]), uint8(col[cY]), uint8(col[cZ]), uint8(255)})
}
