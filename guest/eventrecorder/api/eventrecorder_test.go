/*
   Copyright 2023 The Kubernetes Authors.

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

package api_test

import (
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/api/proto"
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/eventrecorder/api"
)

var (
	eventrecorder api.EventRecorder = api.UnimplementedEventRecorder{}
	pod           proto.Pod
	node          proto.Node
)

func ExampleEventRecorder() {
	eventrecorder.Eventf(pod, pod, "event", "reason", "action", "note")
	eventrecorder.Eventf(node, node, "event", "reason", "action", "note")

	// Output:
	//
}
