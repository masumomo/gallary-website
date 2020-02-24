<!DOCTYPE html>
<html>
<head>
    <title>Using beego Framework</title>
</head>
<body>
   <h1>Hello World</h1>
   <table border="1" style="width:100%;">
        {{ range.employees }}
        <tr>
            <td>{{.ID}}</td>
            <td>{{.FirstName}}</td>
            <td>{{.LastName}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>