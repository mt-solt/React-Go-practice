import './App.css'
import { RandomNumberList } from './components/RandomNumberList'

function App() {
  return (
    <div className="app">
      <header className="app-header">
        <h1>乱数管理システム</h1>
      </header>
      <main className="app-main">
        <RandomNumberList />
      </main>
    </div>
  )
}

export default App
