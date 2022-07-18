package gridgen

import (
	"github.com/aquilax/go-perlin"
	"fmt"
//	"time"
)

var p *perlin.Perlin;

const (
	alpha = 2.
	beta = 2.
	numberOfIterations = int32(3)
)

func init() {
//	seed := int64(time.Now().UnixNano())
	seed := int64(1337)
	p = perlin.NewPerlin(alpha, beta, numberOfIterations, seed)
}

func perlin2(x float64, y float64) float64 {	
	v := p.Noise2D(x/10, y/10)

	fmt.Printf("perlin2: %v:%v %0.4f\n", x, y, v)

	return v
}
