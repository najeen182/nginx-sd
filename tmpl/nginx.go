package tmpl

const PlainHTTPTemplate = `
server {
	listen 80;
	server_name {{ .Domain }};

	location / {
		return 404;
	}

	location /.well-known/acme-challenge/ {
        root /var/www/html;
    }
}
`

const HTTPsTemplate = `
upstream {{ .Domain }}_upstream {
    {{- range .Upstreams }}
    server {{ . }};
    {{- end }}
}
server {
	listen 443 ssl;
	server_name {{ .Domain }};

	location / {
		
	    {{- range $key, $value := .ProxyHeaders }}
        proxy_set_header {{ $key }} {{ $value }};
        {{- end }}
		
	}
	ssl_certificate {{ .Cert }};
	ssl_certificate_key {{ .CertKey }};

}
`
