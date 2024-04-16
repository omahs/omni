import React from 'react'

import Arbiscan from '~/assets/images/Arbiscan.svg'
import Caldera from '~/assets/images/Caldera.svg'
import EigenLayer from '~/assets/images/EigenLayer.svg'
import Espresso from '~/assets/images/Espresso.svg'
import Linea from '~/assets/images/Linea.svg'
import Optimism from '~/assets/images/Optimism.svg'
import Polygon from '~/assets/images/Polygon.svg'
import Scroll from '~/assets/images/Scroll.svg'

interface Props {
  chainId?: string
  name?: string
}

const RollupIcon: React.FC<Props> = ({ chainId, name }) => {
  return <img src={Arbiscan} alt="" />
}

export default RollupIcon
