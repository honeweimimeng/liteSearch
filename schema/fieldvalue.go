package schema

type FieldID uint32

type FieldAndFieldValues struct {
	Field       FieldID
	FieldValues []*FieldValue
}

type FieldValue struct {
	FieldName string `json:"keyName"`
	FieldID   FieldID
	Value     interface{} `json:"value"`
}
