<html>
<head>
<title>login</title>
</head>
<body>
<form action="/login" method="post">
	<input type="checkbox" name="interest" value="football">足球
	<input type="checkbox" name="interest" value="basketball">篮球
	<input type="checkbox" name="interest" value="tennis">网球	
	email:<input type="text" name="email">
	密码:<input type="password" name="password">
	<input type="hidden" name="token" value="{{.}}">
	<input type="submit" value="登录">
</form>
<form enctype="multipart/form-data" action="/upload" method="post">
  <input type="file" name="uploadfile" />
  <input type="hidden" name="token" value="{{.}}"/>
  <input type="submit" value="upload" />
</form>
</body>
</html>