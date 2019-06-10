package Support

type Jsonable interface {
	ToJson() ([]byte, error)
}
