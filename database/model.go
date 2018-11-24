package database

type Model interface {
	ToString() string
	GetID() uint
	GetTypeInString() string
}

func ConvertToModels(arg interface{}) []Model {

	var models []Model

	switch values := arg.(type) {

	case []Category:

		for _, category := range values {
			models = append(models, category)
		}

	case []Wallet:

		for _, wallet := range values {
			models = append(models, wallet)
		}

	case []Source:

		for _, source := range values {
			models = append(models, source)
		}

	default:

	}

	return models
}
