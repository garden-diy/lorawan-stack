// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	fmt "fmt"
	go_thethings_network_lorawan_stack_v3_pkg_types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

func (dst *ClaimEndDeviceRequest) SetFields(src *ClaimEndDeviceRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "target_application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.TargetApplicationIds == nil) && dst.TargetApplicationIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.TargetApplicationIds
				}
				if dst.TargetApplicationIds != nil {
					newDst = dst.TargetApplicationIds
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.TargetApplicationIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.TargetApplicationIds = src.TargetApplicationIds
				} else {
					dst.TargetApplicationIds = nil
				}
			}
		case "target_device_id":
			if len(subs) > 0 {
				return fmt.Errorf("'target_device_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetDeviceId = src.TargetDeviceId
			} else {
				var zero string
				dst.TargetDeviceId = zero
			}
		case "target_network_server_address":
			if len(subs) > 0 {
				return fmt.Errorf("'target_network_server_address' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetNetworkServerAddress = src.TargetNetworkServerAddress
			} else {
				var zero string
				dst.TargetNetworkServerAddress = zero
			}
		case "target_network_server_kek_label":
			if len(subs) > 0 {
				return fmt.Errorf("'target_network_server_kek_label' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetNetworkServerKekLabel = src.TargetNetworkServerKekLabel
			} else {
				var zero string
				dst.TargetNetworkServerKekLabel = zero
			}
		case "target_application_server_address":
			if len(subs) > 0 {
				return fmt.Errorf("'target_application_server_address' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetApplicationServerAddress = src.TargetApplicationServerAddress
			} else {
				var zero string
				dst.TargetApplicationServerAddress = zero
			}
		case "target_application_server_kek_label":
			if len(subs) > 0 {
				return fmt.Errorf("'target_application_server_kek_label' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetApplicationServerKekLabel = src.TargetApplicationServerKekLabel
			} else {
				var zero string
				dst.TargetApplicationServerKekLabel = zero
			}
		case "target_application_server_id":
			if len(subs) > 0 {
				return fmt.Errorf("'target_application_server_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetApplicationServerId = src.TargetApplicationServerId
			} else {
				var zero string
				dst.TargetApplicationServerId = zero
			}
		case "target_net_id":
			if len(subs) > 0 {
				return fmt.Errorf("'target_net_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetNetId = src.TargetNetId
			} else {
				dst.TargetNetId = nil
			}
		case "invalidate_authentication_code":
			if len(subs) > 0 {
				return fmt.Errorf("'invalidate_authentication_code' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.InvalidateAuthenticationCode = src.InvalidateAuthenticationCode
			} else {
				var zero bool
				dst.InvalidateAuthenticationCode = zero
			}

		case "source_device":
			if len(subs) == 0 && src == nil {
				dst.SourceDevice = nil
				continue
			} else if len(subs) == 0 {
				dst.SourceDevice = src.SourceDevice
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "authenticated_identifiers":
					_, srcOk := src.SourceDevice.(*ClaimEndDeviceRequest_AuthenticatedIdentifiers_)
					if !srcOk && src.SourceDevice != nil {
						return fmt.Errorf("attempt to set oneof 'authenticated_identifiers', while different oneof is set in source")
					}
					_, dstOk := dst.SourceDevice.(*ClaimEndDeviceRequest_AuthenticatedIdentifiers_)
					if !dstOk && dst.SourceDevice != nil {
						return fmt.Errorf("attempt to set oneof 'authenticated_identifiers', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ClaimEndDeviceRequest_AuthenticatedIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.SourceDevice.(*ClaimEndDeviceRequest_AuthenticatedIdentifiers_).AuthenticatedIdentifiers
						}
						if dstOk {
							newDst = dst.SourceDevice.(*ClaimEndDeviceRequest_AuthenticatedIdentifiers_).AuthenticatedIdentifiers
						} else {
							newDst = &ClaimEndDeviceRequest_AuthenticatedIdentifiers{}
							dst.SourceDevice = &ClaimEndDeviceRequest_AuthenticatedIdentifiers_{AuthenticatedIdentifiers: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.SourceDevice = src.SourceDevice
						} else {
							dst.SourceDevice = nil
						}
					}
				case "qr_code":
					_, srcOk := src.SourceDevice.(*ClaimEndDeviceRequest_QrCode)
					if !srcOk && src.SourceDevice != nil {
						return fmt.Errorf("attempt to set oneof 'qr_code', while different oneof is set in source")
					}
					_, dstOk := dst.SourceDevice.(*ClaimEndDeviceRequest_QrCode)
					if !dstOk && dst.SourceDevice != nil {
						return fmt.Errorf("attempt to set oneof 'qr_code', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'qr_code' has no subfields, but %s were specified", oneofSubs)
					}
					if src != nil {
						dst.SourceDevice = src.SourceDevice
					} else {
						dst.SourceDevice = nil
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *AuthorizeApplicationRequest) SetFields(src *AuthorizeApplicationRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if (src == nil || src.ApplicationIds == nil) && dst.ApplicationIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.ApplicationIds
				}
				if dst.ApplicationIds != nil {
					newDst = dst.ApplicationIds
				} else {
					newDst = &ApplicationIdentifiers{}
					dst.ApplicationIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIds = src.ApplicationIds
				} else {
					dst.ApplicationIds = nil
				}
			}
		case "api_key":
			if len(subs) > 0 {
				return fmt.Errorf("'api_key' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ApiKey = src.ApiKey
			} else {
				var zero string
				dst.ApiKey = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetInfoByJoinEUIRequest) SetFields(src *GetInfoByJoinEUIRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEui = src.JoinEui
			} else {
				dst.JoinEui = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetInfoByJoinEUIResponse) SetFields(src *GetInfoByJoinEUIResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEui = src.JoinEui
			} else {
				dst.JoinEui = nil
			}
		case "supports_claiming":
			if len(subs) > 0 {
				return fmt.Errorf("'supports_claiming' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SupportsClaiming = src.SupportsClaiming
			} else {
				var zero bool
				dst.SupportsClaiming = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetClaimStatusResponse) SetFields(src *GetClaimStatusResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "end_device_ids":
			if len(subs) > 0 {
				var newDst, newSrc *EndDeviceIdentifiers
				if (src == nil || src.EndDeviceIds == nil) && dst.EndDeviceIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.EndDeviceIds
				}
				if dst.EndDeviceIds != nil {
					newDst = dst.EndDeviceIds
				} else {
					newDst = &EndDeviceIdentifiers{}
					dst.EndDeviceIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.EndDeviceIds = src.EndDeviceIds
				} else {
					dst.EndDeviceIds = nil
				}
			}
		case "home_net_id":
			if len(subs) > 0 {
				return fmt.Errorf("'home_net_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.HomeNetId = src.HomeNetId
			} else {
				dst.HomeNetId = nil
			}
		case "home_ns_id":
			if len(subs) > 0 {
				return fmt.Errorf("'home_ns_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.HomeNsId = src.HomeNsId
			} else {
				dst.HomeNsId = nil
			}
		case "vendor_specific":
			if len(subs) > 0 {
				var newDst, newSrc *GetClaimStatusResponse_VendorSpecific
				if (src == nil || src.VendorSpecific == nil) && dst.VendorSpecific == nil {
					continue
				}
				if src != nil {
					newSrc = src.VendorSpecific
				}
				if dst.VendorSpecific != nil {
					newDst = dst.VendorSpecific
				} else {
					newDst = &GetClaimStatusResponse_VendorSpecific{}
					dst.VendorSpecific = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.VendorSpecific = src.VendorSpecific
				} else {
					dst.VendorSpecific = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CUPSRedirection) SetFields(src *CUPSRedirection, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "target_cups_uri":
			if len(subs) > 0 {
				return fmt.Errorf("'target_cups_uri' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetCupsUri = src.TargetCupsUri
			} else {
				var zero string
				dst.TargetCupsUri = zero
			}
		case "current_gateway_key":
			if len(subs) > 0 {
				return fmt.Errorf("'current_gateway_key' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CurrentGatewayKey = src.CurrentGatewayKey
			} else {
				var zero string
				dst.CurrentGatewayKey = zero
			}
		case "target_cups_trust":
			if len(subs) > 0 {
				return fmt.Errorf("'target_cups_trust' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetCupsTrust = src.TargetCupsTrust
			} else {
				dst.TargetCupsTrust = nil
			}

		case "gateway_credentials":
			if len(subs) == 0 && src == nil {
				dst.GatewayCredentials = nil
				continue
			} else if len(subs) == 0 {
				dst.GatewayCredentials = src.GatewayCredentials
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "client_tls":
					_, srcOk := src.GatewayCredentials.(*CUPSRedirection_ClientTls)
					if !srcOk && src.GatewayCredentials != nil {
						return fmt.Errorf("attempt to set oneof 'client_tls', while different oneof is set in source")
					}
					_, dstOk := dst.GatewayCredentials.(*CUPSRedirection_ClientTls)
					if !dstOk && dst.GatewayCredentials != nil {
						return fmt.Errorf("attempt to set oneof 'client_tls', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *CUPSRedirection_ClientTLS
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.GatewayCredentials.(*CUPSRedirection_ClientTls).ClientTls
						}
						if dstOk {
							newDst = dst.GatewayCredentials.(*CUPSRedirection_ClientTls).ClientTls
						} else {
							newDst = &CUPSRedirection_ClientTLS{}
							dst.GatewayCredentials = &CUPSRedirection_ClientTls{ClientTls: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.GatewayCredentials = src.GatewayCredentials
						} else {
							dst.GatewayCredentials = nil
						}
					}
				case "auth_token":
					_, srcOk := src.GatewayCredentials.(*CUPSRedirection_AuthToken)
					if !srcOk && src.GatewayCredentials != nil {
						return fmt.Errorf("attempt to set oneof 'auth_token', while different oneof is set in source")
					}
					_, dstOk := dst.GatewayCredentials.(*CUPSRedirection_AuthToken)
					if !dstOk && dst.GatewayCredentials != nil {
						return fmt.Errorf("attempt to set oneof 'auth_token', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'auth_token' has no subfields, but %s were specified", oneofSubs)
					}
					if src != nil {
						dst.GatewayCredentials = src.GatewayCredentials
					} else {
						dst.GatewayCredentials = &CUPSRedirection_AuthToken{}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ClaimGatewayRequest) SetFields(src *ClaimGatewayRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "collaborator":
			if len(subs) > 0 {
				var newDst, newSrc *OrganizationOrUserIdentifiers
				if (src == nil || src.Collaborator == nil) && dst.Collaborator == nil {
					continue
				}
				if src != nil {
					newSrc = src.Collaborator
				}
				if dst.Collaborator != nil {
					newDst = dst.Collaborator
				} else {
					newDst = &OrganizationOrUserIdentifiers{}
					dst.Collaborator = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Collaborator = src.Collaborator
				} else {
					dst.Collaborator = nil
				}
			}
		case "target_gateway_id":
			if len(subs) > 0 {
				return fmt.Errorf("'target_gateway_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetGatewayId = src.TargetGatewayId
			} else {
				var zero string
				dst.TargetGatewayId = zero
			}
		case "target_gateway_server_address":
			if len(subs) > 0 {
				return fmt.Errorf("'target_gateway_server_address' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetGatewayServerAddress = src.TargetGatewayServerAddress
			} else {
				var zero string
				dst.TargetGatewayServerAddress = zero
			}
		case "cups_redirection":
			if len(subs) > 0 {
				var newDst, newSrc *CUPSRedirection
				if (src == nil || src.CupsRedirection == nil) && dst.CupsRedirection == nil {
					continue
				}
				if src != nil {
					newSrc = src.CupsRedirection
				}
				if dst.CupsRedirection != nil {
					newDst = dst.CupsRedirection
				} else {
					newDst = &CUPSRedirection{}
					dst.CupsRedirection = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.CupsRedirection = src.CupsRedirection
				} else {
					dst.CupsRedirection = nil
				}
			}
		case "target_frequency_plan_id":
			if len(subs) > 0 {
				return fmt.Errorf("'target_frequency_plan_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TargetFrequencyPlanId = src.TargetFrequencyPlanId
			} else {
				var zero string
				dst.TargetFrequencyPlanId = zero
			}

		case "source_gateway":
			if len(subs) == 0 && src == nil {
				dst.SourceGateway = nil
				continue
			} else if len(subs) == 0 {
				dst.SourceGateway = src.SourceGateway
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "authenticated_identifiers":
					_, srcOk := src.SourceGateway.(*ClaimGatewayRequest_AuthenticatedIdentifiers_)
					if !srcOk && src.SourceGateway != nil {
						return fmt.Errorf("attempt to set oneof 'authenticated_identifiers', while different oneof is set in source")
					}
					_, dstOk := dst.SourceGateway.(*ClaimGatewayRequest_AuthenticatedIdentifiers_)
					if !dstOk && dst.SourceGateway != nil {
						return fmt.Errorf("attempt to set oneof 'authenticated_identifiers', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ClaimGatewayRequest_AuthenticatedIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.SourceGateway.(*ClaimGatewayRequest_AuthenticatedIdentifiers_).AuthenticatedIdentifiers
						}
						if dstOk {
							newDst = dst.SourceGateway.(*ClaimGatewayRequest_AuthenticatedIdentifiers_).AuthenticatedIdentifiers
						} else {
							newDst = &ClaimGatewayRequest_AuthenticatedIdentifiers{}
							dst.SourceGateway = &ClaimGatewayRequest_AuthenticatedIdentifiers_{AuthenticatedIdentifiers: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.SourceGateway = src.SourceGateway
						} else {
							dst.SourceGateway = nil
						}
					}
				case "qr_code":
					_, srcOk := src.SourceGateway.(*ClaimGatewayRequest_QrCode)
					if !srcOk && src.SourceGateway != nil {
						return fmt.Errorf("attempt to set oneof 'qr_code', while different oneof is set in source")
					}
					_, dstOk := dst.SourceGateway.(*ClaimGatewayRequest_QrCode)
					if !dstOk && dst.SourceGateway != nil {
						return fmt.Errorf("attempt to set oneof 'qr_code', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						return fmt.Errorf("'qr_code' has no subfields, but %s were specified", oneofSubs)
					}
					if src != nil {
						dst.SourceGateway = src.SourceGateway
					} else {
						dst.SourceGateway = nil
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *AuthorizeGatewayRequest) SetFields(src *AuthorizeGatewayRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "gateway_ids":
			if len(subs) > 0 {
				var newDst, newSrc *GatewayIdentifiers
				if (src == nil || src.GatewayIds == nil) && dst.GatewayIds == nil {
					continue
				}
				if src != nil {
					newSrc = src.GatewayIds
				}
				if dst.GatewayIds != nil {
					newDst = dst.GatewayIds
				} else {
					newDst = &GatewayIdentifiers{}
					dst.GatewayIds = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.GatewayIds = src.GatewayIds
				} else {
					dst.GatewayIds = nil
				}
			}
		case "api_key":
			if len(subs) > 0 {
				return fmt.Errorf("'api_key' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ApiKey = src.ApiKey
			} else {
				var zero string
				dst.ApiKey = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ClaimEndDeviceRequest_AuthenticatedIdentifiers) SetFields(src *ClaimEndDeviceRequest_AuthenticatedIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEui = src.JoinEui
			} else {
				var zero go_thethings_network_lorawan_stack_v3_pkg_types.EUI64
				dst.JoinEui = zero
			}
		case "dev_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevEui = src.DevEui
			} else {
				var zero go_thethings_network_lorawan_stack_v3_pkg_types.EUI64
				dst.DevEui = zero
			}
		case "authentication_code":
			if len(subs) > 0 {
				return fmt.Errorf("'authentication_code' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.AuthenticationCode = src.AuthenticationCode
			} else {
				var zero string
				dst.AuthenticationCode = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GetClaimStatusResponse_VendorSpecific) SetFields(src *GetClaimStatusResponse_VendorSpecific, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "organization_unique_identifier":
			if len(subs) > 0 {
				return fmt.Errorf("'organization_unique_identifier' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.OrganizationUniqueIdentifier = src.OrganizationUniqueIdentifier
			} else {
				var zero uint32
				dst.OrganizationUniqueIdentifier = zero
			}
		case "data":
			if len(subs) > 0 {
				return fmt.Errorf("'data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Data = src.Data
			} else {
				dst.Data = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *CUPSRedirection_ClientTLS) SetFields(src *CUPSRedirection_ClientTLS, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "cert":
			if len(subs) > 0 {
				return fmt.Errorf("'cert' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Cert = src.Cert
			} else {
				dst.Cert = nil
			}
		case "key":
			if len(subs) > 0 {
				return fmt.Errorf("'key' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Key = src.Key
			} else {
				dst.Key = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ClaimGatewayRequest_AuthenticatedIdentifiers) SetFields(src *ClaimGatewayRequest_AuthenticatedIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "gateway_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'gateway_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.GatewayEui = src.GatewayEui
			} else {
				var zero go_thethings_network_lorawan_stack_v3_pkg_types.EUI64
				dst.GatewayEui = zero
			}
		case "authentication_code":
			if len(subs) > 0 {
				return fmt.Errorf("'authentication_code' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.AuthenticationCode = src.AuthenticationCode
			} else {
				dst.AuthenticationCode = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
