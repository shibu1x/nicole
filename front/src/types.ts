export interface Video {
    Title: string
    WatchURL: string
    ThumbnailURL: string
    OwnerId: string
    IsPaymentRequired: boolean
    IsMuted: boolean
}

export type Rankings = Video[][]