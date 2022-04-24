package schema

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Bench struct {
	NodeName  string             `bson:"node_name,omitempty"`
	TargetIP  string             `bson:"target_ip,omitempty"`
	NetSpeed  int32              `bson:"net_speed,omitempty"`
	Latency   int32              `bson:"latency,omitempty"`
	RamUsage  int32              `bson:"ram_usage,omitempty"`
	CpuUsage  int32              `bson:"cpu_usage,omitempty"`
	TimeStamp primitive.ObjectID `bson:"time_stamp,omitempty"`
}

/* God please for this function name! */
func TestingTimeToPrimitiveObjAndViceVersa_Example() {

	now := primitive.NewDateTimeFromTime((time.Now()))
	nowObj := primitive.NewObjectIDFromTimestamp(time.Now())

	fmt.Println("primitive.NewObjectIDFromTimestamp(time.Now()) ~~> ", nowObj)
	fmt.Println("primitive.NewDateTimeFromTime((time.Now())) ~~> ", now)

}
