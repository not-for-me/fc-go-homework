package markovchain

const (
	rowNum = 26
	colNum = 26
)

// SourceModel for containing trained data
type SourceModel struct {
	name            string
	transitionProps [rowNum][colNum]float64
}

// New func is a constructor function
func New(name string, corpus []byte) *SourceModel {
	return &SourceModel{}
}
