package reverse_string

func ReverseString(input string) (output string) {
	var arr []int32
	for _, i := range input {
		arr = append(arr, i)
	}
	len := len(arr)
	for j := len - 1; j >= 0; j-- {
		output = output + string(arr[j])
	}
	return output
}
