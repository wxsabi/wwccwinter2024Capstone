<hr>

### Tue Jun 4 15:25:20 2024 -0700
#### Author: Joaquín Cisneros <jaycisneros@jacm.io>
#### commit: `74d67efcd71d6e22c372534f9d8f4c75354d74e8`


[Update]: Items can be displayed, users can login/logout

Time spent: 15 hrs


<hr>

### Sun Jun 2 23:48:51 2024 -0700
#### Author: Joaquín Cisneros <jaycisneros@jacm.io>
#### commit: `d65574daff837038be502459da3c6a9a072fec4e`


[Update]: User can now sign up with form

Time spent: 10 hrs

- need to add a success message


<hr>

### Sun Jun 2 02:07:51 2024 -0700
#### Author: Joaquín Cisneros <jaycisneros@jacm.io>
#### commit: `0a07235e89fcd21ed5b0e9a18706637ccfb9e2bd`


[Update]: Index.html now shows all items

Time spent:

- Added JS to show all items in index.html


<hr>

### Thu May 16 15:11:18 2024 -0700
#### Author: Joaquín Cisneros <jaycisneros@jacm.io>
#### commit: `d03f2d608ea71b8453ede36933fc3232850c1a6e`


[Fixed]: Fixed bug for DELETE User

Time spent: 3 hrs

- Fixed bug where a User could not be deleted due to foreign key constraints


<hr>

### Thu May 16 12:11:54 2024 -0700
#### Author: Joaquín Cisneros <jaycisneros@jacm.io>
#### commit: `83a1771394357280ccf2b9abb0ff00dcca55d67c`


[Add]: Added ItemPhotos db table

Time spent: 2 hrs

- Create an ItemPhotos table to show photos of the items based on ItemID => Users(ID)


<hr>

### Wed May 15 22:25:11 2024 -0700
#### Author: Joaquín Cisneros <jaycisneros@jacm.io>
#### commit: `76329609129e4cbc91d637e4c0225aac07bced0f`


[Addition]: Added Photos table to the InitDb() function

Time spent: 1hr

- Added the photos table to the initdb function


<hr>

### Wed May 15 16:10:55 2024 -0700
#### Author: Joaquín Cisneros <jaycisneros@jacm.io>
#### commit: `0e547e275684657fb6475467d959f5a785a391bf`


[Fixed]: Fixed but where signup handler panicked

Time spent: 3 hrs

- Fixed signupHandler.go was missing models.InitDB()
- Wrote logoutHanderl.go but I can't test it until the web interface is built
  because the logic is cookie based.


<hr>

### Wed May 15 12:37:39 2024 -0700
#### Author: Joaquín Cisneros <jaycisneros@jacm.io>
#### commit: `e94209f96055994f26484e513d471dc07cdda177`


[Fixed]: Fixed bugs in user handler

Time spent: 2 hrs

- Fixed bugs in user handler to initDB()
- Fixed PUT method in userHandler.go
- created curl_commands.txt to keep commands handy for http requests.


<hr>

### Mon May 13 20:34:39 2024 -0700
#### Author: Joaquín Cisneros <jaycisneros@jacm.io>
#### commit: `19fd5506f2cc5b28a78e316938f5a86c9d087fc4`


[Update]: fixed bugs in signup and sign in

Time spent: 8hrs

- Made sure to fix bugs and test signup and sign in
- refactored db to make sure sessions are handled in their own table
- fixed bug parsing Time.Now using JSON
- can now generate session token and remember token and browser cookie

NEXT:
- Need to refactor code in signin and signup to accomodate the new database structure
- add code for logout


<hr>

### Sat May 4 17:37:43 2024 -0700
#### Author: Joaquín Cisneros <jaycisneros@jacm.io>
#### commit: `b73a21d48a740d68d3fe8e7455479ed7c63dee9f`


[Update]: updated database for users

Time spent: 4hrs

- updated users table to include more fields that were needed
- wrote singin handler
- started implementing user session routine
- tested db to make sure new table works

NEXT: finish signin handler


<hr>

### Thu May 2 18:59:52 2024 -0700
#### Author: Joaquín Cisneros <kincho@jarvis.local>
#### commit: `6fa94816df0a8f3b891a37d14e5250c24d0b7753`


[Added]: Started adding screenshots of AI queries

Time spent: --

- created AI Qeries folder


<hr>

### Thu May 2 18:55:55 2024 -0700
#### Author: Joaquín Cisneros <kincho@jarvis.local>
#### commit: `3ba1ad5f7e7bc9ba0e8b8b236586222e8653cb0a`


[Refactored]: Removed repetitive InitDb() to models

Time spent: 4 hrs

- refactored the InitDb func
- fixed bug that did not allow POST request to signup user

NEXT: create login session and cookies


<hr>

