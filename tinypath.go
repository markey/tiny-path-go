package tiny_path_go

import (
	"image/color"
	"math/rand"
)


func render(width int, height int, spp int) {
	spp_inv := 1.0 / spp
	image_data := [][]{height}{width, color.Black}

	inv_pixel_count := 1.0 / (width * height)
	i := 0
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			pixel := color(0, 0, 0)
			for s := 0; s < spp; s++ {
				sx := (x + rand.Int() - (width * 0.5)
				sy := (y + rand.Int() - (height * 0.5)
				pixel := radiance(CAMERA.span_ray(sx, sy)) * spp_inv
			}
			image_data[y][x] = pixel
			print_progress(i, width * height)
			i++
		}
	}
}
