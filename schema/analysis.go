package schema

type Analysis struct {
	Task         string `bson:"task,omitempty"`
	TargetIP     string `bson:"target_ip,omitempty"`
	Completed    bool   `bson:"completed,omitempty"`
	Paid         bool   `bson:"paid,omitempty"`
	UnitOutput   string `bson:"unit_output,omitempty"`
	MergedOutput string `bson:"merged_output,omitempty"`
}
