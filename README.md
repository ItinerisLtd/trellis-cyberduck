<div align="center">

# trellis-cyberduck

</div>

<div align="center">

[![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/itinerisltd/trellis-cyberduck?include_prereleases&style=flat-square)](https://github.com/ItinerisLtd/trellis-cyberduck/releases)
[![CircleCI Build Status](https://img.shields.io/circleci/build/gh/ItinerisLtd/trellis-cyberduck?style=flat-square)](https://circleci.com/gh/ItinerisLtd/trellis-cyberduck)
[![License](https://img.shields.io/github/license/itinerisltd/trellis-cyberduck.svg?style=flat-square)](https://github.com/itinerisltd/trellis-cyberduck/blob/master/LICENSE)
[![Hire Itineris](https://img.shields.io/badge/Hire-Itineris-ff69b4.svg?style=flat-square)](https://www.itineris.co.uk/contact/)
[![Twitter Follow @itineris_ltd](https://img.shields.io/twitter/follow/itineris_ltd?style=flat-square&color=1da1f2)](https://twitter.com/itineris_ltd)
[![Twitter Follow @TangRufus](https://img.shields.io/twitter/follow/TangRufus?style=flat-square&color=1da1f2)](https://twitter.com/tangrufus)


</div>

<p align="center">
  <strong> Trellis commands for Cyberduck</strong>
  <br />
  <br />
  Built with ♥ by <a href="https://www.itineris.co.uk/">Itineris</a>
</p>

---

## Why

Ain't nobody got time for opening up [Trellis](https://github.com/roots/trellis) files to check SFTP details. This CLI tool allows you to open SFTP connections to Trellis servers via [Cyberduck](https://cyberduck.io/) with single command.

## Usage

```sh-session
$ trellis-cyberduck open --help
Open SFTP connections to Trellis servers

Usage:
  trellis-cyberduck open <environment> [<site>] [flags]

Examples:
  $ trellis-cyberduck open production example.com
  $ trellis-cyberduck open staging my-site --user admin
  $ trellis-cyberduck open staging my-site --directory project_uploads_path

Flags:
  -d, --directory string   Directory to open. Options: project_root|project_source_path|project_uploads_path|project_current_symlink_path (default "project_root")
  -h, --help               help for open
  -u, --user string        User to connect. Options: web|admin (default "web")
```

If you have [trellis-cli](https://github.com/roots/trellis-cli) v0.10.0 or later installed, you can run it as a plugin command like so `$ trellis cyberduck --help`.

---

<div align="center">

**[Itineris](https://www.itineris.co.uk/) is hiring. [Join Us!](https://www.itineris.co.uk/careers/)**

</div>

---

## Install

It goes without saying, you should install [Cyberduck](https://cyberduck.io/) first.

### macOS and Linux via Homebrew

```bash
brew install itinerisltd/tap/trellis-cyberduck
```

### Manual Install

trellis-cyberduck provides binary releases for a variety of OSes. These binary versions can be manually downloaded and installed.

1. Download the [latest release](https://github.com/itinerisltd/trellis-cyberduck/releases/latest) or any [specific version](https://github.com/itinerisltd/trellis-cyberduck/releases)
1. Unpack it (`tar -zxvf trellis-cyberduck_0.1.0_Darwin_x86_64.tar.gz`)
1. Find the `trellis-cyberduck` binary, and move it to its desired destination (`mv trellis-cyberduck /usr/local/bin/trellis-cyberduck`)
1. Make sure the above path is in your `$PATH`

Extra steps for macOS users:

1. Run the command. You should see macOS is blocking it
1. Grant an exception for the binary by clicking the **Open Anyway** button in the General pane of Security & Privacy preferences. This button is available for about an hour after you try to run the command

## FAQs

### trellis-cli plugin command not working

If you have [trellis-cli](https://github.com/roots/trellis-cli) v0.10.0 or later installed, you can run it as a plugin command like so `$ trellis cyberduck [command]`.

However, at the time of writting, [trellis-cli plugin support](https://github.com/roots/trellis-cli/pull/144) hasn't be released yet.

### Can it be used without trellis-cli

Yes.

### It looks awesome. Where can I find some more goodies like this

- Articles on [Itineris' blog](https://www.itineris.co.uk/blog/)
- Projects on [Itineris' GitHub profile](https://github.com/itinerisltd)
- WordPress plugins on [Itineris](https://profiles.wordpress.org/itinerisltd/#content-plugins) and [TangRufus](https://profiles.wordpress.org/tangrufus/#content-plugins) wp.org profiles
- Follow [@itineris_ltd](https://twitter.com/itineris_ltd) and [@TangRufus](https://twitter.com/tangrufus) on Twitter
- **Hire [Itineris](https://www.itineris.co.uk/services/) to build your next awesome site**

### Where can I give 5-star reviews?

Thanks! Glad you like it. It's important to let us knows somebody is using this project. Please consider:

- tweet something good with mentioning [@itineris_ltd](https://twitter.com/itineris_ltd) and [@TangRufus](https://twitter.com/tangrufus)
- ★ star [the Github repo](https://github.com/itinerisltd/trellis-cyberduck)
- [👀 watch](https://github.com/itinerisltd/trellis-cyberduck/subscription) the Github repo
- write tutorials and blog posts
- **[hire Itineris](https://www.itineris.co.uk/services/)**

## Feedback

**Please provide feedback!** We want to make this project as useful as possible.
Please [submit an issue](https://github.com/itinerisltd/trellis-cyberduck/issues/new) and point out what you do and don't like, or fork the project and [send pull requests](https://github.com/itinerisltd/trellis-cyberduck/pulls/).
**No issue is too small.**

## Security Vulnerabilities

If you discover a security vulnerability within this project, please email us at [dev@itineris.co.uk](mailto:dev@itineris.co.uk).
All security vulnerabilities will be promptly addressed.

## Credits

[trellis-cyberduck](https://github.com/itinerisltd/trellis-cyberduck) is a [Itineris Limited](https://www.itineris.co.uk/) project and created by [Tang Rufus](https://www.typist.tech/).

Full list of contributors can be found [here](https://github.com/itinerisltd/trellis-cyberduck/graphs/contributors).

## License

[trellis-cyberduck](https://github.com/itinerisltd/trellis-cyberduck) is released under the [MIT License](https://opensource.org/licenses/MIT).
