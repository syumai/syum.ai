package server

const mainHTML = `<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <title>syum.ai</title>
  <link rel="stylesheet" href="https://unpkg.com/normalize.css@8.0.1/">
  <style>
.avatar-container, .avatar {
  width: 180px;
  height: 180px;
}
  </style>
</head>
<body>
  <h2>syum.ai</h2>
  <div class="avatar-container">
	  <img src="/image" class="avatar">
  </div>
  <h3>Routes</h3>
  <ul>
    <li><a href="/ascii">/ascii</a></li>
    <li><a href="/image">/image</a></li>
  </ul>
  <h3>Links</h3>
  <ul>
    <li><a href="https://twitter.com/__syumai">Twitter</a></li>
    <li><a href="https://github.com/syumai">GitHub</a></li>
    <li><a href="https://github.com/syumai/syum.ai">Repo</a></li>
  </ul>
</body>
</html>`

const syumaiASCIIArt = `
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMssssssssssssssssssssssssssssyyyyMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMssssssssssssssssssssssssssssyyyyMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMssssssssssss////////ssssyyyyssssyyyyMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMssssssssssss////////ssssyyyyssssyyyyMMMMMMMMMMMMMMMMMMMM
MMMMMMMMhhhhoooohhhhhhhhssssssss--------////ssssssssyyyyhhhhhhhhoooohhhhMMMMMMMM
MMMMMMMMhhhhoooohhhhhhhhssssssss--------////ssssssssyyyyhhhhhhhhoooohhhhMMMMMMMM
MMMMhhhhoooohhhh        ssssssss--------////ssssssssyyyy....    hhhhoooohhhhMMMM
MMMMhhhhoooohhhh        ssssssss--------////ssssssssyyyy....    hhhhoooohhhhMMMM
MMMMhhhhhhhh            ssssssss////////////ssssssssyyyy....        hhhhhhhhMMMM
MMMMhhhhhhhh            ssssssss////////////ssssssssyyyy....        hhhhhhhhMMMM
MMMMhhhhhhhh                ssssssssssssssssssssyyyy....            hhhhhhhhMMMM
MMMMhhhhhhhh                ssssssssssssssssssssyyyy....            hhhhhhhhMMMM
MMMMhhhhhhhh                    ssssssssssssyyyy....                hhhhhhhhMMMM
MMMMhhhhhhhh                    ssssssssssssyyyy....                hhhhhhhhMMMM
MMMMhhhh    hhhhhhhh                ............            hhhhhhhh....hhhhMMMM
MMMMhhhh    hhhhhhhh                ............            hhhhhhhh....hhhhMMMM
MMMMhhhh            hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh        ....hhhhMMMM
MMMMhhhh            hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh        ....hhhhMMMM
MMMMhhhh                                                            ....hhhhMMMM
MMMMhhhh                                                            ....hhhhMMMM
MMMMhhhh                MMMM                    MMMM                ....hhhhMMMM
MMMMhhhh                MMMM                    MMMM                ....hhhhMMMM
MMMMhhhh                MMMM                    MMMM                ....hhhhMMMM
MMMMhhhh                MMMM                    MMMM                ....hhhhMMMM
MMMMhhhh                                                            ....hhhhMMMM
MMMMhhhh                                                            ....hhhhMMMM
MMMMhhhh                ++++++++++++++++++++++++++++                ....hhhhMMMM
MMMMhhhh                ++++++++++++++++++++++++++++                ....hhhhMMMM
MMMMhhhh                    ++++++++++++++++++++                    ....hhhhMMMM
MMMMhhhh                    ++++++++++++++++++++                    ....hhhhMMMM
MMMMhhhh                                                        ........hhhhMMMM
MMMMhhhh                                                        ........hhhhMMMM
MMMMoooohhhhhhhh                                    ............hhhhhhhhssssMMMM
MMMMoooohhhhhhhh                                    ............hhhhhhhhssssMMMM
MMMMhhhhhhhhsssshhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhsssshhhhhhhhMMMM
MMMMhhhhhhhhsssshhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhsssshhhhhhhhMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
`
