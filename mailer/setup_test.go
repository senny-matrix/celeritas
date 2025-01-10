package mailer

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"os"
	"testing"
	"time"
)

var pool *dockertest.Pool
var resource *dockertest.Resource

var mailer = Mail{
	Domain:      "localhost",
	Templates:   "./testdata/mail",
	Host:        "localhost",
	Port:        1026,
	Encryption:  "none",
	FromAddress: "rogers@luso.solutions",
	FromName:    "Rogers",
	Jobs:        make(chan Message, 1),
	Results:     make(chan Result, 1),
}

func TestMain(m *testing.M) {
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalln("Could not connect to docker: " + err.Error())
	}
	pool = p

	opts := dockertest.RunOptions{
		Repository:   "mailhog/mailhog",
		Tag:          "latest",
		Env:          []string{},
		ExposedPorts: []string{"1025/tcp", "8025/tcp"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"1025/tcp": {{HostIP: "0.0.0.0", HostPort: "1026"}},
			"8025/tcp": {{HostIP: "0.0.0.0", HostPort: "8026"}},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Println("Could not start resource: " + err.Error())
		_ = pool.Purge(resource)
		log.Fatalln("Could not start resource: " + err.Error())
	}

	time.Sleep(2 * time.Second)

	go mailer.ListenForMail()

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalln("Could not purge resource: " + err.Error())
	}
	os.Exit(code)
}
