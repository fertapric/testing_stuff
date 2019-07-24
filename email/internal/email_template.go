package internal

const DefaultSubjectTemplate = `
Elixir CI: Build failed for {{.Repo.FullName}} ({{.CheckSuite.HeadBranch}})
`

const DefauleEmailMarkdownTemplate = `
<p>
Check run [{{ .CheckRun.Name }}]({{ .CheckRun.HTMLURL }}) {{ .CheckRun.Status }} in {{ .Duration }}
</p>

<p style="font-size:small;-webkit-text-size-adjust:none;color:#666;">
  &mdash;
  <br />
  You are receiving this because you are the author of the commit.
  <br />
  <a href="{{ .CheckRun.HTMLURL }}">View it on GitHub</a>
</p>
`
