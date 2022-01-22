
func fakeServerInterfaceWrapper() *httpservice.ServerInterfaceWrapper {
	return &httpservice.ServerInterfaceWrapper{
		Handler: httpservice.HTTPHandler{},
	}
}
