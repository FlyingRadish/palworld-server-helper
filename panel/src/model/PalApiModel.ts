export interface RebootParam {
    notifyReboot: boolean
}

export interface BroadcastParam {
    content: string
}

export interface PalPlayer {
    name: string
    uid: string
    steamId: string
}

export interface MemStatus {
    total: number
    used: number
}

export interface RconParam {
    command: string
}

export interface SimpleParam {
    data: string
}