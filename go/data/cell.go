package data

type Cell struct {
	Live     bool
	Parent   *Universe
	Location Coordinate
}

type Coordinate struct {
	X int
	Y int
}
