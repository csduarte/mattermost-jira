package bridge

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
)

// Webhook structure repsents the webhook format set by Atlassian
type Webhook struct {
	WebhookEvent string
	User         struct {
		Name        string
		AvatarUrls  map[string]string
		DisplayName string
	}
	Issue struct {
		Self   string
		Key    string
		Fields struct {
			Issuetype struct {
				IconURL string
				Name    string
			}
			Summary string
		}
	}
	Comment struct {
		Body string
	}
	Changelog struct {
		Items []struct {
			Field      string
			FromString string
			ToString   string
		}
	}
}

// NewWebhookfromJSON decodes io to a webhook struct.
func NewWebhookfromJSON(d io.ReadCloser) (*Webhook, error) {
	decoder := json.NewDecoder(d)
	var w Webhook
	err := decoder.Decode(&w)
	return &w, err
}

// MDUserIcon format the user avatar to a markdown image.
func (w *Webhook) MDUserIcon() string {
	return fmt.Sprintf("![user_icon](%s)", w.User.AvatarUrls["16x16"])
}

// MDUserLink format the user link to a markdown link.
func (w *Webhook) MDUserLink() string {
	u, _ := url.Parse(w.Issue.Self)
	return fmt.Sprintf("[%s](%s://%s/secure/ViewProfile.jspa?name=%s)",
		w.User.DisplayName, u.Scheme, u.Host, w.User.Name)
}

// MDAction format the webhook event to an action stream.
func (w *Webhook) MDAction() string {
	var action string
	switch w.WebhookEvent {
	case "jira:issue_created":
		action = "created"
	case "jira:issue_updated":
		action = "updated"
	case "jira:issue_deleted":
		action = "deleted"
	}
	return action
}

// MDIssueType Get the name value and lowercase the IssueType.
func (w *Webhook) MDIssueType() string {
	return strings.ToLower(w.Issue.Fields.Issuetype.Name)
}

// MDTaskIcon formatss the Task Icon as a markdown image.
func (w *Webhook) MDTaskIcon() string {
	return fmt.Sprintf("![task_icon](%s)", w.Issue.Fields.Issuetype.IconURL)
}

// MDIssueLink formats the issue link as a markdown link.
func (w *Webhook) MDIssueLink() string {
	u, _ := url.Parse(w.Issue.Self)
	return fmt.Sprintf("[%s](%s://%s/browse/%s)",
		w.Issue.Key,
		u.Scheme,
		u.Host,
		w.Issue.Key,
	)
}

// MDSummary Pulls out the summary from the webhook.
func (w *Webhook) MDSummary() string {
	return fmt.Sprintf("%q", w.Issue.Fields.Summary)
}

// MDChangelog Will generate the strikethrough markdown for changes.
func (w *Webhook) MDChangelog() string {
	var changelog string
	if len(w.Changelog.Items) < 1 {
		return changelog
	}
	for _, item := range w.Changelog.Items {
		itemName := strings.ToUpper(string(item.Field[0])) + item.Field[1:]
		if item.FromString == "" {
			item.FromString = "None"
		}
		if itemName == "Description" {
			changelog += fmt.Sprintf(
				"\nNew Description:\n```\n%s\n```",
				item.ToString,
			)
		} else {
			changelog += fmt.Sprintf(
				"\n%s: ~~%s~~ %s",
				itemName,
				item.FromString,
				item.ToString,
			)
		}
	}
	return changelog
}

// MDComment forms the comment body as a new line markdown
func (w *Webhook) MDComment() string {
	var comment string
	if len(w.Comment.Body) > 0 {
		comment = fmt.Sprintf("\nComment:\n```\n%s\n```", w.Comment.Body)
	}
	return comment
}

/*Text format/Expired (Icons are no longer used):
![user_icon](user_icon_link)[UserFirstName UserSecondName](user_link) commented task ![task_icon](task_icon link)[TSK-42](issue_link) "Test task" Status Comments
*/
/*Text format:
[UserFirstName UserSecondName](user_link) commented task [TSK-42](issue_link) "Test task" Status Comments
*/
func (w *Webhook) String() string {
	return fmt.Sprintf("%s %s %s %s %s%s%s",
		w.MDUserLink(), w.MDAction(), w.MDIssueType(),
		w.MDIssueLink(), w.MDSummary(), w.MDChangelog(),
		w.MDComment())
}