### Tue Apr 30 18:51:02 2024 -0700
#### Author: Joaquín Cisneros <kincho@jarvis.local>
#### commit: `78232b880a7e627f71bb49a50f9d6735e2129cf7`


[Added]: Added User signup functinality

Time spent: 4 hrs

- Wrote signupHandler.go that allows user to log in and uses bcrypt to hash the password.
- Will write signinHandler.go later


<hr>

### Thu Apr 25 10:27:53 2024 -0700
#### Author: Joaquín Cisneros <kincho@jarvis.local>
#### commit: `096f8ca226a2811bb4702b0c20605969d134bbd7`


[Added]: Created admin role

Time spent: 3 hrs

- tested database


<hr>

### Thu Apr 18 17:05:32 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `921737e55c7e285841234f69cd01a8a4521aef4d`


[AUTOMATED]Update CHANGELOG.md


<hr>

### Thu Apr 18 17:05:30 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `98d680ffb6889d16037213ca844c0caa716fed67`


[Added]: Created first draft for the Final paper with 7 sections

Time spent: --

- created first draft for final paper
- Once items in the todo list are completed, the web interface will be created
- Updated the to do list:
	- Add image support for users
	- Add image support for items
	- hash the passwords
	- use bcrypt to let users log in
	- create a verbose mode for debugging with an admin toggle
	- figure out user permissions
	- Overall error handling
	- make db table fields mandatory (ID, email,. etc)
	- add a property to struct for regular users and admins.


<hr>

### Thu Apr 18 16:59:35 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `239a8844e0a0637d5355b9fd2b69c977597088e0`


[Updated]: Added support to userHandler.go for GET, POST, PUT, DELETE

Time spent: 5 hrs

- userHandler.go can now write to the database
- A disctinction needs to be made between regulars users and admins... adding it to the todo list


<hr>

### Thu Apr 18 15:42:00 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `db81ee4187005967360f664b32d7517b76444c90`


[Added]: TO-DO.txt file to list things I need to do

Time spent: N/A

- added a todo list so I don't forget some features.
- moved Db    *sql.DB to the variables file.
- refactored models.Db in itemHandler
- fixed a bug where the time of a item creation could not be parsed in itemHandler


<hr>

### Thu Apr 18 13:45:18 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `8a1a7b6271c9c44d0076f1cf1266ff82867a8b1d`


[Update]: Added functionality for http requests for ItemHandler

Time spent: 20 hrs

- added the line _ "github.com/go-sql-driver/mysql" to import the MySQL driver for Go.
- As of this commit, ItemHandler can GET, POST, DELETE and PUT.
- Added support to write to a MySQL database running in the same system
- Most of the time was spent researching how to work with DB in Go
- Next step is to create/update structs and DB manipulation for users


<hr>

### Mon Mar 25 11:41:40 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `59080b23d0c9ef2d9d70d8f233f7491a5c0a1adf`


[AUTOMATED]Update CHANGELOG.md



<hr>

### Mon Mar 25 11:41:39 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `e6a99388480c8fdd64266f4371c9e7fd9a0716b8`


[Refactoring]: Restructured the folder structure to make more sense.

Time spent: 3 hrs

- Added proper comments to the handler files.



<hr>

### Sat Mar 23 17:07:43 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `9737e4bfcda78ceb6727902c592cbf337ae1d213`


[AUTOMATED]Update CHANGELOG.md



<hr>

### Sat Mar 23 17:07:42 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `3c8fa9448824d37bc03f02054e4d555b23a0f1f9`


Time spent: 1hr

[Update]
- This is the final update to the cacp.sh cript... I hope.
- I just couldn't help having too much fun with this.
- I'm particularly proud of this code:

```bash
# Add all changes to staging
git add .

# Check for changes
if git status --porcelain | grep -q '^[MADRC\?\?]'
then
    echo "Found changes to the local repository"
else
    echo "No changes to commit"
    exit 0
fi
```



<hr>

### Sat Mar 23 15:58:21 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `32c44e3e3cb504e805ac7f40886e2d1d453bc1f2`


[AUTOMATED]Update CHANGELOG.md



<hr>

### Sat Mar 23 15:58:20 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `257ce5ba28b722882f49dcf879c9f7daf3f76c55`


Time spent: 1 minute

[Update]:
- Final manual update to changelog.



<hr>

### Sat Mar 23 15:56:42 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `a5f023488d4cad9a59f13f305299c5269338b959`


[AUTOMATED]Update CHANGELOG.md



<hr>

### Sat Mar 23 15:56:41 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `aa2c4a6cd9f8f46e2eb906ebeee012fafa95fead`


Time spent: 1 minute

[Update]:
- Final manual update to changelog.



<hr>

### Sat Mar 23 15:54:46 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `f576a517126d61fa2735de0b31215503100d1333`


[AUTOMATED]Update CHANGELOG.md



<hr>

### Sat Mar 23 15:54:45 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `193039a7d06cc485a01e158e17a7ff70b6362941`


