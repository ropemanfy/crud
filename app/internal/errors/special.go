package errors

var (
	EmptyPeriodData = new("Отсутствуют данные за выбранный период", "internal")
	TooManyDetails  = new("Слишком много данных", "internal")
)
