package client

import "github.com/morzik45/comfy2go/graphapi"

type QueueItem struct {
	PromptID   string                 `json:"prompt_id"`
	Number     int                    `json:"number"`
	NodeErrors map[string]interface{} `json:"node_errors"`
	Messages   chan PromptMessage     `json:"-"`
	Workflow   *graphapi.Graph        `json:"-"`

	mu     sync.Mutex
	closed bool
}

func (qi *QueueItem) SendMessage(msg PromptMessage) {
	qi.mu.Lock()
	defer qi.mu.Unlock()
	if qi.closed {
		return
	}
	select {
	case qi.Messages <- msg:
	default:
		// optionally log or drop silently
	}
}

func (qi *QueueItem) CloseMessages() {
	qi.mu.Lock()
	defer qi.mu.Unlock()
	if !qi.closed {
		close(qi.Messages)
		qi.closed = true
	}
}