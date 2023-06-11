package help

import "net"

func findFreePort() (*net.TCPAddr, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return nil, err
	}
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}
	defer func() { _ = lis.Close() }()
	return lis.Addr().(*net.TCPAddr), nil
}
