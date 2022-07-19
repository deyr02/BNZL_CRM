package customerror

type FieldExist struct{}

func (m *FieldExist) Error() string {
	return "Field already exists"
}

type SystemField struct{}

func (m *SystemField) Error() string {
	return "System Field can not be modified"
}

type NewSystemFiled struct{}

func (m *NewSystemFiled) Error() string {
	return "Can not be added as a system field"
}

type NoRecordFound struct{}

func (m *NoRecordFound) Error() string {

	return "NO record found"
}
