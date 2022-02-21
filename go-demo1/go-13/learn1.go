package go13

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15) //这里的celsius是类型转换
}

// 方法
// kelvin是接收者
// 一个方法只能有一个接收者
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15) //这里的celsius是类型转换
}

func main() {
}
