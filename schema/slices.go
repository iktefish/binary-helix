package schema

type Slices struct {
	ComputationId string   `bson:"computation_id,omitempty"`
	SplitOrder    int32    `bson:"split_order,omitempty"`
	Content       string   `bson:"content,omitempty"`
	AnalysisArt   Analysis `bson:"analysis_art,omitempty"` // NOTE: Art for Artifact
	MergedOutput  []string `bson:"merged_output,omitempty"`
}
