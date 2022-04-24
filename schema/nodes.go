package schema

type Nodes struct {
	NodeName                string `bson:"node_name,omitempty"`
	TargetIP                string `bson:"target_ip,omitempty"`
    TotalCreditAttained     int32  `bson:"total_credit_attained,omitempty"` // NOTE: We can leave this empty!
	TaskCompletionFrequency int32 `bson:"task_completion_frequency,omitempty"`
}
