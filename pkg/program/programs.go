package program

import (
	"github.com/brnck/go-disco/pkg/output"
	"sync"
)

const (
	IterationsDefault     = 50
	WaitTimeDefault       = 100
	RedDefault            = 128
	GreenDefault          = 128
	BlueDefault           = 128
	SpeedDefault          = 100
	RandomDecayDefault    = true
	TrailDecayDefault     = 64
	SizeDefault           = 5
	OneLedPerSceneDefault = true
	HeatKeyDefault        = 0
	CoolingDefault        = 75
	SparkingDefault       = 150
)

type Config struct {
	Key            string `mapstructure:"key"`
	Start          int    `mapstructure:"start"`
	End            int    `mapstructure:"end"`
	Iterations     int    `mapstructure:"iterations"`
	WaitTime       int    `mapstructure:"wait_time"`
	Speed          int    `mapstructure:"speed"`
	Red            int    `mapstructure:"red"`
	Green          int    `mapstructure:"green"`
	Blue           int    `mapstructure:"blue"`
	Map            []Map  `mapstructure:"map"`
	RandomDecay    bool   `mapstructure:"random_decay"`
	TrailDecay     int    `mapstructure:"trail_decay"`
	Size           int    `mapstructure:"size"`
	OneLedPerScene bool   `mapstructure:"one_led_per_scene"`
	HeatKey        int    `mapstructure:"heat_key"`
	Cooling        int    `mapstructure:"cooling"`
	Sparking       int    `mapstructure:"sparking"`
}

type Map struct {
	Min int `mapstructure:"min"`
	Max int `mapstructure:"max"`
}

type Programs struct {
	programs map[string]Program
}

type Program interface {
	getName() string
	Run(output output.Output, c Config, wg *sync.WaitGroup)
}

func New() (*Programs, error) {
	p := &Programs{
		programs: make(map[string]Program),
	}
	if err := registerPrograms(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Programs) AddProgram(program Program) error {
	if _, exists := p.programs[program.getName()]; exists {
		return ErrProgramExists
	}

	p.programs[program.getName()] = program

	return nil
}

func (p *Programs) GetProgram(name string) (Program, error) {
	if _, exists := p.programs[name]; !exists {
		return nil, ErrProgramNotExists
	}

	return p.programs[name], nil
}

func registerPrograms(p *Programs) error {
	if err := p.AddProgram(NewTheaterChase()); err != nil {
		return err
	}
	if err := p.AddProgram(NewTheaterRainbowChase()); err != nil {
		return err
	}
	if err := p.AddProgram(NewRangedRainbowCycle()); err != nil {
		return err
	}
	if err := p.AddProgram(NewStrobe()); err != nil {
		return err
	}
	if err := p.AddProgram(NewMapper()); err != nil {
		return err
	}
	if err := p.AddProgram(NewChaosColors()); err != nil {
		return err
	}
	if err := p.AddProgram(NewChaosFillDown()); err != nil {
		return err
	}
	if err := p.AddProgram(NewColorChase()); err != nil {
		return err
	}
	if err := p.AddProgram(NewColorChaseReverse()); err != nil {
		return err
	}
	if err := p.AddProgram(NewMeteorRain()); err != nil {
		return err
	}
	if err := p.AddProgram(NewStaticRainbow()); err != nil {
		return err
	}
	if err := p.AddProgram(NewSparkle()); err != nil {
		return err
	}
	if err := p.AddProgram(NewSnowSparkle()); err != nil {
		return err
	}
	if err := p.AddProgram(NewStaticColor()); err != nil {
		return err
	}
	if err := p.AddProgram(NewRunningLightsReverse()); err != nil {
		return err
	}
	if err := p.AddProgram(NewRunningLights()); err != nil {
		return err
	}
	if err := p.AddProgram(NewTwinkle()); err != nil {
		return err
	}
	if err := p.AddProgram(NewTwinkleRandom()); err != nil {
		return err
	}
	if err := p.AddProgram(NewFire()); err != nil {
		return err
	}
	if err := p.AddProgram(NewCenterToOutside()); err != nil {
		return err
	}
	if err := p.AddProgram(NewOutsideToCenter()); err != nil {
		return err
	}
	if err := p.AddProgram(NewRightToLeft()); err != nil {
		return err
	}
	if err := p.AddProgram(NewLeftToRight()); err != nil {
		return err
	}

	return nil
}
