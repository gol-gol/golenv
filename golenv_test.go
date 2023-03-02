package golenv

import (
	"os"
	"testing"
)

func TestOverrideIfEnv(t *testing.T) {
	defaultValue := "SHOULD_NOT_EXIST"
	if HasEnv("HOME") {
		if OverrideIfEnv("HOME", defaultValue) == defaultValue {
			t.Fatal("OverrideIfEnv failed, didn't found HOME.")
		}
	}
	if OverrideIfEnv("TMPKEY_SHOULD_NOT_EXIST", defaultValue) != defaultValue {
		t.Fatal("OverrideIfEnv failed, didn't return default for KEY_SHOULD_NOT_EXIST")
	}
}

func TestOverrideIfEnvBool(t *testing.T) {
	if err := os.Setenv("TMPKEY_TRUE", "true"); err != nil {
		t.Fatal(err.Error())
	}
	if !OverrideIfEnvBool("TMPKEY_TRUE", false) {
		t.Fatal("OverrideIfEnvBool failed, for pre-set value.")
	}
	if OverrideIfEnvBool("TMPKEY_FALSE", false) {
		t.Fatal("OverrideIfEnvBool failed, for default value.")
	}
}

func TestOverrideIfEnvInt(t *testing.T) {
	if err := os.Setenv("TMPKEY_ONE", "1"); err != nil {
		t.Fatal(err.Error())
	}
	if OverrideIfEnvInt("TMPKEY_ONE", 0) != 1 {
		t.Fatal("OverrideIfEnvBool failed, for pre-set value.")
	}
	if OverrideIfEnvInt("TMPKEY_TWO", 2) != 2 {
		t.Fatal("OverrideIfEnvBool failed, for default value.")
	}
}

func TestOverrideIfEnvUint64(t *testing.T) {
	if err := os.Setenv("TMPKEY_ONE", "-1"); err != nil {
		t.Fatal(err.Error())
	}
	if OverrideIfEnvUint64("TMPKEY_ONE", 0) < 0 {
		t.Fatal("OverrideIfEnvBool failed, for pre-set value.")
	}
	if OverrideIfEnvUint64("TMPKEY_TWO", 2) != 2 {
		t.Fatal("OverrideIfEnvBool failed, for default value.")
	}
}

func TestHasEnv(t *testing.T) {
	if result := HasEnv("SHOULD_NOT_EXIST"); result {
		t.Fatal("HasEnv failed, found SHOULD_NOT_EXIST.")
	}
	if result := HasEnv("HOME"); !result {
		t.Fatal("HasEnv failed, didn't found HOME.")
	}
}

func TestEnvMap(t *testing.T) {
	envMap := EnvMap()
	if result := envMap["HOME"]; result == "" {
		t.Fatal("EnvMap failed, didn't found HOME.")
	}
	if result := envMap["SHOULD_NOT_EXIST"]; result != "" {
		t.Fatal("EnvMap failed, didn't found HOME.")
	}
}
