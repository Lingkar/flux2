/*
Copyright 2024 The Flux authors

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
	"fmt"
	"github.com/spf13/cobra"

	helmv2 "github.com/fluxcd/helm-controller/api/v2"
)

var diffHrCmd = &cobra.Command{
	Use:               "helmrelease",
	Aliases:           []string{"hr"},
	Short:             "Diff HelmRelease",
	Long:              "TODO:: fill in",
	Example:           "TODO: fill in",
	ValidArgsFunction: resourceNamesCompletionFunc(helmv2.GroupVersion.WithKind(helmv2.HelmReleaseKind)),
	RunE:              diffHrCmdRun,
}

func init() {
	diffHrCmd.Flags().StringVar(&diffKsArgs.path, "path", "", "Path to a local directory that matches the specified Kustomization.spec.path.")

	diffCmd.AddCommand(diffHrCmd)
}

func diffHrCmdRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("%s name is required", helmReleaseType.humanKind)
	}

	// TODO:: Remove when not used anymore
	return fmt.Errorf("Not yet implemented")
}
