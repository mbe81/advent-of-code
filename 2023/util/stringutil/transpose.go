package stringutil

func Transpose(input []string) []string {
	var output []string
	for x := 0; x < len(input[0]); x++ {
		row := ""
		for y := 0; y < len(input); y++ {
			row += string(input[y][x])
		}
		output = append(output, row)
	}

	return output
}
