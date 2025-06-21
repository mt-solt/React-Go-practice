import { useState, useEffect } from 'react'
import type { RandomNumber } from '../types/random'
import { fetchRandomNumbers } from '../api/randomApi'

export const RandomNumberList = () => {
  const [randomNumbers, setRandomNumbers] = useState<RandomNumber[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const loadRandomNumbers = async () => {
    try {
      setLoading(true)
      setError(null)
      const data = await fetchRandomNumbers()
      setRandomNumbers(data)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'エラーが発生しました')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadRandomNumbers()
  }, [])

  return (
    <div className="random-numbers-section">
      <h2>乱数一覧</h2>
      <button onClick={loadRandomNumbers} disabled={loading}>
        {loading ? '読み込み中...' : '乱数を取得'}
      </button>
      
      {error && <p className="error">{error}</p>}
      
      {randomNumbers.length > 0 ? (
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>値</th>
              <th>作成日時</th>
            </tr>
          </thead>
          <tbody>
            {randomNumbers.map((random) => (
              <tr key={random.id}>
                <td>{random.id}</td>
                <td>{random.value}</td>
                <td>{new Date(random.created_at).toLocaleString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p>乱数データがありません</p>
      )}
    </div>
  )
} 