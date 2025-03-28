package app

// NumberResponse represents the response format for the /numbers endpoint
type NumberResponse struct {
	WindowPrevState []int   `json:"windowPrevState"`
	WindowCurrState []int   `json:"windowCurrState"`
	Numbers         []int   `json:"numbers"`
	Avg             float64 `json:"avg"`
}

// NumberType represents valid number type identifiers
type NumberType string

const (
	Prime     NumberType = "p"
	Fibonacci NumberType = "f"
	Even      NumberType = "e"
	Random    NumberType = "r"
)

// IsValid checks if the number type is valid
func (nt NumberType) IsValid() bool {
	switch nt {
	case Prime, Fibonacci, Even, Random:
		return true
	default:
		return false
	}
} 