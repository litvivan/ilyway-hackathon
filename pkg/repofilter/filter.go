package repofilter

import "encoding/json"

type Filter struct {
	fieldsMap map[string]interface{}
}

func NewFromMap(m map[string]interface{}) (Filter, error) {
	return Filter{
		fieldsMap: m,
	}, nil
}

func (f Filter) Make(s interface{}) (interface{}, error) {
	filter := s

	jsonFields, err := json.Marshal(f.fieldsMap)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonFields, &filter)
	if err != nil {
		return nil, err
	}

	return filter, nil
}
