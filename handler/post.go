package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sarkarijobadda/backend/conf"
	"github.com/sarkarijobadda/backend/model"
	"github.com/sarkarijobadda/backend/util"

	"github.com/labstack/echo"
)

// CreatePost method
func (h *Handler) CreatePost(c echo.Context) (err error) {
	p := &model.Post{}

	if err = c.Bind(p); err != nil {
		return
	}

	if len(p.ImageLink) == 0 {
		p.ImageLink = conf.DefaultImageLink
	}

	shortLink, err := util.CreateDynamicLink(p)
	if err != nil {
		print(err)
	}
	p.ShortLink = shortLink
	t := time.Now()
	p.CreatedAt = t
	p.UpdatedAt = t

	sql := `
	INSERT INTO posts(short_link, image_link, type, title, name, info, created_date, updated_date, 
	organisation, total_vacancy, age_limit_as_on, draft, created_at, updated_at) 
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	r, err := h.DB.Exec(sql, p.ShortLink, p.ImageLink, p.Type, p.Title, p.Name, p.Info, p.CreatedDate, p.UpdatedDate, p.Organisation, p.TotalVacancy, p.AgeLimitAsOn, p.Draft, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return
	}
	p.ID, _ = r.LastInsertId()

	sql = `INSERT INTO posts_search(id,title,name,info, organisation) VALUES (?,?,?,?,?)`
	_, err = h.DB.Exec(sql, p.ID, p.Title, p.Name, p.Info, p.Organisation)
	if err != nil {
		fmt.Println(err)
		return
	}

	sql = `INSERT INTO dates(date, title, post_id) VALUES (?,?,?)`
	for _, v := range p.Dates {
		_, err = h.DB.Exec(sql, v.Date, v.Title, p.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	sql = `INSERT INTO links(title, url, post_id) VALUES(?,?,?)`
	for _, v := range p.Links {
		_, err = h.DB.Exec(sql, v.Title, v.URL, p.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	sql = `INSERT INTO items(title, body, post_id) VALUES (?,?,?)`
	for _, v := range p.AgeLimits {
		_, err = h.DB.Exec(sql, v.Title, v.Body, p.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	sql = `INSERT INTO fees (title, body, post_id) VALUES (?,?,?)`
	for _, v := range p.Fees {
		_, err = h.DB.Exec(sql, v.Title, v.Body, p.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	sql = `
	INSERT into vacancies (category, name, gen, obc, bca, bcb, ews,
	sc, st, ph, total, age_limit, eligibility, post_id) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	for _, v := range p.Vacancies {
		_, err = h.DB.Exec(sql, v.Category, v.Name, v.Gen, v.OBC, v.BCA, v.BCB, v.EWS, v.SC, v.ST, v.PH, v.Total, v.AgeLimit, v.Eligibility, p.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	return c.JSON(http.StatusCreated, p)
}

// FetchPost method
func (h *Handler) FetchPost(c echo.Context) (err error) {
	posts := []*model.Post{}

	sql := `SELECT id, short_link, image_link, type, title, name, info, created_date, updated_date, organisation, total_vacancy, as_limit_as_on, draft, trash`

	rows, err := h.DB.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var p model.Post
		err = rows.Scan(&p.ID, &p.ShortLink, &p.ImageLink, &p.Type, &p.Title, &p.Name, &p.Info, &p.CreatedDate, &p.UpdatedDate, &p.Organisation, &p.TotalVacancy, &p.AgeLimitAsOn, &p.Draft, &p.Trash)
		if err != nil {
			log.Println(err)
		}
		posts = append(posts, &p)
	}

	return c.JSON(http.StatusOK, posts)
}

// FetchOnePost method
func (h *Handler) FetchOnePost(c echo.Context) (err error) {
	id := c.Param("id")
	var p model.Post

	sql := `SELECT id, short_link, image_link, type, title, name, info, created_date, updated_date, organisation, total_vacancy, age_limit_as_on, draft, trash FROM posts WHERE id = ?`
	row := h.DB.QueryRow(sql, id)
	err = row.Scan(&p.ID, &p.ShortLink, &p.ImageLink, &p.Type, &p.Title, &p.Name, &p.Info, &p.CreatedDate, &p.UpdatedDate, &p.Organisation, &p.TotalVacancy, &p.AgeLimitAsOn, &p.Draft, &p.Trash)

	if err != nil {
		return
	}

	p.Dates = h.getPostImportantDates(p.ID)
	p.Links = h.getPostImportantLink(p.ID)
	p.Fees = h.getPostApplicationFees(p.ID)
	p.AgeLimits = h.getPostAgeLimits(p.ID)
	p.Vacancies = h.getPostAgeVacancies(p.ID)

	return c.JSON(http.StatusOK, p)
}

func (h *Handler) getPostImportantDates(id int64) (dates []model.ImportantDate) {

	rows, err := h.DB.Query(`SELECT id, title, date FROM dates WHERE post_id = ?`, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var d model.ImportantDate
		err = rows.Scan(&d.ID, &d.Title, &d.Date)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		dates = append(dates, d)
	}
	return dates
}

func (h *Handler) getPostImportantLink(id int64) (links []model.ImportantLink) {

	rows, err := h.DB.Query(`SELECT id, title, url FROM links WHERE post_id = ?`, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var l model.ImportantLink
		err = rows.Scan(&l.ID, &l.Title, &l.URL)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		links = append(links, l)
	}
	return links
}

func (h *Handler) getPostApplicationFees(id int64) (fees []model.ApplicationFee) {

	rows, err := h.DB.Query(`SELECT id, title, body FROM fees WHERE post_id = ?`, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var f model.ApplicationFee
		err = rows.Scan(&f.ID, &f.Title, &f.Body)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fees = append(fees, f)
	}
	return fees
}

func (h *Handler) getPostAgeLimits(id int64) (items []model.GeneralItem) {

	rows, err := h.DB.Query(`SELECT id, title, body FROM items WHERE post_id = ?`, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var i model.GeneralItem
		err = rows.Scan(&i.ID, &i.Title, &i.Body)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		items = append(items, i)
	}
	return items
}

func (h *Handler) getPostAgeVacancies(id int64) (vacancies []model.VacancyItem) {

	rows, err := h.DB.Query(`
	SELECT id, category, name, gen, obc, bca, bcb, ews, sc, st, ph, total, age_limit, eligibility FROM vacancies WHERE post_id = ?`, id)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		var v model.VacancyItem
		err = rows.Scan(&v.ID, &v.Category, &v.Name, &v.Gen, &v.OBC, &v.BCA, &v.BCB, &v.EWS, &v.SC, &v.ST, &v.PH, &v.Total, &v.AgeLimit, &v.Eligibility)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		vacancies = append(vacancies, v)
	}
	return vacancies
}

