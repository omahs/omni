import type { LinksFunction, SerializeFrom } from '@remix-run/node'
import stylesheet from '~/tailwind.css'
import {
  Links,
  LiveReload,
  Meta,
  Outlet,
  Scripts,
  ScrollRestoration,
  json,
  useLoaderData,
} from '@remix-run/react'
import { Client, Provider, cacheExchange, fetchExchange } from 'urql'
import { useEnv } from './lib/use-env'
import Navbar from './components/shared/navbar'
import { Footer } from './components/shared/footer'

export const links: LinksFunction = () => [{ rel: 'stylesheet', href: stylesheet }]
export type LoaderData = SerializeFrom<typeof loader>

export function loader() {
  const ENV = {
    GRAPHQL_URL: process.env.GRAPHQL_URL,
  }
  return json({ ENV })
}

function App() {
  return (
    <html lang="en" data-theme="dark">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body className="bg-surface ">
        <Navbar />
        <Outlet />
        <Footer />
        <ScrollRestoration />
        <Scripts />
        <LiveReload />
      </body>
    </html>
  )
}

export default function AppWithProviders() {
  useLoaderData<typeof loader>()

  const ENV = useEnv()
  let client = new Client({
    url: ENV.GRAPHQL_URL ?? '',
    exchanges: [fetchExchange, cacheExchange],
  })
  console.log(ENV.GRAPHQL_URL)
  return (
    <Provider value={client}>
      <App />
    </Provider>
  )
}
