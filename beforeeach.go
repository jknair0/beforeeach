package beforeeach

type BeforeEach func(func())

func Create(setUp func(), tearDown func()) BeforeEach {
	return func(testFunc func()) {
		setUp()
		testFunc()
		tearDown()
	}
}