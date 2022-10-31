package app

import (
	"fmt"
	"net"
	"os"

	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"

	pb "github.com/airconduct/kruntime/runtimes/golett/api/proto"
	"github.com/airconduct/kruntime/runtimes/golett/internal/service"
	"github.com/airconduct/kruntime/runtimes/golett/internal/utils"
)

func serverCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Server start command",
		RunE:  func(cmd *cobra.Command, args []string) error { return cmd.Help() },
	}
	cmd.AddCommand(
		serverStartCommand(),
	)
	return cmd
}

func serverStartCommand() *cobra.Command {
	opts := &serverStartOptions{
		log: rootLogger.WithName("server"),
	}
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start a golet server",
		RunE:  optionsRun(opts),
	}
	opts.AddFlags(cmd.Flags())
	return cmd
}

type serverStartOptions struct {
	port    int
	address string
	log     logr.Logger
}

func (o *serverStartOptions) AddFlags(flag *pflag.FlagSet) {
	flag.StringVar(&o.address, "address", "127.0.0.1", "127.0.0.1 or 0.0.0.0")
	flag.IntVarP(&o.port, "port", "p", 7770, "server port")
}

func (o *serverStartOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

func (o *serverStartOptions) Run() error {
	ctx := utils.SetupSignalHandler()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", o.address, o.port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	gs := service.New()
	pb.RegisterGoletServer(s, gs)
	defer s.Stop()
	go func() {
		o.log.Info("Listen on", "address", lis.Addr())
		if err := s.Serve(lis); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	<-ctx.Done()
	gs.Shutdown()
	s.Stop()
	return nil
}
