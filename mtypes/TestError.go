package mtypes

type TestError struct{}

func (o *TestError) Error() string {
	return "Тестовая ошибка"
}
