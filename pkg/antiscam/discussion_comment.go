package antiscam

import (
	"encoding/json"
	"fmt"

	"github.com/google/go-github/v50/github"
)

func (a *Antiscam) ProcessDiscussionComment(payload []byte) error {
	var discussion_event github.DiscussionEvent
	var team_event github.TeamEvent
	var issue_comment github.IssueCommentEvent

	// github.Event

	fmt.Printf("raw event payload: %s\n", payload)

	if err := json.Unmarshal(payload, &discussion_event); err != nil {
		return err
	}

	fmt.Printf("discussion event: %s\n", github.Stringify(discussion_event))

	if err := json.Unmarshal(payload, &team_event); err != nil {
		return err
	}

	fmt.Printf("team event: %s\n", github.Stringify(team_event))

	if err := json.Unmarshal(payload, &issue_comment); err != nil {
		return err
	}

	fmt.Printf("issue comment: %s\n", github.Stringify(issue_comment))

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
