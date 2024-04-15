import type { MetaFunction } from '@remix-run/node'
import XBlockDataTable from '~/components/home/blockDataTable'
import XMsgDataTable from '~/components/home/messageDataTable'
import Overview from '~/components/home/overview'

export const meta: MetaFunction = () => {
  return [
    { title: 'Omni Network Explorer' },
    { name: 'description', content: 'Omni Network Explorer' },
  ]
}

export default function Index() {
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
