package repositories

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/Techeer-Hogwarts/search-cron/models"
	"github.com/lib/pq"
	"github.com/meilisearch/meilisearch-go"
)

type IndexRepository interface {
	CreateUserIndex(post *models.UserIndex) (*meilisearch.TaskInfo, error)
	CreateProjectIndex(post *models.ProjectIndex) (*meilisearch.TaskInfo, error)
	CreateStudyIndex(post *models.StudyIndex) (*meilisearch.TaskInfo, error)
	CreateBlogIndex(post *models.BlogIndex) (*meilisearch.TaskInfo, error)
	CreateResumeIndex(post *models.ResumeIndex) (*meilisearch.TaskInfo, error)
	CreateSessionIndex(post *models.SessionIndex) (*meilisearch.TaskInfo, error)
	CreateEventIndex(post *models.EventIndex) (*meilisearch.TaskInfo, error)
	CreateStackIndex() (*meilisearch.TaskInfo, error)
	CreateAllIndex(condition string) (*meilisearch.TaskInfo, error)
	DeleteIndexDocument(id, index string) (*meilisearch.TaskInfo, error)
	DeleteIndex(index string) (*meilisearch.TaskInfo, error)
	DeleteAllIndex() (*meilisearch.TaskInfo, error)
	DeleteAllDocuments(index string) (*meilisearch.TaskInfo, error)
}

type indexRepository struct {
	db          *sql.DB
	meiliclient *meilisearch.ServiceManager
}

func NewIndexRepository(db *sql.DB, meiliclient *meilisearch.ServiceManager) IndexRepository {
	return &indexRepository{db, meiliclient}
}

func (r *indexRepository) CreateUserIndex(post *models.UserIndex) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).Index("user").AddDocuments(post)
	if err != nil {
		return info, err
	}
	return info, nil
}

func (r *indexRepository) CreateProjectIndex(post *models.ProjectIndex) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).Index("project").AddDocuments(post)
	if err != nil {
		return info, err
	}
	return info, nil
}

func (r *indexRepository) CreateStudyIndex(post *models.StudyIndex) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).Index("study").AddDocuments(post)
	if err != nil {
		return info, err
	}
	return info, nil
}

func (r *indexRepository) CreateBlogIndex(post *models.BlogIndex) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).Index("blog").AddDocuments(post, "id")
	if err != nil {
		return info, err
	}
	return info, nil
}

func (r *indexRepository) CreateResumeIndex(post *models.ResumeIndex) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).Index("resume").AddDocuments(post, "id")
	if err != nil {
		return info, err
	}
	return info, nil
}

func (r *indexRepository) CreateSessionIndex(post *models.SessionIndex) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).Index("session").AddDocuments(post)
	if err != nil {
		return info, err
	}
	return info, nil
}

func (r *indexRepository) CreateEventIndex(post *models.EventIndex) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).Index("event").AddDocuments(post)
	if err != nil {
		return info, err
	}
	return info, nil
}

func (r *indexRepository) CreateStackIndex() (*meilisearch.TaskInfo, error) {
	stackData, err := readStackTable(r.db)
	if err != nil {
		return nil, err
	}
	if len(stackData) != 0 {
		info, err := (*r.meiliclient).Index("stack").AddDocuments(stackData, "id")
		if err != nil {
			return info, err
		}
		return info, nil
	}
	return nil, nil
}

