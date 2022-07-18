package gridgen

import (
	"math"
	"math/rand"
)

type vector2 struct {
	x float64
	y float64
}

func interpolate(a0 float64, a1 float64, w float64) float64 {
    /* // You may want clamping by inserting:
     * if (0.0 > w) return a0;
     * if (1.0 < w) return a1;
     */
    return (a1 - a0) * w + a0;
    /* // Use this cubic interpolation [[Smoothstep]] instead, for a smooth appearance:
     * return (a1 - a0) * (3.0 - w * 2.0) * w * w + a0;
     *
     * // Use [[Smootherstep]] for an even smoother result with a second derivative equal to zero on boundaries:
     * return (a1 - a0) * ((w * (w * 6.0 - 15.0) + 10.0) * w * w * w) + a0;
     */
}

/* Create pseudorandom direction vector
 */
func randomGradient(ix int64, iy int64) *vector2 {
    // No precomputed gradients mean this works for any number of grid coordinates
    w := int64(8 * 123443);
    s := w / 2; // rotation width
    
	a := ix
	b := iy

    a *= 3284157443; b ^= a << s | a >> w-s;
    b *= 1911520717; a ^= b << s | b >> w-s;
    a *= 2048419325;

    random := rand.Float64() // in [0, 2*Pi]

    v := &vector2{};
    v.x = math.Cos(random)
	v.y = math.Sin(random);

    return v;
}

// Computes the dot product of the distance and gradient vectors.
func dotGridGradient(ix int64, iy int64, x float64, y float64) float64 {
    // Get gradient from integer coordinates
    gradient := randomGradient(ix, iy);

    // Compute the distance vector
    dx := x - float64(ix)
    dy := y - float64(iy)
	

	dotProduct := dx*gradient.x + dy*gradient.y


	return dotProduct
}

// Compute Perlin noise at coordinates x, y
func perlinNative(x float64, y float64) float64 {
    // Determine grid cell coordinates
    x0 := int64(math.Floor(x))
    x1 := x0 + 1;
    y0 := int64(math.Floor(y))
    y1 := y0 + 1;

    // Determine interpolation weights
    // Could also use higher order polynomial/s-curve here
    sx := x - float64(x0);
    //sy := y - float64(y0);


    // Interpolate between grid point gradients
    n0 := dotGridGradient(x0, y0, x, y);
    n1 := dotGridGradient(x1, y0, x, y);
    ix0 := interpolate(n0, n1, sx);

    n0 = dotGridGradient(x0, y1, x, y);
    n1 = dotGridGradient(x1, y1, x, y);
    ix1 := interpolate(n0, n1, sx);

    value := interpolate(ix0, ix1, 1);
    return value;
}
