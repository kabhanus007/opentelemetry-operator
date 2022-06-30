package main

import (
	"flag"
	"github.com/open-telemetry/opentelemetry-operator/internal/experimental/examples/agent/agent"
	"log"
	"os"
	"os/signal"
)

func main() {
	var agentType string
	flag.StringVar(&agentType, "t", "io.opentelemetry.collector", "Agent Type String")

	var agentVersion string
	flag.StringVar(&agentVersion, "v", "1.0.0", "Agent Version String")

	flag.Parse()

	agent := agent.NewAgent(&agent.Logger{log.Default()}, agentType, agentVersion)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	agent.Shutdown()
}
