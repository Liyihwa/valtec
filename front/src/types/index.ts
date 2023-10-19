export interface Map {
  id: number
  name: string
  url: string
  avatar: string
}

export interface Hero {
  id: number
  name: string
  avatar: string
  skills: Skill[]
}

export interface Skill {
  type: string
  img: string
}
export interface Position {
  id: number
  sx: string
  sy: string
  px: string
  py: string
}
