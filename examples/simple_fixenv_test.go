package examples

import (
	"github.com/rekby/fixenv"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateDir_Fixenv(t *testing.T) {
	e := fixenv.New(t)
	fpath := filepath.Join(Folder(e), "dir")
	err := os.Mkdir(fpath, 0666)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
}

func TestCreateFile_Fixenv(t *testing.T) {
	e := fixenv.New(t)
	fpath := filepath.Join(Folder(e), "file")
	f, err := os.Create(fpath)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}
	_ = f.Close()
}

func Folder(e fixenv.Env) string {
	return fixenv.CacheWithCleanup(e, nil, nil, func() (
		string, // result value
		fixenv.FixtureCleanupFunc, // cleanup func
		error, // error, test will fail if error non nil
	) {
		dir, err := os.MkdirTemp("", "")
		if err != nil {
			return "", nil, err
		}
		e.T().Logf("Directory created: %v", dir)
		clean := func() {
			_ = os.RemoveAll(dir)
			e.T().Logf("Directory removed: %v", dir)
		}
		return dir, clean, nil
	})
}
