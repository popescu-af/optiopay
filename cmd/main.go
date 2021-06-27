package main

import (
	"fmt"
	"net/http"

	"github.com/popescu-af/saas-y/pkg/log"

	"github.com/popescu-af/optiopay/services/main-svc/internal/config"
	"github.com/popescu-af/optiopay/services/main-svc/internal/logic"
	"github.com/popescu-af/optiopay/services/main-svc/internal/service"
)

func main() {
	defer log.Sync()

	log.Info("main-svc started")

	env, err := config.ProcessEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

	impl := logic.NewImpl()
	httpWrapper := service.NewHTTPWrapper(impl)
	router := service.NewRouter(httpWrapper.Paths())

	log.Fatal(fmt.Sprintf("error serving - %v", http.ListenAndServe(fmt.Sprintf(":%s", env.Port), router)))
}
