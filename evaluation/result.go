package evaluation

type Result struct {
	Planets        []*ResultPlanet
	FrontPlanets   []*ResultPlanet
	SupportPlanets []*ResultPlanet
	FrontLimit     int16
}
