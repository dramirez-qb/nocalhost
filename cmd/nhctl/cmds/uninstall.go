/*
* Copyright (C) 2021 THL A29 Limited, a Tencent company.  All rights reserved.
* This source code is licensed under the Apache License Version 2.0.
 */

package cmds

import (
	"nocalhost/cmd/nhctl/cmds/common"
	"nocalhost/internal/nhctl/const"
	"nocalhost/internal/nhctl/controller"
	"nocalhost/internal/nhctl/nocalhost"
	"nocalhost/internal/nhctl/utils"
	"nocalhost/pkg/nhctl/log"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var force bool

func init() {
	uninstallCmd.Flags().StringVarP(&common.NameSpace, "namespace", "n", "", "kubernetes namespace")
	uninstallCmd.Flags().BoolVar(&force, "force", false, "force to uninstall anyway")
	rootCmd.AddCommand(uninstallCmd)
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [NAME]",
	Short: "Uninstall application",
	Long:  `Uninstall application`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.Errorf("%q requires at least 1 argument\n", cmd.CommandPath())
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		applicationName := args[0]
		if applicationName == _const.DefaultNocalhostApplication {
			log.Error(_const.DefaultNocalhostApplicationOperateErr)
			return
		}

		must(common.Prepare())

		appMeta, err := nocalhost.GetApplicationMeta(applicationName, common.NameSpace, common.KubeConfig)
		must(err)

		nid = appMeta.NamespaceId

		if appMeta.IsNotInstall() {
			log.Fatal(appMeta.NotInstallTips())
		}

		log.Info("Uninstalling application...")

		//goland:noinspection ALL
		mustI(appMeta.Uninstall(true), "error while uninstall application")

		p, _ := nocalhost.GetProfileV2(common.NameSpace, applicationName, nid)
		if p != nil {
			for _, sv := range p.SvcProfile {
				for _, pf := range sv.DevPortForwardList {
					log.Infof("Stopping %s-%s's port-forward %d:%d", common.NameSpace, applicationName, pf.LocalPort, pf.RemotePort)
					utils.Should(controller.StopPortForward(common.NameSpace, nid, applicationName, sv.Name, pf))
				}
			}
		}

		if err = nocalhost.CleanupAppFilesUnderNs(common.NameSpace, nid); err != nil {
			log.WarnE(err, "")
		}

		log.Infof("Application \"%s\" is uninstalled", applicationName)
	},
}
