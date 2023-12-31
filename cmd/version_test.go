package cmd

import (
	"bytes"
	"testing"

	"github.com/JintaTechx/DDG/tree/main/pkg/services"
	"github.com/stretchr/testify/assert"
)

func Test_versionCmd(t *testing.T) {
	translate.InitLanguage()

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	rootCmd.SetArgs([]string{"version"})

	err := rootCmd.Execute()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, buf.String(), "pdgen v1.0.0\n")
}
