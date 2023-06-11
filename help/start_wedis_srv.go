package help

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	. "github.com/onsi/gomega"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/weedge/wedis-testcase/pkg/injectors"
)

type WedisSrv struct {
	cmd *exec.Cmd

	// tcp addr for host:port(ipv4)
	addr    *net.TCPAddr
	configs map[string]string

	clean func(bool)
}

func StartSrv(srvConfigs map[string]string) *WedisSrv {
	srv := &WedisSrv{}

	binPath := *srvBinPath
	Expect(len(binPath) != 0).Should(BeTrue())
	cmd := exec.Command(binPath)
	cmd.Args = append(cmd.Args, "srv")
	cmd.Args = append(cmd.Args, "-c", srv.InitConfigs(srvConfigs).Name())

	dir := srvConfigs["server.storeOpts.dataDir"]
	Expect(len(dir) != 0).Should(BeTrue())
	addr := srvConfigs["server.respCmdSrv.addr"]
	Expect(len(addr) != 0).Should(BeTrue())

	stdout, err := os.Create(filepath.Join(dir, "stdout"))
	Expect(err).NotTo(HaveOccurred())
	cmd.Stdout = stdout
	stderr, err := os.Create(filepath.Join(dir, "stderr"))
	Expect(err).NotTo(HaveOccurred())
	cmd.Stderr = stderr

	Expect(cmd.Start()).NotTo(HaveOccurred())
	proc, err := process.NewProcess(int32(cmd.Process.Pid))
	Expect(err).NotTo(HaveOccurred())

	client := injectors.InitRedisClient(injectors.WithRedisAddr(addr))
	Eventually(func(g Gomega, ctx context.Context, proc *process.Process, unexpected ...string) {
		err := client.Ping(ctx).Err()
		g.Expect(err).NotTo(HaveOccurred())
		status, err := proc.Status()
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(status).NotTo(ConsistOf(unexpected))
	}).WithContext(context.Background()).WithArguments(proc, process.Zombie).WithTimeout(10 * time.Second).WithPolling(time.Second).Should(Succeed())

	srv.configs = srvConfigs
	srv.cmd = cmd
	srv.clean = func(keepDir bool) {
		Expect(stdout.Close()).NotTo(HaveOccurred())
		Expect(stderr.Close()).NotTo(HaveOccurred())
		if !*keepDataDir || !keepDir {
			Expect(os.RemoveAll(dir)).NotTo(HaveOccurred())
		}
	}
	return srv
}

func (s *WedisSrv) InitConfigs(srvConfigs map[string]string) (f *os.File) {
	addr, err := findFreePort()
	Expect(err).NotTo(HaveOccurred())
	s.addr = addr
	if srvConfigs["server.respCmdSrv.addr"] == "" {
		srvConfigs["server.respCmdSrv.addr"] =
			fmt.Sprintf("%s:%d", addr.IP.String(), addr.Port)
	}

	dir := *dataDir
	Expect(len(dir) != 0).Should(BeTrue())
	dir, err = os.MkdirTemp(dir, fmt.Sprintf("%d-*", time.Now().UnixMilli()))
	Expect(err).NotTo(HaveOccurred())
	srvConfigs["server.storeOpts.dataDir"] = dir

	f, err = os.Create(filepath.Join(dir, "srv.toml"))
	Expect(err).NotTo(HaveOccurred())
	defer func() {
		Expect(f.Close()).NotTo(HaveOccurred())
	}()

	for key := range srvConfigs {
		_, err := f.WriteString(fmt.Sprintf("%s=%s\n", key, srvConfigs[key]))
		Expect(err).NotTo(HaveOccurred())
	}

	return
}

func (s *WedisSrv) Close() {
	s.close(false)
}

func (s *WedisSrv) close(keepDir bool) {
	Expect(s.cmd.Process.Signal(syscall.SIGTERM)).NotTo(HaveOccurred())

	timer := time.AfterFunc(DefaultContainerGracePeriod, func() {
		Expect(s.cmd.Process.Kill()).NotTo(HaveOccurred())
	})
	defer timer.Stop()
	Expect(s.cmd.Wait()).NotTo(HaveOccurred())
	s.clean(keepDir)
}
