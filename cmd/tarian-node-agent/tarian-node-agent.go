package main

import "C"

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kube-tarian/tarian/pkg/logger"
	"github.com/kube-tarian/tarian/pkg/nodeagent"
	cli "github.com/urfave/cli/v2"

	_ "embed"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = "main"
)

func main() {
	app := getCliApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getCliApp() *cli.App {
	return &cli.App{
		Name:    "Tarian Node Agent",
		Usage:   "The Tarian Node Agent is the component which runs as a daemonset, detecting new processes from all containers in the node.",
		Version: version + " (" + commit + ")",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "log-level",
				Usage: "Log level: debug, info, warn, error",
				Value: "info",
			},
			&cli.StringFlag{
				Name:  "log-encoding",
				Usage: "log-encoding: json, console",
				Value: "console",
			},
		},
		DefaultCommand: "run",
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "Run the node agent",
				Action: run,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "cluster-agent-host",
						Usage: "Host address of the cluster agent to communicate with",
						Value: "tarian-cluster-agent.tarian-system.svc",
					},
					&cli.StringFlag{
						Name:  "cluster-agent-port",
						Usage: "Host port of the cluster agent to communicate with",
						Value: "80",
					},
					&cli.StringFlag{
						Name:  "node-name",
						Usage: "Node name where it is running. This is intended to be set from Downward API",
						Value: "",
					},
					&cli.BoolFlag{
						Name:  "enable-add-constraint",
						Usage: "Enable add constraint RPC. Enable this to allow register mode.",
						Value: false,
					},
				},
			},
		},
	}
}

func run(c *cli.Context) error {
	logger := logger.GetLogger(c.String("log-level"), c.String("log-encoding"))
	nodeagent.SetLogger(logger)

	if !isDebugFsMounted() {
		logger.Infow("debugfs is not mounted, will try to mount")

		err := mountDebugFs()
		if err != nil {
			logger.Fatal(err)
		}

		logger.Infow("successfully mounted debugfs", "path", DebugFSRoot)
	}

	// Check host proc dir
	_, err := os.Stat(nodeagent.HostProcDir)
	if err == nil {
		logger.Infow("Host proc is mounted", "dir", nodeagent.HostProcDir)
	} else if os.IsNotExist(err) {
		logger.Fatalw("Host proc directory is not properly mounted", "dir", nodeagent.HostProcDir)
	}

	agent := nodeagent.NewNodeAgent(c.String("cluster-agent-host") + ":" + c.String("cluster-agent-port"))
	agent.EnableAddConstraint(c.Bool("enable-add-constraint"))
	agent.SetNodeName(c.String("node-name"))

	logger.Infow("tarian-node-agent is running", "node-name", c.String("node-name"))
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		logger.Infow("got sigterm signal, attempting graceful shutdown", "signal", sig)

		agent.GracefulStop()
	}()

	agent.Run()
	logger.Info("tarian-node-agent shutdown gracefully")

	return nil
}
