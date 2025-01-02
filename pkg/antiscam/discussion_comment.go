package antiscam

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-github/v50/github"
)

func (a *Antiscam) ProcessDiscussionComment(payload []byte) error {
	var discussion_comment github.DiscussionComment
	if err := json.Unmarshal(payload, &discussion_comment); err != nil {
		return err
	}

	// discussion_comment.get

	var detections []Detection
	detections = append(detections, checkComment(discussion_comment.GetBody())...)

	// body := fmt.Sprintf("@%s The previous user tried to scam you by providing a fake support link. Don't interact with it.\n", discussion_comment.GetAuthor().GetLogin())

	for _, detection := range detections {
		fmt.Printf("Detected scam in %s: %s\n", detection.Location, detection.DebugInfo)
	}

	// if len(detections) > 0 {
	// 	a.client.Teams.DeleteCommentByID(a.ctx, )

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
