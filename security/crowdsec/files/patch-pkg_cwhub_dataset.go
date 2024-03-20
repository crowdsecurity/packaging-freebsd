--- pkg/cwhub/dataset.go.orig	1979-11-30 00:00:00 UTC
+++ pkg/cwhub/dataset.go
@@ -6,6 +6,7 @@ import (
 	"io"
 	"net/http"
 	"os"
+	"path/filepath"
 	"time"
 
 	"github.com/sirupsen/logrus"
@@ -31,19 +32,40 @@ func downloadFile(url string, destPath string) error {
 		return fmt.Errorf("bad http code %d for %s", resp.StatusCode, url)
 	}
 
-	file, err := os.Create(destPath)
+	// Download to a temporary location to avoid corrupting files
+	// that are currently in use or memory mapped.
+
+	tmpFile, err := os.CreateTemp(filepath.Dir(destPath), filepath.Base(destPath)+".*.tmp")
 	if err != nil {
 		return err
 	}
-	defer file.Close()
 
+	tmpFileName := tmpFile.Name()
+	defer func() {
+		tmpFile.Close()
+		os.Remove(tmpFileName)
+	}()
+
 	// avoid reading the whole file in memory
-	_, err = io.Copy(file, resp.Body)
+	_, err = io.Copy(tmpFile, resp.Body)
 	if err != nil {
 		return err
 	}
 
-	if err = file.Sync(); err != nil {
+	if err = tmpFile.Sync(); err != nil {
+		return err
+	}
+
+	if err = tmpFile.Close(); err != nil {
+		return err
+	}
+
+	// a check on stdout is used while scripting to know if the hub has been upgraded
+	// and a configuration reload is required
+	// TODO: use a better way to communicate this
+	fmt.Printf("updated %s\n", filepath.Base(destPath))
+
+	if err = os.Rename(tmpFileName, destPath); err != nil {
 		return err
 	}
 
