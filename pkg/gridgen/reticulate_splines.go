package gridgen;

import (
	log "github.com/sirupsen/logrus"
	"github.com/greyvar/datlib/gridfiles"
)

func matchRulesToTex(row uint32, col uint32, base *gridfiles.Grid, self *BiomeLayer) (string, int32, bool, bool) {
	n := getTileLevel(base, self.level, row - 1, col)
	e := getTileLevel(base, self.level, row, col + 1)
	w  := getTileLevel(base, self.level, row, col - 1)
	s := getTileLevel(base, self.level, row + 1, col)
	
	nw := getTileLevel(base, self.level, row - 1, col - 1)
	ne := getTileLevel(base, self.level, row - 1, col + 1)
	se := getTileLevel(base, self.level, row + 1, col + 1)
	sw := getTileLevel(base, self.level, row + 1, col - 1)

	// Don't reticulate when all surrounding cells are the same base
	if (n == self.level && w == self.level && e == self.level && s == self.level && nw == self.level && ne == self.level && sw == self.level && se == self.level) {
		return "", 0, false, false
	}

	if (n == self.level && ne < self.level && e == self.level && se < self.level && s == self.level && sw < self.level && w == self.level && nw < self.level) { return self.stepUpCrossroadsTexture, 0, false, false }

	if (n < self.level && w < self.level && e < self.level && s < self.level && nw < self.level && ne < self.level && sw < self.level && se < self.level) { return self.stepUpSurroundedTexture, 0, false, false }

	if (ne < self.level && nw < self.level && n == self.level && s == self.level && w == self.level && e == self.level) { return self.stepUpOutcropTexture, 0, false, false }
	if (ne < self.level && se < self.level && n == self.level && s == self.level && w == self.level && e == self.level) { return self.stepUpOutcropTexture, 90, false, false }
	if (se < self.level && sw < self.level && n == self.level && s == self.level && w == self.level && e == self.level) { return self.stepUpOutcropTexture, 180, false, false }
	if (nw < self.level && sw < self.level && n == self.level && s == self.level && w == self.level && e == self.level) { return self.stepUpOutcropTexture, 270, false, false }

	if (ne < self.level && nw < self.level && n == self.level && e == self.level && w == self.level && s < self.level) { return self.stepUpOutcropThinTexture, 0, true, false }
	if (ne < self.level && se < self.level && n == self.level && e == self.level && s == self.level && w < self.level) { return self.stepUpOutcropThinTexture, 90, true, false }

	if (w == self.level && e == self.level && s == self.level && n == self.level && nw < self.level && se < self.level) { return self.stepUpConcaveCornersOppositeTexture, 0, false, false }
	if (w == self.level && e == self.level && s == self.level && n == self.level && ne < self.level && sw < self.level) { return self.stepUpConcaveCornersOppositeTexture, 90, false, false }

	if (w == self.level && n == self.level && s == self.level && e < self.level && nw < self.level) { return self.stepUpConcaveCornerAndSideTexture, 0, false, false }
	if (w == self.level && n == self.level && e == self.level && s < self.level && ne < self.level) { return self.stepUpConcaveCornerAndSideTexture, 90, false, false }
	if (w < self.level && e == self.level && n == self.level && s == self.level && se < self.level) { return self.stepUpConcaveCornerAndSideTexture, 180, false, false }
	if (w == self.level && n < self.level && e == self.level && s == self.level && sw < self.level) { return self.stepUpConcaveCornerAndSideTexture, 270, false, false }

	if (w < self.level && e == self.level && n == self.level && s == self.level && ne < self.level) { return self.stepUpConcaveCornerAndSideTexture, 0, true, false }
	if (n < self.level && e == self.level && s == self.level && w == self.level && se < self.level) { return self.stepUpConcaveCornerAndSideTexture, 90, true, false }
	if (n == self.level && e < self.level && s == self.level && w == self.level && sw < self.level) { return self.stepUpConcaveCornerAndSideTexture, 180, true, false }
	if (n == self.level && e == self.level && s < self.level && w == self.level && nw < self.level) { return self.stepUpConcaveCornerAndSideTexture, 270, true, false }

	if (w < self.level && e < self.level) {
		if (n < self.level && s == self.level) { return self.stepUpEndCapTexture, 0, false, false }
		if (n == self.level && s == self.level) { return self.stepUpBothSidesTexture, 0, false, false }

		if (s < self.level && n == self.level) { return self.stepUpEndCapTexture, 180, false, false } 
//		if (s == self.level && n == self.level) { return self.stepUpBothSidesTexture, 180 } 
	}

	if (n < self.level && s < self.level) {
		if (e < self.level && w == self.level) { return self.stepUpEndCapTexture, 90, false, false }
		if (e == self.level && w == self.level) { return self.stepUpBothSidesTexture, 90, false, false }

		if (w < self.level && e == self.level) { return self.stepUpEndCapTexture, 270, false, false }
		if (w == self.level && e == self.level) { return self.stepUpBothSidesTexture, 270, false, false }
	}

	if (n < self.level && ne < self.level && e == self.level && se < self.level && s == self.level && w < self.level) { return self.stepUpThinBendTexture, 0, false, false }
	if (n < self.level && e < self.level && s == self.level && w == self.level && sw < self.level && ne < self.level) { return self.stepUpThinBendTexture, 90, false, false }
	if (n == self.level && e < self.level && s < self.level && w == self.level && se < self.level && nw < self.level) { return self.stepUpThinBendTexture, 180, false, false }
	if (n == self.level && e == self.level && s < self.level && w < self.level && sw < self.level && ne < self.level) { return self.stepUpThinBendTexture, 270, false, false }


	if (w < self.level && e < self.level && n == self.level && s == self.level) { return self.stepUpBothSidesTexture, 0, false, false }
	if (n < self.level && s < self.level && w == self.level && e == self.level) { return self.stepUpBothSidesTexture, 90, false, false }

	if (w == self.level && n == self.level && nw < self.level) { return self.stepUpConcaveCornerTexture, 0, false, false }
	if (e == self.level && n == self.level && ne < self.level) { return self.stepUpConcaveCornerTexture, 90, false, false }
	if (e == self.level && s == self.level && se < self.level) { return self.stepUpConcaveCornerTexture, 180, false, false }
	if (w == self.level && s == self.level && sw < self.level) { return self.stepUpConcaveCornerTexture, 270, false, false }

	if (n < self.level && w < self.level && e < self.level && s < self.level) { return self.stepUpSurroundedTexture, 0, false, false }

	if (w < self.level && n < self.level) { return self.stepUpConvexCornerTexture, 0, false, false }
	if (e < self.level && n < self.level) { return self.stepUpConvexCornerTexture, 90, false, false }
	if (e < self.level && s < self.level) { return self.stepUpConvexCornerTexture, 180, false, false }
	if (w < self.level && s < self.level) { return self.stepUpConvexCornerTexture, 270, false, false }

	if (n < self.level) {return self.stepUpTexture, 0, false, false }
	if (e < self.level) { return self.stepUpTexture, 90, false, false }
	if (s < self.level) {return self.stepUpTexture, 180, false, false }
	if (w < self.level) { return self.stepUpTexture, 270, false, false }
	

	log.WithFields(log.Fields{
		"row": row,
		"col": col, 
		"n": n,
		"w": w, 
		"e": e,
		"s": s,
		"ne": ne,
		"nw": nw,
		"se": se,
		"sw": sw,
		"self": self.level,
	}).Warnf("Could not match a reticulation rule")

	return "", 0, false, false
}

