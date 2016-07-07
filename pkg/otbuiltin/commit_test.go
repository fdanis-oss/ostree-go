package otbuiltin

import (
       "testing"
)

func TestCommitSuccess(t *testing.T) {
  // Make a directory in which the repo should exist
  testDir := "/tmp/test-init-repo"
  err := os.Mkdir(testDir, 0777)
  if (err != nil){
    t.Errorf("%s", err)
    return
  }
  defer os.RemoveAll(testDir)

  // Initialize the repo
  inited, err := Init("/tmp/test-init-repo", nil)
  if !inited || err != nil {
    fmt.Println("Cannot test commit: failed to initialize repo")
    return
  }

  //Test commit
  opts := NewCommitOptions()
  opts.Branch="test-branch"
  err = Commit(testDir, ".", opts)
  if err != nil {
    t.Errorf("%s", err)
  }
}
