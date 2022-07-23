package main

import (
	"github.com/greyvar/world-generator/pkg/gridgen"
)

func main() {
	//gridgen.GenerateBiomeReticulationTest()
	
	gridgen.Generate(&gridgen.GenerationConfig{
		RowCount: 16,
		ColCount: 20,
		Seed: int64(1345),
		ReticulateSplines: true,
	});
}