// create all index should retrieve all data from database and create index in meilisearch
func (r *indexRepository) CreateAllIndex(condition string) (*meilisearch.TaskInfo, error) {
	// Retrieve all data from the database
	userData, err := readUserTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	log.Printf("Got user data %v", userData)
	projectData, err := readProjectTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	log.Printf("Got project data %v", projectData)
	studyData, err := readStudyTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	log.Printf("Got study data %v", studyData)
	blogData, err := readBlogTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	log.Printf("Got blog data %v", blogData)
	resumeData, err := readResumeTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	log.Printf("Got resume data %v", resumeData)
	sessionData, err := readSessionTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	log.Printf("Got session data %v", sessionData)
	eventData, err := readEventTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	log.Printf("Got event data %v", eventData)
	if len(userData) != 0 {
		info, err := (*r.meiliclient).Index("user").AddDocuments(userData, "id")
		if err != nil {
			return info, err
		}
	}
	if len(projectData) != 0 {
		info, err := (*r.meiliclient).Index("project").AddDocuments(projectData, "id")
		if err != nil {
			return info, err
		}
	}
	if len(studyData) != 0 {
		info, err := (*r.meiliclient).Index("study").AddDocuments(studyData, "id")
		if err != nil {
			return info, err
		}
	}
	if len(blogData) != 0 {
		info, err := (*r.meiliclient).Index("blog").AddDocuments(blogData, "id")
		if err != nil {
			return info, err
		}
	}
	if len(resumeData) != 0 {
		info, err := (*r.meiliclient).Index("resume").AddDocuments(resumeData, "id")
		if err != nil {
			return info, err
		}
	}
	if len(sessionData) != 0 {
		info, err := (*r.meiliclient).Index("session").AddDocuments(sessionData, "id")
		if err != nil {
			return info, err
		}
	}
	if len(eventData) != 0 {
		info, err := (*r.meiliclient).Index("event").AddDocuments(eventData, "id")
		if err != nil {
			return info, err
		}
	}
	// Update the lastSyncedAt value in SyncDb table
	err = updateLastSyncedAt(r.db)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *indexRepository) DeleteIndexDocument(id, index string) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).Index(index).DeleteDocument(id)
	if err != nil {
		return info, err
	}
	return info, nil
}

func (r *indexRepository) DeleteIndex(index string) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).DeleteIndex(index)
	if err != nil {
		return info, err
	}
	return info, nil
}

func (r *indexRepository) DeleteAllIndex() (*meilisearch.TaskInfo, error) {
	param := &meilisearch.IndexesQuery{}
	index, err := (*r.meiliclient).ListIndexes(param)
	if err != nil {
		return nil, err
	}
	for _, v := range index.Results {
		info, err := (*r.meiliclient).DeleteIndex(v.UID)
		if err != nil {
			return info, err
		}
	}
	return nil, nil
}

func (r *indexRepository) DeleteAllDocuments(index string) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).Index(index).DeleteAllDocuments()
	if err != nil {
		return info, err
	}
	return info, nil
}

func readUserTable(db *sql.DB, condition string) ([]models.UserIndex, error) {
	if condition == "all" {
		rows, err := db.Query(`SELECT id, name, school, email, year, grade, stack, "profileImage" FROM public."User" WHERE "isDeleted" = false`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var users []models.UserIndex
		for rows.Next() {
			var id, year int
			var name, school, email, grade, profileImage string
			var stack pq.StringArray
			err := rows.Scan(&id, &name, &school, &email, &year, &grade, &stack, &profileImage)
			if err != nil {
				return nil, err
			}
			user := models.UserIndex{
				ID:           strconv.Itoa(id),
				Name:         name,
				School:       school,
				Email:        email,
				Year:         strconv.Itoa(year),
				Grade:        grade,
				Stack:        stack,
				ProfileImage: profileImage,
			}

			users = append(users, user)
		}
		return users, nil
	} else if condition == "new" {
		// Fetch the lastSyncedAt value from SyncDb table
		var lastSyncedAt time.Time
		err := db.QueryRow(`SELECT "lastSyncedAt" FROM public."SyncDb" ORDER BY id DESC LIMIT 1`).Scan(&lastSyncedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				// If no rows found, set a default date to fetch all records
				lastSyncedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
			} else {
				return nil, err
			}
		}
		log.Println("lastSyncedAt: ", lastSyncedAt)
		// Fetch new users that are created/updated after lastSyncedAt
		rows, err := db.Query(`
            SELECT id, name, school, email, year, grade, stack, "profileImage"
            FROM public."User"
            WHERE ("createdAt" > $1 OR "updatedAt" > $1) AND "isDeleted" = false
        `, lastSyncedAt)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var users []models.UserIndex
		for rows.Next() {
			var id, year int
			var name, school, email, grade, profileImage string
			var stack pq.StringArray
			err := rows.Scan(&id, &name, &school, &email, &year, &grade, &stack, &profileImage)
			if err != nil {
				return nil, err
			}
			user := models.UserIndex{
				ID:           strconv.Itoa(id),
				Name:         name,
				School:       school,
				Email:        email,
				Year:         strconv.Itoa(year),
				Grade:        grade,
				Stack:        stack,
				ProfileImage: profileImage,
			}

			users = append(users, user)
		}
		return users, nil
	}
	return nil, nil
}

