package internal

const DefaultSubjectTemplate = `
Elixir CI: [{{.Repo.FullName}}] Build failed ({{.CheckSuite.HeadBranch}} - {{.CheckSuite.HeadSHA}})
`

const DefauleEmailMarkdownTemplate = `
<p>
  <h1>({{.CheckSuite.HeadBranch}} - {{.CheckSuite.HeadSHA}})</h1>
  <h2> Foo</h2>
  <h3> Foo</h3>

  <b>{{ .CheckRun.Output.Title }}</b>
  <p style="color: #cb2431"><b>{{ .CheckRun.Conclusion }}</b> ran in {{.Duration}}</p>
  <b>{{ .CheckRun.Output.Title }}</b>

  <p>{{ .CheckRun.Output.Summary }}</p>

  <p>
    {{ .CheckRun.Output.Text }}
  </p>
</p>

<p style="font-size:small;-webkit-text-size-adjust:none;color:#666;">
  &mdash;
  <br />
  You are receiving this because you are the author of the commit.
  <br />
  <a href="{{ .CheckRun.HTMLURL }}">View it on GitHub</a>
</p>
`
