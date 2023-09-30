package support

type MapValueTypes interface {
	~float32 | ~uint16 | ~uint32
}

// MapValue Функция пропорционально переносит значение (value) из текущего диапазона значений
// (fromMin .. fromMax) в новый диапазон (toMin .. toMax), заданный параметрами.
// Поддерживаемые типы см. MapValueTypes
func MapValue[T MapValueTypes, R MapValueTypes](value uint16, fromMin, fromMax T, toMin, toMax R) R {
	floatValue := T(value)
	if floatValue < fromMin {
		return toMin
	}
	if floatValue > fromMax {
		return toMax
	}
	ratio := float32(floatValue-fromMin) / float32(fromMax-fromMin)
	return R(ratio*float32(toMax-toMin)) + toMin
}
