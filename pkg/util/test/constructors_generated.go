// Code generated by generate_constructors.go. DO NOT EDIT.

package test

import (
	pbtypes "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type (
	// RootKeysOption transforms ttnpb.RootKeys and returns it.
	// Implemetations must be pure functions with no side-effects.
	RootKeysOption func(*ttnpb.RootKeys) *ttnpb.RootKeys

	// RootKeysOptionNamespace represents the namespace, on which various RootKeysOption are defined.
	RootKeysOptionNamespace struct{}
)

// WithRootKeyId returns a RootKeysOption, which returns a copy of ttnpb.RootKeys with RootKeyId set to v.
func (RootKeysOptionNamespace) WithRootKeyId(v string) RootKeysOption {
	return func(x *ttnpb.RootKeys) *ttnpb.RootKeys {
		copy := ttnpb.Clone(x)
		copy.RootKeyId = v
		return copy
	}
}

// WithAppKey returns a RootKeysOption, which returns a copy of ttnpb.RootKeys with AppKey set to v.
func (RootKeysOptionNamespace) WithAppKey(v *ttnpb.KeyEnvelope) RootKeysOption {
	return func(x *ttnpb.RootKeys) *ttnpb.RootKeys {
		copy := ttnpb.Clone(x)
		copy.AppKey = v
		return copy
	}
}

// WithNwkKey returns a RootKeysOption, which returns a copy of ttnpb.RootKeys with NwkKey set to v.
func (RootKeysOptionNamespace) WithNwkKey(v *ttnpb.KeyEnvelope) RootKeysOption {
	return func(x *ttnpb.RootKeys) *ttnpb.RootKeys {
		copy := ttnpb.Clone(x)
		copy.NwkKey = v
		return copy
	}
}

