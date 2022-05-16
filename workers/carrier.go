package workers

import (
	// "fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	// "github.com/iktefish/binary-helix/blockchain"
	"github.com/iktefish/binary-helix/client"
	"github.com/iktefish/binary-helix/schema"
	"github.com/iktefish/binary-helix/utils"
)

/* This function will send the splits to the Db and servers simultaneously */
func Carrier(ss []string, an string, extra string) bool {
	// fmt.Println("extra ~~> ", extra)
	var wg sync.WaitGroup
	// wg.Add(len(ss))

	if an == "trial" {
		return true
	}

	if utils.Check_AnalyzerList(an) != true {
		log.Fatal("FAIL: No such analysis present!")
	}

	nodes := utils.Get_ActiveNodes()

	computationId := uuid.New().String()
	var analysisArts []schema.Analysis

	for _, node := range nodes {
		analysisArts = append(analysisArts, schema.Analysis{
			Task:          an,
			TargetIP_Port: node.TargetIP_Port,
			Completed:     false,
			Paid:          false,
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
		wg.Add(1)
		go splitToDb(i, s, &wg, computationId, analysisArts[i])
	}

	// for i, _ := range ss {
	// 	chain := blockchain.InitBlockChain()
	// 	chain.AddBlock(computationId, analysisArts[i].TargetIP_Port)

	// 	for i, block := range chain.Blocks {
	// 		if i == 0 {
	// 			fmt.Println("\tBLOCKCHAIN:\t\n")
	// 		}
	// 		fmt.Printf("\tblock.OldHash\t%v ~~>\t%x\n", i, block.OldHash)
	// 		fmt.Printf("\tblock.C_Id\t%v ~~>\t%x\n", i, block.C_Id)
	// 		fmt.Printf("\tblock.Ip_Port\t%v ~~>\t%x\n", i, block.Ip_Port)
	// 		fmt.Printf("\tblock.Hash\t%v ~~>\t%x\n", i, block.Hash)
	// 		fmt.Println("\n")
	// 	}
	// }
	splitToServer(ss, computationId, analysisArts, extra)

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

func splitToServer(s []string, cId string, aArt []schema.Analysis, extra string) {
	client.TaskServer(s, cId, aArt, extra)
}
