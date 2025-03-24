--- pkg/apiclient/client.go.orig	2025-03-24 11:34:09 UTC
+++ pkg/apiclient/client.go
@@ -72,10 +72,6 @@ func InitLAPIClient(ctx context.Context, apiUrl string
 }
 
 func InitLAPIClient(ctx context.Context, apiUrl string, papiUrl string, login string, password string, scenarios []string) error {
-	if lapiClient != nil {
-		return errors.New("client already initialized")
-	}
-
 	apiURL, err := url.Parse(apiUrl)
 	if err != nil {
 		return fmt.Errorf("parsing api url ('%s'): %w", apiURL, err)