// Compose returns a functional composition of opts as a singular RootKeysOption.
func (RootKeysOptionNamespace) Compose(opts ...RootKeysOption) RootKeysOption {
	return func(x *ttnpb.RootKeys) *ttnpb.RootKeys {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular {*ttnpb.RootKeys{ $optionType }}.
func (f RootKeysOption) Compose(opts ...RootKeysOption) RootKeysOption {
	return func(x *ttnpb.RootKeys) *ttnpb.RootKeys {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// RootKeysOptions is namespace containing ttnpb.RootKeys options.
var RootKeysOptions RootKeysOptionNamespace

// MakeRootKeys constructs a new ttnpb.RootKeys.
func MakeRootKeys(opts ...RootKeysOption) *ttnpb.RootKeys {
	return RootKeysOptions.Compose(opts...)(baseRootKeys)
}

type (
	// SessionKeysOption transforms ttnpb.SessionKeys and returns it.
	// Implemetations must be pure functions with no side-effects.
	SessionKeysOption func(*ttnpb.SessionKeys) *ttnpb.SessionKeys

	// SessionKeysOptionNamespace represents the namespace, on which various SessionKeysOption are defined.
	SessionKeysOptionNamespace struct{}
)

// WithSessionKeyId returns a SessionKeysOption, which returns a copy of ttnpb.SessionKeys with SessionKeyId set to v.
func (SessionKeysOptionNamespace) WithSessionKeyId(v []byte) SessionKeysOption {
	return func(x *ttnpb.SessionKeys) *ttnpb.SessionKeys {
		copy := ttnpb.Clone(x)
		copy.SessionKeyId = v
		return copy
	}
}

// WithFNwkSIntKey returns a SessionKeysOption, which returns a copy of ttnpb.SessionKeys with FNwkSIntKey set to v.
func (SessionKeysOptionNamespace) WithFNwkSIntKey(v *ttnpb.KeyEnvelope) SessionKeysOption {
	return func(x *ttnpb.SessionKeys) *ttnpb.SessionKeys {
		copy := ttnpb.Clone(x)
		copy.FNwkSIntKey = v
		return copy
	}
}

// WithSNwkSIntKey returns a SessionKeysOption, which returns a copy of ttnpb.SessionKeys with SNwkSIntKey set to v.
func (SessionKeysOptionNamespace) WithSNwkSIntKey(v *ttnpb.KeyEnvelope) SessionKeysOption {
	return func(x *ttnpb.SessionKeys) *ttnpb.SessionKeys {
		copy := ttnpb.Clone(x)
		copy.SNwkSIntKey = v
		return copy
	}
}

// WithNwkSEncKey returns a SessionKeysOption, which returns a copy of ttnpb.SessionKeys with NwkSEncKey set to v.
func (SessionKeysOptionNamespace) WithNwkSEncKey(v *ttnpb.KeyEnvelope) SessionKeysOption {
	return func(x *ttnpb.SessionKeys) *ttnpb.SessionKeys {
		copy := ttnpb.Clone(x)
		copy.NwkSEncKey = v
		return copy
	}
}

// WithAppSKey returns a SessionKeysOption, which returns a copy of ttnpb.SessionKeys with AppSKey set to v.
func (SessionKeysOptionNamespace) WithAppSKey(v *ttnpb.KeyEnvelope) SessionKeysOption {
	return func(x *ttnpb.SessionKeys) *ttnpb.SessionKeys {
		copy := ttnpb.Clone(x)
		copy.AppSKey = v
		return copy
	}
}

// Compose returns a functional composition of opts as a singular SessionKeysOption.
func (SessionKeysOptionNamespace) Compose(opts ...SessionKeysOption) SessionKeysOption {
	return func(x *ttnpb.SessionKeys) *ttnpb.SessionKeys {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular {*ttnpb.SessionKeys{ $optionType }}.
func (f SessionKeysOption) Compose(opts ...SessionKeysOption) SessionKeysOption {
	return func(x *ttnpb.SessionKeys) *ttnpb.SessionKeys {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// SessionKeysOptions is namespace containing ttnpb.SessionKeys options.
var SessionKeysOptions SessionKeysOptionNamespace

// MakeSessionKeys constructs a new ttnpb.SessionKeys.
func MakeSessionKeys(opts ...SessionKeysOption) *ttnpb.SessionKeys {
	return SessionKeysOptions.Compose(opts...)(baseSessionKeys)
}

type (
	// SessionOption transforms ttnpb.Session and returns it.
	// Implemetations must be pure functions with no side-effects.
	SessionOption func(*ttnpb.Session) *ttnpb.Session

	// SessionOptionNamespace represents the namespace, on which various SessionOption are defined.
	SessionOptionNamespace struct{}
)

// WithDevAddr returns a SessionOption, which returns a copy of ttnpb.Session with DevAddr set to v.
func (SessionOptionNamespace) WithDevAddr(v []byte) SessionOption {
	return func(x *ttnpb.Session) *ttnpb.Session {
		copy := ttnpb.Clone(x)
		copy.DevAddr = v
		return copy
	}
}

// WithKeys returns a SessionOption, which returns a copy of ttnpb.Session with Keys set to v.
func (SessionOptionNamespace) WithKeys(v *ttnpb.SessionKeys) SessionOption {
	return func(x *ttnpb.Session) *ttnpb.Session {
		copy := ttnpb.Clone(x)
		copy.Keys = v
		return copy
	}
}

// WithLastFCntUp returns a SessionOption, which returns a copy of ttnpb.Session with LastFCntUp set to v.
func (SessionOptionNamespace) WithLastFCntUp(v uint32) SessionOption {
	return func(x *ttnpb.Session) *ttnpb.Session {
		copy := ttnpb.Clone(x)
		copy.LastFCntUp = v
		return copy
	}
}

// WithLastNFCntDown returns a SessionOption, which returns a copy of ttnpb.Session with LastNFCntDown set to v.
func (SessionOptionNamespace) WithLastNFCntDown(v uint32) SessionOption {
	return func(x *ttnpb.Session) *ttnpb.Session {
		copy := ttnpb.Clone(x)
		copy.LastNFCntDown = v
		return copy
	}
}

// WithLastAFCntDown returns a SessionOption, which returns a copy of ttnpb.Session with LastAFCntDown set to v.
func (SessionOptionNamespace) WithLastAFCntDown(v uint32) SessionOption {
	return func(x *ttnpb.Session) *ttnpb.Session {
		copy := ttnpb.Clone(x)
		copy.LastAFCntDown = v
		return copy
	}
}

// WithLastConfFCntDown returns a SessionOption, which returns a copy of ttnpb.Session with LastConfFCntDown set to v.
func (SessionOptionNamespace) WithLastConfFCntDown(v uint32) SessionOption {
	return func(x *ttnpb.Session) *ttnpb.Session {
		copy := ttnpb.Clone(x)
		copy.LastConfFCntDown = v
		return copy
	}
}

// WithStartedAt returns a SessionOption, which returns a copy of ttnpb.Session with StartedAt set to v.
func (SessionOptionNamespace) WithStartedAt(v *pbtypes.Timestamp) SessionOption {
	return func(x *ttnpb.Session) *ttnpb.Session {
		copy := ttnpb.Clone(x)
		copy.StartedAt = v
		return copy
	}
}

// WithQueuedApplicationDownlinks returns a SessionOption, which returns a copy of ttnpb.Session with QueuedApplicationDownlinks set to vs.
func (SessionOptionNamespace) WithQueuedApplicationDownlinks(vs ...*ttnpb.ApplicationDownlink) SessionOption {
	return func(x *ttnpb.Session) *ttnpb.Session {
		copy := ttnpb.Clone(x)
		copy.QueuedApplicationDownlinks = vs
		return copy
	}
}

// Compose returns a functional composition of opts as a singular SessionOption.
func (SessionOptionNamespace) Compose(opts ...SessionOption) SessionOption {
	return func(x *ttnpb.Session) *ttnpb.Session {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular {*ttnpb.Session{ $optionType }}.
func (f SessionOption) Compose(opts ...SessionOption) SessionOption {
	return func(x *ttnpb.Session) *ttnpb.Session {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// SessionOptions is namespace containing ttnpb.Session options.
var SessionOptions SessionOptionNamespace

// MakeSession constructs a new ttnpb.Session.
func MakeSession(opts ...SessionOption) *ttnpb.Session {
	return SessionOptions.Compose(opts...)(baseSession)
}

type (
	// MACStateOption transforms ttnpb.MACState and returns it.
	// Implemetations must be pure functions with no side-effects.
	MACStateOption func(*ttnpb.MACState) *ttnpb.MACState

	// MACStateOptionNamespace represents the namespace, on which various MACStateOption are defined.
	MACStateOptionNamespace struct{}
)

// WithCurrentParameters returns a MACStateOption, which returns a copy of ttnpb.MACState with CurrentParameters set to v.
func (MACStateOptionNamespace) WithCurrentParameters(v *ttnpb.MACParameters) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.CurrentParameters = v
		return copy
	}
}

// WithDesiredParameters returns a MACStateOption, which returns a copy of ttnpb.MACState with DesiredParameters set to v.
func (MACStateOptionNamespace) WithDesiredParameters(v *ttnpb.MACParameters) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.DesiredParameters = v
		return copy
	}
}

// WithDeviceClass returns a MACStateOption, which returns a copy of ttnpb.MACState with DeviceClass set to v.
func (MACStateOptionNamespace) WithDeviceClass(v ttnpb.Class) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.DeviceClass = v
		return copy
	}
}

// WithLorawanVersion returns a MACStateOption, which returns a copy of ttnpb.MACState with LorawanVersion set to v.
func (MACStateOptionNamespace) WithLorawanVersion(v ttnpb.MACVersion) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.LorawanVersion = v
		return copy
	}
}

// WithLastConfirmedDownlinkAt returns a MACStateOption, which returns a copy of ttnpb.MACState with LastConfirmedDownlinkAt set to v.
func (MACStateOptionNamespace) WithLastConfirmedDownlinkAt(v *pbtypes.Timestamp) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.LastConfirmedDownlinkAt = v
		return copy
	}
}

// WithLastDevStatusFCntUp returns a MACStateOption, which returns a copy of ttnpb.MACState with LastDevStatusFCntUp set to v.
func (MACStateOptionNamespace) WithLastDevStatusFCntUp(v uint32) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.LastDevStatusFCntUp = v
		return copy
	}
}

