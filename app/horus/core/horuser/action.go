// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package horuser

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func (h *Horuser) Cordon(nodeName, clusterName, modelName string) (err error) {
	kubeClient := h.kubeClientMap[clusterName]
	if kubeClient == nil {
		klog.Errorf("node Cordon kubeClient by clusterName empty.")
		klog.Infof("nodeName:%v,clusterName:%v,modelName:%v", nodeName, clusterName, modelName)
		return err
	}

	ctxFirst, cancelFirst := h.GetK8sContext()
	defer cancelFirst()
	node, err := kubeClient.CoreV1().Nodes().Get(ctxFirst, nodeName, v1.GetOptions{})
	if err != nil {
		klog.Errorf("node Cordon get err nodeName:%v clusterName:%v modelName:%v", nodeName, clusterName, modelName)
		return err
	}

	node.Spec.Unschedulable = true

	ctxSecond, cancelSecond := h.GetK8sContext()
	defer cancelSecond()
	node, err = kubeClient.CoreV1().Nodes().Update(ctxSecond, node, v1.UpdateOptions{})
	if err != nil {
		klog.Errorf("node Cordon update err nodeName:%v clusterName:%v modelName:%v", nodeName, clusterName, modelName)
		return err
	}
	klog.Infof("node Cordon success nodeName:%v clusterName:%v modelName:%v", nodeName, clusterName, modelName)
	return nil
}

func (h *Horuser) Drain() error {
	return nil
}

func (h *Horuser) Evict() error {
	return nil
}
