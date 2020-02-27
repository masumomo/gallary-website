<!DOCTYPE html>
<html>
  <head>
    <title>Using beego Framework</title> 
    <link href="//maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">
    <link rel="stylesheet" href="https://unpkg.com/aos@next/dist/aos.css" />
    <link rel="stylesheet" href="/static/css/style.css" />
    <script src="//maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js"></script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
  </head>
  <body>
    <div class="wrapper">

      <header class="header">
        <h1>Hello World</h1>
        <a href="./login"><button type="button" class="btn btn-outline-dark">Login</button></a>
      </header>
    
      <div class="main">
        {{ with .Photos }}
          {{ range . }}
            <figure class="snip1577" data-aos="fade-down">
              <img src="data:image/jpg;base64,{{.Src}}" alt="{{ .Name }}">
              <figcaption>
                <h3>{{ .Name }}</h3>
                <h4>M.M</h4>
              </figcaption>
              <a href="#">
              </a>
            </figure>
          {{ end }} 
        {{ end }} 
      </div>

      <aside class="aside aside-1">
        <div class="card">
          <div class="card-body">
            <h5 class="card-title">Something!</h5>
            <p class="card-text">With supporting text below as a natural lead-in to additional content. buraburaburabura...</p>
            <a href="#" class="btn btn-outline-dark">Go somewhere</a>
          </div>
        </div>
      </aside>

    </div>
    <script src="https://unpkg.com/aos@next/dist/aos.js"></script>
    <script>
      AOS.init();
    </script>
  </body>
</html>
