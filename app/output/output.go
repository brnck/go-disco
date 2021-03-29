package output

type Output interface {
	SetLed(i int, color uint32)
	Render() error
	Open() error
	Close()
}