// WithPingSlotPeriodicity returns a MACStateOption, which returns a copy of ttnpb.MACState with PingSlotPeriodicity set to v.
func (MACStateOptionNamespace) WithPingSlotPeriodicity(v *ttnpb.PingSlotPeriodValue) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.PingSlotPeriodicity = v
		return copy
	}
}

// WithPendingApplicationDownlink returns a MACStateOption, which returns a copy of ttnpb.MACState with PendingApplicationDownlink set to v.
func (MACStateOptionNamespace) WithPendingApplicationDownlink(v *ttnpb.ApplicationDownlink) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.PendingApplicationDownlink = v
		return copy
	}
}

// WithQueuedResponses returns a MACStateOption, which returns a copy of ttnpb.MACState with QueuedResponses set to vs.
func (MACStateOptionNamespace) WithQueuedResponses(vs ...*ttnpb.MACCommand) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.QueuedResponses = vs
		return copy
	}
}

// WithPendingRequests returns a MACStateOption, which returns a copy of ttnpb.MACState with PendingRequests set to vs.
func (MACStateOptionNamespace) WithPendingRequests(vs ...*ttnpb.MACCommand) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.PendingRequests = vs
		return copy
	}
}

// WithQueuedJoinAccept returns a MACStateOption, which returns a copy of ttnpb.MACState with QueuedJoinAccept set to v.
func (MACStateOptionNamespace) WithQueuedJoinAccept(v *ttnpb.MACState_JoinAccept) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.QueuedJoinAccept = v
		return copy
	}
}

