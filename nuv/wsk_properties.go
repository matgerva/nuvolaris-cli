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
	"os"
	"path/filepath"
)

const wskAuth = "23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP"
const wskApihost = "http://localhost:3233"

func writeWskPropertiesFile() error {
	content := []byte("AUTH=" + wskAuth + "\nAPIHOST=" + wskApihost)
	path, err := WriteFileToNuvolarisConfigDir(".wskprops", content)
	if err != nil {
		return err
	}
	os.Setenv("WSK_CONFIG_FILE", path)
	return nil
}

func getWhiskPropsPath() (string, error) {
	path, err := GetOrCreateNuvolarisConfigDir()
	if err != nil {
		return "", err
	}
	wpath := filepath.Join(path, ".wskprops")
	return wpath, nil
}
