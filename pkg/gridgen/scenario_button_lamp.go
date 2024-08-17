package gridgen

import (
	"github.com/greyvar/datlib/gridfiles"
	"math/rand"
	log "github.com/sirupsen/logrus"
)

func generateScenarioButtonLamp(g *gridfiles.Grid) {
	for i := 0; i < 10; i++ {
		t :=  getRandomTile(g, "sand.png")

		ei := &gridfiles.GridFileEntityInstance{
			Row: t.Row,
			Col: t.Col,
			Definition: "pressureButton",
		}

		g.Entities = append(g.Entities, *ei)
	}
}

func getRandomTile(g *gridfiles.Grid, texture string) *gridfiles.Tile {
	for {
		randRow := uint32(rand.Intn(int(g.RowCount)))
		randCol := uint32(rand.Intn(int(g.ColCount)))

		tile := g.Tiles[randRow][randCol]

		if tile.Texture == texture {
			return tile;
		} else {
			log.Infof("Random tile is no good %v %v", tile.Texture, texture)
			continue
		}
	}
}