// WithPendingJoinRequest returns a MACStateOption, which returns a copy of ttnpb.MACState with PendingJoinRequest set to v.
func (MACStateOptionNamespace) WithPendingJoinRequest(v *ttnpb.MACState_JoinRequest) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.PendingJoinRequest = v
		return copy
	}
}

// WithRxWindowsAvailable returns a MACStateOption, which returns a copy of ttnpb.MACState with RxWindowsAvailable set to v.
func (MACStateOptionNamespace) WithRxWindowsAvailable(v bool) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.RxWindowsAvailable = v
		return copy
	}
}

// WithRecentUplinks returns a MACStateOption, which returns a copy of ttnpb.MACState with RecentUplinks set to vs.
func (MACStateOptionNamespace) WithRecentUplinks(vs ...*ttnpb.MACState_UplinkMessage) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.RecentUplinks = vs
		return copy
	}
}

// WithRecentDownlinks returns a MACStateOption, which returns a copy of ttnpb.MACState with RecentDownlinks set to vs.
func (MACStateOptionNamespace) WithRecentDownlinks(vs ...*ttnpb.MACState_DownlinkMessage) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.RecentDownlinks = vs
		return copy
	}
}

// WithLastNetworkInitiatedDownlinkAt returns a MACStateOption, which returns a copy of ttnpb.MACState with LastNetworkInitiatedDownlinkAt set to v.
func (MACStateOptionNamespace) WithLastNetworkInitiatedDownlinkAt(v *pbtypes.Timestamp) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.LastNetworkInitiatedDownlinkAt = v
		return copy
	}
}

// WithRejectedAdrDataRateIndexes returns a MACStateOption, which returns a copy of ttnpb.MACState with RejectedAdrDataRateIndexes set to vs.
func (MACStateOptionNamespace) WithRejectedAdrDataRateIndexes(vs ...ttnpb.DataRateIndex) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.RejectedAdrDataRateIndexes = vs
		return copy
	}
}

// WithRejectedAdrTxPowerIndexes returns a MACStateOption, which returns a copy of ttnpb.MACState with RejectedAdrTxPowerIndexes set to vs.
func (MACStateOptionNamespace) WithRejectedAdrTxPowerIndexes(vs ...uint32) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.RejectedAdrTxPowerIndexes = vs
		return copy
	}
}

// WithRejectedFrequencies returns a MACStateOption, which returns a copy of ttnpb.MACState with RejectedFrequencies set to vs.
func (MACStateOptionNamespace) WithRejectedFrequencies(vs ...uint64) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.RejectedFrequencies = vs
		return copy
	}
}

// WithLastDownlinkAt returns a MACStateOption, which returns a copy of ttnpb.MACState with LastDownlinkAt set to v.
func (MACStateOptionNamespace) WithLastDownlinkAt(v *pbtypes.Timestamp) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.LastDownlinkAt = v
		return copy
	}
}

// WithRejectedDataRateRanges returns a MACStateOption, which returns a copy of ttnpb.MACState with RejectedDataRateRanges set to v.
func (MACStateOptionNamespace) WithRejectedDataRateRanges(v map[uint64]*ttnpb.MACState_DataRateRanges) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.RejectedDataRateRanges = v
		return copy
	}
}

// WithLastAdrChangeFCntUp returns a MACStateOption, which returns a copy of ttnpb.MACState with LastAdrChangeFCntUp set to v.
func (MACStateOptionNamespace) WithLastAdrChangeFCntUp(v uint32) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.LastAdrChangeFCntUp = v
		return copy
	}
}

