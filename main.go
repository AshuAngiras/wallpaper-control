package main


import (
	"os"
	"log"
	"time"
	"./wallpaper"
	"flag"
	"math/rand"
	"github.com/kardianos/service"

)

var pixabayAPIKey string
var logger service.Logger

// Program structures.
//  Define Start and Stop methods.
type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() error {
	// Run the Service Forever
	index := rand.Intn(10)
	for {
		err := wallpaper.DownloadAndSet(index, pixabayAPIKey)
		if err != nil {
			log.Panic("Cannot Set the wallaper ERROR : ", err)
			os.Exit(1)
		}
		time.Sleep(time.Second * 20)
		index++
	}
}
func (p *program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}

// Service setup.
//   Define service config.
//   Create the service.
//   Setup the logger.
//   Handle service controls (optional).
//   Run the service.
func main() {
	svcFlag  := flag.String("service", "", "Control the system service.")
	APIKey 	 := flag.String("pixabay-key", "--sample-pixabay--key---", "Pixabay API Key")
	flag.Parse()

	pixabayAPIKey = *APIKey

	svcConfig := &service.Config{
		Name:        "Wallpaper Control Service",
		DisplayName: "Service to Control Wallpapers",
		Description: "Service to Control Wallpapers",
	}

	log.Println("Starting the Wallpaper Change Service")
	
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
