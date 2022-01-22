
// getSetup returns the echo context and a response recorder
// for testing purposes.
func getSetup(t *testing.T,
	req *http.Request, pathParams map[string]interface{},
) (echo.Context, *httptest.ResponseRecorder) {
	t.Helper()

	// create recorder and context
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	// set parameters
	names := []string{}
	values := []string{}

	for name, value := range req.URL.Query() {
		names = append(names, name)
		values = append(values, strings.Join(value, ","))
	}

	for name, value := range pathParams {
		names = append(names, name)
		values = append(values, fmt.Sprintf("%v", value))
	}

	c.SetParamNames(names...)
	c.SetParamValues(values...)

	return c, rec
}
