/*
Copyright 2022 The Kubernetes Authors.

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

package v1beta1

import (
	"sigs.k8s.io/cluster-api-provider-cloudstack/api/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *CloudStackAffinityGroup) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*v1beta2.CloudStackAffinityGroup)
	return Convert_v1beta1_CloudStackAffinityGroup_To_v1beta2_CloudStackAffinityGroup(src, dst, nil)
}

func (dst *CloudStackAffinityGroup) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*v1beta2.CloudStackAffinityGroup)
	return Convert_v1beta2_CloudStackAffinityGroup_To_v1beta1_CloudStackAffinityGroup(src, dst, nil)
}
