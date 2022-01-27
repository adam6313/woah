package enum

// Gender -
type Gender uint32

const (
	// GenderNone -
	GenderNone Gender = iota

	// GenderMale -
	GenderMale

	// GenderFemale -
	GenderFemale
)

// GenderNoneToName -
var GenderNoneToName = map[uint32]string{
	0: "NONE",
	1: "MALE",
	2: "FEMALE",
}

// GenderNoneToValue -
var GenderNoneToValue = map[string]uint32{
	"NONE":   0,
	"MALE":   1,
	"FEMALE": 2,
}
