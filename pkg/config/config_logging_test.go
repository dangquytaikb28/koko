package config

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestSetupDoesNotLogCredentials(t *testing.T) {
	const secret = "credential-canary-do-not-log"
	t.Setenv("BOOTSTRAP_TOKEN", secret)
	t.Setenv("REDIS_PASSWORD", secret)
	t.Setenv("REDIS_SENTINEL_PASSWORD", secret)
	t.Setenv("SECRET_ENCRYPT_KEY", secret)
	var out bytes.Buffer
	previousOutput, previousConfig := log.Writer(), GlobalConfig
	log.SetOutput(&out)
	t.Cleanup(func() { log.SetOutput(previousOutput); GlobalConfig = previousConfig })
	Setup(t.TempDir() + "/missing.yml")
	if strings.Contains(out.String(), secret) {
		t.Fatal("startup log exposed a configuration credential")
	}
	if !strings.Contains(out.String(), "Configuration loaded") {
		t.Fatal("missing safe startup diagnostic")
	}
}
