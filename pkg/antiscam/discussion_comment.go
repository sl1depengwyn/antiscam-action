package antiscam

import (
	"encoding/json"
	"fmt"
	// "github.com/google/go-github/v50/github"
	"github.com/shurcooL/githubv4"
)

func (a *Antiscam) ProcessDiscussionComment(payload []byte) error {
	var event map[string]map[string]string

	if err := json.Unmarshal(payload, &event); err != nil {
		return err
	}

	var mutation struct {
		DeleteDiscussionComment struct {
			ClientMutationID *string
		} `graphql:"deleteDiscussionComment(input: $input)"`
	}

	input := githubv4.DeleteDiscussionCommentInput{
		ID: githubv4.ID(event["comment"]["node_id"]),
	}

	err := a.graphql_client.Mutate(a.ctx, &mutation, input, nil)

	if err != nil {
		fmt.Println("Error deleting discussion comment: %v", err)
	}

	fmt.Println("Discussion comment deleted successfully!")
	// discussion_comment.get

	// var detections []Detection
	// detections = append(detections, checkComment(discussion_comment.GetBody())...)

	// body := fmt.Sprintf("@%s The previous user tried to scam you by providing a fake support link. Don't interact with it.\n", discussion_comment.GetAuthor().GetLogin())

	// for _, detection := range detections {
	// 	fmt.Printf("Detected scam in %s: %s\n", detection.Location, detection.DebugInfo)
	// }

	// if len(detections) > 0 {
	// a.client.Teams.DeleteCommentByID(a.ctx, )

	// 	a.client.Issues.DeleteComment(
	// 		a.ctx,
	// 		event.GetRepo().GetOwner().GetLogin(),
	// 		event.GetRepo().GetName(),
	// 		event.GetComment().GetID(),
	// 	)

	// 	if _, _, err := a.client.Issues.CreateComment(
	// 		a.ctx,
	// 		event.GetRepo().GetOwner().GetLogin(),
	// 		event.GetRepo().GetName(),
	// 		event.GetIssue().GetNumber(),
	// 		&github.IssueComment{
	// 			Body: &body,
	// 		},
	// 	); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
