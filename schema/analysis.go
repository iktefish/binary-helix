package schema

type Analysis struct {
	Task          string `bson:"task"`
	TargetIP_Port string `bson:"target_ip_port"`
	Completed     bool   `bson:"completed"`
	Paid          bool   `bson:"paid"`
}
