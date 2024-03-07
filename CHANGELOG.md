# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.8.0] - 2024-03-06

### Changed

- markdown is now being parsed by a child nodejs process so I can use shikiji
  - shikiji is a js only package that I want to run server side to improve static page render times
  - shikiji is also modified to apply direct light and dark color css elements on parse html, which simplifies css
  - shikiji also parses code syntax much better and more accurately than any other markdown parser I have used before

## [0.7.0] - 2023-12-18

### Changed

- more big ui changes
  - big adjustments to code blocks
    - light and dark mode code syntax colors

## [0.6.2] - 2023-12-03

### Added

- starting work with all things `/admin` route
- added a tsconfig.json file where I might have built javascript
  - currently all javascript is just written inside the html tags as its slim and fast
  - not sure if I will ever use a javascript build step, but I have it in place if needed

## [0.6.1] - 2023-11-28

### Changed

- posts now pull from a database for caching, instead of local files
  - web api will need to be built out now
- minor css changes

## [0.6.0] - 2023-11-28

### Changed

- 'data-theme' attribute is now set on the html tag
- html background svg now overlays on top of a gradient

### Added

- dark and light theme toggles
- css variables for theming (wip)

## [0.5.4] - 2023-11-28

### Changed

- preload of font file was adjusted to quicksand.tff

### Fixed

- sticky navbar was added back; not sure how this code was lost
  - checked commit history and can't find where this css was removed??

## [0.5.3] - 2023-11-27

### Changed

- more ui changes
- move to beta track

## [0.5.2] - 2023-11-24

### Changed

- make compact mode a little tighter
- update theming colors

## [0.5.1] - 2023-11-24

### Added

- compact toggle for the index post cards
- localStorage check on theme
  - this is setting up for the possibility of a dark/light theme toggle

## [0.5.0] - 2023-11-24

### Changed

- all html is rendered in real time with templ now

## [0.4.7] - 2023-11-22

### Changed

- more ui changes

## [0.4.6] - 2023-11-21

### Changed

- more ui changes

## [0.4.5] - 2023-11-12

### Changed

- more ui changes

## [0.4.4] - 2023-11-12

### Changed

- more ui changes

## [0.4.3] - 2023-11-11

### Added

- starting testing on transitioning to templ over pongo
  - allowing for in language syntax control and much easier to use with HTMx for paritals

## [0.4.2] - 2023-11-11

## [0.4.1] - 2023-11-08

## [0.4.0] - 2023-11-08

## [0.3.3] - 2023-09-23

## [0.3.2] - 2023-09-22

## [0.3.1] - 2023-09-20

## [0.3.0] - 2023-09-19

## [0.2.3] - 2023-09-13

## [0.2.2] - 2023-09-13

## [0.2.1] - 2023-09-13

## [0.2.0] - 2023-09-13

## [0.1.2] - 2023-09-12

## [0.1.2] - 2023-09-11

## [0.1.1] - 2023-09-11

## [0.1.0] - 2023-09-06

## [0.0.22] - 2023-09-05

## [0.0.21] - 2023-09-04

## [0.0.20] - 2023-09-03

## [0.0.19] - 2023-09-02

## [0.0.18] - 2023-09-01

## [0.0.17] - 2023-09-01

## [0.0.16] - 2023-09-01

## [0.0.15] - 2023-08-29

## [0.0.14] - 2023-08-29

## [0.0.13] - 2023-08-28

## [0.0.12] - 2023-08-28

#### Changed

- linting with revive now

## [0.0.11] - 2023-08-27

## [0.0.10] - 2023-08-27

## [0.0.9] - 2023-08-27

## [0.0.8] - 2023-08-23

## [0.0.7] - 2023-08-23

## [0.0.6] - 2023-08-21

## [0.0.5] - 2023-08-16

## [0.0.4] - 2023-08-15

## [0.0.3] - 2023-08-14

## [0.0.2] - 2023-08-11

## [0.0.1] - 2023-08-10

- init commit