func readProjectTable(db *sql.DB, condition string) ([]models.ProjectIndex, error) {
	if condition == "all" {
		rows, err := db.Query(`
			SELECT 
				p.id, 
				p.name, 
				p."projectExplain", 
				COALESCE(JSONB_AGG(DISTINCT s.name) FILTER (WHERE s.name IS NOT NULL), '[]') AS "teamStacks",
				COALESCE(JSONB_AGG(r."imageUrl") FILTER (WHERE r."imageUrl" IS NOT NULL), '[]') AS "resultImages"
			FROM public."ProjectTeam" p
			LEFT JOIN public."ProjectMainImage" r 
				ON p.id = r."projectTeamId" AND r."isDeleted" = false
			LEFT JOIN public."TeamStack" ts 
				ON p.id = ts."projectTeamId" AND ts."isDeleted" = false
			LEFT JOIN public."Stack" s 
				ON ts."stackId" = s.id AND s."isDeleted" = false
			WHERE p."isDeleted" = false
			GROUP BY p.id, p.name, p."projectExplain"
		`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var projects []models.ProjectIndex
		for rows.Next() {
			var id int
			var name, projectExplain string
			var resultImagesJSON, teamStacksJSON []byte

			err := rows.Scan(&id, &name, &projectExplain, &teamStacksJSON, &resultImagesJSON)
			if err != nil {
				return nil, err
			}

			// Convert JSONB to Go slice
			var teamStacks, resultImages []string
			if err := json.Unmarshal(teamStacksJSON, &teamStacks); err != nil {
				return nil, err
			}
			if err := json.Unmarshal(resultImagesJSON, &resultImages); err != nil {
				return nil, err
			}

			project := models.ProjectIndex{
				ID:             strconv.Itoa(id),
				Name:           name,
				Title:          name, // Assuming Title is the same as Name
				ProjectExplain: projectExplain,
				// ResultImages:   resultImages, // Now it's just an array of image URLs
				// TeamStacks:     teamStacks,
			}

			projects = append(projects, project)
		}
		return projects, nil
	} else if condition == "new" {
		// Fetch the lastSyncedAt value from SyncDb table
		var lastSyncedAt time.Time
		err := db.QueryRow(`SELECT "lastSyncedAt" FROM public."SyncDb" ORDER BY id DESC LIMIT 1`).Scan(&lastSyncedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				// If no rows found, set a default date to fetch all records
				lastSyncedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
			} else {
				return nil, err
			}
		}
		log.Println("lastSyncedAt: ", lastSyncedAt)
		rows, err := db.Query(`
			SELECT 
				p.id, 
				p.name, 
				p."projectExplain", 
				COALESCE(JSONB_AGG(DISTINCT s.name) FILTER (WHERE s.name IS NOT NULL), '[]') AS "teamStacks",
				COALESCE(JSONB_AGG(r."imageUrl") FILTER (WHERE r."imageUrl" IS NOT NULL), '[]') AS "resultImages"
			FROM public."ProjectTeam" p
			LEFT JOIN public."ProjectMainImage" r 
				ON p.id = r."projectTeamId" AND r."isDeleted" = false
			LEFT JOIN public."TeamStack" ts 
				ON p.id = ts."projectTeamId" AND ts."isDeleted" = false
			LEFT JOIN public."Stack" s 
				ON ts."stackId" = s.id AND s."isDeleted" = false
			WHERE p."isDeleted" = false AND (p."createdAt" > $1 OR p."updatedAt" > $1)
			GROUP BY p.id, p.name, p."projectExplain"
		`, lastSyncedAt)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var projects []models.ProjectIndex
		for rows.Next() {
			var id int
			var name, projectExplain string
			var resultImagesJSON, teamStacksJSON []byte

			err := rows.Scan(&id, &name, &projectExplain, &teamStacksJSON, &resultImagesJSON)
			if err != nil {
				return nil, err
			}

			// Convert JSONB to Go slice
			var teamStacks, resultImages []string
			if err := json.Unmarshal(teamStacksJSON, &teamStacks); err != nil {
				return nil, err
			}
			if err := json.Unmarshal(resultImagesJSON, &resultImages); err != nil {
				return nil, err
			}

			project := models.ProjectIndex{
				ID:             strconv.Itoa(id),
				Name:           name,
				Title:          name, // Assuming Title is the same as Name
				ProjectExplain: projectExplain,
				// ResultImages:   resultImages, // Now it's just an array of image URLs
				// TeamStacks:     teamStacks,
			}

			projects = append(projects, project)
		}
		return projects, nil
	}
	return nil, nil
}

