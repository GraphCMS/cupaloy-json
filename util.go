package cupaloy

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/GraphCMS/cupaloy-json/v2/internal"

	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
)

var spewConfig = spew.ConfigState{
	Indent:                  "  ",
	SortKeys:                true, // maps should be spewed in a deterministic order
	DisablePointerAddresses: true, // don't spew the addresses of pointers
	DisableCapacities:       true, // don't spew capacities of collections
	SpewKeys:                true, // if unable to sort map keys then spew keys to strings and sort those
}

//go:generate $GOPATH/bin/mockery -output=examples -outpkg=examples_test -testonly -name=TestingT

// TestingT is a subset of the interface testing.TB allowing it to be mocked in tests.
type TestingT interface {
	Helper()
	Failed() bool
	Error(args ...interface{})
	Fatal(args ...interface{})
	Name() string
}

func getNameOfCaller() string {
	pc, _, _, _ := runtime.Caller(2) // first caller is the caller of this function, we want the caller of our caller
	fullPath := runtime.FuncForPC(pc).Name()
	packageFunctionName := filepath.Base(fullPath)

	return strings.Replace(packageFunctionName, ".", "-", -1)
}

func envVariableSet(envVariable string) bool {
	_, varSet := os.LookupEnv(envVariable)
	return varSet
}

func (c *Config) snapshotFilePath(testName string) string {
	return filepath.Join(c.subDirName, testName+c.snapshotFileExtension)
}

// Legacy snapshot format where all items were spewed
func takeV1Snapshot(i ...interface{}) string {
	return spewConfig.Sdump(i...)
}

// New snapshot format where some types are written out raw to the file
func takeSnapshot(i ...interface{}) string {
	snapshot := &bytes.Buffer{}
	for _, v := range i {
		switch vt := v.(type) {
		case string:
			snapshot.WriteString(vt)
			snapshot.WriteString("\n")
		case []byte:
			snapshot.Write(vt)
			snapshot.WriteString("\n")
		default:
			spewConfig.Fdump(snapshot, v)
		}
	}

	return snapshot.String()
}

func (c *Config) readSnapshot(snapshotName string) (string, error) {
	snapshotFile := c.snapshotFilePath(snapshotName)
	buf, err := ioutil.ReadFile(snapshotFile)

	if os.IsNotExist(err) {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (c *Config) updateSnapshot(snapshotName string, prevSnapshot string, snapshot string) error {
	// check that subdirectory exists before writing snapshot
	err := os.MkdirAll(c.subDirName, os.ModePerm)
	if err != nil {
		return errors.New("could not create snapshots directory")
	}

	snapshotFile := c.snapshotFilePath(snapshotName)
	_, err = os.Stat(snapshotFile)
	isNewSnapshot := os.IsNotExist(err)

	err = ioutil.WriteFile(snapshotFile, []byte(snapshot), os.FileMode(0644))
	if err != nil {
		return err
	}

	if !c.failOnUpdate {
		//TODO: should a warning still be printed here?
		return nil
	}

	snapshotDiff := diffSnapshots(prevSnapshot, snapshot)

	if isNewSnapshot {
		return internal.ErrSnapshotCreated{
			Name:     snapshotName,
			Contents: snapshot,
		}
	}

	return internal.ErrSnapshotUpdated{
		Name: snapshotName,
		Diff: snapshotDiff,
	}
}

func diffSnapshots(previous, current string) string {
	diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(previous),
		B:        difflib.SplitLines(current),
		FromFile: "Previous",
		FromDate: "",
		ToFile:   "Current",
		ToDate:   "",
		Context:  1,
	})

	return diff
}

// Equal checks equality between 2 nested interfaces, regardless of sorting order.
func Equal(vx, vy interface{}) bool {
	if reflect.TypeOf(vx) != reflect.TypeOf(vy) {
		return false
	}

	switch x := vx.(type) {
	case map[string]interface{}:
		y := vy.(map[string]interface{})

		if len(x) != len(y) {
			return false
		}

		for k, v := range x {
			val2 := y[k]

			if (v == nil) != (val2 == nil) {
				return false
			}

			if !Equal(v, val2) {
				return false
			}
		}

		return true
	case []interface{}:
		y := vy.([]interface{})

		if len(x) != len(y) {
			return false
		}

		var matches int
		flagged := make([]bool, len(y))
		for _, v := range x {
			for i, v2 := range y {
				if Equal(v, v2) && !flagged[i] {
					matches++
					flagged[i] = true
					break
				}
			}
		}
		return matches == len(x)
	default:
		return vx == vy
	}
}

// Unmarshal parses a JSON string into an interface{}
func Unmarshal(b []byte) (interface{}, error) {
	var j interface{}

	err := json.Unmarshal(b, &j)
	if err != nil {
		return nil, err
	}

	return j, nil
}
