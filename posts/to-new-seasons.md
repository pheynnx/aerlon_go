---
Title: To New Seasons
Date: September 13, 2023
Slug: to-new-seasons
Series: blog
Categories:
  - blog
  - golang
  - http
  - aerlon
  - goldmark
Published: true
Featured: true
PostSnippet: My blog/http website is finally published; while still alpha and there is a lot of progress to be made, I want to start pushing content I actually care about. This post will explain some of the technical aspects behind the sites backend design and the work I have done over the years on this concept.
---

Alright... the blog/site is finally in a running state! After many many years and code bases, I think I have finally gotten everything to a state in which I want to progress. This is many years in the making and many many hours of coding to get to what this site is now. I know this sounds ridiculous as the site is simple, but there is a lot happening behind the scenes and I have worked through a lot of different languages and code bases. This site to me is my personal coding project that I don’t think will ever be done. I don’t want to finish it…

As a learner we should always be growing, testing and developing what we learn into our systems. I couldn’t even imagine what this site code base would look like three years ago, and I am excited to see what it becomes in the next three years.
So, to layout everything, this is my personal blog website. It is very much a WIP and will forever be WIP. This project is also a big part of my resume and portfolio. I do currently work in IT, but not in coding/programming. Maybe one day I will work in software engineering, and I think I would like to, but this site is a way to keep growing my skills and show what I know.

With that, expect the site to change a lot of time and things will come and go. It will probably be written in a new language of framework and some point (I have already designed over 20 dramatically different code bases for this site). But for right now I want to share some code snippets of the backend and explain what’s happening here. First, I will lay out some logically processes of the site that won’t change if the code base changes. Then I will lay out some code from the current code base and talk about it. Let’s go!

### Logical Concepts

#### Post system

All the posts on the site are database models/structs and come from a database (I will describe that in the next section). Below is the schema of the post model in the database. A lot of blog sites are just markdown file parsers. That just means someone is generating `.markdown` files and the server reads the files and converts them to `html` to be shown on a webpage. My site currently doesn’t parse markdown files, it used to… I will explain below. One day it may be a dual system and allow for uploading markdown files.

```sql
CREATE TABLE
  public.post (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    date timestamp without time zone NOT NULL,
    slug text NOT NULL,
    title text NOT NULL,
    series text NOT NULL,
    categories text [] NOT NULL,
    markdown text NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT current_timestamp(),
    updated_at timestamp without time zone NOT NULL DEFAULT current_timestamp(),
    published boolean NOT NULL
  );
```

#### Dual Database System

This is where the money is at, and why the website is so fast. No really… it’s very fast! Like stupid fast! I have looked at a lot of blog sites and different design patterns over the years. And one of the things I have noticed is most blog sites parse markdown because it’s so easy to write and is very human readable. You write your post then the site parses all the information from your post and you generate the `html` from it. But one of the issues I have seen with a site that is designed to parse markdown files, is your reading files from some file system. You have to store these files somewhere and whenever a user requests your post information; your code base has to make a call to the file system then parse all that string information. So, I had an idea, what if I put all that string information into a cache that’s really fast, faster than calling a file system over and over. Can I just pull for a set of cache every time there is a request? Then I started thinking about how to manage this data and how this would work with a caching system. Most markdown blog sites just store their markdown along with their source code; when you want to make a new post you just write to your source code, push to GitHub for example, then pull to your server. So that approach is great and all, and it’s very simple but what if I want to quickly edit a post? What if I want to start a post and not push it live to the site yet? What if I don’t want to expose my server to the public so I can ssh in and pull from my source? So now I have two problems to solve, I want to cache my posts and I want an easy way to manage them with flexibility.

And so, the dual database system with a web-based administrator console was birthed. I will touch more on the admin console below, but this is my solution to manage the content easily. On the server database side of things, the site uses a PostgreSQL cloud-based server and a local Redis caching server.

##### PostgreSQL

