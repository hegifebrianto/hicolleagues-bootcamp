package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMuatDotEnv(t *testing.T) {
	p := filepath.Join(t.TempDir(), ".env")
	os.WriteFile(p, []byte("# komentar\nTES_A=satu\nTES_B=\"dua spasi\"\n\nexport TES_C=tiga\nTES_D=dari_file\n"), 0o600)
	t.Setenv("TES_D", "dari_terminal") // terminal harus menang
	os.Unsetenv("TES_A")
	os.Unsetenv("TES_B")
	os.Unsetenv("TES_C")
	muatDotEnv(p)
	for k, want := range map[string]string{"TES_A": "satu", "TES_B": "dua spasi", "TES_C": "tiga", "TES_D": "dari_terminal"} {
		if got := os.Getenv(k); got != want {
			t.Errorf("%s = %q, mau %q", k, got, want)
		}
	}
	muatDotEnv("/tidak/ada/.env") // tidak boleh panic
}
