package generate

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorType string `json:"error_type"`
}

type GeneratedInferenceResponse struct {
	Details            GeneratedInferenceResponseDetails `json:"details"`
	GeneratedInference string                            `json:"generated_text"`
}

func (resp *GeneratedInferenceResponse) ConcatenateTokens() string {
	var concatenated string
	for _, token := range resp.Details.Tokens {
		concatenated += token.Text + " "
	}
	return concatenated
}

type BestOfSequence struct {
	FinishReason       string           `json:"finish_reason"`
	GeneratedInference string           `json:"generated_text"`
	GeneratedTokens    int              `json:"generated_tokens"`
	Prefill            []PrefillElement `json:"prefill"`
	Seed               int              `json:"seed"`
	Tokens             []Token          `json:"tokens"`
}

type PrefillElement struct { // TODO: perhaps a better name can be found from the falcon docs?
	ID      int     `json:"id"`
	LogProb float32 `json:"logprob"`
	Text    string  `json:"text"`
}

type Token struct {
	ID      int     `json:"id"`
	LogProb float32 `json:"logprob"`
	Special bool    `json:"special"`
	Text    string  `json:"text"`
}

type GeneratedInferenceResponseDetails struct {
	BestOfSequences []BestOfSequence `json:"best_of_sequences"`
	FinishReason    string           `json:"finish_reason"`
	GeneratedTokens int              `json:"generated_tokens"`
	Prefill         []PrefillElement `json:"prefill"`
	Seed            int              `json:"seed"`
	Tokens          []Token          `json:"tokens"`
}
