// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package main

import (
	"fmt"
	"time"

	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

func isPodRunning(c *KubeClient, podName string) wait.ConditionFunc {
	return func() (bool, error) {
		fmt.Printf(".")

		pod, err := getPod(c, podName)
		if err != nil {
			return false, err
		}

		switch pod.Status.Phase {
		case coreV1.PodPending:
			return false, nil
		case coreV1.PodRunning:
			return true, nil
		case coreV1.PodFailed, coreV1.PodSucceeded, coreV1.PodUnknown:
			return false, fmt.Errorf("pod cannot start...aborting")
		}
		return false, nil
	}
}

func isNamespaceTerminated(c *KubeClient, namespace string) wait.ConditionFunc {
	return func() (bool, error) {
		fmt.Printf(".")

		_, err := getNamespace(c, namespace)
		if err != nil {
			return true, err
		}
		return false, nil
	}
}

func getPod(c *KubeClient, podName string) (*coreV1.Pod, error) {
	return c.clientset.CoreV1().Pods(c.namespace).Get(c.ctx, podName, metaV1.GetOptions{})
}
func getNamespace(c *KubeClient, namespace string) (*coreV1.Namespace, error) {
	return c.clientset.CoreV1().Namespaces().Get(c.ctx, namespace, metaV1.GetOptions{})
}

func waitForPodRunning(c *KubeClient, podName string, timeoutSec int) error {
	return wait.PollImmediate(time.Second, time.Duration(timeoutSec)*time.Second, isPodRunning(c, podName))
}

func waitForNamespaceToBeTerminated(c *KubeClient, namespace string, timeoutSec int) error {
	return wait.PollImmediate(time.Second, time.Duration(timeoutSec)*time.Second, isNamespaceTerminated(c, namespace))
}
