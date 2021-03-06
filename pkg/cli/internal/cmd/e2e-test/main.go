/*
Copyright 2019 The KubeCarrier Authors.

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

package e2e_test

import (
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"

	kubermatictestutil "k8c.io/utils/pkg/testutil"

	"k8c.io/kubecarrier/pkg/testutil"
)

func NewCommand(log logr.Logger) *cobra.Command {
	cfg := &testutil.FrameworkConfig{}
	cmd := &cobra.Command{
		Use:   "e2e-test",
		Short: "end2end testing utilities",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cfg.Default()
			for p := cmd.Parent(); p != nil; p = p.Parent() {
				if p.PersistentPreRun != nil {
					p.PersistentPreRun(p, args)
				}
				if p.PersistentPreRunE != nil {
					if err := p.PersistentPreRunE(p, args); err != nil {
						panic(err)
					}
				}
			}
		},
	}

	cmd.AddCommand(
		newRunCommand(log, cfg),
		newSetupE2EOperator(log),
	)

	cmd.PersistentFlags().StringVar(&cfg.TestID, "test-id", "", "unique e2e test id")
	cmd.PersistentFlags().StringVar(&cfg.ManagementExternalKubeconfigPath, "management-external-kubeconfig", "", "management cluster external (reachable outside cluster/docker) kubeconfig file")
	cmd.PersistentFlags().StringVar(&cfg.ManagementInternalKubeconfigPath, "management-internal-kubeconfig", "", "management cluster internal (reachable within cluster/docker) kubeconfig file")
	cmd.PersistentFlags().StringVar(&cfg.ServiceExternalKubeconfigPath, "service-external-kubeconfig", "", "service cluster external (reachable outside cluster/docker) kubeconfig file")
	cmd.PersistentFlags().StringVar(&cfg.ServiceInternalKubeconfigPath, "service-internal-kubeconfig", "", "service cluster internal (reachable within cluster/docker) kubeconfig file")
	cmd.PersistentFlags().StringVar((*string)(&cfg.CleanUpStrategy), "clean-up-strategy", string(kubermatictestutil.CleanupOnSuccess), "cleanup strategy after the test ends. Valid values are always, on-success and never")
	return cmd
}
