package coordinator

// coordinatorUserFacing defines the methods that the Coordinator needs to implement for the user
// facing REST API
type coordinatorUserFacing interface {
	startJob()
}

type RESTserver struct {
	c coordinatorUserFacing
}

func (server RESTserver) serveHTTP() {
	server.c.startJob()
}
