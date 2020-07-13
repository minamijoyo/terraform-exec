package tfexec

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/hashicorp/go-version"
)

type Terraform struct {
	execPath   string
	workingDir string
	env        []string
	logger     *log.Logger

	versionLock  sync.Mutex
	execVersion  *version.Version
	provVersions map[string]*version.Version
}

// NewTerraform returns a Terraform struct with default values for all fields.
// If a blank execPath is supplied, NewTerraform will attempt to locate an
// appropriate binary on the system PATH.
func NewTerraform(workingDir string, execPath string) (*Terraform, error) {
	if workingDir == "" {
		return nil, fmt.Errorf("Terraform cannot be initialised with empty workdir")
	}

	if _, err := os.Stat(workingDir); err != nil {
		return nil, fmt.Errorf("error initialising Terraform with workdir %s: %s", workingDir, err)
	}

	if execPath == "" {
		err := fmt.Errorf("NewTerraform: please supply the path to a Terraform executable using execPath, e.g. using the tfinstall package.")
		return nil, &ErrNoSuitableBinary{err: err}

	}
	tf := Terraform{
		execPath:   execPath,
		workingDir: workingDir,
		env:        os.Environ(),
		logger:     log.New(ioutil.Discard, "", 0),
	}

	return &tf, nil
}

func (tf *Terraform) SetEnv(env map[string]string) {
	var tfenv []string

	// always propagate CHECKPOINT_DISABLE env var unless it is
	// explicitly overridden with tf.SetEnv
	if _, ok := env["CHECKPOINT_DISABLE"]; !ok {
		env["CHECKPOINT_DISABLE"] = os.Getenv("CHECKPOINT_DISABLE")
	}

	for k, v := range env {
		tfenv = append(tfenv, k+"="+v)
	}

	tf.env = tfenv
}

func (tf *Terraform) SetLogger(logger *log.Logger) {
	tf.logger = logger
}
