package gridgen;

type Biome struct {
	layers map[string]*BiomeLayer
}

type BiomeLayer struct {
	base string
	level int
	stepUpTexture string
	stepUpConvexCornerTexture string
	stepUpConcaveCornerTexture string
	stepUpBothSidesTexture string
	stepUpEndCapTexture string
	stepUpThinBendTexture string
	stepUpConcaveCornerAndSideTexture string
	stepUpConcaveCornersOppositeTexture string
	stepUpSurroundedTexture string
	stepUpOutcropTexture string
	stepUpOutcropThinTexture string
	stepUpCrossroadsTexture string
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
			alternatives: []string { "water.png" },
		},
		"sand": &BiomeLayer {
			base: "sand",
			level: 1, 
			startAt: 0,
			stopAt: .3,
			stepUpTexture: "sandWater.png",
			stepUpConvexCornerTexture: "sandWaterConvexCorner.png",
			stepUpConcaveCornerTexture: "sandWaterConcaveCorner.png",
			stepUpBothSidesTexture: "sandWaterBothSides.png",
			stepUpEndCapTexture: "sandWaterEndCap.png",
			stepUpThinBendTexture: "sandWaterThinBend.png",
			stepUpConcaveCornerAndSideTexture: "sandWaterConcaveCornerAndSide.png",
			stepUpConcaveCornersOppositeTexture: "sandWaterConcaveCornersOpposite.png",
			stepUpSurroundedTexture: "sandWaterSurrounded.png",
			stepUpOutcropTexture: "sandWaterOutcrop.png",
			stepUpOutcropThinTexture: "sandWaterOutcropThin.png",
			stepUpCrossroadsTexture: "sandWaterCrossroads.png",
		},
		"grass": &BiomeLayer {
			base: "grass",
			level: 2, 
			startAt: .3,
			stopAt: 1,
			stepUpTexture: "grassSand.png",
			stepUpConvexCornerTexture: "grassSandConvexCorner.png",
			stepUpConcaveCornerTexture: "grassSandConcaveCorner.png",
			stepUpEndCapTexture: "grassSandEndCap.png",
			stepUpConcaveCornerAndSideTexture: "grassSandConcaveCornerAndSide.png",
			stepUpConcaveCornersOppositeTexture: "grassSandConcaveCornersOpposite.png",
			stepUpThinBendTexture: "grassSandThinBend.png",
			stepUpBothSidesTexture: "grassSandBothSides.png",
			stepUpSurroundedTexture: "grassSandSurrounded.png",
			stepUpOutcropTexture: "grassSandOutcrop.png",
		},
	}, 
}


