/*
 * Copyright (c) 2019 SUSE LLC. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package caaspctl

import (
	"fmt"
	"path"

	"suse.com/caaspctl/internal/pkg/caaspctl/deployments"
)

func KubeadmInitConfFile() string {
	return "kubeadm-init.conf"
}

func JoinConfDir() string {
	return "kubeadm-join.conf.d"
}

func MasterConfTemplateFile() string {
	return path.Join(JoinConfDir(), "master.conf.template")
}

func WorkerConfTemplateFile() string {
	return path.Join(JoinConfDir(), "worker.conf.template")
}

func MachineConfFile(target string) string {
	return path.Join(JoinConfDir(), fmt.Sprintf("%s.conf", target))
}

func TemplatePathForRole(role deployments.Role) string {
	switch role {
	case deployments.MasterRole:
		return MasterConfTemplateFile()
	case deployments.WorkerRole:
		return WorkerConfTemplateFile()
	}
	return ""
}

func AddonsDir() string {
	return "addons"
}

func CniDir() string {
	return path.Join(AddonsDir(), "cni")
}

func FlannelManifestFile() string {
	return path.Join(CniDir(), "flannel.yaml")
}

func KubeConfigAdminFile() string {
	return "admin.conf"
}

func PkiDir() string {
	return "pki"
}