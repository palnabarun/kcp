/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"

	"github.com/spf13/pflag"

	"github.com/kcp-dev/kcp/cli/cmd/kubectl-ws/cmd"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-ws", pflag.ExitOnError)
	pflag.CommandLine = flags

	workspaceCmd := cmd.KubectlWsCommand()

	if err := workspaceCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
