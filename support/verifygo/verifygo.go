package main

func main() {
	const threads = 1000

	verifyJsonFrom(threads)
	verifyJsonMerge(threads)
	verifyProtoFrom(threads)
	verifyProtoMerge(threads)
	verifyEnum()
}
