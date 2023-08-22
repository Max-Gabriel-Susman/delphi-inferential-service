package generate

type GenerateInferenceRequest struct {
	Inputs     string                        `json:"inputs"`
	Parameters InferenceGenerationParameters `json:"parameters"`
}

type InferenceGenerationParameters struct {
	BestOf               int      `json:"best_of"`
	DecoderInputeDetails bool     `json:"decoder_input_details"`
	Details              bool     `json:"details"`
	DoSample             bool     `json:"do_sample"`
	MaxNewTokens         int      `json:"max_new_tokens"`
	RepetitionPenalty    float32  `json:"repetition_penalty"`
	ReturnFullText       bool     `json:"return_full_text"`
	Seed                 *string  `json:"seed"` // verify type of this example value was null
	Stop                 []string `json:"stop"`
	Temperature          float32  `json:"temperature"`
	TopK                 float32  `json:"top_k"` // verify type, example vaulue was int but peers were floats
	TopP                 float32  `json:"top_p"`
	Truncate             *string  `json:"truncate"` // verify type of this example value was null
	TypicalP             float32  `json:"typical_p"`
	Watermark            bool     `json:"watermark"`
	// make sure we're not missing any fields: TODO
}
