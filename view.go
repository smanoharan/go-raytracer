// view.go: Contains definitions of: 
//	- Ray
//	- Camera
package main

// A Ray is a directed line segment, with a start point.
type Ray struct	{
	Start, Direction Vec3
}

// A Camera is a view window into the scene.
type Camera struct {
	Pos, LookAt, Up Vec3
	Width, Height int // size of the view-window, in pixels.
	FovY entry // the field-of-view angle, in degrees, along Y-axis.
}


