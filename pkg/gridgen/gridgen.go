package gridgen

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"io/ioutil"
	"gopkg.in/yaml.v3"
)

type Tile struct {
	Texture string
	Row int
	Col int
	Rot int
	FlipH bool `yaml:"flipH"`
	FlipV bool `yaml:"flipV"`
}

type Grid struct {
	name string
	Tiles map[int]map[int]*Tile
	RowCount int
	ColCount int
}

type SerializableGrid struct {
	Tiles []*Tile
	RowCount int
	ColCount int
}

func (g *Grid) build() {
	g.Tiles = make(map[int]map[int]*Tile)

	for row := 0; row < g.RowCount; row++ {
		g.Tiles[row] = make(map[int]*Tile)

		for col := 0; col < g.ColCount; col++ {
			t := &Tile{
				Row: row,
				Col: col,
				Texture: "water",
			}

			g.Tiles[row][col] = t
		}
	}
}

type position struct {
	row int
	col int
}

func (g *Grid) cellIterator() []*position {
	ret := make([]*position, 0)

	for row := 0; row < g.RowCount; row++ {
		for col := 0; col < g.ColCount; col++ {
			ret = append(ret, &position{
				row: row,
				col: col,
			})
		}
	}

	return ret
}

func GenerateBiomeReticulationTest() {

	base := &Grid{}
	base.name = "biome reticulation"
	base.RowCount = 16
	base.ColCount = 16
	base.build()

//	drawCross(base, 4, 1, 12, 15, "sand.png")
//	drawCross(base, 6, 3, 10, 13, "grass")

	drawRect(base, 5, 5, 6, 14, "sand")
	drawRect(base, 3, 7, 12, 8, "sand")
	drawRect(base, 3, 10, 12, 12, "sand")

	write(base)
	write(texturedGrid(base))
}

func drawCross(g *Grid, startRow int, startCol int, stopRow int, stopCol int, tex string) {
	drawRect(g, startRow, startCol, stopRow, stopCol, tex)
	drawRect(g, startCol, startRow, stopCol, stopRow, tex)
}

func drawRect(g *Grid, startRow int, startCol int, stopRow int, stopCol int, tex string) {
	for row := startRow; row < stopRow; row++ {
		for col := startCol; col < stopCol; col++ {
			g.Tiles[row][col].Texture = tex
		}
	}
}

func Generate(cfg *GenerationConfig) {
	baseGrid := &Grid{}
	baseGrid.name = "biome"
	baseGrid.RowCount = cfg.RowCount
	baseGrid.ColCount = cfg.ColCount
	baseGrid.build()

	perlinInit(cfg.Seed)

	for _, pos := range baseGrid.cellIterator() {
		baseTexture := getTileBase(pos.row, pos.col)

		if !cfg.ReticulateSplines {
			baseTexture += ".png"
		}

		baseGrid.Tiles[pos.row][pos.col].Texture = baseTexture

	}

	generateScenarioButtonLamp(baseGrid)

	if cfg.ReticulateSplines {
		write(texturedGrid(baseGrid))
	} else {
		write(baseGrid)
	}
}

func texturedGrid(baseGrid *Grid) *Grid {
	texGrid := &Grid{}
	texGrid.RowCount = baseGrid.RowCount
	texGrid.ColCount = baseGrid.ColCount
	texGrid.build()

	for _, pos := range baseGrid.cellIterator() {
		tex, rot, flipH, flipV := reticulateSplines(baseGrid, texGrid, pos.row, pos.col)

		texGrid.Tiles[pos.row][pos.col].Texture = tex
		texGrid.Tiles[pos.row][pos.col].Rot = rot
		texGrid.Tiles[pos.row][pos.col].FlipH = flipH
		texGrid.Tiles[pos.row][pos.col].FlipV = flipV
	}

	return texGrid
}

func reticulateSplines(base *Grid, textured *Grid, row int, col int) (string, int, bool, bool) {
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

func getTileLevel(base *Grid, selfDefault int, row int, col int) int {
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

func write(g *Grid) {	
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

func printGrid(g *Grid) {
	for _, t := range g.Tiles {
		fmt.Printf("%v", t)
	}
}

func getTileBase(row int, col int) string {
	v := perlin2(float64(row), float64(col))

	for _, layer := range currentBiome.layers {
		if v >= layer.startAt && v <= layer.stopAt {
			return layer.base
		}
	}

	log.Warnf("getTileBase failed to find biome layer for: %v", v)
	return "grass"
}

