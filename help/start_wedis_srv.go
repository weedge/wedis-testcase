package help

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/viper"
	"github.com/weedge/wedis-testcase/pkg/injectors"
)

type WedisSrv struct {
	cmd *exec.Cmd

	// tcp addr for host:port(ipv4)
	addr    *net.TCPAddr
	configs map[string]string

	clean func(bool)
}

func init() {
	viper.AutomaticEnv()
}

func StartSrv(srvConfigs map[string]string) *WedisSrv {
	srv := &WedisSrv{}

	binPath := os.Getenv(EnvSrvBinPath)
	Expect(len(binPath) != 0).Should(BeTrue())
	cmd := exec.Command(binPath)
	cmd.Args = append(cmd.Args, "srv")
	cmd.Args = append(cmd.Args, "--config", srv.InitConfigs(srvConfigs).Name())

	dir := srvConfigs["storeCfg.dataDir"]
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
	}).WithContext(context.Background()).WithArguments(proc, process.Zombie).WithTimeout(3 * time.Second).WithPolling(time.Second).Should(Succeed())

	srv.configs = srvConfigs
	srv.cmd = cmd
	srv.clean = func(keepDir bool) {
		Expect(stdout.Close()).NotTo(HaveOccurred())
		Expect(stderr.Close()).NotTo(HaveOccurred())
		if len(os.Getenv(EnvKeepDataDir)) == 0 && !keepDir {
			//println("remove dir", dir)
			Expect(os.RemoveAll(dir)).NotTo(HaveOccurred())
		}
	}
	return srv
}

func (s *WedisSrv) InitConfigs(srvConfigs map[string]string) (f *os.File) {
	addr, err := findFreePort()
	Expect(err).NotTo(HaveOccurred())
	if srvConfigs["server.respCmdSrv.addr"] == "" {
		srvConfigs["server.respCmdSrv.addr"] = fmt.Sprintf("%s:%d", addr.IP.String(), addr.Port)
		srvPort := os.Getenv(EnvSrvPort)
		if len(srvPort) > 0 {
			port, err := strconv.Atoi(srvPort)
			Expect(err).NotTo(HaveOccurred())
			srvConfigs["server.respCmdSrv.addr"] = fmt.Sprintf("%s:%d", addr.IP.String(), port)
			addr.Port = port
		}
	}
	s.addr = addr

	dir := os.Getenv(EnvDataDir)
	Expect(len(dir) != 0).Should(BeTrue())
	dir, err = os.MkdirTemp(dir, fmt.Sprintf("%d-*", time.Now().UnixMilli()))
	Expect(err).NotTo(HaveOccurred())
	// if want test one storager for diff suit test, please config storeCfg.dataDir, keep dir don't remove
	if srvConfigs["storeCfg.dataDir"] == "" {
		srvConfigs["storeCfg.dataDir"] = dir
	}

	f, err = os.Create(filepath.Join(dir, "srv.toml"))
	Expect(err).NotTo(HaveOccurred())
	defer func() {
		Expect(f.Close()).NotTo(HaveOccurred())
	}()

	for key := range srvConfigs {
		_, err := f.WriteString(fmt.Sprintf("%s = \"%s\"\n", key, srvConfigs[key]))
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

func (s *WedisSrv) NewTCPClient() *TCPClient {
	c, err := net.Dial(s.addr.Network(), s.addr.String())
	Expect(err).NotTo(HaveOccurred())
	return newTCPClient(c)
}

func (s *WedisSrv) NewRedisClient() redis.UniversalClient {
	return injectors.InitRedisClient(injectors.WithRedisAddr(s.addr.String()))
}
