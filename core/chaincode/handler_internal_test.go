/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package chaincode

import (
	"github.com/Matrix-Zhang/fabric-gm/core/common/sysccprovider"
	"github.com/Matrix-Zhang/fabric-gm/core/container/ccintf"
	pb "github.com/Matrix-Zhang/fabric-gm/protos/peer"
)

// Helpers to access unexported state.

func SetHandlerChaincodeID(h *Handler, chaincodeID *pb.ChaincodeID) {
	h.chaincodeID = chaincodeID
}

func SetHandlerChatStream(h *Handler, chatStream ccintf.ChaincodeStream) {
	h.chatStream = chatStream
}

func SetHandlerCCInstance(h *Handler, ccInstance *sysccprovider.ChaincodeInstance) {
	h.ccInstance = ccInstance
}

func StreamDone(h *Handler) <-chan struct{} {
	return h.streamDone()
}

func SetStreamDoneChan(h *Handler, ch chan struct{}) {
	h.streamDoneChan = ch
}
