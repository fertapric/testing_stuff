package internal

const DefaultSubjectTemplate = `
Elixir CI: [{{.Repo.FullName}}] Build failed ({{.CheckSuite.HeadBranch}} - {{printf \"%.6s\" .CheckSuite.HeadSHA}})
`

const DefauleEmailMarkdownTemplate = `
<p>
  <h2>{{.CheckSuite.Repository.FullName}} ({{.CheckSuite.HeadBranch}} - {{printf \"%.6s\" .CheckSuite.HeadSHA}})</h2>

  <p>The following checks failed:</p>

  <h3>{{ .CheckRun.Name }}</h3>

  <p><b style="color: #cb2431">{{ .CheckRun.Conclusion }}</b> ran at {{.CheckRun.StartedAt}} in {{.Duration}}</p>

  <b>{{ .CheckRun.Output.Title }}</b>
  <p>{{ .CheckRun.Output.Summary }}</p>
  <p>{{ .CheckRun.Output.Text }}</p>
</p>

<p style="font-size:small;-webkit-text-size-adjust:none;color:#666;">
  &mdash;
  <br />
  You are receiving this because you are the author of the commit.
  <br />
  <a href="{{ .CheckRun.HTMLURL }}">View it on GitHub</a>
</p>
`
