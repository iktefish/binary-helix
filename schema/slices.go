package schema

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Slices struct {
	// ID          primitive.ObjectID `bson:"_id,omitempty"`
	ComputationId string   `bson:"computation_id,omitempty"`
	SplitOrder    int32    `bson:"split_order,omitempty"`
	Content       string   `bson:"content,omitempty"`
	AnalysisArt   Analysis `bson:"analysis_art,omitempty"` // NOTE: Art for Artifact
}
