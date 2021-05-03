package output

import "log"

type stdout struct {
	log *log.Logger
}

func (s stdout) SetLed(i int, color uint32) {
	s.log.Printf("Setting LED %v to Color %v", i, color)
}

func (s stdout) GetLedColor(i int) uint32 {
	s.log.Printf("Gettting LED %v color", i)

	return 255
}

func (s stdout) Render() error {
	s.log.Println("Initializing rendering")

	return nil
}

func (s stdout) Open() error {
	s.log.Println("Opening LED channel")

	return nil
}

func (s stdout) Close() {
	s.log.Println("Closing LED channel")
}

func newStdout(l *log.Logger) (*stdout, error) {
	return &stdout{log: l}, nil
}
