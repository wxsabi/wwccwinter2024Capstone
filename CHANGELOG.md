commit a4dc4ce1b4d3cbc3ca06585cc342eb892e4c7d56
Author: Jay Cisneros <jaycisneros@jacm.io>
Date:   Fri Mar 22 22:31:52 2024 -0700

    Update CHANGELOG.md

 CHANGELOG.md | 188 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 1 file changed, 188 insertions(+)

commit ff153e584f17421f37fd3dbbeff07b15d365ef09
Author: Jay Cisneros <jaycisneros@jacm.io>
Date:   Fri Mar 22 22:28:35 2024 -0700

    [Refactoring]
    - I just learned that you can create a changelog out git log...
    You live, you learn.
    - removed the node module that I was using to generate a changelog
    since the script I wrote works better
    - Updated .gitignore accordingly.

 .directory                                  |   2 +-
 .gitignore                                  |   1 -
 CHANGELOG.md                                |   6 ------
 Rough sketch.png => images/Rough sketch.png | Bin
 4 files changed, 1 insertion(+), 8 deletions(-)

commit cc0857fd83941e1850c40b0c138aacdffa1132b4
Author: Jay Cisneros <jaycisneros@jacm.io>
Date:   Fri Mar 22 16:59:23 2024 -0700

    type(refactor): moved the handler functions for the webserver to their own folder called "handlers". Created a rough sketch of a diagram of the websites main structure. Code refactoring had to be done. - Some more naming and refactoring has to be done for things to be better organized and for separation of concerns. ### Time Spent: 8 hours.

 CHANGELOG.md                |   6 ++
 Rough sketch.png            | Bin 0 -> 53263 bytes
 go.mod                      |   3 +
 go.sum                      |   0
 handlers/indexHandler.go    |  20 ++++
 handlers/offeringHandler.go |  25 +++++
 index.html                  |   4 +-
 main.go                     | 246 ++++++++++++++++++--------------------------
 offering/offering.go        |  15 +++
 user/user.go                |  15 +++
 10 files changed, 185 insertions(+), 149 deletions(-)

commit 80770fc323823074ee39cb97fbbd1f09f5687ebb
Author: Jay Cisneros <jaycisneros@jacm.io>
Date:   Thu Mar 21 16:22:29 2024 -0700

    other(misc): Added several files and modules
    
    - Added main.go with the initial implementation of a web server that listens on port 8888
    - Added the generate-chagelog node module. From here on, commits are to have the following format: "type(category): description [flags]"
    - Added .gitinore file and included the /node_modules folder
    - Added the index.html file with a template for the web server
    - Next step is to create a user class and an item class and modify the web server to accept both in JSON format

 .directory                              |   4 +
 .gitignore                              |   1 +
 CHANGELOG.md                            |  22 -----
 index.html                              |  17 ++++
 main.go                                 | 168 ++++++++++++++++++++++++++++++++
 package-lock.json                       |  45 +++++++++
 package.json                            |   5 +
 wwcc/{CHANGELOG.md => OLD_CHANGELOG.md} |   0
 8 files changed, 240 insertions(+), 22 deletions(-)

commit 6cac1dfcce765da6281767efded80b124d42601e
Author: Jay Cisneros <jaycisneros@jacm.io>
Date:   Thu Mar 21 04:35:54 2024 -0700

    Restructure github files

 CHANGELOG.md                                       | 164 +------------------
 wwcc/.directory                                    |   4 +
 wwcc/CHANGELOG.md                                  | 182 +++++++++++++++++++++
 CS231 Final.docx => wwcc/CS231 Final.docx          | Bin
 CS231 Final.pdf => wwcc/CS231 Final.pdf            | Bin
 .../College Marketplace Project Proposal.docx      | Bin
 wwcc/README.md                                     |  19 +++
 7 files changed, 207 insertions(+), 162 deletions(-)

commit 569de1f9cf90eb01ac3af173dc9962e267712a60
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Thu Mar 14 23:27:49 2024 -0700

    CS231 Final

 CS231 Final.docx | Bin 0 -> 20597 bytes
 CS231 Final.pdf  | Bin 0 -> 64364 bytes
 2 files changed, 0 insertions(+), 0 deletions(-)

commit f3a16b40547a1dd83fd2a87f25f46002d6895973
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Thu Mar 14 23:27:30 2024 -0700

    Delete CS231 Final.pdf

 CS231 Final.pdf | Bin 60357 -> 0 bytes
 1 file changed, 0 insertions(+), 0 deletions(-)

commit e215eed7a2e82d0ac30018dccad8dcc48dd80f31
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Thu Mar 14 23:27:20 2024 -0700

    Delete CS231 Final.docx

 CS231 Final.docx | Bin 22364 -> 0 bytes
 1 file changed, 0 insertions(+), 0 deletions(-)

commit e8e6cdee067aaaed3d720368430a711d455eb8e1
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Thu Mar 14 23:09:45 2024 -0700

    Completed CS231 Final Paper

 CS231 Final.docx | Bin 0 -> 22364 bytes
 CS231 Final.pdf  | Bin 0 -> 60357 bytes
 2 files changed, 0 insertions(+), 0 deletions(-)

commit 32275232c5710f0ea4a5be63c6209220202d3557
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Thu Mar 14 23:08:56 2024 -0700

    Update CHANGELOG.md
    
    Completed CS231 Final Paper

 CHANGELOG.md | 12 ++++++++++++
 1 file changed, 12 insertions(+)

commit 75a561394a42c6616c73f0f05b0534d5650d8a1e
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Wed Mar 6 23:57:16 2024 -0800

    Added Written proposal

 College Marketplace Project Proposal.docx | Bin 0 -> 16993 bytes
 1 file changed, 0 insertions(+), 0 deletions(-)

commit 385d3ec55b066975adb3a4be97171fe89508cf5b
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Wed Mar 6 23:56:43 2024 -0800

    Added written proposal Update CHANGELOG.md
    
    Added:
    - Written Proposal.docx

 CHANGELOG.md | 6 +++++-
 1 file changed, 5 insertions(+), 1 deletion(-)

commit 19832b125506b34e21e4c7857c72eb6d18228544
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Wed Mar 6 23:40:40 2024 -0800

    Update README.md

 README.md | 19 +++++++++++++++++--
 1 file changed, 17 insertions(+), 2 deletions(-)

commit ab2f70ed973ce8fb3fa250600b7570f855aa7d46
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Wed Mar 6 23:35:37 2024 -0800

    Update CHANGELOG.md
    
    Changed from Ad Blocker to College Marketplace

 CHANGELOG.md | 18 ++++++++++++++----
 1 file changed, 14 insertions(+), 4 deletions(-)

commit eefddeebdb73c5970803e45f603d73f6d3334b9d
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Tue Mar 5 23:53:06 2024 -0800

    Added changelog.md file

 CHANGELOG.md | 156 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 1 file changed, 156 insertions(+)

commit 3e9cd2062aa9e61a85033d3332459c8583c6c5a1
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Tue Mar 5 23:52:26 2024 -0800

    Update README.md

 README.md | 2 ++
 1 file changed, 2 insertions(+)

commit bfcd4f3a88e68588db7627f6a6a915dac57f93b6
Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
Date:   Tue Mar 5 23:51:29 2024 -0800

    Initial commit

 README.md | 2 ++
 1 file changed, 2 insertions(+)
