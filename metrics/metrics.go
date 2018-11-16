package metrics

import (
	"errors"
	"fmt"
	"net/http"

	zipkin "gopkg.in/spacemonkeygo/monkit-zipkin.v2"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"
	"gopkg.in/spacemonkeygo/monkit.v2/environment"
	"gopkg.in/spacemonkeygo/monkit.v2/present"
)

// Setup registers the service environment into the  monkit default registry ...
func Setup() {
	environment.Register(monkit.Default)
}

// Start takes the host and port for the metrics server to listen on ...
func Start(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Service listening at: %s\n", addr)

	// Start the metrics server on a go routine (non-blocking async)
	// Endless allows for 0-downtime updates/delpoyments
	go func() {
		http.ListenAndServe(addr, present.HTTP(monkit.Default))
	}()
}

// RegisterZipkin ...
func RegisterZipkin(host string, port int) error {
	// Check host and port
	if host == "" || port == 0 {
		return errors.New("Host and Port must have a valid value")
	}

	// Define a zipkin collector using the given host and port
	addr := fmt.Sprintf("%s:%d", host, port)
	collector, err := zipkin.NewScribeCollector(addr)
	if err != nil {
		return err
	}

	// Register the zipkin collector with the zipkin instance and monkit registry
	zipkin.RegisterZipkin(monkit.Default, collector, zipkin.Options{
		Fraction: 1,
	})

	return nil
}