// WithRecentMacCommandIdentifiers returns a MACStateOption, which returns a copy of ttnpb.MACState with RecentMacCommandIdentifiers set to vs.
func (MACStateOptionNamespace) WithRecentMacCommandIdentifiers(vs ...ttnpb.MACCommandIdentifier) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		copy := ttnpb.Clone(x)
		copy.RecentMacCommandIdentifiers = vs
		return copy
	}
}

// Compose returns a functional composition of opts as a singular MACStateOption.
func (MACStateOptionNamespace) Compose(opts ...MACStateOption) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular {*ttnpb.MACState{ $optionType }}.
func (f MACStateOption) Compose(opts ...MACStateOption) MACStateOption {
	return func(x *ttnpb.MACState) *ttnpb.MACState {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// MACStateOptions is namespace containing ttnpb.MACState options.
var MACStateOptions MACStateOptionNamespace

// MakeMACState constructs a new ttnpb.MACState.
func MakeMACState(opts ...MACStateOption) *ttnpb.MACState {
	return MACStateOptions.Compose(opts...)(baseMACState)
}

type (
	// EndDeviceIdentifiersOption transforms ttnpb.EndDeviceIdentifiers and returns it.
	// Implemetations must be pure functions with no side-effects.
	EndDeviceIdentifiersOption func(*ttnpb.EndDeviceIdentifiers) *ttnpb.EndDeviceIdentifiers

	// EndDeviceIdentifiersOptionNamespace represents the namespace, on which various EndDeviceIdentifiersOption are defined.
	EndDeviceIdentifiersOptionNamespace struct{}
)

// WithDeviceId returns a EndDeviceIdentifiersOption, which returns a copy of ttnpb.EndDeviceIdentifiers with DeviceId set to v.
func (EndDeviceIdentifiersOptionNamespace) WithDeviceId(v string) EndDeviceIdentifiersOption {
	return func(x *ttnpb.EndDeviceIdentifiers) *ttnpb.EndDeviceIdentifiers {
		copy := ttnpb.Clone(x)
		copy.DeviceId = v
		return copy
	}
}

// WithApplicationIds returns a EndDeviceIdentifiersOption, which returns a copy of ttnpb.EndDeviceIdentifiers with ApplicationIds set to v.
func (EndDeviceIdentifiersOptionNamespace) WithApplicationIds(v *ttnpb.ApplicationIdentifiers) EndDeviceIdentifiersOption {
	return func(x *ttnpb.EndDeviceIdentifiers) *ttnpb.EndDeviceIdentifiers {
		copy := ttnpb.Clone(x)
		copy.ApplicationIds = v
		return copy
	}
}

// WithDevEui returns a EndDeviceIdentifiersOption, which returns a copy of ttnpb.EndDeviceIdentifiers with DevEui set to v.
func (EndDeviceIdentifiersOptionNamespace) WithDevEui(v []byte) EndDeviceIdentifiersOption {
	return func(x *ttnpb.EndDeviceIdentifiers) *ttnpb.EndDeviceIdentifiers {
		copy := ttnpb.Clone(x)
		copy.DevEui = v
		return copy
	}
}

// WithJoinEui returns a EndDeviceIdentifiersOption, which returns a copy of ttnpb.EndDeviceIdentifiers with JoinEui set to v.
func (EndDeviceIdentifiersOptionNamespace) WithJoinEui(v []byte) EndDeviceIdentifiersOption {
	return func(x *ttnpb.EndDeviceIdentifiers) *ttnpb.EndDeviceIdentifiers {
		copy := ttnpb.Clone(x)
		copy.JoinEui = v
		return copy
	}
}

// WithDevAddr returns a EndDeviceIdentifiersOption, which returns a copy of ttnpb.EndDeviceIdentifiers with DevAddr set to v.
func (EndDeviceIdentifiersOptionNamespace) WithDevAddr(v []byte) EndDeviceIdentifiersOption {
	return func(x *ttnpb.EndDeviceIdentifiers) *ttnpb.EndDeviceIdentifiers {
		copy := ttnpb.Clone(x)
		copy.DevAddr = v
		return copy
	}
}

// Compose returns a functional composition of opts as a singular EndDeviceIdentifiersOption.
func (EndDeviceIdentifiersOptionNamespace) Compose(opts ...EndDeviceIdentifiersOption) EndDeviceIdentifiersOption {
	return func(x *ttnpb.EndDeviceIdentifiers) *ttnpb.EndDeviceIdentifiers {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular {*ttnpb.EndDeviceIdentifiers{ $optionType }}.
func (f EndDeviceIdentifiersOption) Compose(opts ...EndDeviceIdentifiersOption) EndDeviceIdentifiersOption {
	return func(x *ttnpb.EndDeviceIdentifiers) *ttnpb.EndDeviceIdentifiers {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// EndDeviceIdentifiersOptions is namespace containing ttnpb.EndDeviceIdentifiers options.
var EndDeviceIdentifiersOptions EndDeviceIdentifiersOptionNamespace

// MakeEndDeviceIdentifiers constructs a new ttnpb.EndDeviceIdentifiers.
func MakeEndDeviceIdentifiers(opts ...EndDeviceIdentifiersOption) *ttnpb.EndDeviceIdentifiers {
	return EndDeviceIdentifiersOptions.Compose(opts...)(baseEndDeviceIdentifiers)
}

type (
	// EndDeviceOption transforms ttnpb.EndDevice and returns it.
	// Implemetations must be pure functions with no side-effects.
	EndDeviceOption func(*ttnpb.EndDevice) *ttnpb.EndDevice

	// EndDeviceOptionNamespace represents the namespace, on which various EndDeviceOption are defined.
	EndDeviceOptionNamespace struct{}
)

// WithIds returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Ids set to v.
func (EndDeviceOptionNamespace) WithIds(v *ttnpb.EndDeviceIdentifiers) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.Ids = v
		return copy
	}
}

// WithCreatedAt returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with CreatedAt set to v.
func (EndDeviceOptionNamespace) WithCreatedAt(v *pbtypes.Timestamp) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.CreatedAt = v
		return copy
	}
}

// WithUpdatedAt returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with UpdatedAt set to v.
func (EndDeviceOptionNamespace) WithUpdatedAt(v *pbtypes.Timestamp) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.UpdatedAt = v
		return copy
	}
}

