import { json } from '@remix-run/node'
import React, { useEffect } from 'react'
import { XMsg } from '~/graphql/graphql'
import { ColumnDef } from '@tanstack/react-table'
import SimpleTable from '../shared/simpleTable'
import { useLoaderData } from '@remix-run/react'
import { dateFormatter, hashShortener } from '~/lib/formatting'
import Tag from '../shared/tag'
import RollupIcon from '../shared/rollupIcon'
import { Link } from '@remix-run/react'
import LongArrow from '~/assets/images/LongArrow.svg'
import { loader } from '~/routes/_index'


export default function XMsgDataTable() {
  const data = useLoaderData<typeof loader>()

  const rows = data.xmsgs


  const columnConfig = {
    canFilter: false,
    enableColumnFilter: false,
  }

  const columns = React.useMemo<ColumnDef<any>[]>(
    () => [
      {
        ...columnConfig,
        accessorKey: 'offset',
        header: () => <span>Nounce</span>,
        cell: (value: any) => (
          <Link to="/" className="link font-bold text-b-sm">
            {value.getValue()}
          </Link>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'timeStamp',
        header: () => <span>Age</span>,
        cell: (value: any) => (
          <span className="text-subtlest font-bold text-b-xs">
            {' '}
            {dateFormatter(value.getValue())}
          </span>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'status',
        header: () => <span>Status</span>,
        cell: (value: any) => <Tag status={value.getValue()} />,
      },
      {
        ...columnConfig,
        accessorKey: 'srcLogoUrl',
        header: () => <span></span>,
        cell: (value: any) => <RollupIcon name="arbiscan" />,
      },
      {
        ...columnConfig,
        accessorKey: 'fromAddress',
        header: () => <span>Address</span>,
        cell: (value: any) => (
          <Link to="/" className="link">
            <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
            ico
          </Link>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'blockHash',
        header: () => <span>Block Hash</span>,
        cell: (value: any) => (
          <Link to="/" className="link">
            <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
            ico
          </Link>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'Empty',
        header: () => <span></span>,
        cell: (value: any) => <img src={LongArrow} alt="" />,
      },
      {
        ...columnConfig,
        accessorKey: 'destLogoUrl',
        header: () => <span></span>,
        cell: (value: any) => <RollupIcon name="arbiscan" />,
      },
      {
        ...columnConfig,
        accessorKey: 'destAddress',
        header: () => <span>Address</span>,
        cell: (value: any) => (
          <Link to="/" className="link">
            <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
            ico
          </Link>
        ),
      },
      {
        ...columnConfig,
        accessorKey: 'txHash',
        header: () => <span>Tx Hash</span>,
        cell: (value: any) => (
          <Link to="/" className="link">
            <span className="font-bold text-b-sm">{hashShortener(value.getValue())}</span>
            ico
          </Link>
        ),
      },
    ],
    [],
  )

  return (
    <div className="flex-none">
      <div className="flex flex-col">
        <h5 className="text-default mb-4">XMsgs</h5>
        <div className={''}>filter options</div>
      </div>
      <div>
        <SimpleTable columns={columns} data={rows} />
        count: {data.count}
      </div>
    </div>
  )
}
