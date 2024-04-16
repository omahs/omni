import React from 'react'

interface Props {}

const Card: React.FC<{
  data: { title: string; value: number }
}> = ({ data }) => {
  const { title, value } = data
  return (
    <div className="w-15 h-[140px] min-w-[310px] rounded-lg bg-raised px-6 py-3">
      <span className="text-subtlest mb-[17px] block"> {title}</span>

      <h4 className="text-default">{value}</h4>
    </div>
  )
}

const Overview: React.FC<Props> = ({}) => {
  const cards = [
    { title: 'Total XMsgs sent', value: 79489200 },
    { title: 'Total XMsgs sent', value: 79489200 },
    { title: 'Total XMsgs sent', value: 79489200 },
    { title: 'Total XMsgs sent', value: 79489200 },
  ]
  return (
    <div className="mb-12 mt-8">
      <h5 className="text-default">Omni X-Explorer</h5>
      <div className="flex flex-row flex-wrap gap-3 mt-3">
        {cards.map(card => (
          <Card data={card} />
        ))}
      </div>
    </div>
  )
}

export default Overview
