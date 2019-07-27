package internal

import (
	"context"
	"fmt"
	// "github.com/gomarkdown/markdown"
	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
	"gopkg.in/gomail.v2"
	"log"
	"strings"
	"time"
)

func generateEmail(spec Specification, event github.CheckSuiteEvent, commit github.Commit) (*gomail.Message, error) {
	message := gomail.NewMessage()

	message.SetAddressHeader("From", spec.MailFrom, "Elixir CI")
	message.SetAddressHeader("To", *commit.Committer.Email, *commit.Committer.Name)
	log.Printf("Creating email for %v <%v>...\n", *commit.Committer.Name, *commit.Committer.Email)

	// populate subject:
	subject, err := Render(DefaultSubjectTemplate, event)
	if err != nil {
		return nil, fmt.Errorf("could not render subject template: %v", err)
	}
	message.SetHeader("Subject", subject)

	githubClient := newGitHubClient(spec)
	listCheckRunsResults, _, err := githubClient.Checks.ListCheckRunsCheckSuite(context.Background(), *event.Repo.Owner.Login, *event.Repo.Name, *event.CheckSuite.ID, nil)
	if err != nil {
		return nil, fmt.Errorf("could not get check runs: %v", err)
	}

	contentBuilder := strings.Builder{}

  dataHeader := struct {
      CheckSuite *github.CheckSuite
      Commit github.Commit
      Event github.CheckSuiteEvent
      SHA string
    }{
      CheckSuite: event.CheckSuite,
      Commit: commit,
      Event: event,
      SHA: fmt.Sprintf("%.6s", *event.CheckSuite.HeadSHA),
    }
  headerPart, err := Render(EmailHeaderTemplate, dataHeader)
  if err != nil {
    return nil, fmt.Errorf("could not render header template: %v", err)
  }

  contentBuilder.WriteString(headerPart)
  contentBuilder.WriteString("\n")

	for _, checkRun := range listCheckRunsResults.CheckRuns {
		var duraiton time.Duration

    if *checkRun.Conclusion == "failure" {

  		if checkRun.CompletedAt != nil {
  			duraiton = checkRun.CompletedAt.Sub(checkRun.StartedAt.Time)
  		}

      // detailsBuilder := strings.Builder{}
      // detailsBuilder.WriteString(*checkRun.Output.Text)

  		data := struct {
  			CheckRun *github.CheckRun
        CheckSuite *github.CheckSuite
  			Duration time.Duration
        // Details string
        Commit github.Commit
        Event github.CheckSuiteEvent
  		}{
  			CheckRun: checkRun,
        CheckSuite: event.CheckSuite,
  			Duration: duraiton,
        Commit: commit,
        Event: event,
        // Details: "foo", //string(markdown.ToHTML([]byte(detailsBuilder.String()), nil, nil)),
  		}
  		contentPart, err := Render(EmailCheckTemplate, data)
  		if err != nil {
  			return nil, fmt.Errorf("could not render content template for check run '%s': %v", *checkRun.Name, err)
  		}
  		contentBuilder.WriteString(contentPart)
  		contentBuilder.WriteString("\n")
    }
	}

  dataFooter := struct {
    WhyReceiving string
  }{
    WhyReceiving: "the author of this commit",
  }
  footerPart, err := Render(EmailFooterTemplate, dataFooter)
  if err != nil {
    return nil, fmt.Errorf("could not render footer template: %v", err)
  }

  contentBuilder.WriteString(footerPart)
  contentBuilder.WriteString("\n")

	contentMD := contentBuilder.String()
	message.AddAlternative("text/text", contentMD)
	// contentHTML := markdown.ToHTML([]byte(contentMD), nil, nil)
	message.AddAlternative("text/html", string(contentMD))
	return message, nil
}

func newGitHubClient(spec Specification) *github.Client {
	if spec.GitHubToken == "" {
		return github.NewClient(nil)
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: spec.GitHubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}
