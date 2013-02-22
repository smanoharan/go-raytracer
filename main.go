package main

import (
	"image"
	"image/png"
	"os"
	"time"
)

func main() {

	// camera at (0,1,5) looking towards origin
	eyePos, lookAt, up := Vec3{0, 2, 6}, ZERO_V3, Y_V3
	width, height, fovY := 300, 300, entry(50)
	view := &Camera{eyePos, lookAt, up, width, height, fovY}

	// init ray tracer
	rayTracer := NewRayTracer(view, &RayTracerOptions{5, 2, 8})

	// create materials, scene and lights:
	mat := &Material{Vec3{0.3, 0.3, 0.3}, ZERO_V3, Vec3{0.2, 0.4, 0.2}, Vec3{0.2, 0.35, 0.2}, entry(15)}
	atten := X_V3
	lights := []Light{
		&PointLight{Vec3{0.2, 0.4, 0.2}, Vec3{0, 5, 3}, atten},
		&PointLight{Vec3{0.4, 0.3, 0.3}, Vec3{-6, 1, 3}, atten},
	}
	//	scene1 := []Shape{NewSphere(TWO, &ZERO_V3, mat), NewSphere(ONE, X_V3.scale(FOUR), mat)}
	s2size := 5
	scene2 := make([]Shape, s2size*s2size)
	for i := 0; i < s2size; i++ {
		for j := 0; j < s2size; j++ {
			si, sj, ci, cj := entry(1.5), entry(-2), entry(i-2), entry(j)
			scene2[i*s2size+j] = NewSphere(entry(0.5), &Vec3{si * ci, -TWO, sj * cj}, mat)
		}
	}

	//	saveImg("scene1.png", rayTracer.Draw(scene1, lights))
	start := time.Now()
	saveImg("scene2.png", rayTracer.Draw(scene2, lights))
	end := time.Now()
	println("Done in", end.Sub(start).String())
}

func saveImg(path string, img *image.RGBA) error {
	output, err := os.Create(path)
	if err != nil {
		return err
	}
	defer output.Close()
	return png.Encode(output, img)
}
