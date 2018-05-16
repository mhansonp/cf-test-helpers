package helpers

import (
	"github.com/mhansonp/cf-test-helpers/commandstarter"
	"github.com/mhansonp/cf-test-helpers/helpers/internal"
	"github.com/onsi/gomega/gexec"
)

func Curl(cfg helpersinternal.CurlConfig, args ...string) *gexec.Session {
	cmdStarter := commandstarter.NewCommandStarter()
	return helpersinternal.Curl(cmdStarter, cfg.GetSkipSSLValidation(), args...)
}

func CurlSkipSSL(skip bool, args ...string) *gexec.Session {
	cmdStarter := commandstarter.NewCommandStarter()
	return helpersinternal.Curl(cmdStarter, skip, args...)
}