One of the issues with a caching system is to achieve high speeds of data transfer, information needs to usually be store in volatile memory that isn’t permanent storage. Obviously, this isn’t going to work for a blog that wants to store years of posts haha. So caching was a great way to get the speed I was looking for, but I needed a way to keep this data protected and safe. This is where I looked to cloud based PostgreSQL systems. In this instance I looked to CockroachDB to be my more permanent storage system where all sources of state truth would be stored. They have a great free tier, and it is very easy to work with, so thank you CockroachDB! All my state models are SQL tables in CockroachDB and state changes to data all flow through this database. Here is a snippet of querying the database for a post by id.

```rust
impl Post {
    pub async fn get_published_post_by_id_postgres(
        postgres_pool: &Pool<Postgres>,
        post_id: &str,
    ) -> Result<Self, AppError> {
        let id = Uuid::parse_str(&post_id)?;

        let post = query_as!(Post,
r#"select id as "post_id?", date, slug, title, series, categories, markdown,
published, created_at as "post_created_at?", updated_at as "post_updated_at?"
from post where id = $1 and published = true"#, &id)
            .fetch_one(postgres_pool)
            .await?;

        Ok(post)
    }
}
```

##### Redis

On the flip side, I need the caching system to get the full speed of ram storage and make sure posts are delivered to the end user as fast as possible. Redis is a super great choice for caching string data and it has a very simple API. I even used generics to deserialize the Redis cache into different Rust structs. Here is what that looks like below, and I will explain this a little.

```rust
impl RedisConnection {
    pub async fn set_cache_redis<T: serde::Serialize>(
        &mut self,
        model: &Vec<T>,
    ) -> Result<(), AppError> {
        let serial = serde_json::to_string(model)?;

        let _ = self.redis_connection.set("posts_cache", serial).await?;

        Ok(())
    }

    pub async fn get_cache_redis<T: for<'de> serde::Deserialize<'de>>(
        &mut self,
    ) -> Result<Vec<T>, AppError> {
        let cache: String = self.redis_connection.get("posts_cache").await?;

        let posts: Vec<T> = serde_json::from_str(&cache)?;

        Ok(posts)
    }
}
```

So, for right now my Rust codebase is just using the native Redis driver, but this may switch to BB8 at some point, I will need to test performance with large HTTP request pressure. But if you look at this code snippet above, I have two methods on the Redis connection itself: one that sets the cache with generics and one that gets the cache with generics. So, it’s the very common getter and setter design pattern, just with generics. The generics aren’t really needed for the setting of the cache as only one struct model will every set the cache, but my code base currently has three different model shapes that can be deserialized out of the cache. Also side note, this side is very much alpha… just a production worthy alpha haha… but at some point this `get_cache_redis<T: for<'de> serde::Deserialize<'de>>` method will only take an empty trait for `<T>` that subscribes to `serde::Deserialize` and my models that can be deserialized implement that trait. Currently this method is allowing any `Type` that implements `serde::Deserialize` to be passed in… obviously that is very error prone. So, both the setter and getter methods will have more restrictive generics in the future.

So when the server boots up or changes are made to the database on the PostgreSQL side, the cache in Redis is updated to the latest snapshot of data. Then the frontend site consumes all its data from the cache, which is sometimes six times faster when you compare HTTP requests to the cloud database.

#### Admin console

The admin console on the site is my solution to managing and handling all the database content from a single location. The admin webpage is a combination of backend HTML rendering and a frontend SPA (single page application) framework. The admin page is actually two ‘single page applications’ that are being served at two different endpoints by the Rust Axum server.

```rust
#[derive(Template)]
#[template(path = "compiled/admin.html")]
struct AdminTemplate {}

pub async fn admin_handler() -> Result<impl IntoResponse, AppError> {
    Ok(HtmlTemplate(AdminTemplate {}))
}

#[derive(Template)]
#[template(path = "compiled/admin_login.html")]
struct AdminLoginTemplate {}

pub async fn get_admin_login_handler() -> Result<impl IntoResponse, AppError> {
    Ok(HtmlTemplate(AdminLoginTemplate {}))
}

#[derive(Serialize, Deserialize, Debug)]
#[allow(dead_code)]
pub struct LoginInput {
    password: String,
    pin: String,
}
```

Here you can see the two different endpoints being served on the backend.

WIP
