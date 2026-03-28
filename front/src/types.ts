export interface Video {
    Title: string
    WatchURL: string
    ThumbnailURL: string
    OwnerId: string
    IsPaymentRequired: boolean
    IsMuted: boolean
    IsMutedByOwner: boolean
    IsMutedByTitle: boolean
}

export type Rankings = Video[][]