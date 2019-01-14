// Copyright 2018 ETH Zurich
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sibra_mgmt

import (
	"fmt"

	"github.com/scionproto/scion/go/proto"
)

var _ proto.Cerealizable = (*SteadyRegRep)(nil)

type SteadyRegRep struct {
	Ack []uint16
}

func (s *SteadyRegRep) ProtoId() proto.ProtoIdType {
	return proto.SibraSteadyRegRep_TypeID
}

func (s *SteadyRegRep) String() string {
	return fmt.Sprintf("ack %v", s.Ack)
}