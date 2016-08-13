package cf

import (
	"github.com/cloudfoundry-incubator/cf-test-helpers/helpers"
	"github.com/onsi/gomega/gexec"
)

var Cf = func(args ...string) *gexec.Session {
	return helpers.Run("cf", args...)
}
