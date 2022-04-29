package schema

type Analysis struct {
	Task          string `bson:"task"`
	TargetIP_Port string `bson:"target_ip_port"`
	Completed     bool   `bson:"completed"`
	Paid          bool   `bson:"paid"`
	// UnitOutput    string `bson:"unit_output,omitempty"`
	// MergedOutput  string `bson:"merged_output,omitempty"`
}
