package excel

type ExporterInterface interface {
	GetSheetName() string
	GetEntity() any
	GetProvider() func() ([]any, error)
}

type Exporter[T any] struct {
	SheetName string
	Entity    T
	Provider  func() ([]*T, error)
}

func (e Exporter[T]) GetSheetName() string {
	return e.SheetName
}

func (e Exporter[T]) GetEntity() any {
	return e.Entity
}

func (e Exporter[T]) GetProvider() func() ([]any, error) {
	return func() ([]any, error) {
		entities, err := e.Provider()
		if err != nil {
			return nil, err
		}
		result := make([]any, len(entities))
		for i, entity := range entities {
			result[i] = entity
		}
		return result, nil
	}
}
