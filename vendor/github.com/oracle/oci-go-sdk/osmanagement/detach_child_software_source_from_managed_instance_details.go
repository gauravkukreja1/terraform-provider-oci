// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DetachChildSoftwareSourceFromManagedInstanceDetails Information for detaching a software source from a managed instance
type DetachChildSoftwareSourceFromManagedInstanceDetails struct {

	// OCID for the Software Source
	SoftwareSourceId *string `mandatory:"true" json:"softwareSourceId"`
}

func (m DetachChildSoftwareSourceFromManagedInstanceDetails) String() string {
	return common.PointerString(m)
}
