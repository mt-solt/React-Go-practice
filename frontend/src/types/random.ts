export interface RandomNumber {
  id: number
  value: number
  created_at: string
}

export interface CreateRandomRequest {
  userId: string
}

export interface UpdateRandomRequest {
  id: number
  value: number
}

export interface DeleteRandomRequest {
  id: number
} 