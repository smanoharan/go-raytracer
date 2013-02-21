// view.go: Contains definitions of: 
//	- Ray
//	- Camera
package main

// A Ray is a directed line segment, with a start point.
type Ray struct	{
	start, direction Vec3
}

// A Camera is a view window into the scene.
type Camera struct {
	pos, lookAt, up Vec3
	width, height int // size of the view-window, in pixels.
	fovY entry // the field-of-view angle, in degrees, along Y-axis.
}


