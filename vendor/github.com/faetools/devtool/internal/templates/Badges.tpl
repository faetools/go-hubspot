{{ if .IsMicroservice }}
{{ $token := "YOUR_STATUS_API_TOKEN" }}

{{ with .CicleCIToken }}
	{{ $token = . }}
{{ else }}
TODO: Get [YOUR_STATUS_API_TOKEN here](https://app.circleci.com/settings/project/github/faetools/{{ .Name }}/api) and insert it below.
{{ end }}

[![CircleCI](https://circleci.com/gh/faetools/{{ .Name }}/tree/master.svg?style=shield&circle-token={{ $token }})](https://app.circleci.com/pipelines/github/faetools/{{ .Name }}?branch=master) {{ end -}}

{{- with .QualityGate -}}
[![Sonar Quality Gate](https://img.shields.io/badge/SonarCloud-{{ . }}-{{ if eq . "OK" }}brightgreen{{ else }}red{{ end }})](https://sonarcloud.io/project/overview?id=faetools_{{ $.Name }}) {{ end -}}

{{- if .Coverage -}}
![Go Coverage](https://img.shields.io/badge/Go%20Coverage-{{ printf "%.f" .Coverage }}%25-{{ .CoverageColor }}.svg?longCache=true&style=flat) {{ end }}
![Devtool version](https://img.shields.io/badge/Devtool-{{ .DevtoolVersion }}-brightgreen.svg)
![Maintainer](https://img.shields.io/badge/team-{{ .Team }}-blue)
{{ if .IsPublic -}}
[![Go Report Card](https://goreportcard.com/badge/github/faetools/{{ .Name }})](https://goreportcard.com/report/github/faetools/{{ .Name }})
[![GoDoc Reference](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/faetools/{{ .Name }})
{{- end -}}