func readStudyTable(db *sql.DB, condition string) ([]models.StudyIndex, error) {
	if condition == "all" {
		rows, err := db.Query(`
			SELECT 
				p.id, 
				p.name, 
				p."studyExplain", 
				COALESCE(JSONB_AGG(r."imageUrl") FILTER (WHERE r."imageUrl" IS NOT NULL), '[]') AS "resultImages"
			FROM public."StudyTeam" p
			LEFT JOIN public."StudyResultImage" r 
				ON p.id = r."studyTeamId" AND r."isDeleted" = false
			WHERE p."isDeleted" = false
			GROUP BY p.id, p.name, p."studyExplain"
		`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var studies []models.StudyIndex
		for rows.Next() {
			var id int
			var name, studyExplain string
			var resultImagesJSON []byte
			err := rows.Scan(&id, &name, &studyExplain, &resultImagesJSON)
			if err != nil {
				return nil, err
			}
			var resultImages []string
			if err := json.Unmarshal(resultImagesJSON, &resultImages); err != nil {
				return nil, err
			}
			study := models.StudyIndex{
				ID:           strconv.Itoa(id),
				Name:         name,
				Title:        name,
				StudyExplain: studyExplain,
				// ResultImages: resultImages,
			}

			studies = append(studies, study)
		}
		return studies, nil
	} else if condition == "new" {
		// Fetch the lastSyncedAt value from SyncDb table
		var lastSyncedAt time.Time
		err := db.QueryRow(`SELECT "lastSyncedAt" FROM public."SyncDb" ORDER BY id DESC LIMIT 1`).Scan(&lastSyncedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				// If no rows found, set a default date to fetch all records
				lastSyncedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
			} else {
				return nil, err
			}
		}
		log.Println("lastSyncedAt: ", lastSyncedAt)
		// Fetch new users that are created/updated after lastSyncedAt
		rows, err := db.Query(`
			SELECT 
				p.id, 
				p.name, 
				p."studyExplain", 
				COALESCE(JSONB_AGG(r."imageUrl") FILTER (WHERE r."imageUrl" IS NOT NULL), '[]') AS "resultImages"
			FROM public."StudyTeam" p
			LEFT JOIN public."StudyResultImage" r 
				ON p.id = r."studyTeamId" AND r."isDeleted" = false
			WHERE p."isDeleted" = false AND (p."createdAt" > $1 OR p."updatedAt" > $1)
			GROUP BY p.id, p.name, p."studyExplain"
		`, lastSyncedAt)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var studies []models.StudyIndex
		for rows.Next() {
			var id int
			var name, studyExplain string
			var resultImagesJSON []byte
			err := rows.Scan(&id, &name, &studyExplain, &resultImagesJSON)
			if err != nil {
				return nil, err
			}
			var resultImages []string
			if err := json.Unmarshal(resultImagesJSON, &resultImages); err != nil {
				return nil, err
			}
			study := models.StudyIndex{
				ID:           strconv.Itoa(id),
				Name:         name,
				Title:        name,
				StudyExplain: studyExplain,
				// ResultImages: resultImages,
			}
			studies = append(studies, study)
		}
		return studies, nil
	}
	return nil, nil
}

