/*
*rpcobjects
 */
package rpcobjects

// import "net"

// Args .
type Args struct {
	N, M int
}

//Multiply .
func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}
