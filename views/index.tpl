<!DOCTYPE html>
<html>
  <head>
    <title>Using beego Framework</title>
  </head>
  <body>
    <h1>Hello World</h1>
    <a href="./login">Login</a>
    <table border="1" style="width:100%;">
      	{{ with .Photos }}
          {{ range . }}
      			<tr>
              	<td>{{ .Name }}</td>
              	<td> <img src="data:image/jpg;base64,{{.Src}}" alt="image"></td>
      			</tr>
			{{ end }} 
			{{ end }} 
    </table>
  </body>
</html>