// WithName returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Name set to v.
func (EndDeviceOptionNamespace) WithName(v string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.Name = v
		return copy
	}
}

// WithDescription returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Description set to v.
func (EndDeviceOptionNamespace) WithDescription(v string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.Description = v
		return copy
	}
}

// WithAttributes returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Attributes set to v.
func (EndDeviceOptionNamespace) WithAttributes(v map[string]string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.Attributes = v
		return copy
	}
}

// WithVersionIds returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with VersionIds set to v.
func (EndDeviceOptionNamespace) WithVersionIds(v *ttnpb.EndDeviceVersionIdentifiers) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.VersionIds = v
		return copy
	}
}

// WithServiceProfileId returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ServiceProfileId set to v.
func (EndDeviceOptionNamespace) WithServiceProfileId(v string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.ServiceProfileId = v
		return copy
	}
}

// WithNetworkServerAddress returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with NetworkServerAddress set to v.
func (EndDeviceOptionNamespace) WithNetworkServerAddress(v string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.NetworkServerAddress = v
		return copy
	}
}

// WithNetworkServerKekLabel returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with NetworkServerKekLabel set to v.
func (EndDeviceOptionNamespace) WithNetworkServerKekLabel(v string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.NetworkServerKekLabel = v
		return copy
	}
}

// WithApplicationServerAddress returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ApplicationServerAddress set to v.
func (EndDeviceOptionNamespace) WithApplicationServerAddress(v string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.ApplicationServerAddress = v
		return copy
	}
}

// WithApplicationServerKekLabel returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ApplicationServerKekLabel set to v.
func (EndDeviceOptionNamespace) WithApplicationServerKekLabel(v string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.ApplicationServerKekLabel = v
		return copy
	}
}

// WithApplicationServerId returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ApplicationServerId set to v.
func (EndDeviceOptionNamespace) WithApplicationServerId(v string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.ApplicationServerId = v
		return copy
	}
}

// WithJoinServerAddress returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with JoinServerAddress set to v.
func (EndDeviceOptionNamespace) WithJoinServerAddress(v string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.JoinServerAddress = v
		return copy
	}
}

// WithLocations returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Locations set to v.
func (EndDeviceOptionNamespace) WithLocations(v map[string]*ttnpb.Location) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.Locations = v
		return copy
	}
}

