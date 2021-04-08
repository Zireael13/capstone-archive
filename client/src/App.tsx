import React from 'react'
import './App.css'
import { Route } from 'react-router-dom'
import HomePage from './components/pages/HomePage'
import MoreviewPage from './components/pages/MoreviewPage'
import Form from './components/pages/Form'
import SearchPage from './components/pages/SearchPage'
import LoginPage from './components/pages/LoginPage'
import { QueryClientProvider, QueryClient } from 'react-query'
import { ReactQueryDevtools } from 'react-query/devtools'

function App(): JSX.Element {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <div>
        <Route path="/" exact component={HomePage} />
        <Route path="/view/:id" exact component={MoreviewPage} />
        <Route path="/add" exact component={Form} />
        <Route path="/search" exact component={SearchPage} />
        <Route path="/login" exact component={LoginPage} />
      </div>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  )
}

export default App