func readBlogTable(db *sql.DB, condition string) ([]models.BlogIndex, error) {
	if condition == "all" {
		rows, err := db.Query(`
			SELECT b.id, b.title, b.url, b.date, b."userId", u.name, u."profileImage", b.tags, b.thumbnail
			FROM public."Blog" b
			JOIN public."User" u ON b."userId" = u.id
			WHERE b."isDeleted" = false
		`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var blogs []models.BlogIndex
		for rows.Next() {
			var id, userId int
			var title, url, date, userName, userProfileImage, thumbnail string
			var tags pq.StringArray
			err := rows.Scan(&id, &title, &url, &date, &userId, &userName, &userProfileImage, &tags, &thumbnail)
			if err != nil {
				return nil, err
			}
			blog := models.BlogIndex{
				ID:               strconv.Itoa(id),
				Title:            title,
				URL:              url,
				Date:             date,
				UserID:           strconv.Itoa(userId),
				UserName:         userName,
				UserProfileImage: userProfileImage,
				Stack:            tags,
				Thumbnail:        thumbnail,
			}

			blogs = append(blogs, blog)
		}
		return blogs, nil
	} else if condition == "new" {
		// Fetch the lastSyncedAt value from SyncDb table
		var lastSyncedAt time.Time
		err := db.QueryRow(`SELECT "lastSyncedAt" FROM public."SyncDb" ORDER BY id DESC LIMIT 1`).Scan(&lastSyncedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				// If no rows found, set a default date to fetch all records
				lastSyncedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
			} else {
				return nil, err
			}
		}
		log.Println("lastSyncedAt: ", lastSyncedAt)
		// Fetch new users that are created/updated after lastSyncedAt
		rows, err := db.Query(`
            SELECT b.id, b.title, b.url, b.date, b."userId", u.name, u."profileImage", b.tags, b.thumbnail
			FROM public."Blog" b
			JOIN public."User" u ON b."userId" = u.id
			WHERE (b."createdAt" > $1 OR b."updatedAt" > $1) AND b."isDeleted" = false
        `, lastSyncedAt)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var blogs []models.BlogIndex
		for rows.Next() {
			var id, userId int
			var title, url, date, userName, userProfileImage, thumbnail string
			var tags pq.StringArray
			err := rows.Scan(&id, &title, &url, &date, &userId, &userName, &userProfileImage, &tags, &thumbnail)
			if err != nil {
				return nil, err
			}
			blog := models.BlogIndex{
				ID:               strconv.Itoa(id),
				Title:            title,
				URL:              url,
				Date:             date,
				UserID:           strconv.Itoa(userId),
				UserName:         userName,
				UserProfileImage: userProfileImage,
				Stack:            tags,
				Thumbnail:        thumbnail,
			}

			blogs = append(blogs, blog)
		}
		return blogs, nil
	}
	return nil, nil
}

