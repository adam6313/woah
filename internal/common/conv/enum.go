package conv

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

// ItemStatusEnum -	商品狀態
type ItemStatusEnum uint32

const (
	// 預設值
	ItemStatusNone ItemStatusEnum = iota

	// 上架
	ItemStatusOnSold

	// 下架
	ItemStatusOffSold

	// 範圍檢查點
	ItemStatusEnd
)

// Verify - 驗證是否符合範圍
func (s ItemStatusEnum) Verify() bool {
	return s > ItemStatusNone && s < ItemStatusEnd
}
