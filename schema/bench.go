package schema

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Bench struct {
	NodeName      string             `bson:"node_name,omitempty"`
	TargetIP      string             `bson:"target_ip,omitempty"`
	NetSpeedToS   int32              `bson:"net_speed_to_s,omitempty"`
	NetSpeedFromS int32              `bson:"net_speed_from_s,omitempty"`
	RamTotal      int64              `bson:"ram_usage,omitempty"`
	RamUsed       int64              `bson:"ram_used,omitempty"`
	RamCached     int64              `bson:"ram_cached,omitempty"`
	RamFree       int64              `bson:"ram_free,omitempty"`
	CpuUser       float64            `bson:"cpu_user,omitempty"`
	CpuSystem     float64            `bson:"cpu_system,omitempty"`
	CpuIdle       float64            `bson:"cpu_idle,omitempty"`
	TimeStamp     primitive.ObjectID `bson:"time_stamp,omitempty"`
}

/* God please for this function name! */
func Test_TimeToPrim() {

	now := primitive.NewDateTimeFromTime((time.Now()))
	nowObj := primitive.NewObjectIDFromTimestamp(time.Now())

	fmt.Println("primitive.NewObjectIDFromTimestamp(time.Now()) ~~> ", nowObj)
	fmt.Println("primitive.NewDateTimeFromTime((time.Now())) ~~> ", now)

}
