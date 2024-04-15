import React from 'react'

interface Props {
  status: 'Success' | 'Failure' | 'Pending'
}

const Tag: React.FC<Props> = ({ status }) => {
  const [color, setColor] = React.useState({})

  const getColor = () => {
    switch (status) {
      case 'Success':
        return {
          text: 'text-positive',
          icon: 'text-icon-positive',
          bg: 'bg-positive',
        }
    }
  }

  return <div className={`py-[3.5px] px-[5.5px] rounded-[4px] bg-positive text-positive`}>{status}</div>
}

export default Tag
