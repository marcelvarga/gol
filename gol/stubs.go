package gol

var WorldEvolution = "GolOperations.Evolve"
var AliveCellsEvent = "GolOperations.ReportAliveCellsCount"
var SendFLip = "GolOperations.GetFLip"
var Save = "GolOperations.Save"
var Quit = "GolOperations.Quit"
var Shut = "GolOperations.Kill"
var Pause = "GolOperations.Pause"
var Resume = "GolOperations.Resume"

type Request struct {
	InitialWorld [][]byte
	P            Params
	Events       chan<- Event
}

type Result struct {
	OutputWorld [][]byte
}

type RequestAliveCells struct {
}

type ReportAliveCells struct {
	AliveCellsCountEv AliveCellsCount
}

type RequestCellFlip struct {
}
type GetCellFlip struct {
	Flip CellFlipped
}
type RequestForKey struct {
}
type ReceiveFromKey struct {
	ScreenshotWorld [][]byte
	Turn            int
}