// WithPicture returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Picture set to v.
func (EndDeviceOptionNamespace) WithPicture(v *ttnpb.Picture) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.Picture = v
		return copy
	}
}

// WithSupportsClassB returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with SupportsClassB set to v.
func (EndDeviceOptionNamespace) WithSupportsClassB(v bool) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.SupportsClassB = v
		return copy
	}
}

// WithSupportsClassC returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with SupportsClassC set to v.
func (EndDeviceOptionNamespace) WithSupportsClassC(v bool) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.SupportsClassC = v
		return copy
	}
}

// WithLorawanVersion returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LorawanVersion set to v.
func (EndDeviceOptionNamespace) WithLorawanVersion(v ttnpb.MACVersion) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.LorawanVersion = v
		return copy
	}
}

// WithLorawanPhyVersion returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LorawanPhyVersion set to v.
func (EndDeviceOptionNamespace) WithLorawanPhyVersion(v ttnpb.PHYVersion) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.LorawanPhyVersion = v
		return copy
	}
}

// WithFrequencyPlanId returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with FrequencyPlanId set to v.
func (EndDeviceOptionNamespace) WithFrequencyPlanId(v string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.FrequencyPlanId = v
		return copy
	}
}

// WithMinFrequency returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with MinFrequency set to v.
func (EndDeviceOptionNamespace) WithMinFrequency(v uint64) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.MinFrequency = v
		return copy
	}
}

// WithMaxFrequency returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with MaxFrequency set to v.
func (EndDeviceOptionNamespace) WithMaxFrequency(v uint64) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.MaxFrequency = v
		return copy
	}
}

// WithSupportsJoin returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with SupportsJoin set to v.
func (EndDeviceOptionNamespace) WithSupportsJoin(v bool) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.SupportsJoin = v
		return copy
	}
}

// WithResetsJoinNonces returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ResetsJoinNonces set to v.
func (EndDeviceOptionNamespace) WithResetsJoinNonces(v bool) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.ResetsJoinNonces = v
		return copy
	}
}

// WithRootKeys returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with RootKeys set to v.
func (EndDeviceOptionNamespace) WithRootKeys(v *ttnpb.RootKeys) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.RootKeys = v
		return copy
	}
}

// WithNetId returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with NetId set to v.
func (EndDeviceOptionNamespace) WithNetId(v []byte) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.NetId = v
		return copy
	}
}

// WithMacSettings returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with MacSettings set to v.
func (EndDeviceOptionNamespace) WithMacSettings(v *ttnpb.MACSettings) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.MacSettings = v
		return copy
	}
}

// WithMacState returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with MacState set to v.
func (EndDeviceOptionNamespace) WithMacState(v *ttnpb.MACState) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.MacState = v
		return copy
	}
}

// WithPendingMacState returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with PendingMacState set to v.
func (EndDeviceOptionNamespace) WithPendingMacState(v *ttnpb.MACState) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.PendingMacState = v
		return copy
	}
}

// WithSession returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Session set to v.
func (EndDeviceOptionNamespace) WithSession(v *ttnpb.Session) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.Session = v
		return copy
	}
}

// WithPendingSession returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with PendingSession set to v.
func (EndDeviceOptionNamespace) WithPendingSession(v *ttnpb.Session) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.PendingSession = v
		return copy
	}
}

// WithLastDevNonce returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastDevNonce set to v.
func (EndDeviceOptionNamespace) WithLastDevNonce(v uint32) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.LastDevNonce = v
		return copy
	}
}

// WithUsedDevNonces returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with UsedDevNonces set to vs.
func (EndDeviceOptionNamespace) WithUsedDevNonces(vs ...uint32) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.UsedDevNonces = vs
		return copy
	}
}

// WithLastJoinNonce returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastJoinNonce set to v.
func (EndDeviceOptionNamespace) WithLastJoinNonce(v uint32) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.LastJoinNonce = v
		return copy
	}
}

// WithLastRjCount_0 returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastRjCount_0 set to v.
func (EndDeviceOptionNamespace) WithLastRjCount_0(v uint32) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.LastRjCount_0 = v
		return copy
	}
}

// WithLastRjCount_1 returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastRjCount_1 set to v.
func (EndDeviceOptionNamespace) WithLastRjCount_1(v uint32) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.LastRjCount_1 = v
		return copy
	}
}

