package models

type BaiscIndex struct {
	ID    string `json:"id" form:"id" example:"1"`
	Index string `json:"index" form:"index" example:"index"`
}

type UserIndex struct {
	ID           string        `json:"id" form:"id" example:"1"`
	Name         string        `json:"name" form:"name" example:"윤정은"`
	School       string        `json:"school" form:"school" example:"서울대학교"`
	Email        string        `json:"email" form:"email" example:"test@gmail.com"`
	Year         string        `json:"year" form:"year" example:"7"`
	Grade        string        `json:"grade" form:"grade" example:"3"`
	Stack        []string      `json:"stack" form:"stack" example:"[Go]"`
	ProfileImage string        `json:"profileImage" form:"profileImage" example:"https://example.com/profile.jpg"`
	MainPosition string        `json:"mainPosition" form:"mainPosition" example:"BACKEND"`
	ProjectTeams []UserProject `json:"projectTeams" form:"projectTeams" example:"[{mainImage: 'https://example.com/main.jpg'}]"`
}

type UserProject struct {
	MainImage string `json:"mainImage" form:"mainImage" example:"https://example.com/main.jpg"`
}

type ProjectIndex struct {
	ID              string             `json:"id" form:"id" example:"1"`
	Name            string             `json:"name" form:"name" example:"프로젝트 이름"`
	Title           string             `json:"title" form:"title" example:"내부용. 무시하세요"`
	ProjectExplain  string             `json:"projectExplain" form:"projectExplain" example:"프로젝트 설명"`
	MainImages      []string           `json:"mainImages" form:"mainImages" example:"https://example.com/main.jpg"`
	TeamStacks      []ProjectTeamStack `json:"teamStacks" form:"teamStacks" example:"[{stack: 'Go', isMain: true}]"`
	Recruited       bool               `json:"recruited" form:"recruited" example:"true"`
	Finished        bool               `json:"finished" form:"finished" example:"false"`
	FrontendNum     string             `json:"frontendNum" form:"frontendNum" example:"2"`
	BackendNum      string             `json:"backendNum" form:"backendNum" example:"1"`
	DevOpsNum       string             `json:"devOpsNum" form:"devOpsNum" example:"1"`
	FullstackNum    string             `json:"fullstackNum" form:"fullstackNum" example:"0"`
	DataEngineerNum string             `json:"dataEngineerNum" form:"dataEngineerNum" example:"0"`
	CreatedAt       string             `json:"createdAt" form:"createdAt" example:"2021-01-01"`
}

type ProjectTeamStack struct {
	Stack  string `json:"stack" form:"stack" example:"Go"`
	IsMain bool   `json:"isMain" form:"isMain" example:"true"`
}

type StudyIndex struct {
	ID           string `json:"id" form:"id" example:"1"`
	Name         string `json:"name" form:"name" example:"스터디 이름"`
	Title        string `json:"title" form:"title" example:"내부용. 무시하세요"`
	StudyExplain string `json:"studyExplain" form:"studyExplain" example:"스터디 설명"`
	Recruited    bool   `json:"recruited" form:"recruited" example:"true"`
	RecruitNum   string `json:"recruitNum" form:"recruitNum" example:"5"`
	Finished     bool   `json:"finished" form:"finished" example:"false"`
	CreatedAt    string `json:"createdAt" form:"createdAt" example:"2021-01-01"`
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
	CreatedAt        string `json:"createdAt" form:"createdAt" example:"2021-01-01"`
	UserID           string `json:"userID" form:"userID" example:"1"`
	UserName         string `json:"userName" form:"userName" example:"윤정은"`
	UserProfileImage string `json:"userProfileImage" form:"userProfileImage" example:"https://example.com/profile.jpg"`
	Year             string `json:"year" form:"year" example:"7"`
	Position         string `json:"position" form:"position" example:"BACKEND"`
}

type SessionIndex struct {
	ID        string   `json:"id" form:"id" example:"1"`
	Thumbnail string   `json:"thumbnail" form:"thumbnail" example:"https://example.com/thumbnail.jpg"`
	Title     string   `json:"title" form:"title" example:"세션 제목"`
	Presenter string   `json:"presenter" form:"presenter" example:"윤정은"`
	Date      string   `json:"date" form:"date" example:"2021-01-01"`
	LikeCount string   `json:"likeCount" form:"likeCount" example:"10"`
	VideoURL  string   `json:"videoURL" form:"videoURL" example:"https://example.com/video"`
	FileURL   string   `json:"fileURL" form:"fileURL" example:"https://example.com/file"`
	UserID    string   `json:"userID" form:"userID" example:"1"`
	User      UserInfo `json:"user" form:"user"`
}

type UserInfo struct {
	Name         string `json:"name" form:"name" example:"윤정은"`
	ProfileImage string `json:"profileImage" form:"profileImage" example:"https://example.com/profile.jpg"`
}

type EventIndex struct {
	ID       string `json:"id" form:"id" example:"1"`
	URl      string `json:"url" form:"url" example:"https://example.com/event"`
	Category string `json:"category" form:"category" example:"세션"`
	Title    string `json:"title" form:"title" example:"세션 제목"`
}

type StackIndex struct {
	ID       string `json:"id" form:"id" example:"1"`
	Name     string `json:"name" form:"name" example:"Go"`
	Category string `json:"category" form:"category" example:"BACKEND"`
}

type BootcampIndex struct {
	ID       string `json:"id" form:"id" example:"1"`
	Year     string `json:"year" form:"year" example:"2021"`
	ImageURL string `json:"imageURL" form:"imageURL" example:"https://example.com/bootcamp.jpg"`
	Rank     string `json:"rank" form:"rank" example:"1"`
}
