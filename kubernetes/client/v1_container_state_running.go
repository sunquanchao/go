/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

// ContainerStateRunning is a running state of a container.
type V1ContainerStateRunning struct {

	// Time at which the container was last (re-)started
	StartedAt time.Time `json:"startedAt,omitempty"`
}
