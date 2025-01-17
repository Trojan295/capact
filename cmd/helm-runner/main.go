package main

import (
	"log"

	"capact.io/capact/pkg/runner"
	"capact.io/capact/pkg/runner/helm"
	statusreporter "capact.io/capact/pkg/runner/status-reporter"

	"github.com/vrischmann/envconfig"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func main() {
	var cfg helm.Config
	err := envconfig.InitWithPrefix(&cfg, "RUNNER")
	exitOnError(err, "while loading configuration")

	stop := signals.SetupSignalHandler()

	k8sCfg, err := config.GetConfig()
	exitOnError(err, "while creating k8s config")

	helmRunner := helm.NewRunner(k8sCfg, cfg)

	statusReporter := statusreporter.NewNoop()

	// create and run manager
	mgr, err := runner.NewManager(helmRunner, statusReporter)
	exitOnError(err, "while creating runner manager")

	err = mgr.Execute(stop)
	exitOnError(err, "while executing runner")
}

func exitOnError(err error, context string) {
	if err != nil {
		log.Fatalf("%s: %v", context, err)
	}
}
