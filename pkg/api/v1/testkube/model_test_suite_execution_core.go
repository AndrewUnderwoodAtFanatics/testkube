/*
 * Testkube API
 *
 * Testkube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

import (
	"time"
)

// test suite execution core
type TestSuiteExecutionCore struct {
	// execution id
	Id string `json:"id,omitempty"`
	// test suite execution start time
	StartTime time.Time `json:"startTime,omitempty"`
	// test suite execution end time
	EndTime time.Time                 `json:"endTime,omitempty"`
	Status  *TestSuiteExecutionStatus `json:"status,omitempty"`
}