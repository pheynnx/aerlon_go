---
Title: Aerlon Blog Site
Date: September 20, 2023
Slug: aerlon-blog
Series: blog
Categories:
  - blog
  - golang
  - http
  - aerlon
  - goldmark
Published: true
Featured: true
PostSnippet: ""
---

It has been way to long... and after maybe fifth different iterations of the this site and multiple years... I may, haha, have this as the final backend system for the blog site! I am a little bit of a perfecionist when it comes to coding and have put my hands on many differnet http backend/frontend systems over the last couple of years. Here are some of the frameworks I have worked with.

- flask (python)
- express (node)
- koa (node)
- oak (deno)
- react (js)
- vue (js)
- preact (js)
- solidjs (js)
- svelte (js)
- sveltekit (js)
- nextjs (js)
- nuxtjs (js)
- jester (nim)
- prologue (nim)
- karax (nim)
- asp.net (c#)
- giraffe (f#)
- warp (rust)
- axum (rust)
- fiber (go)
- echo (go)
- net/http (go)

This isn't in any particular order but it's a rough time span of different stacks I have worked with over the years as this project has evolved. The funny part is, if you go look at the code base for this project it's actually in a pretty simple state. One of the biggest lessons I have learned over the last couple years when building this site, simple is usually better. Don't over complex things just be cool, or fancy. The code to the site is itself a portfolio, so I want to show off cool stuff; but clean correct code is better than a bunch of random spaghetti. And believe me when I say I have written some spaghetti code bases haha.

I love the backend and frontend of web design and think that there is beauty in learning how the two work together. There are millions of ways to do backend and frontend, and lots of libraries and frameworks that are doing a little mix of both now! Lots of cool stuff out there! For this website I wanted to build a statically generated blog site that served HTML extremely fast to the client. With that design in mind a SPA would be out the window lol! So I was looking to generate most if not all my HTML on the backend and serve static pages to the end user; so I am looking for server generation. When I first started coding, Nextjs was just a baby at the time and I was also just a baby coder... and that was probably to my benefit. I fell in love with React (gag! cough! jk!) SPAs pretty quickly and unfortunately learned JavaScript the React way before learning JavaScript. And if I would have known or cared about Nextjs back then, then I probably would have only ever used that for everything. So I am going to count my lucky stars there and be glad I started branching out to other technologies. I am nowhere near a 'good' developer, not even sure what that means, but I would like to think I might be more analytical when approaching a web issue, then just, 'THROW REACT AT IT'! I still like React by the way.

So for my website after a lot of work and a lot of messing around, I think I am settling on a Go backend using the net/http standard library package. I am a big fan of Rust and love how it writes and handles different workflows, but it just doesn't have all the libraries I am looking to work with and it does bring an unnecessary complexity. This post isn't about Go or why I choose Go persay, but Go is a language to get applications done simple and correct. Go is like a Toyota Camry, it will get you to where you need to go everytime; it gets the job done and doesn't cost you too much in gas. Rust is like a Maserati; I look good in it and could maybe drive down my driveway before I have to shift gears and break everything. Enough of that though, lets look at how this blog site works a little on the backend.

First I want to look at the current post struct model and how the posts are being parsed. About a year ago I designed my first caching system for the blog posts, in I guess you could call it design version one.

#### Version 1.0 of the caching system

This concept was database reliant, actually dual database reliant. I wanted a way to edit and design posts and store them in a cloud solution but make queries extremely fast for end users. If my posts were all in a cloud database, when a user queried the index root page or a blog post this would take a dramatic amount of time. The cloud database would have to be called, then all the data mapped to a structure of some sort. Then the markdown would need to be parsed to HTML. Lastly, all the data would need to be injected and generated into HTML template strings to be served by the backend. There are a lot of heavy calculation steps here; and waiting for a database connection over the internet to return this data could take awhile; plus if the online service is down... my site is down... or at least not displaying any data. So I got the maybe smart idea to cache all the online database data into a redis cache on the startup of the server, then I could make calls to the redis dataset whenever needed for a very fast in-memory/local solution. Here is what this would look like in order.

1. Start HTTP server
2. Make call to online database
3. Map data to language structures
4. Parse the markdown for the dataset into HTML and update the struct field in place
5. Store these structures as object/json strings in redis
6. When an end user called for a index or a blog post, query redis and generate the HTML template

#### Version 1.5 of the caching system

Then I started realizing the downfalls of this kind of set. I would have to run a redis service on whatever server I was using. I also would have to make sure the online dataset and the redis cache were always synced when changes to the dataset were made. One of the biggest pros to this setup though, was I could build a frontend API to the online database and update posts through the browser in realtime and then update the cache on the server. This would allow for zero reloading of the backend server and I could update posts in real time! And that last part is what kept me on this design for a long time. I wanted the ability to update a post from a `/admin` page on the server at any given time and any location; zero manual server reloads required. I also switched from a redis server to just storing the data sets in a map on the application itself; this eliminated a reliance on a redis service. I wrote a frontend SPA in signal for the `/admin` admin dashboard and connected it to an API on my HTTP server. These API endpoints made direct CRUD/REST changes to the online based PostgreSQL server. Whenever I would make a change to the online database data, I would then just update the in-memory cache map.

#### Version 2.0 of the caching system

And then I stopped and really thought hard about it one day, simple is usually better. Don't try to recreate the wheel. A dual database caching system with a frontend admin console is cool and all... but is it needed? Is it really serving my needs well? Am I in constant need to quickly update a post on the go? Is updating markdown data in an HTML textarea tag a great experience? And so I made the decision to keep it simple and just head back to writing all my posts in .markdown files and storing them alongside my code source. This is a simple and basic solution. No need to overcomplicate everything. Buttttttt! I still wanted my site to be extremely fast to the end user; and I wanted all the posts and the home page index list to be extremely fast as well! So I am still implementing a startup caching system that generates in-memory HTML strings. I am not generating statif HTML files. I am generating in-memory strings in a map; and these strings just happen to be render HTML templates. I did some benchmarking and prototyping over the last couple of months, and I noticed that serving a large string that is stored in a map is faster then serving a file, and its definitly faster then an HTML template engine rendering the HTML on demand. There is a time and place for on demand HTML rendering, but in my case my current set of HTML is better static and prerendered. Going back, serving an HTML file from an HTTP endpoint will always only be as fast as the file I/O of your server or computer. But serving a string from a variable is constrained to the speed of your memory. Now... please don't take my words as gospel, in my current setup this system works for me. There are pros and cons to most different things in all areas of life. Okay enough of that, lets look at some code snippets of the posts and the caching system.

##### Post struct

```go
type Post struct {
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Slug        string    `json:"slug"`
	Series      string    `json:"series"`
	Categories  []string  `json:"categories"`
	Markdown    string    `json:"markdown"`
	Published   bool      `json:"published"`
	Featured    bool      `json:"featured"`
	PostSnippet string    `json:"post_snippet"`
}
```

So the post struct contains all the metadata of the post itself, the markdown of the post, and lastly some state flags. These fields are set by the parsing of .markdown files. Lets look at the parsing of the files and see how the markdown is parsed to HTML and also how the struct fields are set.

```go
func ParseMarkdownAndMeta(content []byte) (*Post, error) {
	var buf bytes.Buffer
	cxt := parser.NewContext()
	err := md.Convert(content, &buf, parser.WithContext(cxt))
	if err != nil {
		return nil, err
	}

	meta := meta.Get(cxt)

	date, err := time.Parse("January 2, 2006", meta["Date"].(string))
	if err != nil {
		return nil, err
	}

	var categories []string
	for _, c := range meta["Categories"].([]interface{}) {
		categories = append(categories, c.(string))
	}

	slices.Sort(categories)

	return &Post{
		Title:       meta["Title"].(string),
		Date:        date,
		Slug:        meta["Slug"].(string),
		Series:      meta["Series"].(string),
		Categories:  categories,
		Markdown:    buf.String(),
		Published:   meta["Published"].(bool),
		Featured:    meta["Featured"].(bool),
		PostSnippet: meta["PostSnippet"].(string),
	}, nil
}
```

So looking at the function above, lets focus on the fourth line that says `err := md.Convert(content, &buf, parser.WithContext(cxt))`. There is a `md` variables with a `Convert()` method on it that then writes to a couple buffers. Lets look at the `md` variable below.

```go
var md = goldmark.New(
	goldmark.WithExtensions(
		meta.Meta,
		extension.GFM,
		highlighting.NewHighlighting(
			highlighting.WithWrapperRenderer(func(w util.BufWriter, context highlighting.CodeBlockContext, entering bool) {
				lang, _ := context.Language()

				if entering {
					if lang == nil {
						w.WriteString("<pre><code>")
						return
					}
					w.WriteString(fmt.Sprintf(`<div class="code-block"><p class="code-block-header"><span class="language-name">%s</span></p><pre class="aer"><code class="language-`, lang))
					w.Write(lang)
					w.WriteString(`" data-lang="`)
					w.Write(lang)
					w.WriteString(`">`)
				} else {
					if lang == nil {
						w.WriteString("</pre></code>")
						return
					}
					w.WriteString(`</code></pre></div>`)
				}
			}),
			highlighting.WithFormatOptions(
				chromahtml.PreventSurroundingPre(true),
				chromahtml.WithClasses(true),
			),
		),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
)
```

So the `md` variables is an interface that returns a goldmark struct. It is from the project 'github.com/yuin/goldmark'. This is a really awesome markdown parser with a lot of flexibility and extensions. I am using the custom meta extension and the custom highlighting extension provided through chroma. The meta extension is really cool and allows me to parse out meta from the top of a .markdown file. The meta on this blog post file looks like this snippet below.

```markdown
---
Title: Aerlon Blog Site
Date: September 20, 2023
Slug: aerlon-blog
Series: blog
Categories:
  - blog
  - golang
  - http
  - aerlon
  - goldmark
Published: true
Featured: true
---
```

The `---` three hypen marks at the top and bottom of the meta data denote where this data starts and ends. The meta extension then parses the data using a YAML format; and as this is pretty much impossible to map to types in Go, these return `interface{}` types values. Going back up three snippets to where I am mapping the data to my Post struct in Go, you can see the returned value from the parsed meta is of the type `map[string]interface{}`. I cut out that part in the snippet below, but you can see where the `Title` value is being set with `meta["Title"].(string)`; the `.(string)` part is called a type assertion in Go.

```go
return &Post{
  Title:       meta["Title"].(string),
  Date:        date,
  Slug:        meta["Slug"].(string),
  Series:      meta["Series"].(string),
  Categories:  categories,
  Markdown:    buf.String(),
  Published:   meta["Published"].(bool),
  Featured:    meta["Featured"].(bool),
  PostSnippet: meta["PostSnippet"].(string),
}
```

##### Work in Progress | 9-20-23
