// Copyright 2013 The Go Circuit Project
// Use of this source code is governed by the license for
// The Go Circuit Project, found in the LICENSE file.
//
// Authors:
//   2013 Petar Maymounkov <p@gocircuit.org>

package locus

import (
	"encoding/gob"

	"github.com/gocircuit/circuit/kit/kinfolk"
	"github.com/gocircuit/circuit/kit/lang"
	"github.com/gocircuit/circuit/use/circuit"
)

// Peer encloses a cross-interface to the kinfolk system of a circuit worker, as well as
// a cross-interface to its exported resource hierarchy.
type Peer struct {
	Kin kinfolk.KinXID // Cross-interface to the kin system at this locus
	Term circuit.X  // Cross-interface to the anchor namespace of this peer, see “kit/anchor”
}

func (i Peer) Key() string {
	return i.Kin.ID.String()
}

func (i Peer) ID() lang.ReceiverID {
	return i.Kin.ID
}

func init() {
	gob.Register(&Peer{})
}
