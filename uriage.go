package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/subcommands"
	_ "github.com/mattn/go-sqlite3"
)

type commentsJSON struct {
	Comments []comment
}

type comment struct {
	ID            int64
	PostID        int64 `json:"post_id"`
	UserID        int64 `json:"user_id"`
	Date          string
	Body          *string
	Name          string
	Deleted       bool
	AncestorPath  string `json:"ancestor_path"`
	ChildrenCount int64  `json:"children_count"`
	Children      []comment
}

type article struct {
	ID                      int64
	PublicationID           int64 `json:"publication_id"`
	Title                   string
	SocialTitle             string `json:"social_title"`
	Slug                    string
	PostDate                string `json:"post_date"`
	Audience                string
	WriteCommentPermissions string `json:"write_comment_permissions"`
	CanonicalURL            string `json:"canonical_url"`
	CoverImage              string `json:"cover_image"`
	Description             string
	Wordcount               int64
	Reactions               map[string]int `json:"reactions"`
	ReactionCount           int64          `json:"reaction_count"`
	CommentCount            int64          `json:"comment_count"`
	ChildCommentCount       int64          `json:"child_comment_count"`
}

type articlesCmd struct{}

func (*articlesCmd) Name() string {
	return "articles"
}

func (*articlesCmd) Synopsis() string {
	return "Get all the articles from ACX."
}

func (*articlesCmd) Usage() string {
	return `articles
	Get all the articles from ACX.
`
}

func (p *articlesCmd) SetFlags(f *flag.FlagSet) {}

func (a *articlesCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	getArticles()
	return subcommands.ExitSuccess
}

