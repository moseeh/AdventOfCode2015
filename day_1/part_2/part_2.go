package part2

func Firstchar(filestr string) int {
	count := 0
	for i, v := range filestr {
		if string(v) == "(" {
			count++
		} else {
			count--
		}
		if count == -1 {
			return i + 1
		}
	}
	return 0
}
