package server

var (
	// IndexTemplate is the go template for the index page.
	IndexTemplate = `<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="apple-touch-icon" sizes="57x57" href="/apple-touch-icon-57x57.png">
    <link rel="apple-touch-icon" sizes="60x60" href="/apple-touch-icon-60x60.png">
    <link rel="apple-touch-icon" sizes="72x72" href="/apple-touch-icon-72x72.png">
    <link rel="apple-touch-icon" sizes="76x76" href="/apple-touch-icon-76x76.png">
    <link rel="apple-touch-icon" sizes="114x114" href="/apple-touch-icon-114x114.png">
    <link rel="apple-touch-icon" sizes="120x120" href="/apple-touch-icon-120x120.png">
    <link rel="apple-touch-icon" sizes="144x144" href="/apple-touch-icon-144x144.png">
    <link rel="apple-touch-icon" sizes="152x152" href="/apple-touch-icon-152x152.png">
    <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon-180x180.png">
    <link rel="icon" type="image/png" href="/favicon-32x32.png" sizes="32x32">
    <link rel="icon" type="image/png" href="/android-chrome-192x192.png" sizes="192x192">
    <link rel="icon" type="image/png" href="/favicon-96x96.png" sizes="96x96">
    <link rel="icon" type="image/png" href="/favicon-16x16.png" sizes="16x16">
    <link rel="manifest" href="/manifest.json">
    <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#272f32">
    <meta name="apple-mobile-web-app-title" content="s">
    <meta name="application-name" content="s">
    <meta name="msapplication-TileColor" content="#272f32">
    <meta name="msapplication-TileImage" content="/mstile-144x144.png">
    <meta name="theme-color" content="#272f32">
    <style>
{{.CSS}}
    </style>
    <title>s</title>
  </head>
  <body>
    <form name="search" action="/search" method="POST" onsubmit="return validateForm()">
      <select name="provider" tabindex="2">
{{range .Providers}}         <option{{.|Selected}}>{{.}}</option>
{{end}}      </select><br />
      <input class="input" type="text" name="q" tabindex="1" placeholder="{{.Placeholder}}" autofocus required /><br />
      <input type="submit" value="[ s ]" tabindex="3" />
    </form>
    <script>
{{.JS}}
    </script>
  </body>
</html>
`

	// IndexJS is the javascript for the main index page.
	IndexJS = `function validateForm() {
    var v = document.forms["search"]["q"].value;
    if (v == null || v == "") {
        return false;
    }
}`

	// IndexCSS is the css for the main index page.
	IndexCSS = `*{font-family:"Tahoma","Geneva",sans-serif;font-size:14pt;text-align:center}
body{margin:0;padding:2em;background-color:#272F32;color:#DAEAEF}
a,a:visited{color:#FFF;font-size:.8em}
a:active,a:hover{color:#9DBDC6}
select{position:absolute;top:1.5em;right:1.5em;text-align:left}
option{text-align:left}
form{margin-top:10em}
input[type=text]{width:100%;max-width:450px;border-bottom:1px solid #DAEAEF}
input{color:#DAEAEF;background-color:#272F32;display:inline-block;padding:0 0 .5em;margin:0 0 .5em;border:0;outline:none;border-radius:0;-webkit-border-radius:0;-webkit-appearance:none;appearance:none;-moz-appearance:none}
input:required{box-shadow:none}
input:invalid{box-shadow:none}
input[type=submit]{background-color:#272F32;font-size:2em;transition:color .2s ease}
input[type=submit]:active,input[type=submit]:hover{color:#9DBDC6}
input::-webkit-input-placeholder{font-style:italic}
input::-moz-placeholder{font-style:italic}
input:-moz-placeholder{font-style:italic}
input:-ms-input-placeholder{font-style:italic}
input:focus::-webkit-input-placeholder{color:transparent}
input:focus::-moz-placeholder{color:transparent}
input:focus:-moz-placeholder{color:transparent}
input:focus:-ms-input-placeholder{color:transparent}`
)
