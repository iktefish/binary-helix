package schema

type Nodes struct {
	NodeName                string `bson:"node_name,omitempty"`
	TargetIP_Port           string `bson:"target_ip_port,omitempty"`
	Active                  bool   `bson:"active,omitempty"`
	TotalCreditAttained     int32  `bson:"total_credit_attained,omitempty"`
	TaskCompletionFrequency int32  `bson:"task_completion_frequency,omitempty"`
}
