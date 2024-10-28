package output

type DatabasePort[Data any] interface {
	Find(collectionOrTable string, where map[string]interface{}) ([]Data, error)
	Insert(collectionOrTable string, data Data) (interface{}, error)
}
