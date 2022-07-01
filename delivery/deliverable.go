package delivery

type Deliverable interface {
	Execute() bool
}
