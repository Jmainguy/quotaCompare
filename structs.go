package main

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// ClusterInfo : Information about the cluster
type ClusterInfo struct {
	NodeInfo                         map[string]NodeInfo
	ClusterAllocatableMemory         resource.Quantity
	ClusterAllocatableCPU            resource.Quantity
	ClusterAllocatablePods           resource.Quantity
	ClusterUsedCPURequests           resource.Quantity
	ClusterUsedCPU                   resource.Quantity
	ClusterUsedMemory                resource.Quantity
	ClusterUsedMemoryRequests        resource.Quantity
	ClusterUsedMemoryLimits          resource.Quantity
	ClusterUsedPods                  int64
	RqclusterAllocatedLimitsMemory   resource.Quantity
	RqclusterAllocatedLimitsCPU      resource.Quantity
	RqclusterAllocatedPods           resource.Quantity
	RqclusterAllocatedRequestsMemory resource.Quantity
	RqclusterAllocatedRequestsCPU    resource.Quantity
	NminusCPU                        resource.Quantity
	NminusMemory                     resource.Quantity
	NminusPods                       resource.Quantity
	NodeLabel                        string
}

// NodeInfo : Information about the node
type NodeInfo struct {
	AllocatableCPU     resource.Quantity
	AllocatableMemory  resource.Quantity
	AllocatablePods    resource.Quantity
	UsedPods           int64
	UsedCPU            resource.Quantity
	UsedMemory         resource.Quantity
	UsedMemoryRequests resource.Quantity
	UsedMemoryLimits   resource.Quantity
	UsedCPURequests    resource.Quantity
	PrintOutput        bool
}

// Capcity : Json to print out about metrics we gathered
type Capcity struct {
	EventKind                                string             `json:"event.kind"`
	EventModule                              string             `json:"event.module"`
	EventProvider                            string             `json:"event.provider"`
	EventType                                string             `json:"event.type"`
	EventVersion                             string             `json:"event.version"`
	ResourceQuotaCPURequestCores             int64              `json:"k8s_quota.resource_quota.cpu_request.cores"`
	ResourceQuotaCPURequestMilliCores        int64              `json:"k8s_quota.resource_quota.cpu_request.millicores"`
	ResourceQuotaMemoryRequest               int64              `json:"k8s_quota.resource_quota.memory_request"`
	ResourceQuotaMemoryLimit                 int64              `json:"k8s_quota.resource_quota.memory_limit"`
	ResourceQuotaPods                        int64              `json:"k8s_quota.resource_quota.pods"`
	SubscriptionFactorMemoryRequestTotal     float64            `json:"k8s_quota.subscription_factor.memory.request.total"`
	SubscriptionFactorMemoryRequestNminusone float64            `json:"k8s_quota.subscription_factor.memory.request.nminusone"`
	SubscriptionFactorCPURequestTotal        float64            `json:"k8s_quota.subscription_factor.cpu.request.total"`
	SubscriptionFactorCPURequestNminusone    float64            `json:"k8s_quota.subscription_factor.cpu.request.nminusone"`
	SubscriptionFactorPodsTotal              float64            `json:"k8s_quota.subscription_factor.pods.total"`
	SubscriptionFactorPodsNminusone          float64            `json:"k8s_quota.subscription_factor.pods.nminusone"`
	AllocatableMemoryTotal                   int64              `json:"k8s_quota.alloctable.memory.total"`
	AllocatableMemoryNminusone               int64              `json:"k8s_quota.alloctable.memory.nminusone"`
	AllocatableCPUTotal                      int64              `json:"k8s_quota.alloctable.cpu.total"`
	AllocatableCPUNminusone                  int64              `json:"k8s_quota.alloctable.cpu.nminusone"`
	AllocatablePodsTotal                     int64              `json:"k8s_quota.alloctable.pods.total"`
	AllocatablePodsNminusone                 int64              `json:"k8s_quota.alloctable.pods.nminusone"`
	ContainerResourceCPURequestCores         int64              `json:"k8s_quota.container_resource.cpu_request.cores"`
	ContainerResourceCPURequestMilliCores    int64              `json:"k8s_quota.container_resource.cpu_request.millicores"`
	ContainerResourceMemoryRequest           int64              `json:"k8s_quota.container_resource.memory_request"`
	ContainerResourceMemoryLimit             int64              `json:"k8s_quota.container_resource.memory_limit"`
	ContainerResourcePods                    int64              `json:"k8s_quota.container_resource.pods"`
	NodeLabel                                string             `json:"k8s_quota.node_label"`
	UtilizationFactorPods                    map[string]float64 `json:"k8s_quota.utilization_factor.pods"`
	UtilizationFactorPodsTotal               float64            `json:"k8s_quota.utilization_factor.pods.total"`
	UtilizationFactorPodsNminusone           float64            `json:"k8s_quota.utilization_factor.pods.nminusone"`
	UtilizationFactorMemoryRequests          map[string]float64 `json:"k8s_quota.utilization_factor.memory_request"`
	UtilizationFactorMemoryRequestsTotal     float64            `json:"k8s_quota.utilization_factor.memory_request.total"`
	UtilizationFactorMemoryRequestsNminusone float64            `json:"k8s_quota.utilization_factor.memory_request.nminusone"`
	UtilizationFactorCPURequests             map[string]float64 `json:"k8s_quota.utilization_factor.cpu_request"`
	UtilizationFactorCPURequestsTotal        float64            `json:"k8s_quota.utilization_factor.cpu_request.total"`
	UtilizationFactorCPURequestsNminusone    float64            `json:"k8s_quota.utilization_factor.cpu_request.nminusone"`
	AvailableMemoryRequestTotal              int64              `json:"k8s_quota.available.memory_request.total"`
	AvailableMemoryRequestNminusone          int64              `json:"k8s_quota.available.memory_request.nminusone"`
	AvailableCPURequestTotal                 int64              `json:"k8s_quota.available.cpu_request.total"`
	AvailableCPURequestNminusone             int64              `json:"k8s_quota.available.cpu_request.nminusone"`
	AvailablePodsTotal                       int64              `json:"k8s_quota.available.pods.total"`
	AvailablePodsNminusone                   int64              `json:"k8s_quota.available.pods.nminusone"`
}
