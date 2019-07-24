package internal

const DefaultSubjectTemplate = `
{{.CheckSuite.App.Name}} check for {{.Repo.FullName}}#{{.CheckSuite.HeadBranch}}
`

const DefauleEmailMarkdownTemplate = `
Check run [{{ .CheckRun.Name }}]({{ .CheckRun.HTMLURL }}) {{ .CheckRun.Status }} in {{ .Duration }}
`
