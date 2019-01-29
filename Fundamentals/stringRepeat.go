package main

func RepeatStr(repetitions int, value string) string {
	msg := ""
	for i := 0; i < repetitions; i++ {
		msg += value
	}
	return msg
}
