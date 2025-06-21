export interface RandomNumber {
  uuid: string
  value: number
  created_at: string
}

export interface CreateRandomRequest {
  userId: string
}

export interface UpdateRandomRequest {
  id: string
  value: number
}

export interface DeleteRandomRequest {
  id: string
} 