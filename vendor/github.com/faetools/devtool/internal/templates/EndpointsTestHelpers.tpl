import "github.com/stretchr/testify/assert"

{{ range . }}
{{$hasParams := .RequiresParamObject -}}
{{$pathParams := .PathParams -}}
{{$opid := .OperationId -}}
{{$funcName := printf "%s%s" (lower (slice .OperationId 0 1)) (slice .OperationId 1) }}

// ({{ .Method }} {{ .Path }})
// {{ $funcName }} helps test {{  .OperationId }}.
func {{ $funcName }}(t *testing.T,
	w *httpservice.ServerInterfaceWrapper,
	{{- range .PathParams }}
	{{ .ParamName }} {{ if ne (lower .Schema.GoType) .Schema.GoType }}clients.{{ end }}{{ .Schema.GoType }},
	{{- end }}
	{{ if $hasParams -}}
	params *clients.{{  .OperationId }}Params,
	{{- end }}
	{{ range $i, $b := .Bodies -}}
   body{{ if ne $i 0 }}{{ $i }}{{ end }} clients.{{$opid}}{{.NameTag}}RequestBody,
	{{- end -}}
) *clients.{{  .OperationId }}Response {
	t.Helper()

	// create the request
	req, err := clients.New{{  .OperationId }}Request(
		os.Getenv("TEST_SERVER"),
	{{- range .PathParams -}}
	{{ .ParamName }},
	{{- end -}}
	{{- if $hasParams -}}
	params,
	{{- end -}}
	{{- range $i, $b := .Bodies -}}
   body{{ if ne $i 0 }}{{ $i }}{{ end }},
	{{- end -}}
	)
	assert.NoError(t, err)

	// get echo.Context and recorder
	c, rec := getSetup(t, req, map[string]interface{}{
		{{ range .PathParams -}}
		 "{{ .ParamName }}": {{ .ParamName }},
		 {{ end -}}
		})

	// make the request
	assert.NoError(t, w.{{  .OperationId }}(c))

	// parse the response
	res := rec.Result()
	resp, err := clients.Parse{{  .OperationId }}Response(res)
	assert.NoError(t, err)
	assert.NoError(t, res.Body.Close())

	return resp
}

{{ end }}
