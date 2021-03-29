package stdout

import "log"

type Stdout struct {
	log *log.Logger
}

func (s *Stdout) SetLed(i int, color uint32) {
	s.log.Printf("Setting LED %v to Color %v", i, color)
}

func (s *Stdout) Render() error {
	s.log.Println("Initializing rendering")

	return nil
}

func (s *Stdout) Open() error {
	s.log.Println("Opening LED channel")

	return nil
}

func (s *Stdout) Close() {
	s.log.Println("Closing LED channel")
}

func Init(l *log.Logger) (*Stdout, error) {
	return &Stdout{log: l}, nil
}
