package tiny_path_go

type Col struct {
	r int
	g int
	b int
}

func render(width int, height int, spp int) {
	spp_inv := 1.0 / spp
	var image_data [height]int

}
