//go:build unit
// +build unit

/*
Copyright 2021 The Flux authors

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
	"testing"
)

func start(t *testing.T, tmpl map[string]string) {
	t.Helper()
	testEnv.CreateObjectFile("./testdata/create_hr/setup-source.yaml", tmpl, t)
	testEnv.CreateObjectFile("./testdata/create_hr/hc-basic.yaml", tmpl, t)
}

func TestBuildHelmRelease(t *testing.T) {
	tests := []struct {
		name       string
		args       string
		resultFile string
		assertFunc string
	}{
		{
			name:       "no args",
			args:       "build helmrelease podinfo",
			resultFile: "invalid resource path \"\"",
			assertFunc: "assertError",
		},
		{
			name:       "build podinfo",
			args:       "build helmrelease podinfo --path ./testdata/create_hr/hc_basic.yaml",
			resultFile: "./testdata/build_hr/hr_build_result.yaml",
			assertFunc: "assertGoldenTemplateFile",
		},
	}

	tmpl := map[string]string{
		"fluxns": allocateNamespace("flux-system"),
	}
	setup(t, tmpl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var assert assertFunc

			switch tt.assertFunc {
			case "assertGoldenTemplateFile":
				assert = assertGoldenTemplateFile(tt.resultFile, tmpl)
			case "assertError":
				assert = assertError(tt.resultFile)
			}

			cmd := cmdTestCase{
				args:   tt.args + " -n " + tmpl["fluxns"],
				assert: assert,
			}

			cmd.runTestCmd(t)
		})
	}
}
