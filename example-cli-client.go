/* Demonstration of OIDC 2.0 Authorization Flow from a CLI program
   Draws heavily from: https://medium.com/@balaajanthan/openid-flow-from-a-cli-ac45de876ead */

package main

import(
  "log"
  "os"
  "path/filepath"
  "github.com/coreos/go-oidc/v3/oidc"
  "golang.org/x/net/context"
  "gopkg.in/ini.v1"
  "example-cli-client/client"
)

var (
  clientName = "example-cli-client"
  clientID = ""
  clientSecret = ""
  providerUrl = ""
)

func readIni() {
  ex, err := os.Executable()

  if err != nil {
    panic(err)
  }

  cfg, err := ini.Load(filepath.Join(filepath.Dir(ex), clientName + ".ini"))

  if err != nil {
    panic(err)
  }

  cs := cfg.Section(clientName)

  clientID = cs.Key("clientID").String()

  if clientID == "" {
    log.Fatal(clientName + ".ini does not specify clientID")
    os.Exit(1)
  }

  clientSecret = cs.Key("clientSecret").String()

  if clientSecret == "" {
    log.Fatal(clientName + ".ini does not specify clientSecret")
    os.Exit(1)
  }

  providerUrl = cs.Key("providerUrl").String()

  if providerUrl == "" {
    log.Fatal(clientName + ".ini does not specify providerUrl")
    os.Exit(1)
  }

  log.Printf(
    "Read configuration:\n" +
    " clientID = %s\n" +
    " clientSecret = %s\n" +
    " providerUrl = %s\n",
    clientID,
    "*REDACTED*",
    providerUrl,
  )
}

func main() {
  const callbackUrl = "http://localhost:8080/callback"
  readIni()

  ctx := context.Background()
  provider, err := oidc.NewProvider(ctx, providerUrl)

  if err != nil {
    log.Fatal(err)
  }

  tokenEp := provider.Endpoint().TokenURL
  authzEp := provider.Endpoint().AuthURL

  client.HandleOpenIDFlow(clientID, clientSecret, callbackUrl, authzEp, tokenEp)
}
