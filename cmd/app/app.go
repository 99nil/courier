// Copyright © 2021 zc2638 <zc2638@qq.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/sync/errgroup"

	"github.com/99nil/courier/handler"

	"github.com/pkgms/go/server"
	"github.com/spf13/cobra"

	"github.com/99nil/courier/pkg/meta"
)

var cfgFile string

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "courier",
		Short:        "courier",
		SilenceUsage: true,
		RunE:         runE,
	}
	cfgFilePath := os.Getenv(meta.EnvPrefix + "_CONFIG")
	if cfgFilePath == "" {
		cfgFilePath = "config/config.yaml"
	}
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFilePath, "config file (default is $HOME/config.yaml)")
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	wg, ctx := errgroup.WithContext(ctx)

	wg.Go(func() error {
		s := server.New(&server.Config{
			Port: 6000,
		})
		s.Handler = handler.NewServerHandler()
		fmt.Println("API server listen on", s.Addr)
		return s.Run(ctx)
	})

	wg.Go(func() error {
		s := server.New(&server.Config{
			Port: 6443,
		})
		s.Handler = handler.NewServerHandler()
		fmt.Println("Client server listen on", s.Addr)
		return s.Run(ctx)
	})
	return wg.Wait()
}
