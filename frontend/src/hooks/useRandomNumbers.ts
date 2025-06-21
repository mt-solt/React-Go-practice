import { useState, useEffect } from 'react'
import type { RandomNumber } from '../types/random'
import { fetchRandomNumbers, createRandom, updateRandom, deleteRandom } from '../api/randomApi'

export const useRandomNumbers = () => {
  const [randomNumbers, setRandomNumbers] = useState<RandomNumber[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  // 乱数一覧取得
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

  // 乱数登録
  const handleCreateRandom = async () => {
    try {
      setLoading(true)
      setError(null)
      await createRandom({ userId: 'test_user' })
      // 登録後に一覧を再取得
      await loadRandomNumbers()
    } catch (err) {
      setError(err instanceof Error ? err.message : '乱数登録に失敗しました')
    } finally {
      setLoading(false)
    }
  }

  // 乱数更新
  const handleUpdateRandom = async (id: string) => {
    try {
      setLoading(true)
      setError(null)
      await updateRandom({ id })
      // 更新後に一覧を再取得
      await loadRandomNumbers()
    } catch (err) {
      setError(err instanceof Error ? err.message : '乱数更新に失敗しました')
    } finally {
      setLoading(false)
    }
  }

  // 乱数削除
  const handleDeleteRandom = async (id: string) => {
    try {
      setLoading(true)
      setError(null)
      await deleteRandom({ id })
      // 削除後に一覧を再取得
      await loadRandomNumbers()
    } catch (err) {
      setError(err instanceof Error ? err.message : '乱数削除に失敗しました')
    } finally {
      setLoading(false)
    }
  }

  // 初回マウント時に一覧を取得
  useEffect(() => {
    loadRandomNumbers()
  }, [])

  // 乱数一覧取得、乱数登録、乱数更新、乱数削除のフックを返却
  return {
    randomNumbers,
    loading,
    error,
    loadRandomNumbers,
    handleCreateRandom,
    handleUpdateRandom,
    handleDeleteRandom
  }
} 