// UpdatePost method
func (h *Handler) UpdatePost(c echo.Context) (err error) {
	id := c.Param("id")

	p := &model.Post{}

	if err = c.Bind(p); err != nil {
		return
	}

	sql := `UPDATE posts SET short_link = ?, image_link = ?, type = ?, title = ?, name = ?, info = ?, created_date = ?, updated_date = ?, organisation = ?, total_vacancy = ?, age_limit_as_on = ?, draft = ?, trash = ?, updated_at = ? WHERE id = ?`

	_, err = h.DB.Exec(sql, p.ShortLink, p.ImageLink, p.Type, p.Title, p.Name, p.Info, p.CreatedDate, p.UpdatedDate, p.Organisation, p.TotalVacancy, p.AgeLimitAsOn, p.Draft, p.Trash, time.Now(), id)
	if err != nil {
		fmt.Println(err)
		return
	}

	sql = `UPDATE dates SET date = ?, title = ? WHERE id = ?`
	for _, v := range p.Dates {
		fmt.Println(v.Date, v.Title, v.ID)
		_, err = h.DB.Exec(sql, v.Date, v.Title, v.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	sql = `UPDATE links SET title = ?, url = ? WHERE id = ?`
	for _, v := range p.Links {
		_, err = h.DB.Exec(sql, v.Title, v.URL, v.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	sql = `UPDATE items SET title = ?, body = ? WHERE id = ?`
	for _, v := range p.AgeLimits {
		_, err = h.DB.Exec(sql, v.Title, v.Body, v.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	sql = `UPDATE fees SET title = ?, body = ? WHERE id = ?`
	for _, v := range p.Fees {
		_, err = h.DB.Exec(sql, v.Title, v.Body, v.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	sql = `UPDATE vacancies SET category = ?, name = ?, gen = ?, obc = ?, bca = ?, bcb = ?, ews = ?, sc = ?, st = ?, ph = ?, total = ?, age_limit = ?, eligibility = ? WHERE id = ?`
	for _, v := range p.Vacancies {
		_, err = h.DB.Exec(sql, v.Category, v.Name, v.Gen, v.OBC, v.BCA, v.BCB, v.EWS, v.SC, v.ST, v.PH, v.Total, v.AgeLimit, v.Eligibility, v.ID)
		if err != nil {
			fmt.Println(err)
		}
	}

	return c.NoContent(http.StatusOK)
}

// DeletePost method
func (h *Handler) DeletePost(c echo.Context) (err error) {
	id := c.Param("id")

	_, err = h.DB.Exec("UPDATE posts SET trash = true WHERE id = ?", id)

	if err != nil {
		return
	}
	return c.NoContent(http.StatusOK)
}

// GetPostShortInfo  getting the short info for the posts
func (h *Handler) GetPostShortInfo(c echo.Context) (err error) {
	qp := c.QueryParam("type")
	var sql string
	if qp == "" {
		sql = "SELECT id, type, title, updated_date FROM posts WHERE trash = false ORDER BY id DESC"
	} else {
		sql = "SELECT id, type, title, updated_date FROM posts WHERE type LIKE ? AND trash = false ORDER BY id DESC"
	}

	posts := []*model.PostShortInfo{}

	rows, err := h.DB.Query(sql, qp)
	if err != nil {
		return
	}

	for rows.Next() {
		var p model.PostShortInfo
		err = rows.Scan(&p.ID, &p.Type, &p.Title, &p.UpdatedDate)
		if err != nil {
			log.Println(err)
		}
		posts = append(posts, &p)
	}
	return c.JSON(http.StatusOK, posts)
}

// SearchPost for searching post
func (h *Handler) SearchPost(c echo.Context) (err error) {
	q := c.QueryParam("q")
	posts := []*model.PostShortInfo{}
	if err != nil {
		log.Println(err)
	}

	rows, err := h.DB.Query(`
	SELECT id, type, title, updated_date FROM posts WHERE trash = false AND id IN
(SELECT id FROM posts_search WHERE posts_search MATCH ?)`, q)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var p model.PostShortInfo
		err = rows.Scan(&p.ID, &p.Type, &p.Title, &p.UpdatedDate)
		if err != nil {
			log.Println(err)
		}
		posts = append(posts, &p)
	}
	return c.JSON(http.StatusOK, posts)
}