func readResumeTable(db *sql.DB, condition string) ([]models.ResumeIndex, error) {
	if condition == "all" {
		rows, err := db.Query(`
			SELECT b.id, b.title, b."createdAt", b."userId", u.name, u."profileImage", u.year, b.position
			FROM public."Resume" b
			JOIN public."User" u ON b."userId" = u.id
			WHERE b."isDeleted" = false
		`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var resumes []models.ResumeIndex
		for rows.Next() {
			var id, userId, year int
			var title, userName, userProfileImage, position string
			var createdAt time.Time
			err := rows.Scan(&id, &title, &createdAt, &userId, &userName, &userProfileImage, &year, &position)
			if err != nil {
				return nil, err
			}
			resume := models.ResumeIndex{
				ID:               strconv.Itoa(id),
				Title:            title,
				CreatedAt:        createdAt.Format(time.RFC3339),
				UserID:           strconv.Itoa(userId),
				UserName:         userName,
				UserProfileImage: userProfileImage,
				Year:             strconv.Itoa(year),
				Position:         position,
			}

			resumes = append(resumes, resume)
		}
		return resumes, nil
	} else if condition == "new" {
		// Fetch the lastSyncedAt value from SyncDb table
		var lastSyncedAt time.Time
		err := db.QueryRow(`SELECT "lastSyncedAt" FROM public."SyncDb" ORDER BY id DESC LIMIT 1`).Scan(&lastSyncedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				// If no rows found, set a default date to fetch all records
				lastSyncedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
			} else {
				return nil, err
			}
		}
		log.Println("lastSyncedAt: ", lastSyncedAt)
		// Fetch new users that are created/updated after lastSyncedAt
		rows, err := db.Query(`
			SELECT b.id, b.title, b."createdAt", b."userId", u.name, u."profileImage", u.year, b.position
			FROM public."Resume" b
			JOIN public."User" u ON b."userId" = u.id
			WHERE (b."createdAt" > $1 OR b."updatedAt" > $1) AND b."isDeleted" = false
		`, lastSyncedAt)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var resumes []models.ResumeIndex
		for rows.Next() {
			var id, userId, year int
			var title, userName, userProfileImage, position string
			var createdAt time.Time
			err := rows.Scan(&id, &title, &createdAt, &userId, &userName, &userProfileImage, &year, &position)
			if err != nil {
				return nil, err
			}
			resume := models.ResumeIndex{
				ID:               strconv.Itoa(id),
				Title:            title,
				CreatedAt:        createdAt.Format(time.RFC3339),
				UserID:           strconv.Itoa(userId),
				UserName:         userName,
				UserProfileImage: userProfileImage,
				Year:             strconv.Itoa(year),
				Position:         position,
			}

			resumes = append(resumes, resume)
		}
		return resumes, nil
	}
	return nil, nil
}

func readSessionTable(db *sql.DB, condition string) ([]models.SessionIndex, error) {
	if condition == "all" {
		// Fetch all sessions
		rows, err := db.Query(`
            SELECT s.id, s.title, s.presenter, s.date, s."likeCount", s."viewCount", s.thumbnail
            FROM public."Session" s
            WHERE s."isDeleted" = false
        `)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var sessions []models.SessionIndex
		for rows.Next() {
			var id int
			var title, presenter, date, thumbnail string
			var likeCount, viewCount int
			err := rows.Scan(&id, &title, &presenter, &date, &likeCount, &viewCount, &thumbnail)
			if err != nil {
				return nil, err
			}

			session := models.SessionIndex{
				ID:        strconv.Itoa(id),
				Title:     title,
				Presenter: presenter,
				Date:      date,
				LikeCount: strconv.Itoa(likeCount),
				Thumbnail: thumbnail,
			}

			sessions = append(sessions, session)
		}
		return sessions, nil
	} else if condition == "new" {
		// Fetch the lastSyncedAt value from SyncDb table
		var lastSyncedAt time.Time
		err := db.QueryRow(`SELECT "lastSyncedAt" FROM public."SyncDb" ORDER BY id DESC LIMIT 1`).Scan(&lastSyncedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				// If no rows found, set a default date to fetch all records
				lastSyncedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
			} else {
				return nil, err
			}
		}

		log.Println("lastSyncedAt: ", lastSyncedAt)

		// If lastSyncedAt is zero (not set), fetch all records (or use a default date)
		if lastSyncedAt.IsZero() {
			// Set a fallback date (you can use any sensible date)
			lastSyncedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
		}

		// Fetch new sessions created or updated after lastSyncedAt
		rows, err := db.Query(`
            SELECT s.id, s.title, s.presenter, s.date, s."likeCount", s."viewCount", s.thumbnail
            FROM public."Session" s
            WHERE s."isDeleted" = false AND (s."createdAt" > $1 OR s."updatedAt" > $1)
        `, lastSyncedAt)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var sessions []models.SessionIndex
		for rows.Next() {
			var id int
			var title, presenter, date, thumbnail string
			var likeCount, viewCount int
			err := rows.Scan(&id, &title, &presenter, &date, &likeCount, &viewCount, &thumbnail)
			if err != nil {
				return nil, err
			}

			session := models.SessionIndex{
				ID:        strconv.Itoa(id),
				Title:     title,
				Presenter: presenter,
				Date:      date,
				LikeCount: strconv.Itoa(likeCount),
				// ViewCount: strconv.Itoa(viewCount),
				Thumbnail: thumbnail,
			}

			sessions = append(sessions, session)
		}

		return sessions, nil
	}

	return nil, nil
}

