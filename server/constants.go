package server

const indexHTML = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1" />
    <title>syum.ai</title>
    <link
      rel="stylesheet"
      href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.5/css/bulma.min.css"
    />
    <link rel="stylesheet" href="./css/index.css" />
  </head>
  <body>
    <section class="hero is-primary is-small">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            syum.ai
          </h1>
        </div>
      </div>
    </section>
    <section class="section">
      <div class="container">
        <div class="content avatar-content">
          <div class="avatar-container">
            <img src="" id="avatar" class="is-invisible" />
          </div>
          <div>
            Share your random-generated syumai<br />
            <a href="#" class="is-small" id="sharedUrl"></a>
          </div>
          <div>
            <button id="reloadButton" class="button">Reload image</button>
          </div>
        </div>
        <h2 class="title is-4">Profile</h2>
        <ul class="content">
          <li>Name: syumai</li>
          <li>Job: Web application developer</li>
          <li>Location: Tokyo, Japan</li>
          <li>Skills: Go, JavaScript, TypeScript, Ruby</li>
          <li>Twitter: <a href="https://twitter.com/__syumai">@__syumai</a></li>
          <li>GitHub: <a href="https://github.com/syumai">@syumai</a></li>
          <li>Mastodon: <a href="https://mstdn.jp/@_syumai">@_syumai@mstdn.jp</a></li>
          <li>Blog: <a href="https://syumai.hateblo.jp/">syumai.hateblo.jp</a></li>
          <li>Articles: <a href="https://zenn.dev/syumai">zenn.dev/syumai</a></li>
        </ul>
        <h2 class="title is-4">Works</h2>

        <div class="columns">
          <div class="column is-4">
            <div class="card">
              <header class="card-header">
                <p class="card-header-title">
                  Deno (TypeScript)
                </p>
              </header>
              <div class="card-content">
                <ul class="content">
                  <li>
                    <a href="https://github.com/syumai/dinatra" target="_blank"
                      >dinatra</a
                    >: Small web app framework for Deno
                  </li>
                  <li>
                    <a href="https://github.com/syumai/dejs" target="_blank"
                      >dejs</a
                    >: EJS engine for Deno
                  </li>
                  <li>
                    <a href="https://github.com/syumai/dem" target="_blank"
                      >dem</a
                    >: Simple module version manager for Deno
                  </li>
                </ul>
              </div>
            </div>
          </div>

          <div class="column">
            <div class="card">
              <header class="card-header">
                <p class="card-header-title">
                  React.js / Vue.js
                </p>
              </header>
              <div class="card-content">
                <ul class="content">
                  <li>
                    <a
                      href="https://github.com/syumai/go-playground-addons"
                      target="_blank"
                      >go-playground-addons</a
                    >: Chrome extension for the Go Playground implemented in
                    React Hooks + TypeScript
                  </li>
                  <li>
                    <a href="https://github.com/syumai/trollo" target="_blank"
                      >trollo</a
                    >: Simple Trello like kanban app implemented in Vue + Vuex
                  </li>
                </ul>
              </div>
            </div>
          </div>

          <div class="column">
            <div class="card">
              <header class="card-header">
                <p class="card-header-title">
                  Go
                </p>
              </header>
              <div class="card-content">
                <ul class="content">
                  <li>
                    <a
                      href="https://github.com/syumai/protocat"
                      target="_blank"
                    >protocat</a
                    >: A CLI tool to concatenate multiple proto files
                  </li>
                  <li>
                    <a
                      href="https://github.com/syumai/go-hyperscript"
                      target="_blank"
                      >go-hyperscript</a
                    >: Simple Virtual DOM implemented in Go + WebAssembly
                  </li>
                  <li>
                    <a
                      href="https://github.com/syumai/go-wasm-todo-list"
                      target="_blank"
                      >go-wasm-todo-list</a
                    >: TODO list web app implemented in Go + WebAssembly
                  </li>
                  <li>
                    <a
                      href="https://github.com/syumai/syumaigen"
                      target="_blank"
                      >syumaigen</a
                    >: CLI tool to generate syumai's avatar image
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>

        <div class="columns">
          <div class="column is-4">
            <div class="card">
              <header class="card-header">
                <p class="card-header-title">
                  Ruby
                </p>
              </header>
              <div class="card-content">
                <ul class="content">
                  <li>
                    <a href="https://github.com/syumai/yabass" target="_blank"
                      >yabass</a
                    >: Static site generator
                  </li>
                </ul>
              </div>
            </div>
          </div>

          <div class="column is-4">
            <div class="card">
              <header class="card-header">
                <p class="card-header-title">
                  Others
                </p>
              </header>
              <div class="card-content">
                <ul class="content">
                  <li>
                    <a href="https://gpgsync.herokuapp.com">Go Playground with coedit mode</a>
                  </li>
                  <li>
                    <a href="/2020">2020</a>: New year content of 2020
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>

        <h2 class="title is-4">Repo</h2>
        <ul class="content">
          <li>
            <a href="https://github.com/syumai/syum.ai"
              >https://github.com/syumai/syum.ai</a
            >
          </li>
        </ul>
      </div>
    </section>
    <script src="./js/index.js" defer></script>
	<script>
      window.intercomSettings = {
    	app_id: "di99phht"
      };
    </script>

    <script>
    (function(){var w=window;var ic=w.Intercom;if(typeof ic==="function"){ic('reattach_activator');ic('update',w.intercomSettings);}else{var d=document;var i=function(){i.c(arguments);};i.q=[];i.c=function(args){i.q.push(args);};w.Intercom=i;var l=function(){var s=d.createElement('script');s.type='text/javascript';s.async=true;s.src='https://widget.intercom.io/widget/di99phht';var x=d.getElementsByTagName('script')[0];x.parentNode.insertBefore(s,x);};if(document.readyState==='complete'){l();}else if(w.attachEvent){w.attachEvent('onload',l);}else{w.addEventListener('load',l,false);}}})();
    </script>
  </body>
</html>
`

const SyumaiASCIIArt = `MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
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