// WithLastDevStatusReceivedAt returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastDevStatusReceivedAt set to v.
func (EndDeviceOptionNamespace) WithLastDevStatusReceivedAt(v *pbtypes.Timestamp) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.LastDevStatusReceivedAt = v
		return copy
	}
}

// WithPowerState returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with PowerState set to v.
func (EndDeviceOptionNamespace) WithPowerState(v ttnpb.PowerState) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.PowerState = v
		return copy
	}
}

// WithBatteryPercentage returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with BatteryPercentage set to v.
func (EndDeviceOptionNamespace) WithBatteryPercentage(v *pbtypes.FloatValue) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.BatteryPercentage = v
		return copy
	}
}

// WithDownlinkMargin returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with DownlinkMargin set to v.
func (EndDeviceOptionNamespace) WithDownlinkMargin(v int32) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.DownlinkMargin = v
		return copy
	}
}

// WithQueuedApplicationDownlinks returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with QueuedApplicationDownlinks set to vs.
func (EndDeviceOptionNamespace) WithQueuedApplicationDownlinks(vs ...*ttnpb.ApplicationDownlink) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.QueuedApplicationDownlinks = vs
		return copy
	}
}

// WithFormatters returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Formatters set to v.
func (EndDeviceOptionNamespace) WithFormatters(v *ttnpb.MessagePayloadFormatters) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.Formatters = v
		return copy
	}
}

// WithProvisionerId returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ProvisionerId set to v.
func (EndDeviceOptionNamespace) WithProvisionerId(v string) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.ProvisionerId = v
		return copy
	}
}

// WithProvisioningData returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ProvisioningData set to v.
func (EndDeviceOptionNamespace) WithProvisioningData(v *pbtypes.Struct) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.ProvisioningData = v
		return copy
	}
}

// WithMulticast returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with Multicast set to v.
func (EndDeviceOptionNamespace) WithMulticast(v bool) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.Multicast = v
		return copy
	}
}

// WithClaimAuthenticationCode returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ClaimAuthenticationCode set to v.
func (EndDeviceOptionNamespace) WithClaimAuthenticationCode(v *ttnpb.EndDeviceAuthenticationCode) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.ClaimAuthenticationCode = v
		return copy
	}
}

// WithSkipPayloadCrypto returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with SkipPayloadCrypto set to v.
func (EndDeviceOptionNamespace) WithSkipPayloadCrypto(v bool) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.SkipPayloadCrypto = v
		return copy
	}
}

// WithSkipPayloadCryptoOverride returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with SkipPayloadCryptoOverride set to v.
func (EndDeviceOptionNamespace) WithSkipPayloadCryptoOverride(v *pbtypes.BoolValue) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.SkipPayloadCryptoOverride = v
		return copy
	}
}

// WithActivatedAt returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with ActivatedAt set to v.
func (EndDeviceOptionNamespace) WithActivatedAt(v *pbtypes.Timestamp) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.ActivatedAt = v
		return copy
	}
}

// WithLastSeenAt returns a EndDeviceOption, which returns a copy of ttnpb.EndDevice with LastSeenAt set to v.
func (EndDeviceOptionNamespace) WithLastSeenAt(v *pbtypes.Timestamp) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		copy := ttnpb.Clone(x)
		copy.LastSeenAt = v
		return copy
	}
}

// Compose returns a functional composition of opts as a singular EndDeviceOption.
func (EndDeviceOptionNamespace) Compose(opts ...EndDeviceOption) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// Compose returns a functional composition of f and opts as a singular {*ttnpb.EndDevice{ $optionType }}.
func (f EndDeviceOption) Compose(opts ...EndDeviceOption) EndDeviceOption {
	return func(x *ttnpb.EndDevice) *ttnpb.EndDevice {
		x = f(x)
		for _, opt := range opts {
			x = opt(x)
		}
		return x
	}
}

// EndDeviceOptions is namespace containing ttnpb.EndDevice options.
var EndDeviceOptions EndDeviceOptionNamespace

// MakeEndDevice constructs a new ttnpb.EndDevice.
func MakeEndDevice(opts ...EndDeviceOption) *ttnpb.EndDevice {
	return EndDeviceOptions.Compose(opts...)(baseEndDevice)
}