Time spent: 1 minute

[Update]:
- Added a line jump to script



<hr>

### Sat Mar 23 15:49:57 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `93014cc14fc32facbaa508d487594c28e8ef876d`


[AUTOMATED]Update CHANGELOG.md



<hr>

### Sat Mar 23 15:49:56 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `3e8ae939bd8cb45606b1e0d6528d07291be237dd`


Time spent: 3 hrs

[Upgrade]:
- Upgraded the cacp.sh script to better parse the CHANGELOG.md file



<hr>

### Sat Mar 23 13:11:24 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `5a45bf8b23db2c1441c53fe21ea66b2d9e86ece9`


[AUTOMATED]Update CHANGELOG.md



<hr>

### Sat Mar 23 13:11:23 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `5bdec4365c6fe0145e7fad621e2421ead37430c7`


Time spent: 2 hrs

[Addition]:
- Created the cacp.sh script to automate the process of updating the git repo and the changelog.



<hr>

### Fri Mar 22 22:31:52 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `a4dc4ce1b4d3cbc3ca06585cc342eb892e4c7d56`


Update CHANGELOG.md



<hr>

### Fri Mar 22 22:28:35 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `ff153e584f17421f37fd3dbbeff07b15d365ef09`


[Refactoring]
- I just learned that you can create a changelog out git log...
You live, you learn.
- removed the node module that I was using to generate a changelog
since the script I wrote works better
- Updated .gitignore accordingly.



<hr>

### Fri Mar 22 16:59:23 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `cc0857fd83941e1850c40b0c138aacdffa1132b4`


type(refactor): moved the handler functions for the webserver to their own folder called "handlers". Created a rough sketch of a diagram of the websites main structure. Code refactoring had to be done. - Some more naming and refactoring has to be done for things to be better organized and for separation of concerns. ### Time Spent: 8 hours.



<hr>

### Thu Mar 21 16:22:29 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `80770fc323823074ee39cb97fbbd1f09f5687ebb`


other(misc): Added several files and modules

- Added main.go with the initial implementation of a web server that listens on port 8888
- Added the generate-chagelog node module. From here on, commits are to have the following format: "type(category): description [flags]"
- Added .gitinore file and included the /node_modules folder
- Added the index.html file with a template for the web server
- Next step is to create a user class and an item class and modify the web server to accept both in JSON format



<hr>

### Thu Mar 21 04:35:54 2024 -0700
#### Author: Jay Cisneros <jaycisneros@jacm.io>
#### commit: `6cac1dfcce765da6281767efded80b124d42601e`


Restructure github files



<hr>

### Thu Mar 14 23:27:49 2024 -0700
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `569de1f9cf90eb01ac3af173dc9962e267712a60`


CS231 Final


<hr>

### Thu Mar 14 23:27:30 2024 -0700
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `f3a16b40547a1dd83fd2a87f25f46002d6895973`


Delete CS231 Final.pdf


<hr>

### Thu Mar 14 23:27:20 2024 -0700
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `e215eed7a2e82d0ac30018dccad8dcc48dd80f31`


Delete CS231 Final.docx


<hr>

### Thu Mar 14 23:09:45 2024 -0700
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `e8e6cdee067aaaed3d720368430a711d455eb8e1`


Completed CS231 Final Paper


<hr>

### Thu Mar 14 23:08:56 2024 -0700
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `32275232c5710f0ea4a5be63c6209220202d3557`


Update CHANGELOG.md

Completed CS231 Final Paper


<hr>

### Wed Mar 6 23:57:16 2024 -0800
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `75a561394a42c6616c73f0f05b0534d5650d8a1e`


Added Written proposal


<hr>

### Wed Mar 6 23:56:43 2024 -0800
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `385d3ec55b066975adb3a4be97171fe89508cf5b`


Added written proposal Update CHANGELOG.md

Added: 
- Written Proposal.docx


<hr>

### Wed Mar 6 23:40:40 2024 -0800
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `19832b125506b34e21e4c7857c72eb6d18228544`


Update README.md


<hr>

### Wed Mar 6 23:35:37 2024 -0800
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `ab2f70ed973ce8fb3fa250600b7570f855aa7d46`


Update CHANGELOG.md

Changed from Ad Blocker to College Marketplace


<hr>

### Tue Mar 5 23:53:06 2024 -0800
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `eefddeebdb73c5970803e45f603d73f6d3334b9d`


Added changelog.md file


<hr>

### Tue Mar 5 23:52:26 2024 -0800
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `3e9cd2062aa9e61a85033d3332459c8583c6c5a1`


Update README.md


<hr>

### Tue Mar 5 23:51:29 2024 -0800
#### Author: wxsabi <35592749+wxsabi@users.noreply.github.com>
#### commit: `bfcd4f3a88e68588db7627f6a6a915dac57f93b6`


Initial commit

