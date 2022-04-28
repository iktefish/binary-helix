package workers

import (
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	// "github.com/iktefish/binary-helix/client"
	"github.com/iktefish/binary-helix/schema"
	"github.com/iktefish/binary-helix/utils"
)

/* This function will send the splits to the Db and servers simultaneously */
func Carrier(ss []string, an string, others ...string) bool {
    fmt.Println("others ~~> ", others)
	var wg sync.WaitGroup
	wg.Add(len(ss))

	if utils.Check_AnalyzerList(an) != true {
		log.Fatal("FAIL: No such analysis present!")
	}

	nodes := utils.Get_ActiveNodes()

	computationId := uuid.New().String()
	var analysisArts []schema.Analysis

	for _, node := range nodes {
		analysisArts = append(analysisArts, schema.Analysis{
			Task:         an,
			TargetIP_Port:     node.TargetIP_Port,
			Completed:    false,
			Paid:         false,
			UnitOutput:   "",
			MergedOutput: "",
		})
	}

	// computationId := uuid.New().String()
	// analysisArt := schema.Analysis{
	// 	Task:         an,
	// 	TargetIP:     "172.17.0.2",
	// 	Completed:    false,
	// 	Paid:         false,
	// 	UnitOutput:   "",
	// 	MergedOutput: "",
	// }

	for i, s := range ss {
		go splitToDb(i, s, &wg, computationId, analysisArts[i])
	}

	wg.Wait()
	return true
}

func splitToDb(i int, s string, wg *sync.WaitGroup, cId string, aArt schema.Analysis) {
	defer wg.Done()

	client, ctx := utils.ConnectDb()
	defer client.Disconnect(ctx)

	slicesDb := client.Database("slices_db")
	slices := slicesDb.Collection("slices")

	slice := schema.Slices{
		ComputationId: cId,
		SplitOrder:    int32(i + 1),
		Content:       s,
		AnalysisArt:   aArt,
	}

	/* Insert slice in slices collection */
	_, err := slices.InsertOne(ctx, slice)
	utils.HandleError(err)
}

func splitToServer(i, int, s string, wg *sync.WaitGroup, cId string, aArt schema.Analysis) {
	defer wg.Done()

	// client.CheckServers()
}
