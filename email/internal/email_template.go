package internal

const DefaultSubjectTemplate = `
Elixir CI: [{{.Repo.FullName}}] Build failed ({{.CheckSuite.HeadBranch}} - {{.CheckSuite.HeadSHA}})
`

const DefauleEmailMarkdownTemplate = `
<p>
  <b>{{ .CheckRun.Output.Title }}</b>

  <p>{{ .CheckRun.Output.Summary }}</p>

  <p>
    <code>
      {{ .CheckRun.Output.Text }}
    </code>
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
