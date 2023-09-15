package models

// RegModels models
var RegModels []interface{}

func RegistModel(model interface{}) {
	RegModels = append(RegModels, model)
}

// GetModels get models
func GetModels() []interface{} {
	return RegModels
}