func readEventTable(db *sql.DB, condition string) ([]models.EventIndex, error) {
	if condition == "all" {
		// Fetch all events
		rows, err := db.Query(`
            SELECT id, category, title, url
            FROM public."Event"
            WHERE "isDeleted" = false
        `)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var events []models.EventIndex
		for rows.Next() {
			var id int
			var title, category, url string
			err := rows.Scan(&id, &category, &title, &url)
			if err != nil {
				return nil, err
			}

			event := models.EventIndex{
				ID:       strconv.Itoa(id),
				URl:      url,
				Title:    title,
				Category: category,
			}

			events = append(events, event)
		}
		return events, nil
	} else if condition == "new" {
		// Fetch the lastSyncedAt value from SyncDb table
		var lastSyncedAt time.Time
		err := db.QueryRow(`SELECT "lastSyncedAt" FROM public."SyncDb" ORDER BY id DESC LIMIT 1`).Scan(&lastSyncedAt)
		if err != nil {
			if err == sql.ErrNoRows {
				// If no rows found, set a default date to fetch all records
				lastSyncedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
			} else {
				return nil, err
			}
		}

		log.Println("lastSyncedAt: ", lastSyncedAt)

		// If lastSyncedAt is zero (not set), fetch all records (or use a default date)
		if lastSyncedAt.IsZero() {
			// Set a fallback date (you can use any sensible date)
			lastSyncedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
		}

		// Fetch new events created or updated after lastSyncedAt
		rows, err := db.Query(`
            SELECT id, category, title, url
            FROM public."Event"
            WHERE "isDeleted" = false AND ("createdAt" > $1 OR "updatedAt" > $1)
        `, lastSyncedAt)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var events []models.EventIndex
		for rows.Next() {
			var id int
			var title, category, url string
			err := rows.Scan(&id, &category, &title, &url)
			if err != nil {
				return nil, err
			}

			event := models.EventIndex{
				ID:       strconv.Itoa(id),
				URl:      url,
				Title:    title,
				Category: category,
			}

			events = append(events, event)
		}

		return events, nil
	}
	return nil, nil
}

func readStackTable(db *sql.DB) ([]models.StackIndex, error) {
	rows, err := db.Query(`
		SELECT id, name, category
		FROM public."Stack"
		WHERE "isDeleted" = false
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stacks []models.StackIndex
	for rows.Next() {
		var id int
		var name, category string
		err := rows.Scan(&id, &name, &category)
		if err != nil {
			return nil, err
		}

		stack := models.StackIndex{
			ID:       strconv.Itoa(id),
			Name:     name,
			Category: category,
		}

		stacks = append(stacks, stack)
	}

	return stacks, nil
}

func updateLastSyncedAt(db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO public."SyncDb" ("lastSyncedAt") VALUES (NOW())`)
	if err != nil {
		return err
	}
	return nil
}
