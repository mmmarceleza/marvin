// Copyright 2023 Undistro Authors
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

package cmd

import (
	"context"
	"flag"
	"os"

	"github.com/fatih/color"
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
)

var noColor bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "marvin",
	Short: "A Kubernetes cluster scanner",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initNoColor)

	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "Disable color output")
	var allFlags flag.FlagSet
	klog.InitFlags(&allFlags)
	allFlags.VisitAll(func(f *flag.Flag) {
		if f.Name == "v" {
			rootCmd.PersistentFlags().AddGoFlag(f)
		}
	})
	rootCmd.SetContext(logr.NewContext(context.Background(), klogr.New()))
}

func initNoColor() {
	color.NoColor = noColor
}
