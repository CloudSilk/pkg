package archive

import "testing"

func TestPackageDir(t *testing.T) {
	filePath, err := PackageFolder("../temp/test", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(filePath)
}
