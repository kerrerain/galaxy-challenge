package base

import (
	"github.com/magleff/galaxy-challenge/models"
	"math"
)

func computeDistance(p1 models.Planet, p2 models.Planet) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}
