package models

type BaiscIndex struct {
	ID    string `json:"id" form:"id" example:"1"`
	Index string `json:"index" form:"index" example:"index"`
}

type UserIndex struct {
	ID           string   `json:"id" form:"id" example:"1"`
	Name         string   `json:"name" form:"name" example:"윤정은"`
	School       string   `json:"school" form:"school" example:"서울대학교"`
	Email        string   `json:"email" form:"email" example:"test@gmail.com"`
	Year         int      `json:"year" form:"year" example:"7"`
	Grade        string   `json:"grade" form:"grade" example:"3"`
	Stack        []string `json:"stack" form:"stack" example:"Go"`
	ProfileImage string   `json:"profileImage" form:"profileImage" example:"https://example.com/profile.jpg"`
}

type ProjectIndex struct {
	ID             string   `json:"id" form:"id" example:"1"`
	Name           string   `json:"name" form:"name" example:"프로젝트 이름"`
	Title          string   `json:"title" form:"title" example:"프로젝트 제목"`
	ProjectExplain string   `json:"projectExplain" form:"projectExplain" example:"프로젝트 설명"`
	ResultImages   []string `json:"resultImages" form:"resultImages" example:"https://example.com/result.jpg"`
	TeamStacks     []string `json:"teamStacks" form:"teamStacks" example:"Go"`
}

type StudyIndex struct {
	ID           string   `json:"id" form:"id" example:"1"`
	Name         string   `json:"name" form:"name" example:"스터디 이름"`
	Title        string   `json:"title" form:"title" example:"스터디 제목"`
	StudyExplain string   `json:"studyExplain" form:"studyExplain" example:"스터디 설명"`
	ResultImages []string `json:"resultImages" form:"resultImages" example:"https://example.com/result.jpg"`
}

type BlogIndex struct {
	ID               string   `json:"id" form:"id" example:"1"`
	Title            string   `json:"title" form:"title" example:"블로그 제목"`
	URL              string   `json:"url" form:"url" example:"https://example.com/blog"`
	Date             string   `json:"date" form:"date" example:"2021-01-01"`
	UserID           string   `json:"userID" form:"userID" example:"1"`
	UserName         string   `json:"userName" form:"userName" example:"윤정은"`
	UserProfileImage string   `json:"userProfileImage" form:"userProfileImage" example:"https://example.com/profile.jpg"`
	Thumbnail        string   `json:"thumbnail" form:"thumbnail" example:"https://example.com/thumbnail.jpg"`
	Stack            []string `json:"stack" form:"stack" example:"Go"`
}

type ResumeIndex struct {
	ID               string `json:"id" form:"id" example:"1"`
	Title            string `json:"title" form:"title" example:"이력서 제목"`
	URL              string `json:"url" form:"url" example:"https://example.com/resume"`
	CreatedAt        string `json:"createdAt" form:"createdAt" example:"2021-01-01"`
	UserID           string `json:"userID" form:"userID" example:"1"`
	UserName         string `json:"userName" form:"userName" example:"윤정은"`
	UserProfileImage string `json:"userProfileImage" form:"userProfileImage" example:"https://example.com/profile.jpg"`
}

type SessionIndex struct {
	ID        string `json:"id" form:"id" example:"1"`
	Thumbnail string `json:"thumbnail" form:"thumbnail" example:"https://example.com/thumbnail.jpg"`
	Title     string `json:"title" form:"title" example:"세션 제목"`
	Presenter string `json:"presenter" form:"presenter" example:"윤정은"`
	Date      string `json:"date" form:"date" example:"2021-01-01"`
	LikeCount int    `json:"likeCount" form:"likeCount" example:"10"`
	ViewCount int    `json:"viewCount" form:"viewCount" example:"100"`
}

type EventIndex struct {
	ID       string `json:"id" form:"id" example:"1"`
	Category string `json:"category" form:"category" example:"세션"`
	Title    string `json:"title" form:"title" example:"세션 제목"`
}
