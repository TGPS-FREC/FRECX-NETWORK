// Copyright 2016 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"crypto/rand"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/FRECNET/params"
)

const (
	ipcAPIs  = "FREx:1.0 FRExlending:1.0 S2PoS:1.0 admin:1.0 debug:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0"
	httpAPIs = "eth:1.0 net:1.0 rpc:1.0 web3:1.0"
)

// Tests that a node embedded within a console can be started up properly and
// then terminated by closing the input stream.
func TestConsoleWelcome(t *testing.T) {
	coinbase := "fre8605cdbbdb6d264aa742e77020dcbc58fcdce182"

	// Start a FRE console, make sure it's cleaned up and terminate the console
	FRE := runFRE(t,
		"--FREx.datadir", tmpdir(t)+"FREx/"+time.Now().String(),
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--etherbase", coinbase,
		"console")

	// Gather all the infos the welcome message needs to contain
	FRE.SetTemplateFunc("goos", func() string { return runtime.GOOS })
	FRE.SetTemplateFunc("goarch", func() string { return runtime.GOARCH })
	FRE.SetTemplateFunc("gover", runtime.Version)
	FRE.SetTemplateFunc("FREver", func() string { return params.Version })
	FRE.SetTemplateFunc("niltime", func() string { return time.Unix(1544771829, 0).Format(time.RFC1123) })
	FRE.SetTemplateFunc("apis", func() string { return ipcAPIs })

	// Verify the actual welcome message to the required template
	FRE.Expect(`
Welcome to the FRE JavaScript console!

instance: FRE/v{{FREver}}/{{goos}}-{{goarch}}/{{gover}}
coinbase: {{.Etherbase}}
at block: 0 ({{niltime}})
 datadir: {{.Datadir}}
 modules: {{apis}}

> {{.InputLine "exit"}}
`)
	FRE.ExpectExit()
}

// Tests that a console can be attached to a running node via various means.
func TestIPCAttachWelcome(t *testing.T) {
	// Configure the instance for IPC attachement
	coinbase := "fre8605cdbbdb6d264aa742e77020dcbc58fcdce182"
	var ipc string
	if runtime.GOOS == "windows" {
		ipc = `\\.\pipe\FRE` + strconv.Itoa(trulyRandInt(100000, 999999))
	} else {
		ws := tmpdir(t)
		defer os.RemoveAll(ws)
		ipc = filepath.Join(ws, "FRE.ipc")
	}
	FRE := runFRE(t,
		"--FREx.datadir", tmpdir(t)+"FREx/"+time.Now().String(),
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--etherbase", coinbase, "--ipcpath", ipc)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, FRE, "ipc:"+ipc, ipcAPIs)

	FRE.Interrupt()
	FRE.ExpectExit()
}

func TestHTTPAttachWelcome(t *testing.T) {
	coinbase := "fre8605cdbbdb6d264aa742e77020dcbc58fcdce182"
	port := strconv.Itoa(trulyRandInt(1024, 65536)) // Yeah, sometimes this will fail, sorry :P
	FRE := runFRE(t,
		"--FREx.datadir", tmpdir(t)+"FREx/"+time.Now().String(),
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--etherbase", coinbase, "--rpc", "--rpcport", port)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, FRE, "http://localhost:"+port, httpAPIs)

	FRE.Interrupt()
	FRE.ExpectExit()
}

func TestWSAttachWelcome(t *testing.T) {
	coinbase := "fre8605cdbbdb6d264aa742e77020dcbc58fcdce182"
	port := strconv.Itoa(trulyRandInt(1024, 65536)) // Yeah, sometimes this will fail, sorry :P

	FRE := runFRE(t,
		"--FREx.datadir", tmpdir(t)+"FREx/"+time.Now().String(),
		"--port", "0", "--maxpeers", "0", "--nodiscover", "--nat", "none",
		"--etherbase", coinbase, "--ws", "--wsport", port)

	time.Sleep(2 * time.Second) // Simple way to wait for the RPC endpoint to open
	testAttachWelcome(t, FRE, "ws://localhost:"+port, httpAPIs)

	FRE.Interrupt()
	FRE.ExpectExit()
}

func testAttachWelcome(t *testing.T, FRE *testFRE, endpoint, apis string) {
	// Attach to a running FRE note and terminate immediately
	attach := runFRE(t, "attach", endpoint)
	defer attach.ExpectExit()
	attach.CloseStdin()

	// Gather all the infos the welcome message needs to contain
	attach.SetTemplateFunc("goos", func() string { return runtime.GOOS })
	attach.SetTemplateFunc("goarch", func() string { return runtime.GOARCH })
	attach.SetTemplateFunc("gover", runtime.Version)
	attach.SetTemplateFunc("FREver", func() string { return params.Version })
	attach.SetTemplateFunc("etherbase", func() string { return FRE.Etherbase })
	attach.SetTemplateFunc("niltime", func() string { return time.Unix(1544771829, 0).Format(time.RFC1123) })
	attach.SetTemplateFunc("ipc", func() bool { return strings.HasPrefix(endpoint, "ipc") })
	attach.SetTemplateFunc("datadir", func() string { return FRE.Datadir })
	attach.SetTemplateFunc("apis", func() string { return apis })

	// Verify the actual welcome message to the required template
	attach.Expect(`
Welcome to the FRE JavaScript console!

instance: FRE/v{{FREver}}/{{goos}}-{{goarch}}/{{gover}}
coinbase: {{etherbase}}
at block: 0 ({{niltime}}){{if ipc}}
 datadir: {{datadir}}{{end}}
 modules: {{apis}}

> {{.InputLine "exit" }}
`)
	attach.ExpectExit()
}

// trulyRandInt generates a crypto random integer used by the console tests to
// not clash network ports with other tests running cocurrently.
func trulyRandInt(lo, hi int) int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(hi-lo)))
	return int(num.Int64()) + lo
}
