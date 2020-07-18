package tfexec

import (
	"context"
	"os"
	"os/exec"
	"strings"
)

const (
	checkpointDisableEnvVar = "CHECKPOINT_DISABLE"
	logEnvVar               = "TF_LOG"
	inputEnvVar             = "TF_INPUT"
	automationEnvVar        = "TF_IN_AUTOMATION"
	logPathEnvVar           = "TF_LOG_PATH"

	varEnvVarPrefix = "TF_VAR_"
)

var prohibitedEnvVars = []string{
	inputEnvVar,
	automationEnvVar,
	logPathEnvVar,
	logEnvVar,
}

func environ() map[string]string {
	env := map[string]string{}
	for _, ev := range os.Environ() {
		parts := strings.SplitN(ev, "=", 2)
		if len(parts) == 0 {
			continue
		}
		k := parts[0]
		v := ""
		if len(parts) == 2 {
			v = parts[1]
		}
		env[k] = v
	}
	return env
}

func (tf *Terraform) buildEnv() []string {
	var menv map[string]string
	if tf.env == nil {
		menv = environ()
	} else {
		menv = make(map[string]string, len(tf.env))
		for k, v := range tf.env {
			menv[k] = v
		}
	}

	if _, ok := menv[checkpointDisableEnvVar]; !ok {
		// always propagate CHECKPOINT_DISABLE env var unless it is
		// explicitly overridden with tf.SetEnv
		menv[checkpointDisableEnvVar] = os.Getenv(checkpointDisableEnvVar)
	}

	menv[logEnvVar] = "" // so logging can't pollute our stderr output
	menv[inputEnvVar] = "0"
	menv[automationEnvVar] = "1"

	env := []string{}
	for k, v := range menv {
		env = append(env, k+"="+v)
	}

	return env
}

func (tf *Terraform) buildTerraformCmd(ctx context.Context, args ...string) *exec.Cmd {
	env := tf.buildEnv()

	cmd := exec.CommandContext(ctx, tf.execPath, args...)
	cmd.Env = env
	cmd.Dir = tf.workingDir

	tf.logger.Printf("Terraform command: %s", cmdString(cmd))

	return cmd
}
