package internal

const DefaultSubjectTemplate = `
Build failed for {{.Repo.FullName}}#{{.CheckSuite.HeadBranch}}
`

const DefauleEmailMarkdownTemplate = `
Build results available in {{ .CheckRun.HTMLURL }}. It took {{ .Duration }}
`
