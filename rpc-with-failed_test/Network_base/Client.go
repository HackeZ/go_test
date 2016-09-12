package Network_base

type ClientEnd struct {
	endname interface{} // this end-point's name
	ch      chan reqMsg // copy of Network.endCh
}
