export interface RandomNumber {
  id: number
  value: number
  created_at: string
}

export interface CreateRandomRequest {
  value: number
}

export interface UpdateRandomRequest {
  id: number
  value: number
}

export interface DeleteRandomRequest {
  id: number
} 