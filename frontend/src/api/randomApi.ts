import type { RandomNumber, CreateRandomRequest, UpdateRandomRequest, DeleteRandomRequest } from '../types/random'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

export const fetchRandomNumbers = async (): Promise<RandomNumber[]> => {
  const response = await fetch(`${API_BASE_URL}/random`)
  if (!response.ok) {
    throw new Error('APIリクエストに失敗しました')
  }

  const resJson = await response.json()
  return resJson.data
}

// ユーザIDをキーに乱数を取得
export const fetchRandomByUser = async (userId: string): Promise<RandomNumber[]> => {
  const response = await fetch(`${API_BASE_URL}/random/user?userId=${userId}`)
  if (!response.ok) {
    throw new Error('APIリクエストに失敗しました')
  }

  return response.json()
}

// 乱数登録
export const createRandom = async (data: CreateRandomRequest): Promise<RandomNumber> => {
  const response = await fetch(`${API_BASE_URL}/create_random`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })
  if (!response.ok) {
    throw new Error('APIリクエストに失敗しました')
  }
  return response.json()
}

// 乱数更新
export const updateRandom = async (data: UpdateRandomRequest): Promise<RandomNumber> => {
  const response = await fetch(`${API_BASE_URL}/update_random`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })
  if (!response.ok) {
    throw new Error('APIリクエストに失敗しました')
  }
  return response.json()
}

// 乱数削除
export const deleteRandom = async (data: DeleteRandomRequest): Promise<void> => {
  const response = await fetch(`${API_BASE_URL}/delete_random`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })
  if (!response.ok) {
    throw new Error('APIリクエストに失敗しました')
  }
} 