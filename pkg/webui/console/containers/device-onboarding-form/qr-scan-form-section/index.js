// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React, { useCallback, useState } from 'react'
import { useSelector } from 'react-redux'

import QRModalButton from '@ttn-lw/components/qr-modal-button'
import { useFormContext } from '@ttn-lw/components/form'
import Link from '@ttn-lw/components/link'
import Icon from '@ttn-lw/components/icon'
import ModalButton from '@ttn-lw/components/button/modal-button'
import ButtonGroup from '@ttn-lw/components/button/group'

import Message from '@ttn-lw/lib/components/message'

import sharedMessages from '@ttn-lw/lib/shared-messages'

import { readDeviceQr } from '@console/lib/device-claiming-parse-qr'

import { selectDeviceBrands } from '@console/store/selectors/device-repository'

import m from '../messages'

const hexToDecimal = hex => parseInt(hex, 16)

const qrDataInitialState = {
  valid: false,
  approved: false,
  data: undefined,
  device: undefined,
}

const DeviceQRScanFormSection = () => {
  const { resetForm, setValues } = useFormContext()
  const brands = useSelector(selectDeviceBrands)
  const [qrData, setQrData] = useState(qrDataInitialState)

  const handleReset = useCallback(() => {
    resetForm()
    setQrData(qrDataInitialState)
  }, [resetForm])

  const getBrand = useCallback(
    vendorId => {
      const brand = brands.find(brand => brand?.lora_alliance_vendor_id === hexToDecimal(vendorId))

      return brand
    },
    [brands],
  )

  const handleQRCodeApprove = useCallback(() => {
    const { device } = qrData
    const brand = getBrand(device.profileID.vendorID)

    setValues(values => ({
      ...values,
      _withQRdata: true,
      ids: {
        ...values.ids,
        join_eui: device.joinEUI,
        dev_eui: device.devEUI,
        device_id: `eui-${device.devEUI.toLowerCase()}`,
      },
      target_device_id: `eui-${device.devEUI.toLowerCase()}`,
      authenticated_identifiers: {
        dev_eui: device.devEUI,
        authentication_code: device.ownerToken ? device.ownerToken : '',
        join_eui: device.joinEUI,
      },
      version_ids: {
        ...values.version_ids,
        brand_id: brand ? brand.brand_id : values.version_ids.brand_id,
      },
    }))

    setQrData({ ...qrData, approved: true })
  }, [getBrand, qrData, setValues])

  const handleQRCodeCancel = useCallback(() => {
    setQrData(qrDataInitialState)
  }, [])

  const handleQRCodeRead = useCallback(
    qrCode => {
      const device = readDeviceQr(qrCode)
      if (device && device.devEUI) {
        const brand = getBrand(device.profileID.vendorID)
        const sheetData = [
          {
            header: m.deviceInfo,
            items: [
              {
                key: sharedMessages.claimAuthCode,
                value: device.ownerToken,
                type: 'code',
                sensitive: true,
              },
              {
                key: sharedMessages.joinEUI,
                value: device.joinEUI,
                type: 'byte',
                sensitive: false,
              },
              { key: sharedMessages.devEUI, value: device.devEUI, type: 'byte', sensitive: false },
              { key: sharedMessages.brand, value: brand?.name },
            ],
          },
        ]
        setQrData({
          ...qrData,
          valid: true,
          data: sheetData,
          device,
        })
      } else {
        setQrData({ ...qrData, data: [], valid: false })
      }
    },
    [getBrand, qrData],
  )

  return (
    <>
      {qrData.approved ? (
        <div className="mb-cs-xs">
          <Icon icon="check" textPaddedRight className="c-success" />
          <Message content={m.scanSuccess} />
        </div>
      ) : (
        <div className="mb-cs-xs">
          <Message content={m.hasEndDeviceQR} />
        </div>
      )}
      <ButtonGroup>
        {qrData.approved ? (
          <ModalButton
            type="button"
            icon="close"
            onApprove={handleReset}
            message={m.resetQRCodeData}
            modalData={{
              title: m.resetQRCodeData,
              noTitleLine: true,
              buttonMessage: m.resetQRCodeData,
              children: <Message content={m.resetConfirm} component="span" />,
              approveButtonProps: {
                icon: 'close',
              },
            }}
          />
        ) : (
          <QRModalButton
            message={m.scanEndDevice}
            onApprove={handleQRCodeApprove}
            onCancel={handleQRCodeCancel}
            onRead={handleQRCodeRead}
            qrData={qrData}
          />
        )}
        <Link.DocLink className="ml-cs-xs" path="/devices/adding-devices" secondary>
          <Message content={m.deviceGuide} />
        </Link.DocLink>
      </ButtonGroup>
      <hr className="mt-cs-m mb-0" />
    </>
  )
}

export default DeviceQRScanFormSection
