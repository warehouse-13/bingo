package command_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
	"github.com/warehouse-13/bingo/pkg/command"
	"github.com/warehouse-13/bingo/pkg/config"
)

func Test_BootstrapFn_succeeds(t *testing.T) {
	g := NewWithT(t)

	var cliName = "foo"
	cfg := &config.Config{CliName: cliName}
	fs := afero.NewMemMapFs()

	g.Expect(command.BootstrapFn(cfg, fs)).To(Succeed())

	expectedFiles := map[string]string{
		"main.go":                "package main",
		"pkg/command/version.go": "package command",
		"pkg/command/example.go": "package command",
		"pkg/flags/flags.go":     "package flags",
		"pkg/config/config.go":   "package config",
		"pkg/version/version.go": "package version",
	}

	for name, contents := range expectedFiles {
		_, err := fs.Stat(name)
		g.Expect(err).NotTo(HaveOccurred())

		ok, err := afero.FileContainsBytes(fs, name, []byte(contents))
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(ok).To(BeTrue(), fmt.Sprintf("File %s does not contain expected contents", name))
	}

	// check this one separately as we want to make sure the name went in
	_, err := fs.Stat("pkg/command/app.go")
	g.Expect(err).NotTo(HaveOccurred())

	ok, err := afero.FileContainsBytes(fs, "pkg/command/app.go", []byte(cliName))
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(ok).To(BeTrue())
}

func Test_BootstrapFn_errorsWhenFilesExist(t *testing.T) {
	g := NewWithT(t)

	tt := []struct {
		filename string
	}{
		{filename: "main.go"},
		{filename: "pkg/command/app.go"},
		{filename: "pkg/command/version.go"},
		{filename: "pkg/command/example.go"},
		{filename: "pkg/flags/flags.go"},
		{filename: "pkg/config/config.go"},
		{filename: "pkg/version/version.go"},
	}

	for _, tc := range tt {
		t.Run("should fail if managed files exist", func(t *testing.T) {
			cfg := &config.Config{CliName: "foo"}

			fs := afero.NewMemMapFs()
			_, err := fs.Create(tc.filename)
			g.Expect(err).NotTo(HaveOccurred())

			g.Expect(command.BootstrapFn(cfg, fs)).NotTo(Succeed())
		})
	}
}
