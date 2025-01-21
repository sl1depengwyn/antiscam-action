package antiscam

import (
	"encoding/json"
	"fmt"
	// "github.com/google/go-github/v50/github"
	"github.com/shurcooL/githubv4"
)

func (a *Antiscam) ProcessDiscussionComment(payload []byte) error {
	var event map[string]interface{}
	if err := json.Unmarshal(payload, &event); err != nil {
		return err
	}

	var detections []Detection
	detections = append(detections, checkComment(event["comment"].(map[string]interface{})["body"].(string))...)

	for _, detection := range detections {
		fmt.Printf("Detected scam in %s: %s\n", detection.Location, detection.DebugInfo)
	}

	if len(detections) > 0 {
		var comment_deletion struct {
			DeleteDiscussionComment struct {
				ClientMutationID *string
			} `graphql:"deleteDiscussionComment(input: $input)"`
		}

		comment_deletion_input := githubv4.DeleteDiscussionCommentInput{
			ID: githubv4.ID(event["comment"].(map[string]interface{})["node_id"].(string)),
		}

		comment_deletion_err := a.graphql_client.Mutate(a.ctx, &comment_deletion, comment_deletion_input, nil)

		if comment_deletion_err != nil {
			return comment_deletion_err
		}

		var comment_addition struct {
			AddDiscussionComment struct {
				Comment struct {
					ID   githubv4.ID
					Body githubv4.String
					URL  githubv4.URI
				}
				ClientMutationID *string
			} `graphql:"addDiscussionComment(input: $input)"`
		}

		body := fmt.Sprintf("@%s The previous user tried to scam you by providing a fake support link. Don't interact with it.\n", event["comment"].(map[string]interface{})["user"].(map[string]interface{})["login"].(string))

		comment_addition_input := githubv4.AddDiscussionCommentInput{
			DiscussionID: githubv4.ID(event["discussion"].(map[string]interface{})["node_id"].(string)),
			Body:         githubv4.String(body),
		}

		comment_addition_err := a.graphql_client.Mutate(a.ctx, &comment_addition, comment_addition_input, nil)
		if comment_addition_err != nil {
			return comment_addition_err
		}
	}

	return nil
}
