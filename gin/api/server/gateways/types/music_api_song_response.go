package types

type MusicApiSongResponse struct {
    id int
    title string
    url string  
    description string
}

var Songs = []MusicApiSongResponse{
    {id: 1, title: "Song A", url: "https://bash.com/test.mp4", description: "this is song A"},
    {id: 2, title: "Song B", url: "https://bash.com/test.mp4", description: "this is song B"},
    {id: 3, title: "Song C", url: "https://bash.com/test.mp4", description: "this is song C"},
    {id: 4, title: "Song D", url: "https://bash.com/test.mp4", description: "this is song D"},
    {id: 5, title: "Song E", url: "https://bash.com/test.mp4", description: "this is song E"},
}

