RayTracer
============

A simple RayTracer using the BlinnPhong shader, implemented in Go.

Supported features:
-------------------
* Light sources: point and directional
* Lighting/Material properties: Diffuse colour, Specular colour and Shininess 
* Primitives: Sphere only
* Soft shadows
* Anti-aliasing

Usage:
------
1. Instantiate a Camera object:

    camera := &Camera{cameraPosition, lookAtDirection, upVector, imageWidth, imageHeight, fovY}

	* `cameraPosition`: a 3D vector, the position of the camera in space.
	* `lookAtDirection`: a 3D vector, the direction in which the camera is pointing (aka the forward direction).
	* `upVector`: a 3D vector, the direction considered to be straight up, typically the +Y axis `(0,1,0)` or +Z axis `(0,0,1)`.
	* `imageWidth`: an int, the width of the output image, in pixels.
	* `imageHeight`: an int, the height of the output image, in pixels.
	* `fovY`: a float, the field-of-view angle in the Y-axis of the image, in degrees (not radians).

2. Instantiate the ray tracer:

    raytracer := NewRayTracer(camera, &RayTracerOptions{ recursiveRayLimit, samplingFactor, numShadowRays })

	* `camera`: a Camera (see step 1).
	* `recursiveRayLimit`: an int, the maximum number of times to bounce each ray off surfaces.
	A larger limit results in more reflections being visible, but increases runtime (keeping all other parameters constant, 
	runtime is in `O(n!)` where `n` is the `recursiveRayLimit`). Recommended value: 5
	* `samplingFactor`: an int, for anti-aliasing. Each pixel in the output image results in `samplingFactor x samplingFactor`
	(slightly different) rays being traced (i.e. runtime grows quadractically with respect to the `samplingFactor`.
	* `numShadowRays`: an int, for soft shadowing. For each shadow computation, `numShadowRays` are traced. Runtime grows linearly with `numShadowRays`.

3. Instantiate the lights (a list of point and/or directional lights):

    lights := []Light{
        &PointLight{ color, position, attenuation }, 
		&DirectionalLight{ color, direction }, 
        // ... add as many lights as necessary
    }

	* `color`: a 3D vector, the colour of the light.
	* `position`: a 3D vector, the position in space, of the light.
	* `attenutation`: a 3D vector, controls how quickly the light intensity decreases over distance.
	* `direction`: a 3D vector, the direction in which the directional light travels.

4. Instantiate the scene, including the materials:

    scene := []Shape{
	   NewSphere(radius, position, &Material{ ambient, emission, diffuse, specular, shininess }),
	   // ... add as many shapes to the scene as necessary
    }

	* `radius`: a float, the radius of the sphere.
	* `position`: a 3D vector, the position in space of the sphere.
	* `ambient`, `emission`, `diffuse`, `specular`: all 3D vectors, the various colour properties of the material.
	* `shininess`: float, controls how shiny the material is.


5. Render the scene into an image:

    image := raytracer.Draw(scene, lights)

6. Save the image to disk, using Go's standard file I/O routines. A sample helper function `saveImg(path, image)` has been provided in `main.go`.

TODO
----
* Quad and Triangle primitives.
* Sample scenes.
