package schema

type Analysis struct {
	Task         string `bson:"task"`
	TargetIP     string `bson:"target_ip"`
	Completed    bool   `bson:"completed"`
	Paid         bool   `bson:"paid"`
	UnitOutput   string `bson:"unit_output,omitempty"`
	MergedOutput string `bson:"merged_output,omitempty"`
}
