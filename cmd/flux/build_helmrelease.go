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
	"log"

	"github.com/spf13/cobra"

	"github.com/fluxcd/flux2/v2/internal/build"
	helmv2 "github.com/fluxcd/helm-controller/api/v2"
)

var buildHrCmd = &cobra.Command{
	Use:               "helmrelease",
	Aliases:           []string{"hr"},
	Short:             "Build HelmRelease",
	Long:              "TODO:: fill in",
	Example:           "TODO: fill in",
	ValidArgsFunction: resourceNamesCompletionFunc(helmv2.GroupVersion.WithKind(helmv2.HelmReleaseKind)),
	RunE:              buildHrCmdRun,
}

type buildHrFlags struct {
	path string
}

var buildHrArgs buildHrFlags

func init() {
	buildHrCmd.Flags().StringVar(&buildHrArgs.path, "path", "", "Path to a local directory that matches the specified Kustomization.spec.path.")
	buildCmd.AddCommand(buildHrCmd)
}

func buildHrCmdRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("%s name is required", helmReleaseType.humanKind)
	}
	// name := args[0]

	if buildHrArgs.path == "" {
		return fmt.Errorf("invalid resource path %q", buildKsArgs.path)
	}
	log.Printf("check do i see this?")

	err := build.CheckIfFileIsHR(buildHrArgs.path)
	if err != nil {
		return err
	}

	// TODO:: Remove when not used anymore
	return fmt.Errorf("Not yet implemented")
}
