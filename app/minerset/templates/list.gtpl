<html>
<body>
{{ range .List }}
    {{ .Title }}

    <a href="{{ .Url }}">{{ .Title }}</a>

    {{ .Url }}
{{ end }}
</body>
</html>