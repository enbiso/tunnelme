package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/enbiso/tunnelme/utils"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Execute tunneling server",
	Long:  "Excute tunneling server. This has to be executed in a publicly accessible machine.",
	Run: func(cmd *cobra.Command, args []string) {
		NewServer().Start()
	},
}

// Server data structure
type Server struct {
	remoteAddrs     []string
	remoteListeners []net.Listener
	remoteAddrIndex int
	remoteAddrAlloc remoteAddrAllocMode
	ControlAddr     string
	DataAddr        string
	dataConns       map[string]net.Conn
}

type remoteAddrAllocMode int

const (
	onDemand  remoteAddrAllocMode = 0
	onStartup remoteAddrAllocMode = 1
)

// NewServer constructor
func NewServer() *Server {
	s := new(Server)
	s.remoteAddrIndex = 0
	s.dataConns = map[string]net.Conn{}
	s.DataAddr = ":9002"
	s.ControlAddr = ":9001"
	return s
}

// Start the server
func (s Server) Start() {

	log.Info("TunnelMe Server")
	ctrlLn, _ := net.Listen("tcp", s.ControlAddr)
	dataLn, _ := net.Listen("tcp", s.DataAddr)

	go func() {
		for {
			dataConn, _ := dataLn.Accept()
			remortID, _ := bufio.NewReader(dataConn).ReadString('\n')
			log.Debug("Data connection accepted")
			s.dataConns[strings.TrimSuffix(remortID, "\n")] = dataConn
		}
	}()

	for {
		ctrlConn, _ := ctrlLn.Accept()
		log.Debug("Tunnel client connected")

		clientLn, _ := s.getRemoteListener()
		fmt.Fprintf(ctrlConn, clientLn.Addr().String()+"\n")

		go func() {
			for {
				clientConn, _ := clientLn.Accept()
				id := utils.UUID()
				fmt.Fprintf(ctrlConn, id+"\n")
				for s.dataConns[id] == nil {
					time.Sleep(200 * time.Millisecond)
				}
				dataConn := s.dataConns[id]
				log.Debug("Consumer connected")
				go utils.Pipe(clientConn, dataConn)
			}
		}()
	}
}

func (s Server) getRemoteListener() (net.Listener, error) {
	if s.remoteAddrAlloc == onDemand {
		if len(s.remoteAddrs) < s.remoteAddrIndex {
			addr := s.remoteAddrs[s.remoteAddrIndex]
			s.remoteAddrIndex++
			return net.Listen("tcp", addr)
		}
		return nil, errors.New("Remote address out of bound")
	} else if s.remoteAddrAlloc == onStartup {
		if len(s.remoteListeners) < s.remoteAddrIndex {
			ln := s.remoteListeners[s.remoteAddrIndex]
			s.remoteAddrIndex++
			return ln, nil
		}
	}
	return nil, errors.New("Invalid allocation method")
}

// OnDemandAddrAlloc setup
func (s Server) OnDemandAddrAlloc(startPort int, count int) {
	s.remoteAddrAlloc = onDemand
	s.remoteAddrs = []string{}

	for index := 0; index < count; index++ {
		s.remoteAddrs[index] = fmt.Sprintf(":%d", (startPort + index))
	}
}

// OnStartupAddrAlloc setup
func (s Server) OnStartupAddrAlloc(startPort int, count int) {
	s.remoteAddrAlloc = onStartup
	s.remoteListeners = []net.Listener{}
	for index := 0; index < count; index++ {
		addr := fmt.Sprintf(":%d", (startPort + index))
		s.remoteListeners[index], _ = net.Listen("tcp", addr)
	}
}
