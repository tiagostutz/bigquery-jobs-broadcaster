package main

import "time"

const (
	EventarcType       = "type.googleapis.com/google.cloud.audit.AuditLog"
	EventarcMethodName = "jobservice.jobcompleted"
)

type EventarcPayload struct {
	ProtoPayload struct {
		Type   string `json:"@type"`
		Status struct {
		} `json:"status"`
		AuthenticationInfo struct {
			PrincipalEmail string `json:"principalEmail"`
		} `json:"authenticationInfo"`
		RequestMetadata struct {
			CallerIP                string `json:"callerIp"`
			CallerSuppliedUserAgent string `json:"callerSuppliedUserAgent"`
		} `json:"requestMetadata"`
		ServiceName  string `json:"serviceName"`
		MethodName   string `json:"methodName"`
		ResourceName string `json:"resourceName"`
		ServiceData  struct {
			Type              string `json:"@type"`
			JobCompletedEvent struct {
				EventName string `json:"eventName"`
				Job       struct {
					JobName struct {
						ProjectID string `json:"projectId"`
						JobID     string `json:"jobId"`
						Location  string `json:"location"`
					} `json:"jobName"`
					JobConfiguration struct {
						Query struct {
							Query            string `json:"query"`
							DestinationTable struct {
								ProjectID string `json:"projectId"`
								DatasetID string `json:"datasetId"`
								TableID   string `json:"tableId"`
							} `json:"destinationTable"`
							CreateDisposition string `json:"createDisposition"`
							WriteDisposition  string `json:"writeDisposition"`
							DefaultDataset    struct {
							} `json:"defaultDataset"`
							QueryPriority string `json:"queryPriority"`
							StatementType string `json:"statementType"`
						} `json:"query"`
					} `json:"jobConfiguration"`
					JobStatus struct {
						State string `json:"state"`
						Error struct {
							Message string `json:"message"`
						} `json:"error"`
					} `json:"jobStatus"`
					JobStatistics struct {
						CreateTime          time.Time `json:"createTime"`
						StartTime           time.Time `json:"startTime"`
						EndTime             time.Time `json:"endTime"`
						TotalProcessedBytes string    `json:"totalProcessedBytes"`
						TotalBilledBytes    string    `json:"totalBilledBytes"`
						BillingTier         int       `json:"billingTier"`
						TotalSlotMs         string    `json:"totalSlotMs"`
						ReferencedTables    []struct {
							ProjectID string `json:"projectId"`
							DatasetID string `json:"datasetId"`
							TableID   string `json:"tableId"`
						} `json:"referencedTables"`
						TotalTablesProcessed int `json:"totalTablesProcessed"`
						ReservationUsage     []struct {
							Name   string `json:"name"`
							SlotMs string `json:"slotMs"`
						} `json:"reservationUsage"`
						QueryOutputRowCount string `json:"queryOutputRowCount"`
					} `json:"jobStatistics"`
				} `json:"job"`
			} `json:"jobCompletedEvent"`
		} `json:"serviceData"`
	}
}
