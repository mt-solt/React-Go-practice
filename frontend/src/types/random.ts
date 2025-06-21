export interface RandomNumber {
  uuid: string
  value: number
  user_id: string
}

export interface CreateRandomRequest {
  userId: string
}

export interface UpdateRandomRequest {
  id: string
}

export interface DeleteRandomRequest {
  id: string
} 