import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)


{{ range . }}
{{$hasParams := .RequiresParamObject -}}
{{$opid := .OperationId -}}
{{$funcName := printf "%s%s" (lower (slice .OperationId 0 1)) (slice .OperationId 1) }}

// ({{ .Method }} {{ .Path }})
// Test{{ .OperationId }} tests {{  .OperationId }}.
func Test{{ .OperationId }}(t *testing.T) {
	t.Parallel()
	w := fakeServerInterfaceWrapper()

	resp := {{ $funcName }}(t, w,
	{{- range .PathParams -}}
		"some {{ .ParamName }}",
	{{- end -}}
	{{- if $hasParams -}}
	&clients.{{  .OperationId }}Params{},
	{{- end -}}
	{{- range $i, $b := .Bodies -}}
   clients.{{$opid}}{{.NameTag}}RequestBody{},
	{{- end -}}
	)
	assert.Equal(t, http.StatusNotImplemented, resp.StatusCode())
	{{ range .GetResponseTypeDefinitions -}}
	assert.Empty(t, resp.{{ .TypeDefinition.TypeName }})
	{{ end -}}

}

{{ end }}
