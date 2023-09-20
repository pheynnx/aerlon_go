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
- solidjs (js)
- svelte (js)
- sveltekit (js)
- nextjs (js)
- nuxtjs (js)
- giraffe (f#)
- jester (nim)
- prologue (nim)
- karax (nim)
- warp (rust)
- axum (rust)
- asp.net (c#)
- fiber (go)
- echo (go)
- net/http (go)

This isn't in any particular order but it's a rough time span of different stacks I have worked with over the years as this project has evolved. The funny part is, if you go look at the code base for this project it's actually in a pretty simple state. One of the biggest lessons I have learned over the last couple years when building this site, simple is usually better. Don't over complex things just be cool, or fancy. The code to the site is itself a portfolio, so I want to show off cool stuff; but clean correct code is better than a bunch of random spaghetti. And believe me when I say I have written some spaghetti code bases haha.

I love the backend and frontend of web design and think that there is beauty in learning how the two work together. There are millions of ways to do backend and frontend, and lots of libraries and frameworks that are doing a little mix of both now! Lots of cool stuff out there! For this website I wanted to build a statically generated blog site that served HTML extremely fast to the client. With that design in mind a SPA would be out the window lol! So I was looking to generate most if not all my HTML on the backend and serve static pages to the end user; so I am looking for server generation. When I first started coding, Nextjs was just a baby at the time and I was also just a baby coder... and that was probably to my benefit. I fell in love with React (gag! cough! jk!) SPAs pretty quickly and unfortunately learned JavaScript the React way before learning JavaScript. And if I would have known or cared about Nextjs back then, then I probably would have only ever used that for everything. So I am going to count my lucky stars there and be glad I started branching out to other technologies. I am nowhere near a 'good' developer, not even sure what that means, but I would like to think I might be more analytical when approaching a web issue, then just, 'THROW REACT AT IT'! I still like React by the way.

So for my website after a lot of work and a lot of messing around, I think I am settling on a Go backend using the net/http standard library package. I am a big fan of Rust and love how it writes and handles different workflows, but it just doesn't have all the libraries I am looking to work with and it does bring an unnecessary complexity. This post isn't about Go or why I choose Go persay, but Go is a language to get applications done simple and correct. Go is like a Toyota Camry, it will get you do where you need to go everytime; it gets the job done and doesn't cost you too much in gas. Rust is like a maserati; I look good in it and could maybe drive down my driveway before I have to shift gears and break everything.
