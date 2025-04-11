package conversion

import "strconv"

func StringsToFloats(strings []string) ([]float64, error) {
	//convert strings to float64
	floats := []float64{}
	for _, str := range strings {
		float, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, float)
	}
	return floats, nil
}
