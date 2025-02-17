package repositories

import (
	"database/sql"
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

// create all index should retrieve all data from database and create index in meilisearch
func (r *indexRepository) CreateAllIndex(condition string) (*meilisearch.TaskInfo, error) {
	// Retrieve all data from the database
	userData, err := readUserTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	projectData, err := readProjectTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	studyData, err := readStudyTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	blogData, err := readBlogTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	resumeData, err := readResumeTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	sessionData, err := readSessionTable(r.db, condition)
	if err != nil {
		return nil, err
	}
	eventData, err := readEventTable(r.db, condition)
	if err != nil {
		return nil, err
	}

	info, err := (*r.meiliclient).Index("user").AddDocuments(userData)
	if err != nil {
		return info, err
	}
	info, err = (*r.meiliclient).Index("project").AddDocuments(projectData)
	if err != nil {
		return info, err
	}
	info, err = (*r.meiliclient).Index("study").AddDocuments(studyData)
	if err != nil {
		return info, err
	}
	info, err = (*r.meiliclient).Index("blog").AddDocuments(blogData)
	if err != nil {
		return info, err
	}
	info, err = (*r.meiliclient).Index("resume").AddDocuments(resumeData)
	if err != nil {
		return info, err
	}
	info, err = (*r.meiliclient).Index("session").AddDocuments(sessionData)
	if err != nil {
		return info, err
	}
	info, err = (*r.meiliclient).Index("event").AddDocuments(eventData)
	if err != nil {
		return info, err
	}
	return info, nil
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
				Year:         year,
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
				Year:         year,
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
		rows, err := db.Query(`SELECT id, name, "projectExplain", "resultImages", "teamStacks" FROM public."ProjectTeam" WHERE "isDeleted" = false`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var projects []models.ProjectIndex
		for rows.Next() {
			var id int
			var name, projectExplain string
			var resultImages, teamStacks pq.StringArray
			err := rows.Scan(&id, &name, &projectExplain, &resultImages, &teamStacks)
			if err != nil {
				return nil, err
			}
			project := models.ProjectIndex{
				ID:             strconv.Itoa(id),
				Name:           name,
				Title:          name,
				ProjectExplain: projectExplain,
				ResultImages:   resultImages,
				TeamStacks:     teamStacks,
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
		// Fetch new users that are created/updated after lastSyncedAt
		rows, err := db.Query(`
            SELECT id, name, "projectExplain", "resultImages", "teamStacks"
			FROM public."ProjectTeam"
            WHERE ("createdAt" > $1 OR "updatedAt" > $1) AND "isDeleted" = false
        `, lastSyncedAt)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var projects []models.ProjectIndex
		for rows.Next() {
			var id int
			var name, projectExplain string
			var resultImages, teamStacks pq.StringArray
			err := rows.Scan(&id, &name, &projectExplain, &resultImages, &teamStacks)
			if err != nil {
				return nil, err
			}
			project := models.ProjectIndex{
				ID:             strconv.Itoa(id),
				Name:           name,
				Title:          name,
				ProjectExplain: projectExplain,
				ResultImages:   resultImages,
				TeamStacks:     teamStacks,
			}

			projects = append(projects, project)
		}
		return projects, nil
	}
	return nil, nil
}

func readStudyTable(db *sql.DB, condition string) ([]models.StudyIndex, error) {
	if condition == "all" {
		rows, err := db.Query(`SELECT id, name, "studyExplain", "resultImages" FROM public."StudyTeam" WHERE "isDeleted" = false`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var studies []models.StudyIndex
		for rows.Next() {
			var id int
			var name, studyExplain string
			var resultImages pq.StringArray
			err := rows.Scan(&id, &name, &studyExplain, &resultImages)
			if err != nil {
				return nil, err
			}
			study := models.StudyIndex{
				ID:           strconv.Itoa(id),
				Name:         name,
				Title:        name,
				StudyExplain: studyExplain,
				ResultImages: resultImages,
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
            SELECT id, name, "studyExplain", "resultImages"
			FROM public."StudyTeam"
            WHERE ("createdAt" > $1 OR "updatedAt" > $1) AND "isDeleted" = false
        `, lastSyncedAt)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var studies []models.StudyIndex
		for rows.Next() {
			var id int
			var name, studyExplain string
			var resultImages pq.StringArray
			err := rows.Scan(&id, &name, &studyExplain, &resultImages)
			if err != nil {
				return nil, err
			}
			study := models.StudyIndex{
				ID:           strconv.Itoa(id),
				Name:         name,
				Title:        name,
				StudyExplain: studyExplain,
				ResultImages: resultImages,
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
			SELECT b.id, b.title, b.url, b."createdAt", b."userId", u.name, u."profileImage"
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
			var id, userId int
			var title, url, userName, userProfileImage string
			var createdAt time.Time
			err := rows.Scan(&id, &title, &url, &createdAt, &userId, &userName, &userProfileImage)
			if err != nil {
				return nil, err
			}
			resume := models.ResumeIndex{
				ID:               strconv.Itoa(id),
				Title:            title,
				URL:              url,
				CreatedAt:        createdAt.Format(time.RFC3339),
				UserID:           strconv.Itoa(userId),
				UserName:         userName,
				UserProfileImage: userProfileImage,
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
			SELECT b.id, b.title, b.url, b."createdAt", b."userId", u.name, u."profileImage"
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
			var id, userId int
			var title, url, userName, userProfileImage string
			var createdAt time.Time
			err := rows.Scan(&id, &title, &url, &createdAt, &userId, &userName, &userProfileImage)
			if err != nil {
				return nil, err
			}
			resume := models.ResumeIndex{
				ID:               strconv.Itoa(id),
				Title:            title,
				URL:              url,
				CreatedAt:        createdAt.Format(time.RFC3339),
				UserID:           strconv.Itoa(userId),
				UserName:         userName,
				UserProfileImage: userProfileImage,
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
				LikeCount: likeCount,
				ViewCount: viewCount,
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
				LikeCount: likeCount,
				ViewCount: viewCount,
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
		// Fetch all sessions
		rows, err := db.Query(`
            SELECT id, category, title
            FROM public."Session"
            WHERE "isDeleted" = false
        `)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var sessions []models.EventIndex
		for rows.Next() {
			var id int
			var title, category string
			err := rows.Scan(&id, &category, &title)
			if err != nil {
				return nil, err
			}

			session := models.EventIndex{
				ID:       strconv.Itoa(id),
				Title:    title,
				Category: category,
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
            SELECT id, category, title
            FROM public."Session"
            WHERE "isDeleted" = false AND ("createdAt" > $1 OR "updatedAt" > $1)
        `, lastSyncedAt)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var sessions []models.EventIndex
		for rows.Next() {
			var id int
			var title, category string
			err := rows.Scan(&id, &category, &title)
			if err != nil {
				return nil, err
			}

			session := models.EventIndex{
				ID:       strconv.Itoa(id),
				Title:    title,
				Category: category,
			}

			sessions = append(sessions, session)
		}

		return sessions, nil
	}
	return nil, nil
}
