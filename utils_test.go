/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
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
 */

package main

import "testing"

// ShutdownCallback simulates a successful and failure exit here.
func TestShutdownCallbackSuccess(t *testing.T) {
	// Register two callbacks that return success
	registerObjectStorageShutdown(func() errCode {
		return exitSuccess
	})
	registerShutdown(func() errCode {
		return exitSuccess
	})

	shutdownSignal = make(chan bool, 1)
	shutdownSignal <- true
	// Start executing callbacks and exitFunc receives a success.
	dummySuccess := func(code int) {
		if code != int(exitSuccess) {
			t.Fatalf("Expected %d, got %d instead.", code, exitSuccess)
		}
	}
	monitorShutdownSignal(dummySuccess)
}
