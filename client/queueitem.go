package client

import "github.com/morzik45/comfy2go/graphapi"

type QueueItem struct {
	PromptID   string                 `json:"prompt_id"`
	Number     int                    `json:"number"`
	NodeErrors map[string]interface{} `json:"node_errors"`
	Messages   chan PromptMessage     `json:"-"`
	Workflow   *graphapi.Graph        `json:"-"`
}
