<!DOCTYPE html>
<html>
  <head>
    <title>Using beego Framework</title>
  </head>
  <body>
    <h1>Upload</h1>
    <h2>I made believe to be logged in :)</h2>
    <table border="1" style="width:100%;">
      <form enctype="multipart/form-data" action="./upload" method="post">
        <input type="file" name="myFile" />
        <input type="submit" value="upload" />
      </form>
    </table>
  </body>
</html>
