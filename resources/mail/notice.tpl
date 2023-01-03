Dear there：
0:RenewNo, 1:RenewYes, 3:RenewErr
{{ range $User := .Users}}刚刚帮您检查账户 {{ $User.UserName }} 所有域名情況如下：
{{ range $Domain := $User.Domains}}ID:{{ $Domain.ID}}，{{ $Domain.DomainName }}还有{{ $Domain.Days }}天到期，状态RenewState {{ $Domain.RenewState }}
{{ end }}{{ end }}
更多信息可以參考 "https://www.freenom.com/" Freenom 官网~
freenomBot