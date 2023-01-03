FreenomBot:
{{ range $User := .Users}}
{{ $User.UserName }}:
{{ range $Domain := $User.Domains}}{{ $Domain.DomainName }}剩余{{ $Domain.Days }}天
{{ end }}
{{ end }}
