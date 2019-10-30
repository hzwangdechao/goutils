// Code generated by GCG. DO NOT EDIT.
// Go Code Generator 0.0.7 (https://github.com/OhYee/gcg)

package fp

import ()

func MapString(f func(string) string, input []string) (output []string) {
	output = make([]string, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapInt(f func(int) int, input []int) (output []int) {
	output = make([]int, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapInt8(f func(int8) int8, input []int8) (output []int8) {
	output = make([]int8, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapInt16(f func(int16) int16, input []int16) (output []int16) {
	output = make([]int16, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapInt32(f func(int32) int32, input []int32) (output []int32) {
	output = make([]int32, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapInt64(f func(int64) int64, input []int64) (output []int64) {
	output = make([]int64, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapUint8(f func(uint8) uint8, input []uint8) (output []uint8) {
	output = make([]uint8, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapUint16(f func(uint16) uint16, input []uint16) (output []uint16) {
	output = make([]uint16, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapUint32(f func(uint32) uint32, input []uint32) (output []uint32) {
	output = make([]uint32, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapUint64(f func(uint64) uint64, input []uint64) (output []uint64) {
	output = make([]uint64, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapFloat32(f func(float32) float32, input []float32) (output []float32) {
	output = make([]float32, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapFloat64(f func(float64) float64, input []float64) (output []float64) {
	output = make([]float64, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func MapByte(f func(byte) byte, input []byte) (output []byte) {
	output = make([]byte, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}
func Map(f func(any) any, input []any) (output []any) {
	output = make([]any, 0)
	for _, data := range input {
		output = append(output, f(data))
	}
	return
}