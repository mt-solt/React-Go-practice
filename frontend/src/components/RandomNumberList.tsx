import { useRandomNumbers } from '../hooks/useRandomNumbers'

export const RandomNumberList = () => {
  const {
    randomNumbers,
    loading,
    error,
    loadRandomNumbers,
    handleCreateRandom
  } = useRandomNumbers()

  return (
    <div className="random-numbers-section">
      <h2>乱数一覧</h2>
      <div className="button-group">
        <button onClick={loadRandomNumbers} disabled={loading}>
          {loading ? '読み込み中...' : '乱数を取得'}
        </button>
        <button onClick={handleCreateRandom} disabled={loading}>
          {loading ? '登録中...' : '乱数を登録'}
        </button>
      </div>
      
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