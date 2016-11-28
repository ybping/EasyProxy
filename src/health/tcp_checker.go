package health

import (
	tcp "github.com/tevino/tcp-shaker"
	"log"
	"sync"
	"time"
)

var (
	globalTCPChecker         *TCPChecker
	initGlobalTCPCheckerOnce sync.Once
)

// TCPChecker doing health checking by TCP handshaking.
type TCPChecker struct {
	*tcp.Checker
}

// Check checks given host with timeout.
func (c *TCPChecker) Check(addr string, timeout time.Duration) error {
	return c.CheckAddr(addr, timeout)
}

// GetTCPChecker return singleton TCPChecker.
func GetTCPChecker() *TCPChecker {
	initGlobalTCPCheckerOnce.Do(func() {
		globalTCPChecker = &TCPChecker{
			tcp.NewChecker(true),
		}
		if err := globalTCPChecker.InitChecker(); err != nil {
			log.Fatal("Init TCPChecker failed:", err)
		}
		log.Println("Init TCPChecker successed")
	})
	return globalTCPChecker
}
