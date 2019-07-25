package internal

const DefaultSubjectTemplate = `
Elixir CI: [{{.Repo.FullName}}] Build failed ({{.CheckSuite.HeadBranch}} - {{.CheckSuite.HeadSHA}})
`

const DefauleEmailMarkdownTemplate = `
<p>
  <h2>{{.Event.Repo.FullName}} ({{.CheckSuite.HeadBranch}})</h2>

  <p><a href="{{.Commit.HTMLURL}}">{{.Commit.SHA}})</a> - {{.Commit.Message}}</p>

  <p>The following checks failed:</p>

  <h3>{{ .CheckRun.Name }}</h3>

  <p><b style="color: #cb2431">{{ .CheckRun.Conclusion }}</b> ran at <b>{{.CheckRun.StartedAt}}</b> in <b>{{.Duration}}</b></p>

  <p><b>{{ .CheckRun.Output.Title }}</b></p>
  <p>{{ .CheckRun.Output.Summary }}</p>
  <p>{{ .CheckRun.Output.Text }}</p>
</p>

<p style="font-size:small;-webkit-text-size-adjust:none;color:#666;">
  &mdash;
  <br />
  You are receiving this because you are the author of the commit.
  <br />
  <a href="{{ .CheckSuite.HTMLURL }}">View it on GitHub</a>
</p>
`
