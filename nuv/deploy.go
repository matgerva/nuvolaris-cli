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
	_ "embed"
	"fmt"
	"io/ioutil"
)

//go:embed embed/nuvolaris.yml
var NuvolarisYml []byte

type DeployCmd struct {
	Args              []string `optional:"" name:"args" help:"kind subcommand args"`
	NoPreflightChecks bool     `help:"Disable preflight checks."`
}

// AfterApply is an hook that gets called after parsing the command but before Run is executed
// used to run preflight checks
func (d DeployCmd) AfterApply() error {
	if d.NoPreflightChecks {
		return nil
	}
	homedir, _ := GetHomeDir()

	err := RunPreflightChecks(homedir)

	if err != nil {
		return err
	}
	return nil
}

func (*DeployCmd) Run() error {
	fmt.Println("Deploying Nuvolaris...")
	ioutil.WriteFile("nuvolaris.yml", NuvolarisYml, 0600)
	Task()
	return nil
}
