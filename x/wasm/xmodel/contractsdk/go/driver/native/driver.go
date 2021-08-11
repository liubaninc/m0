package native

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/liubaninc/m0/x/wasm/xmodel/contractsdk/go/code"
	pbrpc "github.com/liubaninc/m0/x/wasm/xmodel/contractsdk/go/pbrpc"
	"google.golang.org/grpc"
)

const (
	mchainPingTimeout = "MCHAIN_PING_TIMEOUT"
	mchainCodePort    = "MCHAIN_CODE_PORT"
	mchainChainAddr   = "MCHAIN_CHAIN_ADDR"
)

type driver struct {
}

// New returns a native driver
func New() code.Driver {
	return new(driver)
}

func (d *driver) Serve(contract code.Contract) {
	chainAddr := os.Getenv(mchainChainAddr)
	codePort := os.Getenv(mchainCodePort)

	if chainAddr == "" {
		panic("empty MCHAIN_CHAIN_ADDR env")
	}

	if codePort == "" {
		panic("empty MCHAIN_CODE_PORT env")
	}

	nativeCodeService := newNativeCodeService(chainAddr, contract)
	rpcServer := grpc.NewServer()
	pbrpc.RegisterNativeCodeServer(rpcServer, nativeCodeService)

	var listener net.Listener
	listener, err := net.Listen("tcp", "127.0.0.1:"+codePort)
	if err != nil {
		panic(err)
	}

	go rpcServer.Serve(listener)

	sigch := make(chan os.Signal, 2)
	signal.Notify(sigch, os.Interrupt, syscall.SIGTERM, syscall.SIGPIPE)

	timer := time.NewTicker(1 * time.Second)
	running := true
	pingTimeout := getPingTimeout()
	for running {
		select {
		case sig := <-sigch:
			running = false
			log.Print("receive signal ", sig)
		case <-timer.C:
			lastping := nativeCodeService.LastpingTime()
			if time.Since(lastping) > pingTimeout {
				log.Print("ping timeout")
				running = false
			}
		}
	}
	rpcServer.GracefulStop()
	nativeCodeService.Close()
	log.Print("native code ended")
}

func getPingTimeout() time.Duration {
	envtimeout := os.Getenv(mchainPingTimeout)
	if envtimeout == "" {
		return 3 * time.Second
	}
	timeout, err := strconv.Atoi(envtimeout)
	if err != nil {
		return 3 * time.Second
	}
	return time.Duration(timeout) * time.Second
}
