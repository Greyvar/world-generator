package gridgen


import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/greyvar/datlib/gridfiles"
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

func GenerateBiomeReticulationTest() {
	base := &gridfiles.Grid{}
	base.Filename = "biome reticulation"
	base.RowCount = 16
	base.ColCount = 16
	base.Build()

//	drawCross(base, 4, 1, 12, 15, "sand.png")
//	drawCross(base, 6, 3, 10, 13, "grass")

	drawRect(base, 5, 5, 6, 14, "sand")
	drawRect(base, 3, 7, 12, 8, "sand")
	drawRect(base, 3, 10, 12, 12, "sand")

	gridfiles.Write(base)
}

func drawCross(g *gridfiles.Grid, startRow int, startCol int, stopRow int, stopCol int, tex string) {
	drawRect(g, startRow, startCol, stopRow, stopCol, tex)
	drawRect(g, startCol, startRow, stopCol, stopRow, tex)
}

func drawRect(g *gridfiles.Grid, startRow int, startCol int, stopRow int, stopCol int, tex string) {
	for _, cell := range g.CellIterator() {
		g.Tiles[cell.Row][cell.Col].Texture = tex
	}
}

func Generate(cfg *GenerationConfig) {
	base := &gridfiles.Grid{}
	base.RowCount = cfg.RowCount
	base.ColCount = cfg.ColCount
	base.Build()

	perlinInit(cfg.Seed)

	for _, pos := range base.CellIterator() {
		baseTexture := getTileBase(pos.Row, pos.Col)

		if !cfg.ReticulateSplines {
			baseTexture += ".png"
		}

		base.Tiles[pos.Row][pos.Col].Texture = baseTexture

	}

	toWrite := base

	if cfg.ReticulateSplines {
		toWrite = texturedGrid(base)
	}

	generateScenarioButtonLamp(toWrite)
	log.Infof("ents: %v", base.Entities)

	gridfiles.Write(toWrite)
}

func texturedGrid(base *gridfiles.Grid) *gridfiles.Grid {
	texgrid := &gridfiles.Grid{}
	texgrid.RowCount = base.RowCount
	texgrid.ColCount = base.ColCount
	texgrid.Build()

	for _, pos := range base.CellIterator() {
		tex, rot, flipH, flipV := reticulateSplines(base, texgrid, pos.Row, pos.Col)

		texgrid.Tiles[pos.Row][pos.Col].Texture = tex
		texgrid.Tiles[pos.Row][pos.Col].Rot = rot
		texgrid.Tiles[pos.Row][pos.Col].FlipH = flipH
		texgrid.Tiles[pos.Row][pos.Col].FlipV = flipV
	}

	return texgrid
}

func reticulateSplines(base *gridfiles.Grid, textured *gridfiles.Grid, row uint32, col uint32) (string, int32, bool, bool) {
	selfTile := base.Tiles[row][col]
	self := currentBiome.layers[selfTile.Texture]
	log.Infof("RS Tex: %+v \n", selfTile)

	// Dont reticulate the map edges
	//if (row == 0 || row == base.RowCount - 1) { return selfTile.Texture + ".png", 0 }
	//if (col == 0 || col == base.ColCount - 1) { return selfTile.Texture + ".png", 0 }

	tex, rot, flipH, flipV := matchRulesToTex(row, col, base, self)

	if tex == "missing-step-up" {
		log.Errorf("Missing step up texture")
		return "missingTextureStepUp.png", 0, false, false

	} else if tex == "missing-step-down" {
		log.Errorf("Missing step down texture")
		return "missingTextureStepDown.png", 0, false, false
	} else if tex == "" {
		return selfTile.Texture + ".png", 0, flipH, flipV
	}


	return tex, rot, flipH, flipV
}

func getTileLevel(base *gridfiles.Grid, selfDefault int, row uint32, col uint32) int {
	if row < 0 || row >= base.RowCount || col < 0 || col >= base.ColCount {
		return selfDefault
	}

	tileLayer := currentBiome.layers[base.Tiles[row][col].Texture]

	if tileLayer == nil {
		return selfDefault
	} else {
		return tileLayer.level
	}
}

func write(g *gridfiles.Grid) {	
	sg := &SerializableGrid{
		RowCount: g.RowCount,
		ColCount: g.ColCount,
	}

	for row := 0; row < g.RowCount; row++ {
		for col := 0; col < g.ColCount; col++ {
			sg.Tiles = append(sg.Tiles, g.Tiles[row][col])
		}
	}


	yml, err := yaml.Marshal(sg)
	
	if err != nil {
		log.Errorf("%v", err)
	}

	err = ioutil.WriteFile("../server/dat/worlds/gen/grids/0.grid", yml, 0644)

	if err != nil {
		log.Errorf("%v", err)
	}

}

func printGrid(g *gridfiles.Grid) {
	for _, t := range g.Tiles {
		fmt.Printf("%v", t)
	}
}

func getTileBase(row uint32, col uint32) string {
	v := perlin2(float64(row), float64(col))

	for _, layer := range currentBiome.layers {
		if v >= layer.startAt && v <= layer.stopAt {
			return layer.base
		}
	}

	log.Warnf("getTileBase failed to find biome layer for: %v", v)
	return "grass"
}

