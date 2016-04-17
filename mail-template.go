package main

var MailTemplate = `
Hi Jorin!
Here's your weekly server status summary.

Kind regards,
Your computer



Most accessed resources:

{{ range .Top 20 }}
{{ .One }} - {{ .Two }}{{ end }}

Total: {{ .Total }}
`
