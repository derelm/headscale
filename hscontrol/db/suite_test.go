package db

import (
	"log"
	"net/netip"
	"os"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

var _ = check.Suite(&Suite{})

type Suite struct{}

var (
	tmpDir string
	db     *HSDatabase
)

func (s *Suite) SetUpTest(c *check.C) {
	s.ResetDB(c)
}

func (s *Suite) TearDownTest(c *check.C) {
	// os.RemoveAll(tmpDir)
}

func (s *Suite) ResetDB(c *check.C) {
	// if len(tmpDir) != 0 {
	// 	os.RemoveAll(tmpDir)
	// }

	var err error
	tmpDir, err = os.MkdirTemp("", "headscale-db-test-*")
	if err != nil {
		c.Fatal(err)
	}

	log.Printf("database path: %s", tmpDir+"/headscale_test.db")

	db, err = NewHeadscaleDatabase(
		"sqlite3",
		tmpDir+"/headscale_test.db",
		false,
		[]netip.Prefix{
			netip.MustParsePrefix("10.27.0.0/23"),
		},
		"",
	)
	if err != nil {
		c.Fatal(err)
	}
}
