# Wikipass

Wikipass is a tool created for generating very secure and uncrackable passwords. \
It uses MediaWiki (which is a Wikipedia API) to recieve contents of a randomly selected Wikipedia article and convert it into a large set of random words manipulated by some rules later. A password is created by linking a few of those words with a hyphen into one long string of characters that should be resistant to brute-force and dictionary attacks. It cannot be social engineered either since the words are chosen randomly. I highly recommend watching [this video](https://www.youtube.com/watch?v=3NjQ9b3pgIg) by Computerphile if you haven't already. \
With that being said, I cannot guarantee complete security during the generation itself due to reasons listed below:
 * Article contents are sent from MediaWiki (which I know very little about how it's implemented) over HTTP. That might be a vulnerability in public networks.
  * Generated passwords are stored in a file encrypted with 256 bit version of the AES algorithm (Advanced Encryption Standard) which by itself is very secure; however, I haven't put much thought into safe memory use so it is possible that some parts or even the whole key won't be wiped from RAM properly.
  * Some parts of the software use pseudorandom functions. That should not be a big problem because it is used mostly for quick decision-making (eg. "Should this character be modified by the rule set?") but keep that in mind.

Because of those reasons I definitely do not recommend using this program in public networks. In private ones where no one else has access to your computer it should be fine. \
Wikipass also happens to be my final project for the *Programming languages* course I am taking at my University.

# Download
To download this project use the *git clone* command from the terminal:
```bash
$ git clone https://github.com/wedkarz02/wikipass.git
```
or use the "Download ZIP" option from the [Github](https://github.com/wedkarz02/wikipass) page and extract the files once downloaded.

# Requirements
 * [Go v1.20+](https://go.dev/dl/)
 * [raylib-go](https://github.com/gen2brain/raylib-go)
 * [Linux OS](https://ubuntu.com/download)

To download [raylib-go](https://github.com/gen2brain/raylib-go) run the *getrl.sh* script **with sudo privileges** from the terminal:
```bash
$ sudo ./scripts/getrl.sh
```

# Quick Setup
Build the project by 'cd-ing' into the *wikipass/* directory and using the *go build* command from the terminal:
```bash
$ cd wikipass/
$ go build
```
Please keep in mind that the first use of [raylib-go](https://github.com/gen2brain/raylib-go) might take a while to compile. Wikipass requires internet connection while generating new passwords.\
Open the app by running the compiled executable file:
```bash
$ ./wikipass
```

# License
Wikipass is available under the MIT license. See the [LICENSE](https://github.com/wedkarz02/wikipass/blob/main/LICENSE) file for more info.

