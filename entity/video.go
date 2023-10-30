package entity

type Video struct {
	Video_Id    string `json:"video_id"`
	Title       string `json:"title" binding:"min=2,max=100"`
	Description string `json:"description" binding:"max=200"`
	URL         string `json:"url" binding:"required,url"`
	Duration    int    `json:"duration"`
	//Author      User   `json:"author" binding:"required"`
}