func main() {
	// f, err := os.ReadFile("comments.json")
	// if err != nil {
	// 	panic(err)
	// }
	// var coj commentsJSON
	// err = json.Unmarshal(f, &coj)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", coj)

	// dir := "articles"

	// f, err := os.Open(dir)
	// if err != nil {
	// 	fmt.Println("Directory open error:", err)
	// 	return
	// }
	// defer f.Close()

	// files, err := f.Readdir(-1) // -1 to read all files
	// if err != nil {
	// 	fmt.Println("Readdir error:", err)
	// 	return
	// }

	// var commentFile commentsJSON

	// for _, file := range files {
	// 	path := filepath.Join(dir, file.Name())
	// 	fmt.Println(path)

	// 	f, err := os.ReadFile(path)
	// 	if err != nil {
	// 		fmt.Println("File read error:", err)
	// 		return
	// 	}

	// 	var articles []article
	// 	err = json.Unmarshal(f, &articles)
	// 	if err != nil {
	// 		fmt.Println("JSON unmarshal error:", err)
	// 		return
	// 	}

	// 	allArticles = append(allArticles, articles...)
	// }

	// fmt.Println(len(allArticles))

	// for _, article := range allArticles {
	// 	getComments(article.ID)
	// 	time.Sleep(1 * time.Second)
	// }

	// a, err := os.Create("cpu.prof")
	// if err != nil {
	// 	panic(err)
	// }
	// if err := pprof.StartCPUProfile(a); err != nil {
	// 	panic(err)
	// }
	// defer pprof.StopCPUProfile()

	// db, err := sql.Open("sqlite3", "./comments.db")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// schema := `
	// CREATE TABLE IF NOT EXISTS comments (
	//     ID INTEGER PRIMARY KEY,
	//     PostID INTEGER,
	//     UserID INTEGER,
	//     Date TEXT,
	//     Body TEXT,
	//     Name TEXT,
	//     AncestorPath TEXT,
	//     ChildrenCount INTEGER
	// );`

	// _, err = db.Exec(schema)
	// if err != nil {
	// 	panic(err)
	// }

	// stmt, err := db.Prepare("INSERT INTO comments (ID, PostID, UserID, Date, Body, Name, AncestorPath, ChildrenCount) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	// if err != nil {
	// 	panic(err)
	// }

	// dir := "comments"

	// f, err := os.Open(dir)
	// if err != nil {
	// 	fmt.Println("Directory open error:", err)
	// 	return
	// }
	// defer f.Close()

	// files, err := f.Readdir(-1) // -1 to read all files
	// if err != nil {
	// 	fmt.Println("Readdir error:", err)
	// 	return
	// }

	// sort.SliceStable(files, func(i, j int) bool {
	// 	return files[i].Name() < files[j].Name()
	// })

	// for _, file := range files {
	// 	start := time.Now()

	// 	tx, err := db.Begin()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	st := tx.Stmt(stmt)

	// 	path := filepath.Join(dir, file.Name())
	// 	fmt.Printf("Processing file: %s\n", path)
	// 	com := flattenCommentJSON(path)

	// 	for _, c := range com {
	// 		_, err := st.Exec(c.ID, c.PostID, c.UserID, c.Date, c.Body, c.Name, c.AncestorPath, c.ChildrenCount)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}

	// 	tx.Commit()
	// 	fmt.Printf("Time taken: %s\n", time.Since(start))
	// }
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&articlesCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

func flattenCommentJSON(path string) []comment {
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var commentFile commentsJSON
	err = json.Unmarshal(f, &commentFile)
	if err != nil {
		panic(err)
	}

	return flattenComments(commentFile.Comments)
}

func flattenComments(comments []comment) []comment {
	// fmt.Printf("Comments: %+v\n", comments)
	if len(comments) == 0 {
		return []comment{}
	}

	var output []comment

	for _, c := range comments {
		output = append(output, c)
		output = append(output, flattenComments(c.Children)...)
	}

	return output
}

func insertComments(db *sql.DB, comments []comment) error {
	stmt, err := db.Prepare("INSERT INTO comments (ID, PostID, UserID, Date, Body, Name, AncestorPath, ChildrenCount) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, c := range comments {
		_, err := stmt.Exec(c.ID, c.PostID, c.UserID, c.Date, c.Body, c.Name, c.AncestorPath, c.ChildrenCount)
		if err != nil {
			panic(err)
		}

		if err := insertComments(db, c.Children); err != nil {
			panic(err)
		}
	}

	return nil
}

func getCommentsFromArticles() {
	if err := os.MkdirAll("articles", 0755); err != nil {
		fmt.Println("Directory creation error:", err)
		return
	}

	if err := os.MkdirAll("comments", 0755); err != nil {
		fmt.Println("Directory creation error:", err)
		return
	}

	dir := "articles"

	f, err := os.Open(dir)
	if err != nil {
		fmt.Println("Directory open error:", err)
		return
	}
	defer f.Close()

	files, err := f.Readdir(-1) // -1 to read all files
	if err != nil {
		fmt.Println("Readdir error:", err)
		return
	}

	var allArticles []article

	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		fmt.Println(path)

		f, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("File read error:", err)
			return
		}

		var articles []article
		err = json.Unmarshal(f, &articles)
		if err != nil {
			fmt.Println("JSON unmarshal error:", err)
			return
		}

		allArticles = append(allArticles, articles...)
	}

	fmt.Println(len(allArticles))

	for _, article := range allArticles {
		getComments(article.ID)
		time.Sleep(1 * time.Second)
	}
}

func getComments(articleID int64) {
	url := fmt.Sprintf("https://www.astralcodexten.com/api/v1/post/%d/comments?token=&all_comments=true&sort=oldest_first", articleID)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	filePath := fmt.Sprintf("comments/article_%d.json", articleID)
	if err := os.WriteFile(filePath, body, 0644); err != nil {
		fmt.Println("File write error:", err)
		return
	}
}

func getArticles() {
	if err := os.MkdirAll("articles", 0755); err != nil {
		fmt.Println("Directory creation error:", err)
		return
	}

	offset := 0
	baseURL := `https://www.astralcodexten.com/api/v1/archive?sort=new&search=&offset=%d&limit=12`

	for {
		url := fmt.Sprintf(baseURL, offset)
		fmt.Println(url)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(body))

		if string(body) == "[]" {
			break
		}

		filePath := fmt.Sprintf("articles/article_offset_%d.json", offset)
		if err := os.WriteFile(filePath, body, 0644); err != nil {
			fmt.Println("File write error:", err)
			return
		}

		offset += 12

		time.Sleep(1 * time.Second)
	}
}
