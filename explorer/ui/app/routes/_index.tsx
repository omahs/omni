import type { LoaderFunction, MetaFunction } from '@remix-run/node'
import XBlockDataTable from '~/components/home/blockDataTable'
import XMsgDataTable from '~/components/home/messageDataTable'
import Overview from '~/components/home/overview'
import { json } from '@remix-run/node'
import { gqlClient } from '~/entry.server'
import { useRevalidator } from '@remix-run/react'
import { useInterval } from '~/hooks/useInterval'
import { xblockcount } from '~/components/queries/block'

export const meta: MetaFunction = () => {
  return [
    { title: 'Omni Network Explorer' },
    { name: 'description', content: 'Omni Network Explorer' },
  ]
}

export const loader: LoaderFunction = async ({ request }) => {
  const res = await gqlClient.query(xblockcount, {})

  const pollData = async () => {
    return json({
      count: Number(res?.data?.xblockcount || '0x'),
      xmsgs: [
        {
          offset: 335,
          timeStamp: new Date(),
          status: 'Success',
          srcLogoUrl: 'https://picsum.photos/24/24',
          fromAddress: '0xbb0xbb23525f97',
          blockHash: '0xbb0xbb23525f97',
          destLogoUrl: 'https://picsum.photos/24/24',
          destAddress: '0xbb0xbb23525f97',
          txHash: '0xbb0xbb23525f97',
        },
        {
          offset: 335,
          timeStamp: new Date(),
          status: 'Failure',
          srcLogoUrl: 'https://picsum.photos/24/24',
          fromAddress: '0xbb0xbb23525f97',
          blockHash: '0xbb0xbb23525f97',
          destLogoUrl: 'https://picsum.photos/24/24',
          destAddress: '0xbb0xbb23525f97',
          txHash: '0xbb0xbb23525f97',
        },
        {
          offset: 335,
          timeStamp: new Date(),
          status: 'Success',
          srcLogoUrl: 'https://picsum.photos/24/24',
          fromAddress: '0xbb0xbb23525f97',
          blockHash: '0xbb0xbb23525f97',
          destLogoUrl: 'https://picsum.photos/24/24',
          destAddress: '0xbb0xbb23525f97',
          txHash: '0xbb0xbb23525f97',
        },
      ],
    })
  }

  return await pollData()
}

export default function Index() {
  const revalidator = useRevalidator()

  useInterval(() => {
    console.log('Revalidating')
    revalidator.revalidate()
  }, 5000)

  return (
    <div className="px-20">
      <div className="flex h-full w-full flex-col">
        <Overview />
        <div className="w-full">
          <XMsgDataTable />
        </div>
        <div className="grow"></div>
      </div>
    </div>
  )
}
