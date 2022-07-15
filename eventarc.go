package main

import "time"

const (
	EventarcTypeAuditLog           = "type.googleapis.com/google.cloud.audit.AuditLog"
	EventarcMethodNameJobCompleted = "jobservice.jobcompleted"
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
			CallerNetwork           string `json:"callerNetwork"`
			RequestAttributes       struct {
			} `json:"requestAttributes"`
			DestinationAttributes struct {
			} `json:"destinationAttributes"`
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
						ReferencedViews      []struct {
							ProjectID string `json:"projectId"`
							DatasetID string `json:"datasetId"`
							TableID   string `json:"tableId"`
						} `json:"referencedViews"`
						TotalViewsProcessed int `json:"totalViewsProcessed"`
						ReservationUsage    []struct {
							Name   string `json:"name"`
							SlotMs string `json:"slotMs"`
						} `json:"reservationUsage"`
						QueryOutputRowCount string `json:"queryOutputRowCount"`
					} `json:"jobStatistics"`
				} `json:"job"`
			} `json:"jobCompletedEvent"`
		} `json:"serviceData"`
	} `json:"protoPayload"`
	InsertID string `json:"insertId"`
	Resource struct {
		Type   string `json:"type"`
		Labels struct {
			ProjectID string `json:"project_id"`
		} `json:"labels"`
	} `json:"resource"`
	Timestamp        time.Time `json:"timestamp"`
	Severity         string    `json:"severity"`
	LogName          string    `json:"logName"`
	ReceiveTimestamp time.Time `json:"receiveTimestamp"`
}
