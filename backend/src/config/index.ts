import development from './env/development.js'
import production from './env/production.js'
import merge from 'lodash.merge'

process.env.NODE_ENV = process.env.NODE_ENV || 'development'

const stage = process.env.STAGE || 'local'
let envConfig = {}

if (stage === 'production') {
  envConfig = production
} else {
  envConfig = development
}

const config = merge(
  {
    stage,
    env: process.env.NODE_ENV,
    origin: 'http://localhost:5173',
    port: 5000,
    secrets: {
    },
  },
  envConfig,
)

export default config
