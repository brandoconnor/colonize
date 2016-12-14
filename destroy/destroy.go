package destroy

import (
	//	"bytes"
	//	//"fmt"
	//	"io/ioutil"
	"os"
	//	"regexp"
	//	"strings"
	//
	"github.com/craigmonson/colonize/config"
	"github.com/craigmonson/colonize/log"
	"github.com/craigmonson/colonize/util"
)

func Run(c *config.ColonizeConfig, l log.Logger, skipRemote bool, remoteAfterApply bool) error {
	os.Chdir(c.TmplPath)

	if skipRemote {
		l.Log("Skipping remote setup")
	} else {
		l.Log("Running remote setup")
		util.RunCmd("./" + c.CombinedRemoteFilePath)
	}

	l.Log("Executing terraform destroy")
	return util.RunCmd(
		"terraform",
		"destroy",
		"-force",
		"-var-file", c.CombinedValsFilePath,
		"-var-file", c.CombinedDerivedValsFilePath,
	)
}