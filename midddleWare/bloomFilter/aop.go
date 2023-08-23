package bloomFilter

func before() {

}

func excuteWithAOP(f func()) {

	before()

	f()
}
