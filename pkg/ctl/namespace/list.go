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

package namespace

import (
	"github.com/olekukonko/tablewriter"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func getNamespacesFromTenant(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Get the list of namespaces of a tenant"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []pulsar.Example

	list := pulsar.Example{
		Desc:    "Get the list of namespaces of a tenant",
		Command: "pulsarctl namespaces list <tenant name>",
	}

	examples = append(examples, list)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "+------------------+\n" +
			"|  NAMESPACE NAME  |\n" +
			"+------------------+\n" +
			"| public/default   |\n" +
			"| public/functions |\n" +
			"+------------------+",
	}

	notTenantName := pulsar.Output{
		Desc: "you must specify a tenant name, please check if the tenant name is provided",
		Out:  "[✖]  only one argument is allowed to be used as a name",
	}

	notExistTenantName := pulsar.Output{
		Desc: "the tenant name not exist, please check if tenant name exists",
		Out:  "[✖]  code: 404 reason: Tenant does not exist",
	}

	out = append(out, successOut, notTenantName, notExistTenantName)
	desc.CommandOutput = out

	vc.SetDescription(
		"list",
		"Get the list of namespaces of a tenant",
		desc.ToString(),
		"list",
	)

	vc.SetRunFuncWithNameArg(func() error {
		return doListNamespaces(vc)
	})
}

func doListNamespaces(vc *cmdutils.VerbCmd) error {
	tenant := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	listNamespaces, err := admin.Namespaces().GetNamespaces(tenant)
	if err == nil {
		table := tablewriter.NewWriter(vc.Command.OutOrStdout())
		table.SetHeader([]string{"Namespace Name"})
		for _, ns := range listNamespaces {
			table.Append([]string{ns})
		}
		table.Render()
	}
	return err
}