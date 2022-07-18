package gridgen

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"io/ioutil"
	"gopkg.in/yaml.v3"
)

type Biome struct {
	layers map[string]*BiomeLayer
}

type BiomeLayer struct {
	base string
	level int
	stepUpTexture string
	stepDownTexture string
	alternatives []string
	startAt float64
	stopAt float64
}

var currentBiome = &Biome {
	layers: map[string]*BiomeLayer {
		"water": &BiomeLayer {
			base: "water",
			level: 0,
			startAt: -1,
			stopAt: 0,
			stepDownTexture: "construct.png",
			stepUpTexture: "waterSand.png",
			alternatives: []string { "water.png" },
		},
		"sand": &BiomeLayer {
			base: "sand",
			level: 1, 
			startAt: 0,
			stopAt: .2,
			stepDownTexture: "waterSand.png",
			stepUpTexture: "sandGrass.png",
		},
		"grass": &BiomeLayer {
			base: "grass",
			level: 2, 
			startAt: .2,
			stopAt: 1,
			stepDownTexture: "grassSand.png",
			stepUpTexture: "construct.png",
		},
	},
}

type Tile struct {
	Texture string
	X int
	Y int
	Rot int
}

type Grid struct {
	name string
	Tiles map[int]map[int]*Tile
	rowCount int
	colCount int
}

type SerializableGrid struct {
	Tiles []*Tile
}

func (g *Grid) build() {
	g.Tiles = make(map[int]map[int]*Tile)

	for row := 0; row < g.rowCount; row++ {
		g.Tiles[row] = make(map[int]*Tile)

		for col := 0; col < g.colCount; col++ {

		}
	}
}


func Generate() {
	baseGrid := &Grid{}
	baseGrid.name = "biome"
	baseGrid.rowCount = 16
	baseGrid.colCount = 16
	baseGrid.build()

	for row := 0; row < baseGrid.rowCount; row++ {
		for col := 0; col < baseGrid.colCount; col++ {
			t := &Tile{}
			t.X = row
			t.Y = col
			t.Texture = getTileBase(row, col)

			baseGrid.Tiles[row][col] = t
		}
	}

	texGrid := &Grid{}
	texGrid.rowCount = baseGrid.rowCount
	texGrid.colCount = baseGrid.colCount
	texGrid.build()

	for row := 0; row < baseGrid.rowCount; row++ {
		for col := 0; col < baseGrid.colCount; col++ {
			t := &Tile{}
			t.X = row
			t.Y = col
			t.Texture, t.Rot = reticulateSplines(baseGrid, texGrid, col, row)

			texGrid.Tiles[row][col] = t
		}
	}

	write(texGrid)
}

func reticulateSplines(base *Grid, textured *Grid, row int, col int) (string, int) {
	selfTile := base.Tiles[row][col]
	self := currentBiome.layers[selfTile.Texture]
	
//	return self + ".png", 0

	// Dont reticulate the map edges
	if (row == 0 || row == base.rowCount - 1) { return selfTile.Texture + ".png", 0 }
	if (col == 0 || col == base.colCount - 1) { return selfTile.Texture + ".png", 0 }

	above := currentBiome.layers[base.Tiles[row - 1][col].Texture]
	left  := currentBiome.layers[base.Tiles[row][col - 1].Texture]
	right := currentBiome.layers[base.Tiles[row][col + 1].Texture]
	below := currentBiome.layers[base.Tiles[row + 1][col].Texture]

	if (above.level < self.level) {return self.stepDownTexture, 0 }
	if (above.level > self.level) {return self.stepDownTexture, 180 }

	log.WithFields(log.Fields{
		"above": above.level,
		"left": left.level, 
		"right": right.level,
		"below": below.level,
		"self": self.level,
	}).Warnf("Could not match a reticulation rule")

	return selfTile.Texture + ".png", 0
}

func write(g *Grid) {	
	sg := &SerializableGrid{}

	for row := 0; row < g.rowCount; row++ {
		for col := 0; col < g.colCount; col++ {
			sg.Tiles = append(sg.Tiles, g.Tiles[row][col])
		}
	}


	fmt.Println("%v", sg)

	yml, err := yaml.Marshal(sg)
	
	fmt.Println("%v", err)

	err = ioutil.WriteFile("/home/jread/sandbox/Development/greyvar/greyvar-server/dat/worlds/gen/grids/0.grid", yml, 0644)

	fmt.Println("%v", err)
}

func printGrid(g *Grid) {
	for _, t := range g.Tiles {
		fmt.Println("%0.0f", t)
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

