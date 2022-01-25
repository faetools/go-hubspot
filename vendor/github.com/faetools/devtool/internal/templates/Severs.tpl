package {{ .PackageName }}

{{ $len := (len .Servers) }}

{{ if eq $len 0 }}
// NOTE: No servers were defined in the openapi.yaml.
{{ else }}

	{{ range $i, $s := .Servers }}

		{{/* create sensible name for the server */}}
		{{ $name := printf "Server%s" (camelCase .Description) }}
		{{ if not .Description }}{{ $name = printf "Server%d" $i }}{{ end }}

		{{/* init default server */}}
		{{ if eq $i 0 }}
			// DefaultServer is the default server to be used.
			{{ if .Variables -}}
				var DefaultServer = {{ $name }}(
					{{- range $k, $v := .Variables }}"",{{- end }})
			{{ else if gt $len 1 -}}
				const DefaultServer = {{ $name }}
			{{ else -}}
				const DefaultServer = "{{ .URL }}"
			{{ end }}
		{{ end }}

		{{/* init other servers */}}
		{{ if .Variables }}
			// {{ $name }} is one server to chose from.
			{{- range $k, $v := .Variables }}{{ with .Description -}}
				// - {{ $k }}: {{ . }}
			{{ end -}}{{ end -}}
			func {{ $name }}(
			{{- range $k, $v := .Variables -}}
				{{ $k }} string,
			{{- end -}}
			) string {
				out := "{{ .URL }}"
			{{- range $k, $v := .Variables }}
				{{ with .Default -}}
					if {{ $k }} == "" {
						{{ $k }} = "{{ . }}"
					}
				{{ end -}}
				out = strings.ReplaceAll(out, "{{ printf "{%s}" $k }}", {{ $k }})
			{{- end }}
				return out
			}

		{{ else if gt $len 1 -}}
			// {{ $name }} is one server to chose from.
			const {{ $name }} = "{{ .URL }}"
		{{ end }}

	{{ end }}{{/* .Servers */}}

{{ end }}{{/* if eq $len 0 */}}
