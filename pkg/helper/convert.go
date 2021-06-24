package helper

func CreateFloatP32FromFloat(x float32) *float32 {
	f := float32(x)
	return &f
}
func CreateIntPFromInt(x int) *int {
	return &x
}
func CreateStringPFromString(x string) *string {
	return &x
}
