package unreal

import (
	"sync"

	"channeld.clewcat.com/channeld/pkg/channeld"
	"channeld.clewcat.com/channeld/pkg/common"
	"channeld.clewcat.com/channeld/pkg/unrealpb"
	"go.uber.org/zap"
)

// Channel data messages need to implement this interface to be able to collect the states for handover
type ChannelDataCollector interface {
	common.Message
	CollectStates(netId uint32, src common.Message) error
}

// Stores stubs for providing the handover data. The stub will be removed when the source server answers the handover context.
var handoverDataProviders map[uint64]chan common.Message = make(map[uint64]chan common.Message)
var handoverDataProvidersLock sync.RWMutex

func getHandoverStub(netId uint32, srcChannelId uint32) uint64 {
	return uint64(srcChannelId)<<32 | uint64(netId)
}

func getHandoverDataProvider(netId uint32, srcChannelId uint32) (chan common.Message, bool) {
	handoverDataProvidersLock.RLock()
	defer handoverDataProvidersLock.RUnlock()
	provider, exists := handoverDataProviders[getHandoverStub(netId, srcChannelId)]
	return provider, exists
}

func addHandoverDataProvider(netId uint32, srcChannelId uint32, provider chan common.Message) {
	handoverDataProvidersLock.Lock()
	defer handoverDataProvidersLock.Unlock()
	handoverDataProviders[getHandoverStub(netId, srcChannelId)] = provider
}

func removeHandoverDataProvider(netId uint32, srcChannelId uint32) {
	handoverDataProvidersLock.Lock()
	defer handoverDataProvidersLock.Unlock()
	delete(handoverDataProviders, getHandoverStub(netId, srcChannelId))
}

func resetHandoverDataProviders() {
	handoverDataProvidersLock.Lock()
	defer handoverDataProvidersLock.Unlock()
	handoverDataProviders = make(map[uint64]chan common.Message)
}

func CheckSpatialInfoChange(netId uint32, newLoc, oldLoc *unrealpb.FVector, spatialNotifier common.SpatialInfoChangedNotifier) {
	var newX, newY, newZ float32

	if newLoc.X != nil {
		newX = *newLoc.X
	} else {
		// Use GetX/Y() to avoid violation memory access!
		newX = oldLoc.GetX()
	}

	if newLoc.Y != nil {
		newY = *newLoc.Y
	} else {
		newY = oldLoc.GetY()
	}

	if newLoc.Z != nil {
		newZ = *newLoc.Z
	} else {
		newZ = oldLoc.GetZ()
	}

	if newX != oldLoc.GetX() || newY != oldLoc.GetY() || newZ != oldLoc.GetZ() {
		// Swap the Y and Z as UE uses the Z-Up rule but channeld uses the Y-up rule.
		oldInfo := common.SpatialInfo{
			X: float64(oldLoc.GetX()),
			Y: float64(oldLoc.GetZ()),
			Z: float64(oldLoc.GetY()),
		}
		newInfo := common.SpatialInfo{
			X: float64(newX),
			Y: float64(newZ),
			Z: float64(newY),
		}

		spatialNotifier.Notify(
			oldInfo,
			newInfo,
			func(srcChannelId common.ChannelId, dstChannelId common.ChannelId, handoverData chan common.Message) {
				/* No matter if it's cross-server, always ask for the handover context.
				// Handover happens within the spatial server - no need to ask for the handover context.
				if channeld.GetChannel(srcChannelId).IsSameOwner(channeld.GetChannel(dstChannelId)) {
					defer allSpawnedObjLock.RLocker().Unlock()
					allSpawnedObjLock.RLock()
					handoverData <- &unrealpb.HandoverData{
						Context: []*unrealpb.HandoverContext{
							{
								Obj:          allSpawnedObj[netId],
								ClientConnId: oldActorState.OwningConnId,
							},
						},
					}
				} else*/{
					addHandoverDataProvider(netId, uint32(srcChannelId), handoverData)
					channeld.GetChannel(srcChannelId).SendToOwner(uint32(unrealpb.MessageType_HANDOVER_CONTEXT), &unrealpb.GetHandoverContextMessage{
						NetId:        netId,
						SrcChannelId: uint32(srcChannelId),
						DstChannelId: uint32(dstChannelId),
					})
					channeld.RootLogger().Info("getting handover context from src server",
						zap.Uint32("srcChannelId", uint32(srcChannelId)),
						zap.Float32("oldX", oldLoc.GetX()), zap.Float32("oldY", oldLoc.GetY()),
						zap.Float32("newX", newX), zap.Float32("newY", newY),
					)
				}
			},
		)
	}
}