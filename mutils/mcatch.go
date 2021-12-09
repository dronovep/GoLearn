package mutils

func mcatch() {
	message := recover()
	if message != nil {
		println("Беда:", message)
	}
}
